package main

import (
	notify "computer_network/protobuf_learning/oneof/proto"
	"fmt"
	"github.com/iancoleman/strcase"
	fieldmask_utils "github.com/mennanov/fieldmask-utils"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
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

func optionalDemo() {
	req := &notify.Book{
		Title: "22",
		Price: proto.Int64(99),
	}
	//判断price是否被赋值
	if req.Price == nil {
		fmt.Println("price 没被赋值")
	} else {
		fmt.Println("price=", req.GetPrice())
	}
}

func fieldMaskDemo() {
	path := []string{"title", "price", "info.b"}
	req := &notify.UpdateBookRequest{
		Op: "admin",
		Book: &notify.Book{
			Title: "23",
			Price: proto.Int64(98),
			Info: &notify.Book_Info{
				B: "bbb",
			},
		},
		UpdateMask: &fieldmaskpb.FieldMask{Paths: path},
	}
	mask, _ := fieldmask_utils.MaskFromProtoFieldMask(req.UpdateMask, strcase.ToCamel)
	var bookDst = make(map[string]interface{})
	// 将数据读取到map[string]interface{}
	// fieldmask-utils支持读取到结构体等，更多用法可查看文档。
	fieldmask_utils.StructToMap(mask, req.Book, bookDst)
	// do update with bookDst
	fmt.Printf("bookDst:%#v\n", bookDst)
}

func main() {
	oneofdemo()
	optionalDemo()
	fieldMaskDemo()
}
