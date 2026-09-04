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

// Create submits a transcription and returns either its exact terminal bytes or
// an accepted Task that can be resumed through Subscribe.
func (r *SpeechToText) Create(ctx context.Context, params SpeechToTextParams, opts ...option.RequestOption) (*core.HybridCreateResponse[json.RawMessage], error) {
	requestOptions, _ := option.ResolveRequestOptions(opts...)
	body, err := speechToTextBody(params)
	if err != nil {
		return nil, err
	}
	return core.CreateHybrid[json.RawMessage](ctx, r.http, speechToTextPath, body, requestOptions)
}

// Subscribe follows an accepted transcription Task and returns its exact
// terminal response bytes.
func (r *SpeechToText) Subscribe(ctx context.Context, acceptance *core.HybridAcceptance, opts ...option.RequestOption) (json.RawMessage, error) {
	requestOptions, pollingOptions := option.ResolveRequestOptions(opts...)
	response, err := core.SubscribeHybrid[json.RawMessage](ctx, r.http, acceptance, requestOptions, pollingOptions)
	if err != nil {
		return nil, err
	}
	return *response, nil
}

// Run returns the exact transcription bytes whether the endpoint completes
// directly or first returns 202 Accepted.
func (r *SpeechToText) Run(ctx context.Context, params SpeechToTextParams, opts ...option.RequestOption) (json.RawMessage, error) {
	requestOptions, pollingOptions := option.ResolveRequestOptions(opts...)
	body, err := speechToTextBody(params)
	if err != nil {
		return nil, err
	}
	response, err := core.RunHybrid[json.RawMessage](ctx, r.http, speechToTextPath, body, requestOptions, pollingOptions)
	if err != nil {
		return nil, err
	}
	return *response, nil
}

func speechToTextBody(params SpeechToTextParams) (core.MultipartBody, error) {
	// Fill the documented default before validating: an empty Model is dropped
	// by CompactParams, and the contract's model enum would then reject the
	// request instead of applying the default.
	if params.Model == "" {
		params.Model = DefaultModel
	}
	compact := core.CompactParams(params)
	if err := core.ValidateParams(contractSchema["speech-to-text"], compact); err != nil {
		return core.MultipartBody{}, err
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
	return core.MultipartBody{
		Fields: fields,
		RepeatedFields: map[string][]string{
			"languages[]":               params.Languages,
			"keywords[]":                params.Keywords,
			"timestamp_granularities[]": params.TimestampGranularities,
		},
		Files: map[string]core.MultipartFile{
			"file": {Path: params.File, FileName: filepath.Base(params.File), ContentType: contentType},
		},
	}, nil
}
