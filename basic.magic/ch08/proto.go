package main

import (
    "github.com/aixgl/ch08/message"
    "fmt"
    "rpc/proto"
)

func encodeProto() {
    fmt.Println("--------encodeProto--------")
	person := &pb.Person{
		Name:"yz",
		Age:11,
		Emails:[]string{"782365461@qq.com","123456@163.com"},
		Phones:[]*pb.PhoneNumber{
			&pb.PhoneNumber{
				Number:"123456",
				Type:pb.PhoneType_HOME,
			},
			&pb.PhoneNumber{
				Number:"123456",
				Type:pb.PhoneType_MOBILE,
			},
			&pb.PhoneNumber{
				Number:"123456",
				Type:pb.PhoneType_WORK,
			},
		},
	}
 
	//marshal:  obj---[]byte
	data,err := proto.Marshal(person)
	if err != nil {
		fmt.Println(err, data)
    }
}