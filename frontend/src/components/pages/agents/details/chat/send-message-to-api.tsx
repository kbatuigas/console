import { config } from 'config';
import type { ChatMessage } from 'database/chat-db';

interface ChatApiResponse {
  message: string;
  success: boolean;
  error?: string;
}

interface ChatApiRequest {
  question: string;
}

// Limit chat history to last 30 messages
export const CHAT_HISTORY_MESSAGE_LIMIT = 30;

interface SendMessageToApiProps {
  message: string;
  chatHistory: ChatMessage[];
  agentUrl?: string;
}

export const sendMessageToApi = async ({
  message,
  chatHistory,
  agentUrl,
}: SendMessageToApiProps): Promise<ChatApiResponse> => {
  try {
    const recentHistory = chatHistory.slice(-CHAT_HISTORY_MESSAGE_LIMIT);

    const payload: ChatApiRequest = {
      question: message,
    };

    const response = await fetch(`${agentUrl}`, {
      method: 'POST',
      headers: {
        Authorization: `Bearer ${config.jwt}`,
        'Content-Type': 'application/json',
        // 'Access-Control-Allow-Origin': '*',
        // 'Access-Control-Allow-Methods': 'POST, GET, OPTIONS',
        // 'Access-Control-Allow-Headers': 'Content-Type',
      },
      body: JSON.stringify(payload),
    });

    if (!response.ok) {
      throw new Error(`API responded with status: ${response.status}`);
    }

    const reader = response.body?.getReader();
    const { value } = (await reader?.read()) || {};
    const text = new TextDecoder().decode(value);

    try {
      return {
        message: text,
        success: true,
      } as ChatApiResponse;
    } catch (err) {
      console.error('Error parsing API response:', err);
      return {
        success: false,
        message: 'Failed to parse server response',
        error: err instanceof Error ? err.message : 'Unknown error',
      };
    }
  } catch (error) {
    console.error('Error sending message to API:', error);
    return {
      success: false,
      message: 'Failed to send message to server. Please try again later.',
      error: error instanceof Error ? error.message : 'Unknown error',
    };
  }
};
