package proxy

import (
	"context"

	lsp "go.lsp.dev/protocol"
	"go.uber.org/zap"
)

// Server is responsible for rewriting messages that are
// originated from the text editor, and need to be sent to gopls.
//
// Since the editor is working on `templ` files, and `gopls` works
// on Go files, the job of this code is to rewrite incoming requests
// to adjust the file names from `*.templ` to `*_templ.go` and to
// remap the line/character positions in the `templ` files to their
// corresponding locations in the Go file.
//
// This allows gopls to operate as usual.
//
// This code also rewrites the responses back from gopls to do the
// inverse operation - to put the file names back, and readjust any
// character positions.
type Server struct {
	Log    *zap.Logger
	Client lsp.Client
}

func NewServer(log *zap.Logger) (s *Server, init func(lsp.Client)) {
	s = &Server{
		Log: log,
	}
	return s, func(client lsp.Client) {
		s.Client = client
	}
}

func (p *Server) Initialize(ctx context.Context, params *lsp.InitializeParams) (result *lsp.InitializeResult, err error) {
	p.Log.Info("client -> server: Initialize")
	defer p.Log.Info("client -> server: Initialize end")

	init := &lsp.InitializeResult{
		ServerInfo: &lsp.ServerInfo{
			Name: "Quanty Lang",
		},
		Capabilities: lsp.ServerCapabilities{
			TextDocumentSync: lsp.TextDocumentSyncKindIncremental,
			CompletionProvider: &lsp.CompletionOptions{
				ResolveProvider: true,
			},
		},
	}

	return init, nil

	// Add the '<' and '{' trigger so that we can do snippets for tags.
	// if result.Capabilities.CompletionProvider == nil {
	// 	result.Capabilities.CompletionProvider = &lsp.CompletionOptions{}
	// }
	// result.Capabilities.CompletionProvider.TriggerCharacters = append(result.Capabilities.CompletionProvider.TriggerCharacters, "{", "<")
	// // Remove all the gopls commands.
	// if result.Capabilities.ExecuteCommandProvider == nil {
	// 	result.Capabilities.ExecuteCommandProvider = &lsp.ExecuteCommandOptions{}
	// }
	// result.Capabilities.ExecuteCommandProvider.Commands = []string{}
	// result.Capabilities.DocumentFormattingProvider = true
	// return result, err
}

func (p *Server) Initialized(ctx context.Context, params *lsp.InitializedParams) (err error) {
	p.Log.Info("client -> server: Initialized")
	defer p.Log.Info("client -> server: Initialized end")

	p.Client.LogMessage(ctx, &lsp.LogMessageParams{
		Type:    lsp.MessageTypeInfo,
		Message: "Hello My Client!",
	})

	return
}

func (p *Server) Shutdown(ctx context.Context) (err error) {
	p.Log.Info("client -> server: Shutdown")
	defer p.Log.Info("client -> server: Shutdown end")
	return
}

func (p *Server) Exit(ctx context.Context) (err error) {
	p.Log.Info("client -> server: Exit")
	defer p.Log.Info("client -> server: Exit end")
	return
}

func (p *Server) WorkDoneProgressCancel(ctx context.Context, params *lsp.WorkDoneProgressCancelParams) (err error) {
	p.Log.Info("client -> server: WorkDoneProgressCancel")
	defer p.Log.Info("client -> server: WorkDoneProgressCancel end")
	return
}

func (p *Server) LogTrace(ctx context.Context, params *lsp.LogTraceParams) (err error) {
	p.Log.Info("client -> server: LogTrace", zap.String("message", params.Message))
	defer p.Log.Info("client -> server: LogTrace end")
	return
}

func (p *Server) SetTrace(ctx context.Context, params *lsp.SetTraceParams) (err error) {
	p.Log.Info("client -> server: SetTrace")
	defer p.Log.Info("client -> server: SetTrace end")

	return
}

func (p *Server) CodeAction(ctx context.Context, params *lsp.CodeActionParams) (result []lsp.CodeAction, err error) {
	p.Log.Info("client -> server: CodeAction")
	defer p.Log.Info("client -> server: CodeAction end")
	return
}

func (p *Server) CodeLens(ctx context.Context, params *lsp.CodeLensParams) (result []lsp.CodeLens, err error) {
	p.Log.Info("client -> server: CodeLens")
	defer p.Log.Info("client -> server: CodeLens end")
	return
}

func (p *Server) CodeLensResolve(ctx context.Context, params *lsp.CodeLens) (result *lsp.CodeLens, err error) {
	p.Log.Info("client -> server: CodeLensResolve")
	defer p.Log.Info("client -> server: CodeLensResolve end")
	return
}

func (p *Server) ColorPresentation(ctx context.Context, params *lsp.ColorPresentationParams) (result []lsp.ColorPresentation, err error) {
	p.Log.Info("client -> server: ColorPresentation ColorPresentation")
	defer p.Log.Info("client -> server: ColorPresentation end")
	return
}

func (p *Server) Completion(ctx context.Context, params *lsp.CompletionParams) (result *lsp.CompletionList, err error) {
	p.Log.Info("client -> server: Completion")
	defer p.Log.Info("client -> server: Completion end")
	return
}

func (p *Server) CompletionResolve(ctx context.Context, params *lsp.CompletionItem) (result *lsp.CompletionItem, err error) {
	p.Log.Info("client -> server: CompletionResolve")
	defer p.Log.Info("client -> server: CompletionResolve end")
	return
}

func (p *Server) Declaration(ctx context.Context, params *lsp.DeclarationParams) (result []lsp.Location /* Declaration | DeclarationLink[] | null */, err error) {
	p.Log.Info("client -> server: Declaration")
	defer p.Log.Info("client -> server: Declaration end")
	return
}

func (p *Server) Definition(ctx context.Context, params *lsp.DefinitionParams) (result []lsp.Location /* Definition | DefinitionLink[] | null */, err error) {
	p.Log.Info("client -> server: Definition")
	defer p.Log.Info("client -> server: Definition end")
	return
}

func (p *Server) DidChange(ctx context.Context, params *lsp.DidChangeTextDocumentParams) (err error) {
	p.Log.Info("client -> server: DidChange", zap.Any("params", params))
	defer p.Log.Info("client -> server: DidChange end")
	return
}

func (p *Server) DidChangeConfiguration(ctx context.Context, params *lsp.DidChangeConfigurationParams) (err error) {
	p.Log.Info("client -> server: DidChangeConfiguration")
	defer p.Log.Info("client -> server: DidChangeConfiguration end")
	return
}

func (p *Server) DidChangeWatchedFiles(ctx context.Context, params *lsp.DidChangeWatchedFilesParams) (err error) {
	p.Log.Info("client -> server: DidChangeWatchedFiles")
	defer p.Log.Info("client -> server: DidChangeWatchedFiles end")
	return
}

func (p *Server) DidChangeWorkspaceFolders(ctx context.Context, params *lsp.DidChangeWorkspaceFoldersParams) (err error) {
	p.Log.Info("client -> server: DidChangeWorkspaceFolders")
	defer p.Log.Info("client -> server: DidChangeWorkspaceFolders end")
	return
}

func (p *Server) DidClose(ctx context.Context, params *lsp.DidCloseTextDocumentParams) (err error) {
	p.Log.Info("client -> server: DidClose")
	defer p.Log.Info("client -> server: DidClose end")
	return
}

func (p *Server) DidOpen(ctx context.Context, params *lsp.DidOpenTextDocumentParams) (err error) {
	p.Log.Info("client -> server: DidOpen", zap.String("uri", string(params.TextDocument.URI)))
	defer p.Log.Info("client -> server: DidOpen end")
	return
}

func (p *Server) DidSave(ctx context.Context, params *lsp.DidSaveTextDocumentParams) (err error) {
	p.Log.Info("client -> server: DidSave")
	defer p.Log.Info("client -> server: DidSave end")
	return
}

func (p *Server) DocumentColor(ctx context.Context, params *lsp.DocumentColorParams) (result []lsp.ColorInformation, err error) {
	p.Log.Info("client -> server: DocumentColor")
	defer p.Log.Info("client -> server: DocumentColor end")
	return
}

func (p *Server) DocumentHighlight(ctx context.Context, params *lsp.DocumentHighlightParams) (result []lsp.DocumentHighlight, err error) {
	p.Log.Info("client -> server: DocumentHighlight")
	defer p.Log.Info("client -> server: DocumentHighlight end")
	return
}

func (p *Server) DocumentLink(ctx context.Context, params *lsp.DocumentLinkParams) (result []lsp.DocumentLink, err error) {
	p.Log.Info("client -> server: DocumentLink", zap.String("uri", string(params.TextDocument.URI)))
	defer p.Log.Info("client -> server: DocumentLink end")
	return
}

func (p *Server) DocumentLinkResolve(ctx context.Context, params *lsp.DocumentLink) (result *lsp.DocumentLink, err error) {
	p.Log.Info("client -> server: DocumentLinkResolve")
	defer p.Log.Info("client -> server: DocumentLinkResolve end")
	return
}

func (p *Server) DocumentSymbol(ctx context.Context, params *lsp.DocumentSymbolParams) (result []interface{} /* []SymbolInformation | []DocumentSymbol */, err error) {
	p.Log.Info("client -> server: DocumentSymbol")
	defer p.Log.Info("client -> server: DocumentSymbol end")
	//TODO: Rewrite the request and response, but for now, ignore it.
	return
}

func (p *Server) ExecuteCommand(ctx context.Context, params *lsp.ExecuteCommandParams) (result interface{}, err error) {
	p.Log.Info("client -> server: ExecuteCommand")
	defer p.Log.Info("client -> server: ExecuteCommand end")
	return
}

func (p *Server) FoldingRanges(ctx context.Context, params *lsp.FoldingRangeParams) (result []lsp.FoldingRange, err error) {
	p.Log.Info("client -> server: FoldingRanges")
	defer p.Log.Info("client -> server: FoldingRanges end")
	return []lsp.FoldingRange{}, nil
}

func (p *Server) Formatting(ctx context.Context, params *lsp.DocumentFormattingParams) (result []lsp.TextEdit, err error) {
	p.Log.Info("client -> server: Formatting")
	defer p.Log.Info("client -> server: Formatting end")
	return
}

func (p *Server) Hover(ctx context.Context, params *lsp.HoverParams) (result *lsp.Hover, err error) {
	p.Log.Info("client -> server: Hover")
	defer p.Log.Info("client -> server: Hover end")
	return
}

func (p *Server) Implementation(ctx context.Context, params *lsp.ImplementationParams) (result []lsp.Location, err error) {
	p.Log.Info("client -> server: Implementation")
	defer p.Log.Info("client -> server: Implementation end")
	return
}

func (p *Server) OnTypeFormatting(ctx context.Context, params *lsp.DocumentOnTypeFormattingParams) (result []lsp.TextEdit, err error) {
	p.Log.Info("client -> server: OnTypeFormatting")
	defer p.Log.Info("client -> server: OnTypeFormatting end")
	return
}

func (p *Server) PrepareRename(ctx context.Context, params *lsp.PrepareRenameParams) (result *lsp.Range, err error) {
	p.Log.Info("client -> server: PrepareRename")
	defer p.Log.Info("client -> server: PrepareRename end")
	return
}

func (p *Server) RangeFormatting(ctx context.Context, params *lsp.DocumentRangeFormattingParams) (result []lsp.TextEdit, err error) {
	p.Log.Info("client -> server: RangeFormatting")
	defer p.Log.Info("client -> server: RangeFormatting end")
	return
}

func (p *Server) References(ctx context.Context, params *lsp.ReferenceParams) (result []lsp.Location, err error) {
	p.Log.Info("client -> server: References")
	defer p.Log.Info("client -> server: References end")
	return
}

func (p *Server) Rename(ctx context.Context, params *lsp.RenameParams) (result *lsp.WorkspaceEdit, err error) {
	p.Log.Info("client -> server: Rename")
	defer p.Log.Info("client -> server: Rename end")
	return
}

func (p *Server) SignatureHelp(ctx context.Context, params *lsp.SignatureHelpParams) (result *lsp.SignatureHelp, err error) {
	p.Log.Info("client -> server: SignatureHelp")
	defer p.Log.Info("client -> server: SignatureHelp end")
	return
}

func (p *Server) Symbols(ctx context.Context, params *lsp.WorkspaceSymbolParams) (result []lsp.SymbolInformation, err error) {
	p.Log.Info("client -> server: Symbols")
	defer p.Log.Info("client -> server: Symbols end")
	return
}

func (p *Server) TypeDefinition(ctx context.Context, params *lsp.TypeDefinitionParams) (result []lsp.Location, err error) {
	p.Log.Info("client -> server: TypeDefinition")
	defer p.Log.Info("client -> server: TypeDefinition end")
	return
}

func (p *Server) WillSave(ctx context.Context, params *lsp.WillSaveTextDocumentParams) (err error) {
	p.Log.Info("client -> server: WillSave")
	defer p.Log.Info("client -> server: WillSave end")
	return
}

func (p *Server) WillSaveWaitUntil(ctx context.Context, params *lsp.WillSaveTextDocumentParams) (result []lsp.TextEdit, err error) {
	p.Log.Info("client -> server: WillSaveWaitUntil")
	defer p.Log.Info("client -> server: WillSaveWaitUntil end")
	return
}

func (p *Server) ShowDocument(ctx context.Context, params *lsp.ShowDocumentParams) (result *lsp.ShowDocumentResult, err error) {
	p.Log.Info("client -> server: ShowDocument")
	defer p.Log.Info("client -> server: ShowDocument end")
	return
}

func (p *Server) WillCreateFiles(ctx context.Context, params *lsp.CreateFilesParams) (result *lsp.WorkspaceEdit, err error) {
	p.Log.Info("client -> server: WillCreateFiles")
	defer p.Log.Info("client -> server: WillCreateFiles end")
	return
}

func (p *Server) DidCreateFiles(ctx context.Context, params *lsp.CreateFilesParams) (err error) {
	p.Log.Info("client -> server: DidCreateFiles")
	defer p.Log.Info("client -> server: DidCreateFiles end")
	return
}

func (p *Server) WillRenameFiles(ctx context.Context, params *lsp.RenameFilesParams) (result *lsp.WorkspaceEdit, err error) {
	p.Log.Info("client -> server: WillRenameFiles")
	defer p.Log.Info("client -> server: WillRenameFiles end")
	return
}

func (p *Server) DidRenameFiles(ctx context.Context, params *lsp.RenameFilesParams) (err error) {
	p.Log.Info("client -> server: DidRenameFiles")
	defer p.Log.Info("client -> server: DidRenameFiles end")
	return
}

func (p *Server) WillDeleteFiles(ctx context.Context, params *lsp.DeleteFilesParams) (result *lsp.WorkspaceEdit, err error) {
	p.Log.Info("client -> server: WillDeleteFiles")
	defer p.Log.Info("client -> server: WillDeleteFiles end")
	return
}

func (p *Server) DidDeleteFiles(ctx context.Context, params *lsp.DeleteFilesParams) (err error) {
	p.Log.Info("client -> server: DidDeleteFiles")
	defer p.Log.Info("client -> server: DidDeleteFiles end")
	return
}

func (p *Server) CodeLensRefresh(ctx context.Context) (err error) {
	p.Log.Info("client -> server: CodeLensRefresh")
	defer p.Log.Info("client -> server: CodeLensRefresh end")
	return
}

func (p *Server) PrepareCallHierarchy(ctx context.Context, params *lsp.CallHierarchyPrepareParams) (result []lsp.CallHierarchyItem, err error) {
	p.Log.Info("client -> server: PrepareCallHierarchy")
	defer p.Log.Info("client -> server: PrepareCallHierarchy end")
	return
}

func (p *Server) IncomingCalls(ctx context.Context, params *lsp.CallHierarchyIncomingCallsParams) (result []lsp.CallHierarchyIncomingCall, err error) {
	p.Log.Info("client -> server: IncomingCalls")
	defer p.Log.Info("client -> server: IncomingCalls end")
	return
}

func (p *Server) OutgoingCalls(ctx context.Context, params *lsp.CallHierarchyOutgoingCallsParams) (result []lsp.CallHierarchyOutgoingCall, err error) {
	p.Log.Info("client -> server: OutgoingCalls")
	defer p.Log.Info("client -> server: OutgoingCalls end")
	return
}

func (p *Server) SemanticTokensFull(ctx context.Context, params *lsp.SemanticTokensParams) (result *lsp.SemanticTokens, err error) {
	p.Log.Info("client -> server: SemanticTokensFull")
	defer p.Log.Info("client -> server: SemanticTokensFull end")
	return
}

func (p *Server) SemanticTokensFullDelta(ctx context.Context, params *lsp.SemanticTokensDeltaParams) (result interface{} /* SemanticTokens | SemanticTokensDelta */, err error) {
	p.Log.Info("client -> server: SemanticTokensFullDelta")
	defer p.Log.Info("client -> server: SemanticTokensFullDelta end")
	return
}

func (p *Server) SemanticTokensRange(ctx context.Context, params *lsp.SemanticTokensRangeParams) (result *lsp.SemanticTokens, err error) {
	p.Log.Info("client -> server: SemanticTokensRange")
	defer p.Log.Info("client -> server: SemanticTokensRange end")
	return
}

func (p *Server) SemanticTokensRefresh(ctx context.Context) (err error) {
	p.Log.Info("client -> server: SemanticTokensRefresh")
	defer p.Log.Info("client -> server: SemanticTokensRefresh end")
	return
}

func (p *Server) LinkedEditingRange(ctx context.Context, params *lsp.LinkedEditingRangeParams) (result *lsp.LinkedEditingRanges, err error) {
	p.Log.Info("client -> server: LinkedEditingRange")
	defer p.Log.Info("client -> server: LinkedEditingRange end")
	return
}

func (p *Server) Moniker(ctx context.Context, params *lsp.MonikerParams) (result []lsp.Moniker, err error) {
	p.Log.Info("client -> server: Moniker")
	defer p.Log.Info("client -> server: Moniker end")
	return
}

func (p *Server) Request(ctx context.Context, method string, params interface{}) (result interface{}, err error) {
	p.Log.Info("client -> server: Request")
	defer p.Log.Info("client -> server: Request end")
	return
}
