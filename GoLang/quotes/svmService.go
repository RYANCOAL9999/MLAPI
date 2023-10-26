package main

import (
	"context"
	"encoding/json"
	"flag"
	"io"
	"log"
	"time"

	pb "github.com/RYANCOAL9999/AI_Go_Proto_File"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var svmClient pb.SVMServiceClient

var svmAddr = flag.String("addr", "localhost:50054", "the address to connect to")

func init() {
	conn, err := grpc.Dial(*svmAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Could not connect to gRPC server: %v", err)
	}
	svmClient = pb.NewSVMServiceClient(conn)
}

func linearSVCResponse(c *gin.Context) {

	bodyAsByteArray, err := io.ReadAll(c.Request.Body)

	if err != nil {
		log.Fatalf("could not linear SVC: %v", err)
		c.JSON(400, gin.H{"error": "Bad Request"})
		return
	}

	var data pb.LinearRequest

	err = json.Unmarshal(bodyAsByteArray, &data)

	if err != nil {
		log.Fatalf("could not linear SVC: %v", err)
		c.JSON(404, gin.H{"error": "missing parameter"})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := svmClient.LinearSVCEvent(ctx, &pb.LinearRequest{
		XDropData:    data.XDropData,
		YDropData:    data.YDropData,
		Size:         data.Size,
		Random:       data.Random,
		Key:          data.Key,
		SampleWeight: data.SampleWeight,
		Kwargs:       data.Kwargs,
	})

	if err != nil {
		log.Fatalf("could not linear SVC: %v", err)
		c.JSON(500, gin.H{"error": "could not linear SVC"})
		return
	}

	c.JSON(200, gin.H{"dataFrame": r})

}

func linearSVRResponse(c *gin.Context) {

	bodyAsByteArray, err := io.ReadAll(c.Request.Body)

	if err != nil {
		log.Fatalf("could not linear SVR: %v", err)
		c.JSON(400, gin.H{"error": "Bad Request"})
		return
	}

	var data pb.LinearRequest

	err = json.Unmarshal(bodyAsByteArray, &data)

	if err != nil {
		log.Fatalf("could not linear SVR: %v", err)
		c.JSON(404, gin.H{"error": "missing parameter"})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := svmClient.LinearSVREvent(ctx, &pb.LinearRequest{
		XDropData:    data.XDropData,
		YDropData:    data.YDropData,
		Size:         data.Size,
		Random:       data.Random,
		Key:          data.Key,
		SampleWeight: data.SampleWeight,
		Kwargs:       data.Kwargs,
	})

	if err != nil {
		log.Fatalf("could not linear SVR: %v", err)
		c.JSON(500, gin.H{"error": "could not linear SVR"})
		return
	}

	c.JSON(200, gin.H{"dataFrame": r})

}

func svcResponse(c *gin.Context) {

	bodyAsByteArray, err := io.ReadAll(c.Request.Body)

	if err != nil {
		log.Fatalf("could not svc: %v", err)
		c.JSON(400, gin.H{"error": "Bad Request"})
		return
	}

	var data pb.SVCRequest

	err = json.Unmarshal(bodyAsByteArray, &data)

	if err != nil {
		log.Fatalf("could not svc: %v", err)
		c.JSON(404, gin.H{"error": "missing parameter"})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := svmClient.SVCEvent(ctx, &pb.SVCRequest{
		XDropData:    data.XDropData,
		YDropData:    data.YDropData,
		Size:         data.Size,
		Random:       data.Random,
		Key:          data.Key,
		SampleWeight: data.SampleWeight,
		Kwargs:       data.Kwargs,
	})

	if err != nil {
		log.Fatalf("could not svc: %v", err)
		c.JSON(500, gin.H{"error": "could not svc"})
		return
	}

	c.JSON(200, gin.H{"dataFrame": r})

}
