package lsp

import (
	"fmt"

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

type SourceMap string
type SourceMaps map[string]SourceMap

func Run() {
	// This increases logging verbosity (optional)

	sourceMaps := make(SourceMaps, 0)

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

		// triggerChars := []string{" "}

		capabilities.CompletionProvider = &protocol.CompletionOptions{
			// TriggerCharacters: triggerChars,
			ResolveProvider: boolPtr(true),
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
		if _, ok := sourceMaps[params.TextDocument.URI]; !ok {
			sourceMaps[params.TextDocument.URI] = SourceMap(params.TextDocument.Text)
		}
		return nil
	}

	handler.TextDocumentDidChange = func(context *glsp.Context, params *protocol.DidChangeTextDocumentParams) error {
		return nil
	}

	handler.TextDocumentDidClose = func(context *glsp.Context, params *protocol.DidCloseTextDocumentParams) error {
		delete(sourceMaps, params.TextDocument.URI)
		return nil
	}

	handler.TextDocumentDidSave = func(context *glsp.Context, params *protocol.DidSaveTextDocumentParams) error {
		return nil
	}

	handler.TextDocumentCompletion = func(context *glsp.Context, params *protocol.CompletionParams) (interface{}, error) {
		if txt, ok := sourceMaps[params.TextDocument.URI]; ok {
			caretPosition := params.Position.IndexIn(string(txt))
			return fakeCompletion(txt, caretPosition), nil
		}

		return nil, nil
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

func fakeCompletion(txt SourceMap, caretPosition int) []protocol.CompletionItem {

	return []protocol.CompletionItem{
		{
			Label: fmt.Sprintf("Hello %d", caretPosition),
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
