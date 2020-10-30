package main

import (
	"context"
	"github.com/wcse/monorepo/backend/api/feed"
	"google.golang.org/grpc"
	"log"
	"time"
)

const (
	feedAddress = "127.0.0.1:8081"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	feedConn, err := grpc.DialContext(ctx, feedAddress, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		panic(err)
	}

	client := feed.NewFeedClient(feedConn)

	reply, err := client.Post(ctx, &feed.PostRequest{
		AudioUrl:   "https://www.soundhelix.com/examples/mp3/SoundHelix-Song-1.mp3",
		Caption:    "This is my very first post",
		Transcript: "I'm Nemo",
		Privacy:    feed.Privacy_PRIVATE,
	})
	if err != nil {
		panic(err)
	}

	log.Println(reply)
}
