package main

import (
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
	"grpc-example-golang/generate"
	"log"
	"net"
)

type PersonServiceImpl struct {
	dataSource []*generate.Person
	generate.UnimplementedPersonServiceServer
}

func (s *PersonServiceImpl) AddPerson(ctx context.Context, person *generate.Person) (*generate.Person, error) {
	s.dataSource = append(s.dataSource, person)
	return person, nil
}

func (s *PersonServiceImpl) ListPerson(e *empty.Empty, stream generate.PersonService_ListPersonServer) error {
	for _, person := range s.dataSource {
		if err := stream.Send(person); err != nil {
			return err
		}
	}
	return nil
}

func newServer() *PersonServiceImpl {
	return &PersonServiceImpl{
		dataSource: []*generate.Person{
			{Name: "小明", Id: 1, Email: "123@qq.com", Phones: []*generate.Person_PhoneNumber{
				{Number: "12345678910", Type: 0},
				{Number: "01987654321", Type: 1},
			}},
			{Name: "小红", Id: 2, Email: "321@qq.com", Phones: []*generate.Person_PhoneNumber{
				{Number: "11111111111", Type: 0},
				{Number: "22222222222", Type: 1},
			}},
		},
	}
}

func main() {
	s := grpc.NewServer()

	lis, err := net.Listen("tcp", "localhost:50051")
	if err != nil {
		log.Fatalln(err)
	}

	generate.RegisterPersonServiceServer(s, newServer())

	log.Fatalln(s.Serve(lis))
}
