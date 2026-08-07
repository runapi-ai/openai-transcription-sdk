import { beforeEach, describe, expect, it, vi } from 'vitest';
import type { HttpClient } from '@runapi.ai/core';
import { SpeechToText } from '../../src/resources/speech-to-text';

describe('OpenAI transcription resource', () => {
  const mockHttp: HttpClient = { request: vi.fn(), upload: vi.fn() };
  beforeEach(() => vi.clearAllMocks());

  it('posts repeated multipart fields and preserves raw text', async () => {
    vi.mocked(mockHttp.request).mockResolvedValueOnce('hello from audio\n');
    const result = await new SpeechToText(mockHttp).run({
      file: new Blob(['audio'], { type: 'audio/mpeg' }),
      filename: 'interview.mp3',
      model: 'gpt-transcribe',
      languages: ['en', 'zh'],
      keywords: ['RunAPI', 'release notes'],
      response_format: 'text',
    });
    const [, path, options] = vi.mocked(mockHttp.request).mock.calls[0]!;
    const form = options!.body as FormData;
    expect(path).toBe('/v1/audio/transcriptions');
    expect(form.getAll('languages[]')).toEqual(['en', 'zh']);
    expect(form.getAll('keywords[]')).toEqual(['RunAPI', 'release notes']);
    expect((form.get('file') as File).name).toBe('interview.mp3');
    expect(result).toBe('hello from audio\n');
  });

  it('defaults model and preserves JSON', async () => {
    vi.mocked(mockHttp.request).mockResolvedValueOnce({ text: 'hello' });
    const result = await new SpeechToText(mockHttp).run({ file: new Blob(['audio']) });
    const form = vi.mocked(mockHttp.request).mock.calls[0]![2]!.body as FormData;
    expect(form.get('model')).toBe('whisper-1');
    expect((form.get('file') as File).name).toBe('audio.mp3');
    expect(result).toEqual({ text: 'hello' });
  });

  it('keeps a File name when filename is omitted', async () => {
    vi.mocked(mockHttp.request).mockResolvedValueOnce({ text: 'hello' });
    const file = new Blob(['audio']) as Blob & { name: string };
    Object.defineProperty(file, 'name', { value: 'meeting.webm' });
    await new SpeechToText(mockHttp).run({ file });

    const form = vi.mocked(mockHttp.request).mock.calls[0]![2]!.body as FormData;
    expect((form.get('file') as File).name).toBe('meeting.webm');
  });
});
