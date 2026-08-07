import { BaseClient, type ClientOptions } from '@runapi.ai/core';
import { SpeechToText } from './resources/speech-to-text';

/** OpenAI-compatible audio transcription client. */
export class OpenaiTranscriptionClient extends BaseClient {
  public readonly speechToText: SpeechToText;

  constructor(options: ClientOptions = {}) {
    super(options);
    this.speechToText = new SpeechToText(this.http);
  }
}
