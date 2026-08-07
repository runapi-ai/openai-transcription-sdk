package openaitranscription

import (
	"context"
	"encoding/json"
	"reflect"
	"testing"

	"github.com/runapi-ai/core-sdk/go/core"
)

type stubHTTPClient struct {
	method   string
	path     string
	body     any
	response json.RawMessage
}

func (s *stubHTTPClient) Request(_ context.Context, method, path string, opts *core.HTTPRequestOptions) (json.RawMessage, error) {
	s.method = method
	s.path = path
	s.body = opts.Body
	return s.response, nil
}

func TestSpeechToTextRunSendsRepeatedMultipartFieldsAndPreservesRawResponse(t *testing.T) {
	stub := &stubHTTPClient{response: json.RawMessage("hello from audio\n")}
	client := NewClientWithHTTP(stub)
	temperature := 0.25

	response, err := client.SpeechToText.Run(context.Background(), SpeechToTextParams{
		File:           "/tmp/interview.mp3",
		Model:          "gpt-transcribe",
		Languages:      []string{"en", "zh"},
		Keywords:       []string{"RunAPI", "release notes"},
		Temperature:    &temperature,
		ResponseFormat: "text",
	})
	if err != nil {
		t.Fatal(err)
	}
	if stub.method != "POST" || stub.path != speechToTextPath {
		t.Fatalf("unexpected request: %s %s", stub.method, stub.path)
	}
	body := stub.body.(core.MultipartBody)
	if !reflect.DeepEqual(body.RepeatedFields["languages[]"], []string{"en", "zh"}) || !reflect.DeepEqual(body.RepeatedFields["keywords[]"], []string{"RunAPI", "release notes"}) {
		t.Fatalf("unexpected repeated fields: %#v", body.RepeatedFields)
	}
	if body.Fields["model"] != "gpt-transcribe" || body.Fields["temperature"] != "0.25" {
		t.Fatalf("unexpected fields: %#v", body.Fields)
	}
	if body.Files["file"].FileName != "interview.mp3" || body.Files["file"].ContentType != "audio/mpeg" {
		t.Fatalf("unexpected file: %#v", body.Files["file"])
	}
	if string(response) != "hello from audio\n" {
		t.Fatalf("unexpected raw response: %q", response)
	}
}

// An omitted Model is dropped by CompactParams, so without the default the
// contract's model enum rejects the request instead of applying it.
func TestSpeechToTextRunDefaultsModelWhenOmitted(t *testing.T) {
	stub := &stubHTTPClient{response: json.RawMessage(`{"text":"ok"}`)}
	client := NewClientWithHTTP(stub)

	if _, err := client.SpeechToText.Run(context.Background(), SpeechToTextParams{
		File: "/tmp/interview.mp3",
	}); err != nil {
		t.Fatal(err)
	}

	body := stub.body.(core.MultipartBody)
	if body.Fields["model"] != DefaultModel {
		t.Fatalf("expected model %q, got %#v", DefaultModel, body.Fields["model"])
	}
}
