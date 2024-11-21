package main

import (
	notify "computer_network/protobuf_learning/oneof/proto"
	"fmt"
)

func oneofdemo() {
	//req := &notify.Notify{
	//	Text: "更新了",
	//	NoticeWay: &notify.Notify_Email{
	//		Email: "123@qq.com",
	//	},
	//}
	req := &notify.Notify{
		Text: "更新了",
		NoticeWay: &notify.Notify_Phone{
			Phone: "10086",
		},
	}

	//客户端判断通知的方式
	switch req.NoticeWay.(type) {
	case *notify.Notify_Email:
		fmt.Printf("text=%v, email=%v", req.GetText(), req.GetEmail())
	case *notify.Notify_Phone:
		fmt.Printf("text=%v, phone=%v", req.GetText(), req.GetPhone())
	}
}

func main() {
	oneofdemo()
}
