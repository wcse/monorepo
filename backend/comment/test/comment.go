package main

import (
	"context"
	"github.com/google/uuid"
	"github.com/wcse/monorepo/backend/api/comment"
	"google.golang.org/grpc"
	"log"
	"time"
)

const (
	commentAddress = "127.0.0.1:8082"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	commentConn, err := grpc.DialContext(ctx, commentAddress, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		panic(err)
	}

	client := comment.NewCommentClient(commentConn)

	reply, err := client.Write(ctx, &comment.CommentRequest{
		FeedId:  uuid.New().String(),
		Content: "This is my very first comment",
	})
	if err != nil {
		panic(err)
	}

	log.Println(reply)
}
