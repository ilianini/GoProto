package main

import (
	"log"

	"github.com/golang/protobuf/proto"
	pb "github.com/ilianini/GoProto/protofiles"
)

func main() {

	p := &pb.Persona{
		Nombre: "Esteban",
		Edad:   23,
		Peso:   70,
		SeguidoresSociales: []*pb.SeguidoresSociales{
			{Youtube: 10, Twitter: 0},
		},
	}

	out, err := proto.Marshal(p)
	if err != nil {
		log.Fatalln("Failed to encode address book:", err)
	}

}
