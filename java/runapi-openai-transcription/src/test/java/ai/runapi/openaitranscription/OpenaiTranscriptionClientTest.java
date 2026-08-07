package ai.runapi.openaitranscription;

import static org.junit.jupiter.api.Assertions.assertEquals;
import static org.junit.jupiter.api.Assertions.assertTrue;

import ai.runapi.core.http.HttpRequest;
import ai.runapi.core.http.HttpResponse;
import ai.runapi.core.http.HttpTransport;
import ai.runapi.openaitranscription.types.SpeechToTextParams;
import ai.runapi.openaitranscription.types.TranscriptionResponse;
import java.io.ByteArrayOutputStream;
import java.nio.charset.StandardCharsets;
import java.nio.file.Files;
import java.nio.file.Path;
import java.util.Collections;
import org.junit.jupiter.api.Test;

final class OpenaiTranscriptionClientTest {
  @Test void postsRepeatedMultipartFieldsAndPreservesRawResponse() throws Exception {
    Path audio = Files.createTempFile("runapi-transcription", ".mp3");
    Files.write(audio, "audio".getBytes(StandardCharsets.UTF_8));
    CapturingTransport transport = new CapturingTransport(new HttpResponse(200, "hello from audio\n", Collections.singletonMap("Content-Type", Collections.singletonList("text/plain"))));
    OpenaiTranscriptionClient client = OpenaiTranscriptionClient.builder().apiKey("k").transport(transport).build();
    TranscriptionResponse response = client.speechToText().run(SpeechToTextParams.builder().file(audio).model("gpt-transcribe").languageHint("en").languageHint("zh").keyword("RunAPI").keyword("release notes").responseFormat("text").build());
    ByteArrayOutputStream bytes = new ByteArrayOutputStream();
    transport.request.getBody().writeTo(bytes);
    String multipart = bytes.toString("UTF-8");
    assertEquals("/v1/audio/transcriptions", transport.request.getPath());
    assertEquals(2, occurrences(multipart, "name=\"languages[]\""));
    assertEquals(2, occurrences(multipart, "name=\"keywords[]\""));
    assertTrue(multipart.contains("name=\"file\"; filename=\"" + audio.getFileName() + "\""));
    assertEquals("hello from audio\n", response.body());
    assertEquals("text/plain", response.contentType());
  }

  private static int occurrences(String text, String needle) { int count = 0, index = 0; while ((index = text.indexOf(needle, index)) >= 0) { count++; index += needle.length(); } return count; }
  private static final class CapturingTransport implements HttpTransport {
    private final HttpResponse response; private HttpRequest request;
    private CapturingTransport(HttpResponse response) { this.response = response; }
    public HttpResponse send(HttpRequest value) { request = value; return response; }
    public void close() {}
  }
}
