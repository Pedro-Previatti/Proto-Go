package main

import (
	"log"

	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

func toJSON(pb proto.Message) string {
	out, err := protojson.Marshal(pb)
	if err != nil {
		log.Fatalln("Unable to convert proto to json", err)
		return ""
	}

	return string(out)
}

func fromJSON(in string, pb proto.Message) {
	if err := protojson.Unmarshal([]byte(in), pb); err != nil {
		log.Fatalln("Unable to convert json to proto", err)
	}
}
