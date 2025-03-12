package parser

import (
	"fmt"
	"stella-lsp/lsp"
)

// this should do much more when dealing with an actual language

type State struct {
	// Map of file names to contents
	Documents map[string]string
}

func NewState() State {
	return State{Documents: map[string]string{}}
}

func (state *State) OpenDocument(uri, text string) {
	state.Documents[uri] = text
}

func (state *State) UpdateDocument(uri, text string) {
	state.Documents[uri] = text
}

func (state *State) Hover(id int, uri string, position lsp.Position) lsp.HoverResponse {
	// look up the type in the type analysis code
	document := state.Documents[uri]

	return lsp.HoverResponse{
		Response: lsp.Response{
			RPC: "2.0",
			ID:  &id,
		},
		Result: lsp.HoverResult{
			Contents: fmt.Sprintf("The document has %d characters.", len(document)),
		},
	}
}

func (state *State) Definition(id int, uri string, position lsp.Position) lsp.DefinitionResponse {
	// look for the definition of the word underneath the cursor
	document := state.Documents[uri]
	_ = document
	return lsp.DefinitionResponse{
		Response: lsp.Response{
			RPC: "2.0",
			ID:  &id,
		},
		Result: lsp.Location{
			URI: uri,
			Range: lsp.Range{
				Start: lsp.Position{
					Line:      position.Line - 1,
					Character: 0,
				},
				End: lsp.Position{
					Line:      position.Line - 1,
					Character: 0,
				},
			},
		},
	}
}
