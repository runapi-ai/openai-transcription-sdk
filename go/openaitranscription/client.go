package openaitranscription

import (
	"context"
	"encoding/json"
	"mime"
	"path/filepath"
	"strconv"

	"github.com/runapi-ai/core-sdk/go/base"
	"github.com/runapi-ai/core-sdk/go/core"
	"github.com/runapi-ai/core-sdk/go/option"
)

const speechToTextPath = "/v1/audio/transcriptions"

// Client provides OpenAI-compatible audio transcription.
type Client struct {
	base.Base
	SpeechToText *SpeechToText
}

// NewClient creates an OpenAI Transcription client with the given options.
func NewClient(opts ...option.ClientOption) (*Client, error) {
	resolved, err := option.ResolveClientOptions(opts...)
	if err != nil {
		return nil, err
	}
	httpClient, err := core.NewHTTPClient(resolved)
	if err != nil {
		return nil, err
	}
	return NewClientWithHTTP(httpClient), nil
}

// NewClientWithHTTP creates an OpenAI Transcription client with a pre-configured transport.
func NewClientWithHTTP(httpClient core.HTTPClient) *Client {
	return &Client{Base: base.New(httpClient), SpeechToText: &SpeechToText{http: httpClient}}
}

// SpeechToText transcribes uploaded audio files.
type SpeechToText struct{ http core.HTTPClient }

// DefaultModel is the transcription model used when the caller omits one.
const DefaultModel = "whisper-1"

// Run transcribes an audio file synchronously and returns the exact response bytes.
func (r *SpeechToText) Run(ctx context.Context, params SpeechToTextParams, opts ...option.RequestOption) (json.RawMessage, error) {
	requestOptions, _ := option.ResolveRequestOptions(opts...)
	// Fill the documented default before validating: an empty Model is dropped
	// by CompactParams, and the contract's model enum would then reject the
	// request instead of applying the default.
	if params.Model == "" {
		params.Model = DefaultModel
	}
	compact := core.CompactParams(params)
	if err := core.ValidateParams(contractSchema["speech-to-text"], compact); err != nil {
		return nil, err
	}

	fields := make(map[string]string)
	if params.Model != "" {
		fields["model"] = params.Model
	}
	if params.Prompt != "" {
		fields["prompt"] = params.Prompt
	}
	if params.Language != "" {
		fields["language"] = params.Language
	}
	if params.Temperature != nil {
		fields["temperature"] = strconv.FormatFloat(*params.Temperature, 'g', -1, 64)
	}
	if params.ResponseFormat != "" {
		fields["response_format"] = params.ResponseFormat
	}
	if params.Stream != nil {
		fields["stream"] = strconv.FormatBool(*params.Stream)
	}

	contentType := mime.TypeByExtension(filepath.Ext(params.File))
	if contentType == "" {
		contentType = "application/octet-stream"
	}
	body := core.MultipartBody{
		Fields: fields,
		RepeatedFields: map[string][]string{
			"languages[]":               params.Languages,
			"keywords[]":                params.Keywords,
			"timestamp_granularities[]": params.TimestampGranularities,
		},
		Files: map[string]core.MultipartFile{
			"file": {Path: params.File, FileName: filepath.Base(params.File), ContentType: contentType},
		},
	}

	return r.http.Request(ctx, "POST", speechToTextPath, &core.HTTPRequestOptions{Body: body, Request: requestOptions})
}
