# frozen_string_literal: true

module RunApi
  module OpenaiTranscription
    CONTRACT = {
      "speech-to-text" => {
        "models" => ["gpt-transcribe", "whisper-1"],
        "fields_by_model" => {
          "gpt-transcribe" => {
            "file" => {
              "required" => true
            },
            "model" => {
              "enum" => ["gpt-transcribe"]
            },
            "response_format" => {
              "enum" => ["json", "text"]
            },
            "temperature" => {
              "min" => 0,
              "max" => 1
            }
          },
          "whisper-1" => {
            "file" => {
              "required" => true
            },
            "model" => {
              "enum" => ["whisper-1"]
            },
            "response_format" => {
              "enum" => ["json", "text", "srt", "verbose_json", "vtt"]
            },
            "temperature" => {
              "min" => 0,
              "max" => 1
            },
            "timestamp_granularities" => {
              "max_items" => 2
            }
          }
        }
      }
    }.freeze
  end
end
