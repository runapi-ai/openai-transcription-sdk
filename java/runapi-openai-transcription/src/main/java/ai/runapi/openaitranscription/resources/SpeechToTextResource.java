package ai.runapi.openaitranscription.resources;

import ai.runapi.core.ApiRequestExecutor;
import ai.runapi.core.ClientOptions;
import ai.runapi.core.RequestOptions;
import ai.runapi.core.http.HttpMethod;
import ai.runapi.core.http.HttpRequest;
import ai.runapi.core.http.HttpResponse;
import ai.runapi.core.http.HttpTransport;
import ai.runapi.core.http.MultipartRequestBody;
import ai.runapi.openaitranscription.types.SpeechToTextParams;
import ai.runapi.openaitranscription.types.TranscriptionResponse;
import java.io.IOException;
import java.nio.file.Files;
import java.util.Map;

/** Synchronous audio transcription operations. */
public final class SpeechToTextResource {
  public static final String ENDPOINT = "/v1/audio/transcriptions";
  private final ApiRequestExecutor executor;

  public SpeechToTextResource(HttpTransport transport, ClientOptions options) { executor = new ApiRequestExecutor(transport, options); }
  public TranscriptionResponse run(SpeechToTextParams params) { return run(params, RequestOptions.none()); }

  public TranscriptionResponse run(SpeechToTextParams params, RequestOptions options) {
    Map<String, Object> values = params.toMap();
    MultipartRequestBody.Builder body = MultipartRequestBody.builder();
    for (Map.Entry<String, Object> field : values.entrySet()) {
      if (!"file".equals(field.getKey()) && !(field.getValue() instanceof Iterable)) body.field(field.getKey(), String.valueOf(field.getValue()));
    }
    append(body, "languages[]", params.languages());
    append(body, "keywords[]", params.keywords());
    append(body, "timestamp_granularities[]", params.timestampGranularities());
    body.file("file", params.file(), params.file().getFileName().toString(), contentType(params));
    HttpResponse response = executor.send(HttpRequest.builder(HttpMethod.POST, ENDPOINT).body(body.build()).options(options).build());
    return new TranscriptionResponse(response.getStatusCode(), response.getBody(), response.firstHeader("Content-Type"));
  }

  private static void append(MultipartRequestBody.Builder body, String name, Iterable<String> values) { for (String value : values) body.field(name, value); }
  private static String contentType(SpeechToTextParams params) {
    try { String detected = Files.probeContentType(params.file()); return detected == null ? "application/octet-stream" : detected; }
    catch (IOException ignored) { return "application/octet-stream"; }
  }
}
