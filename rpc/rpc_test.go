package rpc_test

import (
	"stella-lsp/rpc"
	"testing"
)

type EncodingExample struct {
	Testing bool
}

func TestEncodeMessage(t *testing.T) {
	expected := "Content-Length: 16\r\n\r\n{\"Testing\":true}"
	actual := rpc.EncodeMessage(EncodingExample{true})
	if actual != expected {
		t.Fatal("Expected", expected, "but got", actual)
	}
}

func TestDecodeMessage(t *testing.T) {
	incomingMessage := "Content-Length: 15\r\n\r\n{\"Method\":\"hi\"}"
	method, content, err := rpc.DecodeMessage([]byte(incomingMessage))
	if err != nil {
		t.Fatal(err)
	}

	contentLength := len(content)

	if contentLength != 15 {
		t.Fatal("Expected", 15, "but got", contentLength)
	}
	if method != "hi" {
		t.Fatal("Expected hi but got", method)
	}
}
