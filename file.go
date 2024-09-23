package main

import (
	"fmt"
	"log"
	"os"

	"google.golang.org/protobuf/proto"
)

func writeProto(file string, pb proto.Message) {
	out, err := proto.Marshal(pb)
	if err != nil {
		log.Fatalln("Can't serialize to bytes", err)
		return
	}

	if err = os.WriteFile(file, out, 0644); err != nil {
		log.Fatalln("Can't write to file", err)
		return
	}

	fmt.Println("Data was written")
}

func readProto(file string, pb proto.Message) {
	in, err := os.ReadFile(file)
	if err != nil {
		log.Fatalln("Can't read file", err)
		return
	}

	if err := proto.Unmarshal(in, pb); err != nil {
		log.Fatalln("Unable to unmarshal", err)
		return
	}

	fmt.Println(pb)
}
