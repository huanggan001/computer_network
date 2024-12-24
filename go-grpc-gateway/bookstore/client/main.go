package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	books "grpc/bookstore/pb"
	"log"
	"time"
)

func main() {
	creds, errr := credentials.NewClientTLSFromFile("./cert/server.pem", "grpc_example")
	if errr != nil {
		log.Fatalln(errr)
		return
	}
	conn, err := grpc.Dial("localhost:8080", grpc.WithTransportCredentials(creds))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := books.NewBookstoreClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := client.ListBooksByShelf(ctx, &books.ListBooksRequest{
		Id: 1,
	})
	if err != nil {
		log.Fatalf("could not list books: %v", err)
	}
	fmt.Println(r)
}
