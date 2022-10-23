package main

// var (
// 	version = "dev"
// 	commit  = "none"
// 	date    = "unknown"
// 	builtBy = "unknown"
// )

import (
	"github.com/tliron/glsp"
	protocol "github.com/tliron/glsp/protocol_3_16"
	"github.com/tliron/glsp/server"
	"github.com/tliron/kutil/logging"

	// Must include a backend implementation. See kutil's logging/ for other options.
	_ "github.com/tliron/kutil/logging/simple"
)

const lsName = "quanty-lang"

var version string = "0.0.1"
var handler protocol.Handler

func main() {
	// This increases logging verbosity (optional)
	logging.Configure(1, nil)

	handler = protocol.Handler{}

	server := server.NewServer(&handler, lsName, false)

	handler.Initialize = func(context *glsp.Context, params *protocol.InitializeParams) (interface{}, error) {
		// To see the logs with coc.nvim, run :CocCommand workspace.showOutput
		// https://github.com/neoclide/coc.nvim/wiki/Debug-language-server#using-output-channel
		if params.Trace != nil {
			protocol.SetTraceValue(*params.Trace)
		}

		capabilities := handler.CreateServerCapabilities()

		change := protocol.TextDocumentSyncKindIncremental
		capabilities.TextDocumentSync = protocol.TextDocumentSyncOptions{
			OpenClose: boolPtr(true),
			Change:    &change,
			Save:      boolPtr(true),
		}

		triggerChars := []string{"."}

		capabilities.CompletionProvider = &protocol.CompletionOptions{
			TriggerCharacters: triggerChars,
			ResolveProvider:   boolPtr(true),
		}

		capabilities.ReferencesProvider = &protocol.ReferenceOptions{}

		return protocol.InitializeResult{
			Capabilities: capabilities,
			ServerInfo: &protocol.InitializeResultServerInfo{
				Name:    lsName,
				Version: &version,
			},
		}, nil
	}

	handler.Initialized = func(context *glsp.Context, params *protocol.InitializedParams) error {
		return nil
	}

	handler.Shutdown = func(context *glsp.Context) error {
		protocol.SetTraceValue(protocol.TraceValueOff)
		return nil
	}

	handler.SetTrace = func(context *glsp.Context, params *protocol.SetTraceParams) error {
		protocol.SetTraceValue(params.Value)
		return nil
	}

	handler.TextDocumentDidOpen = func(context *glsp.Context, params *protocol.DidOpenTextDocumentParams) error {
		return nil
	}

	handler.TextDocumentDidChange = func(context *glsp.Context, params *protocol.DidChangeTextDocumentParams) error {
		return nil
	}

	handler.TextDocumentDidClose = func(context *glsp.Context, params *protocol.DidCloseTextDocumentParams) error {
		return nil
	}

	handler.TextDocumentDidSave = func(context *glsp.Context, params *protocol.DidSaveTextDocumentParams) error {
		return nil
	}

	handler.TextDocumentCompletion = func(context *glsp.Context, params *protocol.CompletionParams) (interface{}, error) {
		return fakeCompletion(), nil
	}

	handler.CompletionItemResolve = func(context *glsp.Context, params *protocol.CompletionItem) (*protocol.CompletionItem, error) {
		// if path, ok := params.Data.(string); ok {
		// 	content, err := ioutil.ReadFile(path)
		// 	if err != nil {
		// 		return params, err
		// 	}
		// 	params.Documentation = protocol.MarkupContent{
		// 		Kind:  protocol.MarkupKindMarkdown,
		// 		Value: string(content),
		// 	}
		// }

		return params, nil
	}

	server.RunStdio()
}

func fakeCompletion() []protocol.CompletionItem {
	return []protocol.CompletionItem{
		{
			Label: "Hello",
		},
		{
			Label: "Quanty",
		},
		{
			Label: "Team",
		},
	}
}

func boolPtr(v bool) *bool {
	b := v
	return &b
}

func stringPtr(v string) *string {
	s := v
	return &s
}

// func main() {
// 	usr, err := user.Current()
// 	if err != nil {
// 		panic(err)
// 	}
// 	fmt.Printf(`
// Hello %s!
// This is the Quanty Query Language!

// Version: %s
// Commit: %s
// Built at %s by %s`,
// 		usr.Username, version, commit, date, builtBy)
// 	fmt.Println()

// app := server.Run()

// log.Fatal(app.Listen(":8081"))

// 	src := `
// component Main {
// 	div {
// 		p {
// 			span {
// 				strong {
// 					"{{ . }}!"
// 				}
// 			}
// 		}
// 		span
// 	}
// 	"hello world!"
// }

// component Nav {
// 	ul {
// 		li {
// 			"First"
// 		}
// 		li {
// 			"Second"
// 		}
// 		li {
// 			"{{ . }}"
// 		}
// 	}
// }
// `

// 	l := html.New("locale string")

// 	sdl.Run(src, l)

// 	tmpl := l.Lookup("Nav")

// 	tmpl.Execute(os.Stdout, "COUCOU")
// 	fmt.Println()

// err := lsp.Run()
// if err != nil {
// 	panic(err)
// }
// }

// func main() {
// 	server := lsp.NewServer(&lsp.Options{CompletionProvider: &defines.CompletionOptions{
// 		TriggerCharacters: &[]string{"."},
// 	}})
// 	server.OnHover(func(ctx context.Context, req *defines.HoverParams) (result *defines.Hover, err error) {
// 		logs.Println(req)
// 		return &defines.Hover{Contents: defines.MarkupContent{Kind: defines.MarkupKindPlainText, Value: "hello world"}}, nil
// 	})

// 	server.OnCompletion(func(ctx context.Context, req *defines.CompletionParams) (result *[]defines.CompletionItem, err error) {
// 		logs.Println(req)
// 		d := defines.CompletionItemKindText
// 		return &[]defines.CompletionItem{defines.CompletionItem{
// 			Label:      "code",
// 			Kind:       &d,
// 			InsertText: strPtr("Hello"),
// 		}}, nil
// 	})

// 	server.Run()
// }
