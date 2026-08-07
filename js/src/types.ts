export type TranscriptionModel = 'whisper-1' | 'gpt-transcribe';
export type TranscriptionResponseFormat = 'json' | 'text' | 'srt' | 'verbose_json' | 'vtt';

export interface SpeechToTextParams {
  file: Blob;
  filename?: string;
  model?: TranscriptionModel;
  prompt?: string;
  language?: string;
  languages?: string[];
  keywords?: string[];
  temperature?: number;
  response_format?: TranscriptionResponseFormat;
  stream?: boolean;
  timestamp_granularities?: Array<'word' | 'segment'>;
}

export type SpeechToTextResponse = Record<string, unknown> | string;
