package push

import (
	"JiaoNiBan-push/services/tpns"
	"encoding/json"
	"log"
	"os"
	"strconv"

	amqp "github.com/rabbitmq/amqp091-go"
)

var conn *amqp.Connection
var ch *amqp.Channel
var q amqp.Queue

func wrapTagRules(tags []string) []tpns.TagRule {
	tagItem := tpns.TagItem{
		Tags: tags,
	}
	var tagRule = tpns.TagRule{
		TagItems: []tpns.TagItem{tagItem},
		Operator: tpns.TagOperationOr,
	}
	return []tpns.TagRule{tagRule}
}

func Init() {
	amqp_addr := os.Getenv("AMQP_ADDR")
	conn, err := amqp.Dial(amqp_addr)
	if err != nil {
		panic(err)
	}
	ch, err = conn.Channel()
	if err != nil {
		panic(err)
	}
	q, err = ch.QueueDeclare(queueid, false, false, false, false, nil)
	if err != nil {
		panic(err)
	}

}

func Close() {
	conn.Close()
	ch.Close()
}

func ListenAndPush() {
	msgs, _ := ch.Consume("", queueid, false, false, false, false, nil)
	accessId, _ := strconv.Atoi(os.Getenv("TPNS_ACCESSID"))
	secretKey := os.Getenv("TPNS_SECRETKEY")
	client := tpns.NewClient(tpns.ShanghaiHost, uint32(accessId), secretKey)
	done := make(chan bool)

	go func() {
		for dq := range msgs {
			var r RecvMessage
			err := json.Unmarshal(dq.Body, &r)
			if err != nil {
				panic(err)
			}

			var content string

			if len(r.Desc) > 40 {
				content = r.Desc[:40]
			} else {
				content = r.Desc
			}

			req := tpns.NewRequest(
				tpns.WithAudience(tpns.AudienceAll),
				tpns.WithPlatform(tpns.PlatformAndroid),
				tpns.WithTitle(r.Title),
				tpns.WithContent(content),
				tpns.WithEnvironment(tpns.Develop),
				tpns.WithMessageType(tpns.Notify),
				tpns.WithTagRules(wrapTagRules(TagMap[r.Author])),
			)

			resp, err := client.Do(req)
			if err != nil {
				panic(err)
			}
			log.Println("SUCCESS: ", resp.GetPushId())
			dq.Ack(true)
		}
	}()
	<-done
}
