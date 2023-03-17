package handler

import (
	"ai_helper/biz/common"
	"ai_helper/biz/model"
	"ai_helper/biz/service"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"io/ioutil"
	"net/http"
)

// SendMessageToRobot [Post] /api/im/send_message_to_robot
func SendMessageToRobot(c *gin.Context) {
	// 给机器人发送消息
	req := &model.SendMessageRequest{}
	if err := c.ShouldBindBodyWith(req, binding.JSON); err != nil {
		c.JSON(400, err.Error())
		return
	}
	var ss service.SendMessageToRobotService
	resp, err := ss.Execute(c, req)
	if err != nil {
		common.WriteError(c, err)
		return
	}

	给NLP机器人发消息(req)
	c.JSON(200, resp)
}

func 给NLP机器人发消息(smReq *model.SendMessageRequest) {
	url := "https://miner.picp.net/chatBot"
	method := "GET"

	var buf bytes.Buffer
	var err error

	err = json.NewEncoder(&buf).Encode(smReq)
	if err != nil {
		fmt.Printf("err: %+v\n", err)
		return
	}

	client := &http.Client{}
	req, err := http.NewRequest(method, url, &buf)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
}
