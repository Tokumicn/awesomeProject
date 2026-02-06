# IntelliQ 业务流程图

## 1. 系统架构概览

```mermaid
graph TB
    Client[前端/客户端]
    Flask[Flask App<br/>app.py]
    API[API Routes]
    Chatbot[ChatbotModel<br/>models/chatbot_model.py]
    Processor[CommonProcessor<br/>scene_processor/impl/]
    LLM[大模型服务]
    SceneAPI[场景API服务]
    
    Client -->|HTTP POST| Flask
    Flask -->|路由分发| API
    API -->|处理请求| Chatbot
    Chatbot -->|场景识别| LLM
    Chatbot -->|槽位填充| Processor
    Processor -->|调用场景| SceneAPI
```

## 2. API接口流程

### 2.1 `/multi_question` 接口流程

```mermaid
sequenceDiagram
    participant Client
    participant Flask as Flask App
    participant Chatbot as ChatbotModel
    participant LLM
    
    Client->>Flask: POST /multi_question<br/>{question: "用户问题"}
    Flask->>Flask: 验证参数
    alt 参数错误
        Flask-->>Client: 400 Bad Request
    end
    
    Flask->>Chatbot: process_multi_question(user_input)
    Chatbot->>Chatbot: recognize_intent(user_input)
    Chatbot->>LLM: 发送场景选项+用户输入
    LLM-->>Chatbot: 返回匹配的场景序号
    Chatbot->>Chatbot: 更新 current_purpose
    
    Chatbot->>Chatbot: 获取处理器 CommonProcessor
    Chatbot->>Chatbot: processor.process(user_input, chat_history)
    
    Chatbot-->>Flask: response
    Flask-->>Client: JSON {answer: "回复内容"}
```

### 2.2 `/api/llm_chat` 流式接口流程

```mermaid
sequenceDiagram
    participant Client
    participant Flask as Flask App
    participant Chatbot as ChatbotModel
    participant LLM as 大模型服务
    
    Client->>Flask: POST /api/llm_chat<br/>{user_input, session_id, messages}
    
    Flask->>Flask: 获取/创建会话
    Flask->>Flask: 检查 Accept Header<br/>判断是否为流式请求
    
    alt 流式请求
        Flask->>Chatbot: process_multi_question(user_input)
        Chatbot-->>Flask: 完整响应
        
        Flask->>Flask: generate() 生成器
        Note over Flask: 逐字符/标点<br/>分块发送 SSE
        
        loop 流式输出
            Flask-->>Client: data: "部分内容"
        end
        Flask-->>Client: data: [DONE]
    else 非流式请求
        Flask->>Chatbot: process_multi_question(user_input)
        Chatbot-->>Flask: response
        Flask-->>Client: JSON {response, session_id}
    end
```

## 3. 核心业务处理流程

### 3.1 多轮对话处理主流程

```mermaid
flowchart TD
    Start([用户输入]) --> AddHistory[添加用户消息到<br/>chat_history]
    AddHistory --> CheckScene{是否有<br/>current_purpose?}
    
    CheckScene -->|否| RecognizeIntent[识别意图<br/>recognize_intent]
    CheckScene -->|是| CheckSwitch[检测场景切换意图<br/>detect_scene_switch]
    
    RecognizeIntent --> RecognizeProcess[构建场景选项<br/>发送给LLM判断<br/>extract_continuous_digits]
    RecognizeProcess --> HasScene{识别到场景?}
    HasScene -->|是| SetPurpose[设置current_purpose<br/>last_recognized_scene]
    HasScene -->|否| NoSceneResponse[生成无场景回复<br/>generate_no_scene_response]
    
    SetPurpose --> SlotFilling[槽位填充阶段<br/>is_slot_filling=true]
    
    CheckSwitch -->|是| ClearScene[清除当前场景<br/>clear_current_scene]
    CheckSwitch -->|否| GetProcessor[获取处理器<br/>get_processor_for_scene]
    
    ClearScene --> RecognizeIntent
    
    SlotFilling --> GetProcessor
    GetProcessor --> Process[处理器处理<br/>CommonProcessor.process]
    
    Process --> CheckComplete{槽位已填完?}
    CheckComplete -->|是| CallAPI[调用场景API<br/>call_scene_api]
    CheckComplete -->|否| AskUser[询问缺失信息<br/>ask_user_for_missing_data]
    
    CallAPI --> ProcessResult[处理API结果<br/>process_api_result]
    ProcessResult --> ClearScene2[清除当前场景]
    
    AskUser --> AddAssistant[添加助手回复到<br/>chat_history]
    ClearScene2 --> AddAssistant
    NoSceneResponse --> AddAssistant
    
    AddAssistant --> Return([返回响应])
```

### 3.2 槽位填充详细流程

```mermaid
flowchart LR
    subgraph CommonProcessor
        Process[process]
        ExtractInfo[提取新信息<br/>get_slot_update_message<br/>send_message<br/>extract_json_from_string]
        UpdateSlot[更新槽位<br/>update_slot]
        CheckSlot{is_slot_fully_filled?}
        Complete[respond_with_complete_data]
        AskMore[ask_user_for_missing_data]
    end
    
    subgraph SlotUpdate
        BuildPrompt[构建提示词<br/>scene_name + slot_dynamic_example<br/>+ slot_template + user_input]
        SendLLM[发送LLM<br/>传递chat_history]
        ExtractJSON[提取JSON<br/>处理扁平结构]
    end
    
    subgraph CompleteFlow
        FormatLog[格式化日志]
        GetSceneKey[_get_scene_key]
        PrepareData[准备槽位数据<br/>_get_slot_key]
        CallAPI[call_scene_api]
        ProcessAPI[process_api_result]
    end
    
    Process --> ExtractInfo
    ExtractInfo --> BuildPrompt
    BuildPrompt --> SendLLM
    SendLLM --> ExtractJSON
    ExtractJSON --> UpdateSlot
    UpdateSlot --> CheckSlot
    
    CheckSlot -->|是| Complete
    CheckSlot -->|否| AskMore
    
    Complete --> FormatLog
    FormatLog --> GetSceneKey
    GetSceneKey --> PrepareData
    PrepareData --> CallAPI
    CallAPI --> ProcessAPI
```

## 4. 意图识别流程

```mermaid
flowchart TD
    Start([用户输入]) --> BuildOptions[构建场景选项列表<br/>purpose_description]
    
    subgraph OptionsPrompt
        ListScenes[列出所有场景<br/>scene_templates.items]
        AddIndex[添加序号 1,2,3...]
        AddZero[添加选项0:<br/>无场景/无法判断]
    end
    
    BuildOptions --> OptionsPrompt
    OptionsPrompt --> CreatePrompt[创建识别提示词<br/>包含: last_scene_info<br/>options_prompt<br/>user_input]
    
    CreatePrompt --> SendLLM[发送给LLM<br/>send_message<br/>带上chat_history]
    
    SendLLM --> ExtractChoice[提取数字<br/>extract_continuous_digits]
    
    ExtractChoice --> Choice{用户选择?}
    Choice -->|非0| NewScene{新场景≠<br/>current_purpose?}
    Choice -->|0或null| NoScene[无法识别意图]
    
    NewScene -->|是| UpdateScene[更新current_purpose<br/>更新last_recognized_scene<br/>清除处理器<br/>is_slot_filling=false]
    NewScene -->|否| KeepScene[保持当前场景]
    
    NoScene --> CheckFilling{有场景且<br/>正在补槽?}
    CheckFilling -->|是| KeepScene
    CheckFilling -->|否| ClearScene[清空current_purpose<br/>is_slot_filling=false]
    
    UpdateScene --> Print[打印场景名]
    KeepScene --> Print
    ClearScene --> End([结束])
    Print --> End
```

## 5. 场景切换检测流程

```mermaid
sequenceDiagram
    participant Chatbot as ChatbotModel
    participant LLM as 大模型服务
    
    Note over Chatbot: 已有current_purpose时调用
    
    Chatbot->>Chatbot: 获取当前场景名<br/>current_scene_name
    
    Chatbot->>Chatbot: 构建切换检测提示词<br/>scene_prompts.scene_switch_detection
    Note right of Chatbot: 参数: current_scene_name<br/>user_input
    
    Chatbot->>LLM: 发送检测请求<br/>send_message<br/>带chat_history
    LLM-->>Chatbot: 返回判断结果
    
    Chatbot->>Chatbot: 提取数字<br/>extract_continuous_digits
    
    alt 返回1
        Chatbot->>Chatbot: 记录日志<br/>"检测到用户意图切换场景"
        Chatbot-->>Chatbot: 返回True
    else 返回其他
        Chatbot-->>Chatbot: 返回False
    end
```

## 6. 会话管理流程

```mermaid
stateDiagram-v2
    [*] --> 会话管理
    
    会话管理 --> 获取会话: get_or_create_session
    获取会话 --> 会话存在: session_id存在
    获取会话 --> 创建会话: session_id为空或不存在
    
    创建会话 --> 生成ID: uuid.uuid4().hex[:8]
    生成ID --> 初始化会话: messages=[]<br/>context={}<br/>created_at=None
    
    会话存在 --> 返回会话数据
    初始化会话 --> 返回会话数据
    
    返回会话数据 --> 处理请求
    
    处理请求 --> 重置会话: /api/reset_session
    重置会话 --> 删除会话数据: del sessions[session_id]
    删除会话数据 --> [*]
    
    处理请求 --> [*]
```

## 7. 数据流图

```mermaid
flowchart LR
    subgraph 输入层
        UserInput[用户输入]
        SessionID[会话ID]
        Messages[消息历史]
    end
    
    subgraph 处理层
        Intent[意图识别]
        SlotFill[槽位填充]
        SceneSwitch[场景切换检测]
    end
    
    subgraph 数据层
        SlotData[槽位数据<br/>scene_slots]
        ChatHistory[聊天记录<br/>chat_history]
        Sessions[会话存储<br/>sessions]
        Templates[场景模板<br/>scene_templates]
    end
    
    subgraph 输出层
        Response[文本响应]
        Stream[流式响应]
        APIResult[API调用结果]
    end
    
    UserInput --> Intent
    SessionID --> Sessions
    Messages --> ChatHistory
    
    Intent --> SlotFill
    SlotFill --> SlotData
    Intent --> SceneSwitch
    
    SlotData --> APIResult
    Templates --> Intent
    Templates --> SlotFill
    
    APIResult --> Response
    ChatHistory --> Response
    Response --> Stream
```

## 8. 错误处理流程

```mermaid
flowchart TD
    Start([请求开始]) --> Validate[参数验证]
    
    Validate --> Error{参数错误?}
    Error -->|是| Return400[返回 400 Bad Request<br/>{error: "描述"}]
    
    Error -->|否| Process[业务处理]
    
    Process --> Exception{发生异常?}
    Exception -->|是| LogError[记录错误日志<br/>logging.error]
    Exception -->|否| Success[返回成功响应]
    
    LogError --> IsStream{流式请求?}
    IsStream -->|是| StreamError[返回SSE错误<br/>data: [ERROR] message]
    IsStream -->|否| Return500[返回 500 Internal Server Error]
    
    Return400 --> End([结束])
    Success --> End
    StreamError --> End
    Return500 --> End
```

## 9. 完整的聊天处理时序图

```mermaid
sequenceDiagram
    participant User as 用户
    participant Client as 前端客户端
    participant Flask as Flask后端
    participant Chatbot as ChatbotModel
    participant Processor as CommonProcessor
    participant LLM as 大模型API
    participant SceneAPI as 场景业务API
    
    Note over User,SceneAPI: 场景：用户咨询流量套餐
    
    User->>Client: "我想查流量套餐"
    Client->>Flask: POST /api/llm_chat<br/>{user_input: "我想查流量套餐"}
    
    Flask->>Chatbot: process_multi_question()
    Chatbot->>Chatbot: recognize_intent()
    
    Chatbot->>LLM: 发送场景选项+用户输入
    Note right of Chatbot: "1. 宽带报修<br/>2. 流量查询<br/>3. 套餐办理<br/>..."
    LLM-->>Chatbot: 返回 "2"
    
    Chatbot->>Chatbot: current_purpose="flow_query"
    Chatbot->>Chatbot: is_slot_filling=true
    
    Chatbot->>Processor: get_processor_for_scene()
    Chatbot->>Processor: process(user_input, chat_history)
    
    Processor->>LLM: get_slot_update_message()<br/>询问需要哪些槽位
    LLM-->>Processor: 返回JSON: {phone_number: "138xxxx"}
    
    Processor->>Processor: update_slot()
    Processor->>Processor: is_slot_fully_filled()?
    
    Note over Processor: 假设缺少必要槽位
    Processor->>LLM: ask_user_for_missing_data()<br/>询问缺失信息
    LLM-->>Processor: "请问您的手机号是多少？"
    
    Processor-->>Chatbot: response
    Chatbot->>Chatbot: chat_history.append(assistant_msg)
    Chatbot-->>Flask: response
    Flask-->>Client: {response: "请问您的手机号是多少？", session_id}
    Client-->>User: "请问您的手机号是多少？"
    
    User->>Client: "13812345678"
    Client->>Flask: POST /api/llm_chat<br/>{user_input, session_id}
    
    Flask->>Chatbot: process_multi_question()
    Chatbot->>Chatbot: detect_scene_switch()
    Note over Chatbot: 检查是否切换场景
    
    Chatbot->>Processor: process()
    Processor->>LLM: get_slot_update_message()
    LLM-->>Processor: {phone_number: "13812345678"}
    Processor->>Processor: update_slot()<br/>槽位完整！
    
    Processor->>Processor: respond_with_complete_data()
    Processor->>SceneAPI: call_scene_api()<br/>POST /flow_query<br/>{phone_number: "13812345678"}
    SceneAPI-->>Processor: 返回套餐信息
    
    Processor->>LLM: process_api_result()<br/>生成友好回复
    LLM-->>Processor: "您当前套餐包含10GB流量..."
    
    Processor-->>Chatbot: response
    Chatbot->>Chatbot: clear_current_scene()<br/>场景处理完成
    Chatbot-->>Flask: response
    Flask-->>Client: 完整响应
    Client-->>User: "您当前套餐包含10GB流量..."
```

---

## 文件结构说明

| 文件路径 | 职责 |
|---------|------|
| `app.py` | Flask应用入口，定义API路由，处理HTTP请求/响应，管理会话 |
| `app_config.py` | 配置类，管理环境变量、API前缀、CORS等 |
| `models/chatbot_model.py` | 核心业务逻辑，意图识别、场景切换检测、流程控制 |
| `scene_processor/scene_processor.py` | 场景处理器抽象基类 |
| `scene_processor/impl/common_processor.py` | 通用场景处理器实现，槽位填充逻辑 |
| `utils/helpers.py` | 工具函数：加载配置、调用API、槽位操作等 |
| `scene_config/scene_templates.json` | 场景配置模板 |
| `scene_config/scene_prompts.py` | 场景相关的提示词模板 |
