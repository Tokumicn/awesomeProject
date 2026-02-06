import React from 'react';

export default function MessageBubble({ message, msg, isStreaming }) {
  // 兼容不同的参数名
  const msgData = message || msg;
  const isAI = msgData.role === 'ai';
  
  return (
    <div className={`message-bubble ${isAI ? 'message-bubble-ai' : 'message-bubble-user'}`}>
      <div className="message-content">
        <div className="message-role">
          {isAI && <span style={{ marginRight: 4 }}>💬</span>}
          {isAI ? 'AI助手' : '你'}
        </div>
        <div className="message-text">
          {msgData.text}
          {isStreaming && (
            <span className="typing-indicator">
              <span className="typing-dot"></span>
              <span className="typing-dot"></span>
              <span className="typing-dot"></span>
            </span>
          )}
        </div>
        {msgData.status && msgData.status !== 'streaming' && (
          <div className="message-status">
            {msgData.status === 'welcome' && '✨'}
            {msgData.status === 'completed' && '✓'}
            {msgData.status === 'thinking' && '🤔'}
            {msgData.status}
          </div>
        )}
      </div>
    </div>
  );
} 