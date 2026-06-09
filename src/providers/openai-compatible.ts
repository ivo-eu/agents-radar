/**
 * Base class for OpenAI-compatible providers.
 *
 * Shared by OpenAI, GitHub Copilot, and OpenRouter providers.
 */

import OpenAI from "openai";
import type { LlmProvider } from "./types.ts";

export abstract class OpenAICompatibleProvider implements LlmProvider {
  abstract readonly name: string;
  protected readonly client: OpenAI;
  protected readonly model: string;
  protected readonly stream: boolean;

  constructor(opts: { apiKey?: string; baseURL?: string; model: string; stream?: boolean }) {
    this.model = opts.model;
    this.stream = opts.stream ?? false;
    this.client = new OpenAI({
      apiKey: opts.apiKey,
      baseURL: opts.baseURL,
    });
  }

  async call(prompt: string, maxTokens: number): Promise<string> {
    if (this.stream) {
      const response = await this.client.chat.completions.create({
        model: this.model,
        max_completion_tokens: maxTokens,
        messages: [{ role: "user", content: prompt }],
        stream: true,
      });
      let text = "";
      for await (const chunk of response) {
        text += chunk.choices[0]?.delta?.content ?? "";
      }
      if (!text) throw new Error(`Unexpected empty response from ${this.name}`);
      return text;
    }

    const response = await this.client.chat.completions.create({
      model: this.model,
      max_completion_tokens: maxTokens,
      messages: [{ role: "user", content: prompt }],
    });
    const text = response.choices[0]?.message?.content;
    if (!text) throw new Error(`Unexpected empty response from ${this.name}`);
    return text;
  }
}
