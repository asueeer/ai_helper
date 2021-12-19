struct CreateConversationRequest{
    2: string type;          // 会话类型. "helper": 客服
}

struct CreateConversationData{
    1: i64 conv_id; // 会话id
}

struct CreateConversationResponse{
    1: Meta meta;
    2: CreateConversationData data;
}

struct SendMessageRequest{
    1: i64 receiver_id;       // 接收方id
    2: i64 conv_id;           // 会话id
    3: string role;           // 用户身份, "visitor" 游客; "sys_helper": 系统客服(机器人自动回复); "be_helper": 后台客服
    4: MsgContent content;    // 消息内容
    5: string type;           // 消息类型. "text": 文本; "rich_text": 富文本; "image": 图片; "audio": 语音; "video": 视频.
    6: i32 status;            // 消息状态, 保留字段
    /*
        为了保证消息的有序性, 前端发消息时应该把相应的时间戳发过来
    */
    7: i64 timestamp;
}

struct MsgContent{
    1: string text;
    2: string rich_text;
    3: string img_url;
    4: string audio_url;
    5: string video_url;
}

struct SendMessageData{
    2: i64 message_id;        // 该消息的唯一标识id
    3: i64 conv_id;           // 该消息的会话id
}

struct SendMessageResponse{
    1: Meta meta;
    2: SendMessageData data;
}

struct LoadConversationsRequest{
    1: i64 limit;
    2: i64 cursor; // 相当于是offset

    3: i32 direction; // 拉取的方向; 枚举值 +1: 由现在到未来; -1: 由现在到过去; （默认为-1）

    4: string status; // 枚举值; "chatting": 聊天中; "waiting": 等待; "end": 结束; "all": 查询全部. （默认为chatting）
}

struct LoadConversationsData{
    1: list<Conversation> conversations;
    2: i64 new_cursor;
    3: bool has_more;
    4: i64 total; // 会话总数
}

struct LoadConversationsResponse{
    1: Meta meta;
    2: LoadConversationsData data;
}

struct Conversation{
    1: i64 conv_id;                   // 会话id
    2: string type;                   // 会话类型
    3: i32 unread;                    // 未读消息数
    4: i64 timestamp;                 // 时间戳
    5: list<Participant> paticipants; // 参与者
    6: Message last_msg;              // 最近的一条消息
    7: string conv_icon;              // 会话头像
    8: string conv_title;             // 会话的title
    9: string status;                 // 会话单的状态
}

struct Participant{
    1: i64 user_id;     // 用户id
    2: string head_url; // 用户头像
}

struct LoadConversationDetailRequest{
    1: i64 conv_id; // 会话id
    2: i64 cursor; // 0: 拉取最近的会话信息;
    3: i64 limit; // 默认为50条

    4: string role; // 用户身份; "visitor" 游客; "be_helper": 后台客服

    5: i32 direction; // 拉取的方向; 枚举值 +1: 由现在到未来; -1: 由现在到过去 （默认为-1）
}

struct LoadConversationDetailData{
    1: list<Message> messages;
    2: bool has_more;  // 是否包含更多会话
    3: i64 new_cursor; // 下一次拉取前, 需要传给后端的时间戳
    4: string status; // 当前会话的状态
}

struct LoadConversationDetailResponse{
    1: Meta meta;
    2: LoadConversationDetailData data;
}


struct Message{
    1: i64 sender_id;        // 发送方id
    2: i64 receiver_id;      // 接收方id
    3: string content;       // 消息内容
    4: string type;          // 消息类型
    5: string status;        // 消息状态
    6: string timestamp;     // 时间戳
    7: string sender_name;   // 发送方的用户名
}

struct Meta{
    1: i32 code;    // 业务自定义的返回值, 0为成功, 其他情况为失败
    2: string msg;
}


struct AcceptConversationRequest {
    1: string conv_id; // 会话id
}

struct AcceptConversationResponse {
    1: Meta meta;
}

struct EndConversationRequest{
    1: string conv_id;
}

struct EndConversationResponse{
    1: Meta meta;
}

struct SendMessageToRobotRequest {
    1: string content;    // 消息内容
    2: i64 timestamp;
}
struct SendMessageToRobotResponse {
    1: SendMessageToRobotData data;
    2: Meta meta;
}

struct SendMessageToRobotData{
    1: String resp_content; // 机器人的回复内容
}

service ImService{
    CreateConversationResponse CreateConversation(1: CreateConversationRequest req) (api.post="/im/create_conversation"); // 创建会话
    LoadConversationDetailResponse LoadConversation(1: LoadConversationDetailRequest req) (api.post="/im/load_conversation_detail"); // 加载会话详情
    LoadConversationsResponse LoadConversations(1: LoadConversationsRequest req) (api.post="/im/load_conversations"); // 加载会话列表
    SendMessageResponse SendMessage(1: SendMessageRequest req) (api.post="im/send_message"); // 发送消息

    // 二期-增加工单
    AcceptConversationResponse AcceptConversation(1: AcceptConversationRequest req) (api.post="/im/accept_conversation"); // 客服人员接收会话
    EndConversationResponse EndConversation(1: EndConversationRequest req) (api.post="im/end_conversation"); // 结束会话工单

    // 机器人
    SendMessageToRobotResponse SendMessageToRobot(1: SendMessageToRobotRequest req) (api.post="/im/send_message_to_robot"); // 给机器人发消息
}

