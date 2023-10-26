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

var neighborsAddr = flag.String("addr", "localhost:50052", "the address to connect to")

var neighborsClient pb.NeighborsServiceClient

func init() {
	conn, err := grpc.Dial(*neighborsAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Could not connect to gRPC server: %v", err)
	}
	neighborsClient = pb.NewNeighborsServiceClient(conn)
}

func nearestNeighborsRespone(c *gin.Context) {

	bodyAsByteArray, err := io.ReadAll(c.Request.Body)

	if err != nil {
		log.Fatalf("could not nearest neighbors: %v", err)
		c.JSON(400, gin.H{"error": "Bad Request"})
		return
	}

	var data pb.NearestNeighborsRequest

	err = json.Unmarshal(bodyAsByteArray, &data)

	if err != nil {
		log.Fatalf("could not nearest neighbors: %v", err)
		c.JSON(404, gin.H{"error": "missing parameter"})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := neighborsClient.NearestNeighborsEvent(ctx, &pb.NearestNeighborsRequest{
		XDropData: data.XDropData,
		YDropData: data.YDropData,
		Size:      data.Size,
		Random:    data.Random,
		Key:       data.Key,
		Kwargs:    data.Kwargs,
	})

	if err != nil {
		log.Fatalf("could not nearest neighbors: %v", err)
		c.JSON(500, gin.H{"error": "could not nearest neighbors"})
		return
	}

	responseData, err := json.Marshal(r)

	if err != nil {
		log.Fatalf("response encoding nearest neighbors: %v", err)
		c.JSON(505, gin.H{"error": "could not nearest neighbors"})
		return
	}

	c.JSON(200, responseData)

}

func kdTreeResponse(c *gin.Context) {

	bodyAsByteArray, err := io.ReadAll(c.Request.Body)

	if err != nil {
		log.Fatalf("could not kd tree: %v", err)
		c.JSON(400, gin.H{"error": "Bad Request"})
		return
	}

	var data pb.KDTreeRequest

	err = json.Unmarshal(bodyAsByteArray, &data)

	if err != nil {
		log.Fatalf("could not kd tree: %v", err)
		c.JSON(404, gin.H{"error": "missing parameter"})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := neighborsClient.KDTreeEvent(ctx, &pb.KDTreeRequest{
		XDropData:    data.XDropData,
		YDropData:    data.YDropData,
		Size:         data.Size,
		Random:       data.Random,
		Key:          data.Key,
		SampleWeight: data.SampleWeight,
		K:            data.K,
		Instrument:   data.Instrument,
	})

	if err != nil {
		log.Fatalf("could not kd tree: %v", err)
		c.JSON(500, gin.H{"error": "could not kd Tree"})
		return
	}

	responseData, err := json.Marshal(r)

	if err != nil {
		log.Fatalf("response encoding kd tree: %v", err)
		c.JSON(505, gin.H{"error": "could not kd tree"})
		return
	}

	c.JSON(200, responseData)

}

func nearestCentroidResponse(c *gin.Context) {

	bodyAsByteArray, err := io.ReadAll(c.Request.Body)

	if err != nil {
		log.Fatalf("could not nearest centroid: %v", err)
		c.JSON(400, gin.H{"error": "Bad Request"})
		return
	}

	var data pb.NearestCentroidRequest

	err = json.Unmarshal(bodyAsByteArray, &data)

	if err != nil {
		log.Fatalf("could not nearest centroid: %v", err)
		c.JSON(404, gin.H{"error": "missing parameter"})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := neighborsClient.NearestCentroidEvent(ctx, &pb.NearestCentroidRequest{
		XDropData:       data.XDropData,
		YDropData:       data.YDropData,
		Size:            data.Size,
		Random:          data.Random,
		Key:             data.Key,
		Metric:          data.Metric,
		ShrinkThreshold: data.ShrinkThreshold,
	})

	if err != nil {
		log.Fatalf("could not nearest centroid: %v", err)
		c.JSON(500, gin.H{"error": "could not nearest centroid"})
		return
	}

	responseData, err := json.Marshal(r)

	if err != nil {
		log.Fatalf("response encoding nearest centroid: %v", err)
		c.JSON(505, gin.H{"error": "could not nearest centroid"})
		return
	}

	c.JSON(200, responseData)

}
