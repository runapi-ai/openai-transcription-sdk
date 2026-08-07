import type { ActionSchema, HttpClient, RequestOptions } from '@runapi.ai/core';
import { validateParams } from '@runapi.ai/core';
import { contract } from '../contract_gen';
import type { SpeechToTextParams, SpeechToTextResponse } from '../types';

const ENDPOINT = '/v1/audio/transcriptions';

/** Transcribe an uploaded audio file synchronously. */
export class SpeechToText {
  constructor(private readonly http: HttpClient) {}

  async run(params: SpeechToTextParams, options?: RequestOptions): Promise<SpeechToTextResponse> {
    const model = params.model ?? 'whisper-1';
    const fileName = params.filename ?? this.fileName(params.file);
    const validationParams = { ...params, file: fileName, model } as unknown as Record<string, unknown>;
    validateParams(contract['speech-to-text'] as ActionSchema, validationParams);

    const form = new FormData();
    form.append('file', params.file, fileName);
    form.append('model', model);
    this.appendOptional(form, 'prompt', params.prompt);
    this.appendOptional(form, 'language', params.language);
    this.appendOptional(form, 'temperature', params.temperature);
    this.appendOptional(form, 'response_format', params.response_format);
    this.appendOptional(form, 'stream', params.stream);
    params.languages?.forEach((value) => form.append('languages[]', value));
    params.keywords?.forEach((value) => form.append('keywords[]', value));
    params.timestamp_granularities?.forEach((value) => form.append('timestamp_granularities[]', value));

    return this.http.request<SpeechToTextResponse>('POST', ENDPOINT, { body: form, ...options });
  }

  private appendOptional(form: FormData, name: string, value: unknown): void {
    if (value !== undefined && value !== null && value !== '') form.append(name, String(value));
  }

  private fileName(file: Blob): string {
    const named = file as Blob & { name?: unknown };
    return typeof named.name === 'string' && named.name !== '' ? named.name : 'audio.mp3';
  }
}
