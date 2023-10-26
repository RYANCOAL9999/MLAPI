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

var polynomialAddress = flag.String("addr", "localhost:50053", "the address to connect to")

var polynomialClient pb.PolynomialServiceClient

func init() {
	conn, err := grpc.Dial(*polynomialAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Could not connect to gRPC server: %v", err)
	}
	polynomialClient = pb.NewPolynomialServiceClient(conn)
}

func polynomialFeaturesResponse(c *gin.Context) {

	bodyAsByteArray, err := io.ReadAll(c.Request.Body)

	if err != nil {
		log.Fatalf("could not polynomial Features: %v", err)
		c.JSON(400, gin.H{"error": "Bad Request"})
		return
	}

	var data pb.PolynomialFeaturesFitTransformRequest

	err = json.Unmarshal(bodyAsByteArray, &data)

	if err != nil {
		log.Fatalf("could not polynomial Features: %v", err)
		c.JSON(404, gin.H{"error": "missing parameter"})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := polynomialClient.PolynomialFeaturesEvent(ctx, &pb.PolynomialFeaturesRequest{
		XDropData:    data.XDropData,
		YDropData:    data.YDropData,
		Size:         data.Size,
		Random:       data.Random,
		Key:          data.Key,
		SampleWeight: data.SampleWeight,
		Degree:       data.Degree,
		Instrument:   data.Instrument,
	})

	if err != nil {
		log.Fatalf("could not polynomial Features: %v", err)
		c.JSON(500, gin.H{"error": "could not polynomial Features"})
		return
	}

	c.JSON(200, gin.H{"dataFrame": r})

}

func polynomialFeaturesFitTransformResponse(c *gin.Context) {

	bodyAsByteArray, err := io.ReadAll(c.Request.Body)

	if err != nil {
		log.Fatalf("could not polynomial Features Fit Transform: %v", err)
		c.JSON(400, gin.H{"error": "Bad Request"})
		return
	}

	var data pb.PolynomialFeaturesFitTransformRequest

	err = json.Unmarshal(bodyAsByteArray, &data)

	if err != nil {
		log.Fatalf("could not polynomial Features Fit Transform: %v", err)
		c.JSON(404, gin.H{"error": "missing parameter"})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := polynomialClient.PolynomialFeaturesFitTransformEvent(ctx, &pb.PolynomialFeaturesFitTransformRequest{
		XDropData:    data.XDropData,
		YDropData:    data.YDropData,
		Size:         data.Size,
		Random:       data.Random,
		Key:          data.Key,
		SampleWeight: data.SampleWeight,
		Degree:       data.Degree,
		Instrument:   data.Instrument,
	})

	if err != nil {
		log.Fatalf("could not polynomial Features Fit Transform: %v", err)
		c.JSON(500, gin.H{"error": "could not polynomial Features Fit Transform"})
		return
	}

	c.JSON(200, gin.H{"dataFrame": r})

}
