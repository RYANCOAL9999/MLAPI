package main

import (
	"context"
	"encoding/json"
	"flag"
	"log"
	"time"

	pb "github.com/RYANCOAL9999/AI_Go_Proto_File"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var generalAddr = flag.String("addr", "localhost:50051", "the address to connect to")

var generalClient pb.GeneralServiceClient

func init() {
	conn, err := grpc.Dial(*generalAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Could not connect to gRPC server: %v", err)
	}
	generalClient = pb.NewGeneralServiceClient(conn)
}

func infoResponse(c *gin.Context) {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := generalClient.InfoEvent(ctx, &pb.InfoRequest{})

	if err != nil {
		log.Fatalf("could not info: %v", err)
		c.JSON(500, gin.H{"error": "could not info"})
		return
	}

	responseData, err := json.Marshal(r)

	if err != nil {
		log.Fatalf("response encoding info: %v", err)
		c.JSON(505, gin.H{"error": "could not info"})
		return
	}

	c.JSON(200, responseData)

}

func describleResponse(c *gin.Context) {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := generalClient.DescriblerEvent(ctx, &pb.DescribeRequest{})

	if err != nil {
		log.Fatalf("could not describle: %v", err)
		c.JSON(500, gin.H{"error": "could not describle"})
		return
	}

	responseData, err := json.Marshal(r)

	if err != nil {
		log.Fatalf("response encoding info: %v", err)
		c.JSON(505, gin.H{"error": "could not describle"})
		return
	}

	c.JSON(200, responseData)

}

func headerResponse(c *gin.Context) {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := generalClient.HeaderEvent(ctx, &pb.HeaderRequest{})

	if err != nil {
		log.Fatalf("could not header: %v", err)
		c.JSON(500, gin.H{"error": "could not header"})
		return
	}

	responseData, err := json.Marshal(r)

	if err != nil {
		log.Fatalf("response encoding header: %v", err)
		c.JSON(505, gin.H{"error": "could not header"})
		return
	}

	c.JSON(200, responseData)

}
