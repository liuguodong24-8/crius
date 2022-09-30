package util

import (
	"encoding/json"
	"fmt"
	"time"

	uuid "github.com/satori/go.uuid"
	"gitlab.omytech.com.cn/micro-service/Crius/pkgs/message"
	crius "gitlab.omytech.com.cn/micro-service/Crius/util"
	"gitlab.omytech.com.cn/micro-service/merchant-basic/internal/config"
)

// MQTTEntity ...
var MQTTEntity *message.MqttEntity

// MQTTConnect mqtt连接
func MQTTConnect() error {
	cfg := message.MqttConfig{
		Client:                config.Setting.MQTT.Client,
		Broker:                config.Setting.MQTT.Broker,
		UserName:              config.Setting.MQTT.Username,
		Password:              config.Setting.MQTT.Password,
		ReconnectingHandler:   message.DefaultReconnectingHandler,
		ConnectionLostHandler: message.DefaultConnectionLostHandler,
	}
	var err error
	MQTTEntity, err = message.NewMqtt(cfg)
	return err
}

// PublishMemberTasks 会员任务
func PublishMemberTasks(tasks []message.TaskCategory, memberID uuid.UUID) {
	for _, task := range tasks {
		payload := message.TaskMessage{
			Category: task,
			MemberID: memberID,
			Time:     time.Now(),
		}
		data, _ := json.Marshal(payload)
		crius.Logger.Info(fmt.Sprintf("MQTT publish:%+v", string(data)))
		err := MQTTEntity.Publish(config.Setting.MQTT.TaskTopic, 1, false, data)
		if err != nil {
			crius.Logger.Error(fmt.Sprintf("MQTT err:%+v", err))
		}
	}
}
