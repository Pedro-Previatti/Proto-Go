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

func main() {
	fmt.Print(doSimple())
}
