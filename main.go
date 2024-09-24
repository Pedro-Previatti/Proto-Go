package main

import (
	"fmt"
	"reflect"

	pb "github.com/Pedro-Previatti/Proto-Go/proto"
	"google.golang.org/protobuf/proto"
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

func doMap() *pb.MapExample {
	return &pb.MapExample{
		Ids: map[string]*pb.IdWrapper{
			"key200": {Id: 200},
			"key201": {Id: 201},
			"key202": {Id: 202},
		},
	}
}

func doFile(p proto.Message) {
	path := "simple.bin"

	writeProto(path, p)
	message := &pb.Simple{}
	readProto(path, message)
	// fmt.Println(message)
}

func doToJSON(p proto.Message) string {
	json := toJSON(p)
	// fmt.Println(json)
	return json
}

func doFromJSON(json string, t reflect.Type) proto.Message {
	message := reflect.New(t).Interface().(proto.Message)
	fromJSON(json, message)
	return message
}

func doBook(p *pb.AddressBook) {
	path := "address_book.bin"

	writeBook(path, p)
	message := &pb.AddressBook{}
	readBook(path, message)

	fmt.Print(doToJSON(message))
}

func bookHelper() *pb.AddressBook {
	return &pb.AddressBook{
		People: []*pb.Person{
			{
				Id:    0001,
				Name:  "Test One",
				Email: "test1@email.com",
				Phones: []*pb.Person_PhoneNumber{
					{Number: "00 0 0000-0000", Type: pb.PhoneType_PHONE_TYPE_MOBILE},
					{Number: "00 0 1111-1111", Type: pb.PhoneType_PHONE_TYPE_HOME},
					{Number: "00 0 2222-2222", Type: pb.PhoneType_PHONE_TYPE_WORK},
				},
			},
			{
				Id:    0002,
				Name:  "Test Two",
				Email: "test2@email.com",
				Phones: []*pb.Person_PhoneNumber{
					{Number: "00 0 3333-3333", Type: pb.PhoneType_PHONE_TYPE_MOBILE},
					{Number: "00 0 4444-4444", Type: pb.PhoneType_PHONE_TYPE_HOME},
					{Number: "00 0 5555-5555", Type: pb.PhoneType_PHONE_TYPE_WORK},
				},
			},
		},
	}
}

func main() {
	// fmt.Println("Simple Proto Message:")
	// fmt.Println(doSimple())

	// fmt.Println("Complex Proto Message:")
	// fmt.Println(doComplex())

	// fmt.Println("Proto Message with Enumeration:")
	// fmt.Println(doEnumeration())

	// fmt.Println("Proto OneOfs Id:")
	// doOneOf(&pb.Result_Id{Id: 1})
	// fmt.Println("Proto OneOfs Message:")
	// doOneOf(&pb.Result_Message{Message: "Some message"})

	// fmt.Println("Map Proto Message:")
	// fmt.Println(doMap())

	// fmt.Println("Proto Read and Write to File:")
	// doFile(doSimple())

	// fmt.Println("Proto to JSON:")
	// json := doToJSON(doComplex())
	// fmt.Println(json)

	// fmt.Println("Proto to JSON:")
	// proto := doFromJSON(json, reflect.TypeOf(pb.Complex{}))
	// fmt.Println(proto)

	// fmt.Println("Unknown type check")
	// fmt.Println(doFromJSON(`{"id": 1, "unknown": "test"}`, reflect.TypeOf(pb.Simple{})))

	fmt.Println("Book:")
	doBook(bookHelper())
}
