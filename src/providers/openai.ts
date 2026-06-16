/**
 * OpenAI provider — wraps the openai SDK.
 *
 * Env vars:
 *   OPENAI_API_KEY   - API key
 *   OPENAI_BASE_URL  - endpoint override (optional)
 *   OPENAI_MODEL     - model name (default: gpt-4o)
 *   OPENAI_THINKING  - true/false for DeepSeek thinking mode (optional)
 */

import { OpenAICompatibleProvider } from "./openai-compatible.ts";

function parseThinkingMode(value?: string): "enabled" | "disabled" | undefined {
  const normalized = value?.trim().toLowerCase();
  if (!normalized) return undefined;
  if (["true", "1", "yes", "enabled", "on"].includes(normalized)) return "enabled";
  if (["false", "0", "no", "disabled", "off"].includes(normalized)) return "disabled";
  throw new Error('Invalid OPENAI_THINKING value. Use "true", "false", "enabled", or "disabled".');
}

export class OpenAIProvider extends OpenAICompatibleProvider {
  readonly name = "openai";

  constructor(opts?: {
    apiKey?: string;
    baseURL?: string;
    model?: string;
    stream?: boolean;
    thinking?: "enabled" | "disabled";
  }) {
    super({
      apiKey: opts?.apiKey ?? process.env["OPENAI_API_KEY"],
      baseURL: opts?.baseURL ?? process.env["OPENAI_BASE_URL"],
      model: opts?.model ?? process.env["OPENAI_MODEL"] ?? "gpt-4o",
      stream: opts?.stream ?? process.env["OPENAI_STREAM"] === "true",
      thinking: opts?.thinking ?? parseThinkingMode(process.env["OPENAI_THINKING"]),
    });
  }
}
