package main

import "JiaoNiBan-push/services/push"

// func test_tpns() {
// 	var tagItem = tpns.TagItem{
// 		Tags: []string{"STUDENT_ALL"},
// 	}
// 	var tagRule = tpns.TagRule{
// 		TagItems: []tpns.TagItem{tagItem},
// 		Operator:tpns.TagOperationAnd
// 	}
// 	accessId, _ := strconv.Atoi(os.Getenv("TPNS_ACCESSID"))
// 	secretKey := os.Getenv("TPNS_SECRETKEY")
// 	client := tpns.NewClient(tpns.ShanghaiHost, uint32(accessId), secretKey)
// 	req := tpns.NewRequest(tpns.WithAudience(tpns.AudienceTag),
// 		tpns.WithMessageType(tpns.Notify),
// 		tpns.WithPlatform(tpns.PlatformAndroid),
// 		tpns.WithTitle("testing from go server!"),
// 		tpns.WithContent("testing testing testing maybe available"),
// 		tpns.WithTagRules([]tpns.TagRule{tagRule}),
// 		tpns.WithEnvironment(tpns.Develop),
// 	)
// 	res, err := client.Do(req)
// 	if err != nil {
// 		panic(err)
// 	}
// 	fmt.Println(res.GetPushId())
// }

func main() {
	push.Init()
	defer push.Close()
	push.ListenAndPush()
}
