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

func main() {
	fmt.Println("Simple Proto Message")
	fmt.Println(doSimple())

	fmt.Println("Complex Proto Message")
	fmt.Println(doComplex())
}
