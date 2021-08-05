package main

import (
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
	"grpc-example-golang/generate"
	"io"
	"log"
)

func run(client generate.PersonServiceClient) {
	person := generate.Person{
		Name:  "小强",
		Id:    3,
		Email: "996@icu.com",
		Phones: []*generate.Person_PhoneNumber{
			{Number: "12345678910", Type: 0},
			{Number: "01987654321", Type: 1},
		},
		LastUpdated: nil,
	}
	addPerson, err := client.AddPerson(context.Background(), &person)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(addPerson)

	listPerson, err := client.ListPerson(context.Background(), &empty.Empty{})
	if err != nil {
		log.Fatalln(err)
	}
	for {
		person, err := listPerson.Recv()
		if err != nil {
			if err == io.EOF {
				break
			} else {
				log.Fatalln(err)
			}
		}
		log.Println(person)
	}
}

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalln(err)
	}
	// defer log.Fatalln(conn.Close())

	client := generate.NewPersonServiceClient(conn)

	run(client)
}
