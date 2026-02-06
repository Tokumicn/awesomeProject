import React from 'react';

export default function KnowledgeCitation({ citations }) {
  if (!citations || citations.length === 0) return null;
  return (
    <div className="knowledge-citation">
      <div style={{ fontWeight: 600, marginBottom: 8 }}>
        📚 知识库引用
      </div>
      <ul style={{ margin: 0, paddingLeft: 20 }}>
        {citations.map((c, i) => (
          <li key={i} style={{ marginBottom: 4 }}>{c}</li>
        ))}
      </ul>
    </div>
  );
} 