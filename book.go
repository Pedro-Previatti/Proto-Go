package main

import (
	"log"
	"os"

	pb "github.com/Pedro-Previatti/Proto-Go/proto"
	"google.golang.org/protobuf/proto"
)

func writeBook(file string, p *pb.AddressBook) {
	out, err := proto.Marshal(p)
	if err != nil {
		log.Fatalln("Failed to encode address book:", err)
	}
	if err := os.WriteFile(file, out, 0644); err != nil {
		log.Fatalln("Failed to write address book:", err)
	}
}

func readBook(file string, p *pb.AddressBook) {
	in, err := os.ReadFile(file)
	if err != nil {
		log.Fatalln("Error reading file:", err)
	}
	if err := proto.Unmarshal(in, p); err != nil {
		log.Fatalln("Failed to parse address book:", err)
	}
}
