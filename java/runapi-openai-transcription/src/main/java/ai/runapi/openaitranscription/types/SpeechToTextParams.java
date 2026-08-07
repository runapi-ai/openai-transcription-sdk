package ai.runapi.openaitranscription.types;

import java.nio.file.Path;
import java.util.ArrayList;
import java.util.Collections;
import java.util.LinkedHashMap;
import java.util.List;
import java.util.Map;
import java.util.Objects;
import org.jspecify.annotations.Nullable;

/** Parameters for a synchronous audio transcription request. */
public final class SpeechToTextParams {
  private final Path file;
  private final String model;
  private final @Nullable String prompt;
  private final @Nullable String language;
  private final List<String> languages;
  private final List<String> keywords;
  private final @Nullable Double temperature;
  private final @Nullable String responseFormat;
  private final @Nullable Boolean stream;
  private final List<String> timestampGranularities;

  private SpeechToTextParams(Builder b) {
    file = Objects.requireNonNull(b.file, "file");
    model = b.model == null ? "whisper-1" : b.model;
    prompt = b.prompt; language = b.language;
    languages = Collections.unmodifiableList(new ArrayList<String>(b.languages));
    keywords = Collections.unmodifiableList(new ArrayList<String>(b.keywords));
    temperature = b.temperature; responseFormat = b.responseFormat; stream = b.stream;
    timestampGranularities = Collections.unmodifiableList(new ArrayList<String>(b.timestampGranularities));
  }

  public static Builder builder() { return new Builder(); }
  public Path file() { return file; }
  public String model() { return model; }
  public @Nullable String prompt() { return prompt; }
  public @Nullable String language() { return language; }
  public List<String> languages() { return languages; }
  public List<String> keywords() { return keywords; }
  public @Nullable Double temperature() { return temperature; }
  public @Nullable String responseFormat() { return responseFormat; }
  public @Nullable Boolean stream() { return stream; }
  public List<String> timestampGranularities() { return timestampGranularities; }

  public Map<String, Object> toMap() {
    Map<String, Object> values = new LinkedHashMap<String, Object>();
    values.put("file", file.toString()); values.put("model", model);
    if (prompt != null) values.put("prompt", prompt);
    if (language != null) values.put("language", language);
    if (!languages.isEmpty()) values.put("languages", languages);
    if (!keywords.isEmpty()) values.put("keywords", keywords);
    if (temperature != null) values.put("temperature", temperature);
    if (responseFormat != null) values.put("response_format", responseFormat);
    if (stream != null) values.put("stream", stream);
    if (!timestampGranularities.isEmpty()) values.put("timestamp_granularities", timestampGranularities);
    return values;
  }

  public static final class Builder {
    private Path file; private String model; private String prompt; private String language;
    private final List<String> languages = new ArrayList<String>();
    private final List<String> keywords = new ArrayList<String>();
    private Double temperature; private String responseFormat; private Boolean stream;
    private final List<String> timestampGranularities = new ArrayList<String>();
    private Builder() {}
    public Builder file(Path value) { file = Objects.requireNonNull(value, "file"); return this; }
    public Builder file(String value) { file = java.nio.file.Paths.get(nonBlank(value, "file")); return this; }
    public Builder model(String value) { model = nonBlank(value, "model"); return this; }
    public Builder prompt(String value) { prompt = value; return this; }
    public Builder language(String value) { language = value; return this; }
    public Builder languageHint(String value) { languages.add(nonBlank(value, "language")); return this; }
    public Builder keyword(String value) { keywords.add(nonBlank(value, "keyword")); return this; }
    public Builder temperature(double value) { temperature = value; return this; }
    public Builder responseFormat(String value) { responseFormat = nonBlank(value, "responseFormat"); return this; }
    public Builder stream(boolean value) { stream = value; return this; }
    public Builder timestampGranularity(String value) { timestampGranularities.add(nonBlank(value, "timestampGranularity")); return this; }
    public SpeechToTextParams build() { return new SpeechToTextParams(this); }
  }

  private static String nonBlank(String value, String name) {
    String checked = Objects.requireNonNull(value, name).trim();
    if (checked.isEmpty()) throw new IllegalArgumentException(name + " must not be blank");
    return checked;
  }
}
