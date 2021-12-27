package main

import "JiaoNiBan-push/services/push"

func main() {
	push.Init()
	defer push.Close()
	push.ListenAndPush()
}
