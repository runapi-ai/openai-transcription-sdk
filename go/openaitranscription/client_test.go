package openaitranscription

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"reflect"
	"testing"
	"time"

	"github.com/runapi-ai/core-sdk/go/core"
	"github.com/runapi-ai/core-sdk/go/option"
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

func (s *stubHTTPClient) RequestWithResponse(ctx context.Context, method, path string, opts *core.HTTPRequestOptions) (*core.HTTPResponse, error) {
	payload, err := s.Request(ctx, method, path, opts)
	if err != nil {
		return nil, err
	}
	return &core.HTTPResponse{Body: payload, StatusCode: http.StatusOK, Header: make(http.Header)}, nil
}

func TestSpeechToTextRunPreservesAcceptedRawResponse(t *testing.T) {
	file := filepath.Join(t.TempDir(), "interview.mp3")
	if err := os.WriteFile(file, []byte("audio"), 0o600); err != nil {
		t.Fatal(err)
	}
	requests := 0
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requests++
		switch requests {
		case 1:
			if r.Method != http.MethodPost || r.URL.Path != speechToTextPath || r.Header.Get("Idempotency-Key") == "" {
				t.Fatalf("unexpected create request: %s %s", r.Method, r.URL.Path)
			}
			w.Header().Set("Location", "/api/v1/tasks/transcription_task")
			w.Header().Set("Retry-After", "0")
			w.WriteHeader(http.StatusAccepted)
			_, _ = w.Write([]byte(`{"id":"transcription_task","status":"processing"}`))
		case 2:
			_, _ = w.Write([]byte(`{"id":"transcription_task","status":"completed","response":{"status":200,"content_type":"text/vtt; charset=utf-8","headers":{},"body":"WEBVTT\n\n00:00.000 --> 00:01.000\nHello\n"}}`))
		default:
			t.Fatalf("unexpected request %d", requests)
		}
	}))
	defer server.Close()

	client, err := NewClient(option.WithAPIKey("test-key"), option.WithBaseURL(server.URL))
	if err != nil {
		t.Fatal(err)
	}
	result, err := client.SpeechToText.Run(context.Background(), SpeechToTextParams{File: file, ResponseFormat: "vtt"}, option.WithPollInterval(time.Millisecond))
	if err != nil {
		t.Fatal(err)
	}
	if string(result) != "WEBVTT\n\n00:00.000 --> 00:01.000\nHello\n" {
		t.Fatalf("unexpected terminal bytes: %q", result)
	}
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
		ResponseFormat: "text"})
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
		File: "/tmp/interview.mp3"}); err != nil {
		t.Fatal(err)
	}

	body := stub.body.(core.MultipartBody)
	if body.Fields["model"] != DefaultModel {
		t.Fatalf("expected model %q, got %#v", DefaultModel, body.Fields["model"])
	}
}
