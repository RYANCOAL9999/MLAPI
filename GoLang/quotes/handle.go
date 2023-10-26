package main

import (
	"context"
	"log"
	"time"

	pb "github.com/RYANCOAL9999/AI_Go_Proto_File"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

func headEvent(c *gin.Context, conn *grpc.ClientConn) {

	client := pb.NewGeneralServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)

	defer cancel()

	r, err := client.HeaderEvent(ctx, &pb.HeaderRequest{})

	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	c.JSON(200, gin.H{"dataFrame": r})

}
