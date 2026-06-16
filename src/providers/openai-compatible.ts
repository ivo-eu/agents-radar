/**
 * Base class for OpenAI-compatible providers.
 *
 * Shared by OpenAI, GitHub Copilot, and OpenRouter providers.
 */

import OpenAI from "openai";
import type {
  ChatCompletionCreateParamsNonStreaming,
  ChatCompletionCreateParamsStreaming,
} from "openai/resources/chat/completions";
import type { LlmProvider } from "./types.ts";

type ThinkingMode = "enabled" | "disabled";
type ThinkingParam = { thinking?: { type: ThinkingMode } };

export abstract class OpenAICompatibleProvider implements LlmProvider {
  abstract readonly name: string;
  protected readonly client: OpenAI;
  protected readonly model: string;
  protected readonly stream: boolean;
  protected readonly thinking?: ThinkingMode;

  constructor(opts: {
    apiKey?: string;
    baseURL?: string;
    model: string;
    stream?: boolean;
    thinking?: ThinkingMode;
  }) {
    this.model = opts.model;
    this.stream = opts.stream ?? false;
    this.thinking = opts.thinking;
    this.client = new OpenAI({
      apiKey: opts.apiKey,
      baseURL: opts.baseURL,
    });
  }

  private createParams(
    prompt: string,
    maxTokens: number,
  ): ChatCompletionCreateParamsNonStreaming & ThinkingParam {
    const params = {
      model: this.model,
      max_completion_tokens: maxTokens,
      messages: [{ role: "user" as const, content: prompt }],
      ...(this.thinking ? { thinking: { type: this.thinking } } : {}),
    };
    return params as ChatCompletionCreateParamsNonStreaming & ThinkingParam;
  }

  async call(prompt: string, maxTokens: number): Promise<string> {
    if (this.stream) {
      const response = await this.client.chat.completions.create({
        ...this.createParams(prompt, maxTokens),
        stream: true,
      } as ChatCompletionCreateParamsStreaming & ThinkingParam);
      let text = "";
      for await (const chunk of response) {
        text += chunk.choices[0]?.delta?.content ?? "";
      }
      if (!text) throw new Error(`Unexpected empty response from ${this.name}`);
      return text;
    }

    const response = await this.client.chat.completions.create(this.createParams(prompt, maxTokens));
    const text = response.choices[0]?.message?.content;
    if (!text) throw new Error(`Unexpected empty response from ${this.name}`);
    return text;
  }
}
