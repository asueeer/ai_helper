package model

//struct SendMessageToRobotRequest {
//1: string content;    // 消息内容
//2: i64 timestamp;
//}
//struct SendMessageToRobotResponse {
//1: SendMessageToRobotData data;
//2: Meta meta;
//}
//
//struct SendMessageToRobotData{
//1: String resp_content; // 机器人的回复内容
//}

type SendMessageToRobotRequest struct {
	Content   string `json:"string"`
	Timestamp int64  `json:"timestamp"`
}

type SendMessageToRobotResponse struct {
	SendMessageData SendMessageToRobotData `json:"data"`
	Meta            Meta                   `json:"meta"`
}

type SendMessageToRobotData struct {
	RespContent string `json:"resp_content"`
}
