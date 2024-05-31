package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strings"

	"github.com/SalemLU/GoServerTest/getProducer"
	"google.golang.org/grpc"
)

type myGetProducerServer struct {
	getProducer.UnimplementedGetProducerServer
}

func (m *Movies) SearchDirector(title string, year int64) string {
	for i := 0; i < len(m.Movies); i++ {
		if strings.EqualFold(m.Movies[i].Title, title) {
			if m.Movies[i].Year == year {
				fmt.Println("Found movie")
				return m.Movies[i].Director
			}
		}
	}

	return "not found"
}

func (s myGetProducerServer) Create(ctx context.Context, req *getProducer.CreateRequest) (*getProducer.CreateResponse, error) {

	jsonFile, err := os.Open("db.json")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened db.json")

	byteValue, _ := io.ReadAll(jsonFile)
	var movies Movies
	json.Unmarshal(byteValue, &movies)

	result := movies.SearchDirector(req.Film.Title, req.Film.Year)

	defer jsonFile.Close()

	return &getProducer.CreateResponse{
		Director: []byte(result),
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":8089")
	if err != nil {
		log.Fatalf("cannot create listener: %s", err)
	}

	serverRegistrar := grpc.NewServer()
	service := &myGetProducerServer{}

	getProducer.RegisterGetProducerServer(serverRegistrar, service)
	err = serverRegistrar.Serve(lis)
	if err != nil {
		log.Fatalf("impossible to serve: %s", err)
	}

}
