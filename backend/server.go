package main

import (
	"context"
	_ "github.com/go-sql-driver/mysql"
	"google.golang.org/grpc"
	"log"
	"net"
	"sync"
	"wetalk/monorepo/backend/api/identity"
	"wetalk/monorepo/backend/auth"
	"wetalk/monorepo/backend/ent"
)

func main() {

	lis, err := net.Listen("tcp", ":9001")

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	entClient, err := ent.Open("mysql", "root:hoatong12@tcp(127.0.0.1:3306)/wetalk?parseTime=true")

	if err != nil {
		log.Fatal(err)
	}

	if err := entClient.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	s := &auth.Server{
		Client: entClient,
	}

	grpcServer := grpc.NewServer()

	identity.RegisterAuthServer(grpcServer, s)

	var wg sync.WaitGroup

	wg.Add(1)

	go func() {
		defer wg.Done()
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %s", err)
		}
	}()
	wg.Wait()

}
