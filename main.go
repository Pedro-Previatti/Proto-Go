package main

import (
	"fmt"

	pb "github.com/Pedro-Previatti/Proto-Go/proto"
)

func doSimple() *pb.Simple {
	return &pb.Simple{
		Id:         42,
		IsSimple:   true,
		Name:       "Simple Name",
		SampleList: []int32{1, 2, 3, 4, 5, 6},
	}
}

func doComplex() *pb.Complex {
	return &pb.Complex{
		OneDummy: &pb.Dummy{Id: 100, Name: "Dummy 100"},
		MultipleDummies: []*pb.Dummy{
			{Id: 101, Name: "Dummy 101"},
			{Id: 102, Name: "Dummy 102"},
		},
	}
}

func doEnumeration() *pb.Enumeration {
	return &pb.Enumeration{
		EyeColor: pb.EyeColor_GREEN,
	}
}

func doOneOf(message interface{}) {
	switch x := message.(type) {
	case *pb.Result_Id:
		fmt.Println(message.(*pb.Result_Id).Id)
	case *pb.Result_Message:
		fmt.Println(message.(*pb.Result_Message).Message)
	default:
		fmt.Errorf("Message has unexpected type: %v", x)
	}
}

func main() {
	fmt.Println("Simple Proto Message:")
	fmt.Println(doSimple())

	fmt.Println("Complex Proto Message:")
	fmt.Println(doComplex())

	fmt.Println("Proto Message with Enumeration:")
	fmt.Println(doEnumeration())

	fmt.Println("Proto OneOfs Id:")
	doOneOf(&pb.Result_Id{Id: 1})
	fmt.Println("Proto OneOfs Message:")
	doOneOf(&pb.Result_Message{Message: "Some message"})
}
