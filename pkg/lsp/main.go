package lsp

import (
	"context"
	"os"
	"os/signal"

	"github.com/Flo-Leroux/quanty-lang/pkg/lsp/proxy"
	"go.lsp.dev/jsonrpc2"
	"go.lsp.dev/protocol"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Arguments struct {
	Log string
}

// func Run(args Arguments) error {
func Run() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt)
	defer func() {
		signal.Stop(signalChan)
		cancel()
	}()
	go func() {
		select {
		case <-signalChan: // First signal, cancel context.
			cancel()
		case <-ctx.Done():
		}
		<-signalChan // Second signal, hard exit.
		os.Exit(2)
	}()

	args := Arguments{
		Log: "info",
	}

	return run(ctx, args)
}

func run(ctx context.Context, args Arguments) (err error) {
	log := zap.NewNop()
	if args.Log != "" {
		cfg := zap.NewProductionConfig()
		cfg.EncoderConfig.EncodeTime = zapcore.RFC3339TimeEncoder
		cfg.OutputPaths = []string{
			args.Log,
		}
		log, err = cfg.Build()
		if err != nil {
			log.Error("failed to create logger: %v\n", zap.Error(err))
			os.Exit(1)
		}
	}
	defer log.Sync()
	log.Info("lsp: starting up...")
	defer func() {
		if r := recover(); r != nil {
			log.Fatal("handled panic", zap.Any("recovered", r))
		}
	}()

	log.Info("creating client")
	// _, clientInit := proxy.NewClient(log)

	log.Info("creating proxy")
	// Create the proxy to sit between.
	serverProxy, serverInit := proxy.NewServer(log)

	// Create templ server.
	log.Info("creating templ server")
	quantyStream := jsonrpc2.NewStream(stdrwc{log: log})
	_, quantyConn, quantyClient := protocol.NewServer(context.Background(), serverProxy, quantyStream, log)
	defer quantyConn.Close()

	// Allow both the server and the client to initiate outbound requests.
	// clientInit(quantyClient)
	serverInit(quantyClient)

	log.Info("listening")

	select {
	case <-ctx.Done():
		log.Info("context closed")
	case <-quantyConn.Done():
		log.Info("quantyConn closed")
	}
	log.Info("shutdown complete")
	return
}
