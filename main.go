package main

import (
	"bufio"
	"encoding/json"
	"io"
	"log"
	"os"
	"stella-lsp/lsp"
	"stella-lsp/parser"
	"stella-lsp/rpc"
)

func main() {
	logger := getLogger("./log.txt")
	logger.Println("Logger Started")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(rpc.Split)

	writer := os.Stdout
	state := parser.NewState()

	for scanner.Scan() {
		message := scanner.Bytes()
		method, contents, err := rpc.DecodeMessage(message)
		if err != nil {
			logger.Println("Error decoding message:", err)
			continue
		}

		handleMessage(logger, writer, state, method, contents)
	}
}

func handleMessage(logger *log.Logger, writer io.Writer, state parser.State, method string, contents []byte) {
	logger.Printf("Message Received with method: %s\n", method)

	switch method {
	case "initialize":
		var request lsp.InitializeRequest
		if err := json.Unmarshal(contents, &request); err != nil {
			logger.Println("Error decoding initialize request: ", err)
			return
		}
		logger.Printf("Initialize request received")

		message := lsp.NewInitializeResponse(request.ID)
		writeResponse(writer, message)
		logger.Println("Initialize request handled")

	case "textDocument/didOpen":
		var request lsp.DidOpenTextDocumentNotification
		if err := json.Unmarshal(contents, &request); err != nil {
			logger.Println("Error decoding textDocument/didOpen request: ", err)
			return
		}
		logger.Printf("Got DidOpen notification for: %s", request.Params.TextDocument.URI)
		state.OpenDocument(request.Params.TextDocument.URI, request.Params.TextDocument.Text)

	case "textDocument/didChange":
		var request lsp.DidChangeTextDocumentNotification
		if err := json.Unmarshal(contents, &request); err != nil {
			logger.Println("Error decoding textDocument/didChange request: ", err)
			return
		}
		logger.Printf("Got DidChange notification for: %s", request.Params.TextDocument.URI)
		for _, change := range request.Params.ContentChanges {
			state.UpdateDocument(request.Params.TextDocument.URI, change.Text)
		}

	case "textDocument/hover":
		var request lsp.HoverRequest
		if err := json.Unmarshal(contents, &request); err != nil {
			logger.Println("Error decoding textDocument/hover request: ", err)
			return
		}
		logger.Printf("Got Hover notification for: %s", request.Params.TextDocument.URI)

		message := state.Hover(request.ID, request.Params.TextDocument.URI, request.Params.Position)
		writeResponse(writer, message)

	case "textDocument/definition":
		var request lsp.DefinitionRequest
		if err := json.Unmarshal(contents, &request); err != nil {
			logger.Println("Error decoding textDocument/definition request: ", err)
		}
		logger.Printf("Got Definition notification for: %s", request.Params.TextDocument.URI)

		message := state.Definition(request.ID, request.Params.TextDocument.URI, request.Params.Position)
		writeResponse(writer, message)

	default:
		logger.Printf("\nUnhandled message with method: %s", method)
	}
}

func writeResponse(writer io.Writer, message any) {
	reply := rpc.EncodeMessage(message)
	writer.Write([]byte(reply))
}

func getLogger(filename string) *log.Logger {
	logFile, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		panic("Failed to open log file")
	}
	return log.New(logFile, "[stella-lsp] ", log.Ldate|log.Ltime|log.Lshortfile)
}
