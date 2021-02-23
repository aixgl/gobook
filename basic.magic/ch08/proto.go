package main

import (
	"fmt"
	"github.com/aixgl/ch08/message"
	"github.com/golang/protobuf/proto"
)

func encodeProto() {
	fmt.Println("--------encodeProto--------")
	person := &message.Person{
		Name:   "xyj",
		Age:    30,
		Emails: []string{"00000000@qq.com", "88888888@163.com"},
		Phones: []*message.PhoneNumber{
			&message.PhoneNumber{
				Number: "011-23029321",
				Type:   message.PhoneType_HOME,
			},
			&message.PhoneNumber{
				Number: "13800000000",
				Type:   message.PhoneType_MOBILE,
			},
			&message.PhoneNumber{
				Number: "010-88888888",
				Type:   message.PhoneType_WORK,
			},
		},
	}

	//marshal:  obj---[]byte
	data, err := proto.Marshal(person)
	if err != nil {
		fmt.Println("error", err)
	}
	fmt.Println("data", data)

	// decode
	person1 := &message.Person{}
	err = proto.Unmarshal(data, person1)
	fmt.Println("error info:", err)
	fmt.Println("person:", person1)
}
