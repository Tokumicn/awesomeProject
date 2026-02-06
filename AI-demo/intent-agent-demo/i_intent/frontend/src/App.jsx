import React, { useState } from 'react';
import { Layout, Tabs, Button, Space } from 'antd';
import { PlusOutlined, RobotOutlined, MessageOutlined, CustomerServiceOutlined } from '@ant-design/icons';
import ChatBox from './components/ChatBox';
import AIChatBox from './components/AIChatBox';
import './App.css';

const { Header, Content } = Layout;

function createNewChat(id, type = 'business') {
  return {
    id,
    title: type === 'business' ? `业务对话${id}` : `AI对话${id}`,
    type: type,
    messages: [],
    participants: { user: '你', ai: type === 'business' ? '业务助手' : 'AI助手' },
  };
}

export default function App() {
  const [chats, setChats] = useState([
    createNewChat(1, 'business'),
    createNewChat(2, 'ai')
  ]);
  const [activeKey, setActiveKey] = useState('1');
  const [nextId, setNextId] = useState(3);

  const addChat = (type = 'business') => {
    const id = String(nextId);
    setChats([...chats, createNewChat(id, type)]);
    setActiveKey(id);
    setNextId(nextId + 1);
  };

  const removeChat = (targetKey) => {
    let newActiveKey = activeKey;
    let lastIndex = -1;
    chats.forEach((chat, i) => {
      if (chat.id === targetKey) lastIndex = i - 1;
    });
    const newChats = chats.filter(chat => chat.id !== targetKey);
    if (newChats.length && newActiveKey === targetKey) {
      newActiveKey = newChats[lastIndex >= 0 ? lastIndex : 0].id;
    }
    setChats(newChats);
    setActiveKey(newActiveKey);
  };

  const updateChat = (id, updater) => {
    setChats(chats => chats.map(chat => chat.id === id ? updater(chat) : chat));
  };

  const renderChatComponent = (chat) => {
    if (chat.type === 'ai') {
      return (
        <AIChatBox
          chat={chat}
          updateChat={updater => updateChat(chat.id, updater)}
        />
      );
    } else {
      return (
        <ChatBox
          chat={chat}
          updateChat={updater => updateChat(chat.id, updater)}
        />
      );
    }
  };

  return (
    <Layout className="app-container">
      <Header className="main-header">
        <div className="header-logo">
          <CustomerServiceOutlined style={{ fontSize: 28, color: '#3b82f6' }} />
          <span className="header-title">电信业务AI问答系统</span>
        </div>
        <Space className="header-actions">
          <Button 
            icon={<RobotOutlined />} 
            type="primary" 
            onClick={() => addChat('business')}
            className="header-btn header-btn-primary"
          >
            新业务对话
          </Button>
          <Button 
            icon={<MessageOutlined />} 
            onClick={() => addChat('ai')}
            className="header-btn header-btn-default"
          >
            新AI对话
          </Button>
        </Space>
      </Header>
      <Content className="main-content">
        <div className="tabs-container">
          <Tabs
            type="editable-card"
            activeKey={activeKey}
            onChange={setActiveKey}
            onEdit={removeChat}
            hideAdd
            items={chats.map(chat => ({
              key: chat.id,
              label: (
                <span>
                  {chat.type === 'business' ? <RobotOutlined /> : <MessageOutlined />}
                  {' '}{chat.title}
                </span>
              ),
              children: renderChatComponent(chat),
            }))}
          />
        </div>
      </Content>
    </Layout>
  );
} 