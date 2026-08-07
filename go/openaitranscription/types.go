// Package openaitranscription provides OpenAI-compatible audio transcription through RunAPI.
package openaitranscription

// SpeechToTextParams configures a synchronous audio transcription request.
type SpeechToTextParams struct {
	File                   string   `json:"file" help:"required; local audio file path; max 25 MB"`
	Model                  string   `json:"model,omitempty" help:"optional; whisper-1 or gpt-transcribe; defaults to whisper-1"`
	Prompt                 string   `json:"prompt,omitempty" help:"optional; text that guides transcription"`
	Language               string   `json:"language,omitempty" help:"optional; language hint; cannot be combined with languages"`
	Languages              []string `json:"languages,omitempty" help:"optional for gpt-transcribe; repeated language hints; cannot be combined with language"`
	Keywords               []string `json:"keywords,omitempty" help:"optional for gpt-transcribe; repeated keyword hints"`
	Temperature            *float64 `json:"temperature,omitempty" help:"optional; sampling temperature from 0 to 1"`
	ResponseFormat         string   `json:"response_format,omitempty" help:"optional; response format supported by the selected model"`
	Stream                 *bool    `json:"stream,omitempty" help:"optional; request streaming when supported"`
	TimestampGranularities []string `json:"timestamp_granularities,omitempty" help:"optional for whisper-1 verbose_json; word or segment"`
}
