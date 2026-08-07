package ai.runapi.openaitranscription.types;

import org.jspecify.annotations.Nullable;

/** Exact transcription response body and metadata. */
public final class TranscriptionResponse {
  private final int statusCode;
  private final String body;
  private final @Nullable String contentType;

  public TranscriptionResponse(int statusCode, String body, @Nullable String contentType) {
    this.statusCode = statusCode; this.body = body; this.contentType = contentType;
  }

  public int statusCode() { return statusCode; }
  public String body() { return body; }
  public @Nullable String contentType() { return contentType; }
}
