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

var linearAddr = flag.String("addr", "localhost:50051", "the address to connect to")

var linearClient pb.LinearServiceClient

func init() {
	conn, err := grpc.Dial(*linearAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Could not connect to gRPC server: %v", err)
	}
	linearClient = pb.NewLinearServiceClient(conn)
}

func linearRegressionResponse(c *gin.Context) {

	bodyAsByteArray, err := io.ReadAll(c.Request.Body)

	if err != nil {
		log.Fatalf("could not linear Regression: %v", err)
		c.JSON(400, gin.H{"error": "Bad Request"})
		return
	}

	var data pb.LinearRegressionRequest

	err = json.Unmarshal(bodyAsByteArray, &data)

	if err != nil {
		log.Fatalf("could not linear Regression: %v", err)
		c.JSON(404, gin.H{"error": "missing parameter"})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := linearClient.LinearRegressionEvent(ctx, &pb.LinearRegressionRequest{
		XDropData:    data.XDropData,
		YDropData:    data.YDropData,
		Size:         data.Size,
		Random:       data.Random,
		Key:          data.Key,
		SampleWeight: data.SampleWeight,
		Kwargs:       data.Kwargs,
	})

	if err != nil {
		log.Fatalf("could not linear Regression: %v", err)
		c.JSON(500, gin.H{"error": "could not linear Regression"})
		return
	}

	c.JSON(200, gin.H{"dataFrame": r})

}

func ridgeRespone(c *gin.Context) {

	bodyAsByteArray, err := io.ReadAll(c.Request.Body)

	if err != nil {
		log.Fatalf("could not ridge : %v", err)
		c.JSON(400, gin.H{"error": "Bad Request"})
		return
	}

	var data pb.LinearRidgeRequest

	err = json.Unmarshal(bodyAsByteArray, &data)

	if err != nil {
		log.Fatalf("could not ridge: %v", err)
		c.JSON(404, gin.H{"error": "missing parameter"})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := linearClient.LinearRidgeEvent(ctx, &pb.LinearRidgeRequest{
		XDropData:    data.XDropData,
		YDropData:    data.YDropData,
		Size:         data.Size,
		Random:       data.Random,
		Key:          data.Key,
		SampleWeight: data.SampleWeight,
		Kwargs:       data.Kwargs,
	})

	if err != nil {
		log.Fatalf("could not Ridge: %v", err)
		c.JSON(500, gin.H{"error": "could not Ridge"})
		return
	}

	c.JSON(200, gin.H{"dataFrame": r})

}

func ridgeCVReponse(c *gin.Context) {

	bodyAsByteArray, err := io.ReadAll(c.Request.Body)

	if err != nil {
		log.Fatalf("could not ridge CV: %v", err)
		c.JSON(400, gin.H{"error": "Bad Request"})
		return
	}

	var data pb.LinearRidgeCVRequest

	err = json.Unmarshal(bodyAsByteArray, &data)

	if err != nil {
		log.Fatalf("could not ridge CV: %v", err)
		c.JSON(404, gin.H{"error": "missing parameter"})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := linearClient.LinearRidgeCVEvent(ctx, &pb.LinearRidgeCVRequest{})

	if err != nil {
		log.Fatalf("could not ridge CV: %v", err)
		c.JSON(500, gin.H{"error": "could not ridge CV"})
		return
	}

	c.JSON(200, gin.H{"dataFrame": r})

}

func lassoReponse(c *gin.Context) {

	bodyAsByteArray, err := io.ReadAll(c.Request.Body)

	if err != nil {
		log.Fatalf("could not lasso: %v", err)
		c.JSON(400, gin.H{"error": "Bad Request"})
		return
	}

	var data pb.LassoExpressionRequest

	err = json.Unmarshal(bodyAsByteArray, &data)

	if err != nil {
		log.Fatalf("could not lasso: %v", err)
		c.JSON(404, gin.H{"error": "missing parameter"})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := linearClient.LassoExpressionEvent(ctx, &pb.LassoExpressionRequest{
		Alpha:        data.Alpha,
		XDropData:    data.XDropData,
		YDropData:    data.YDropData,
		Size:         data.Size,
		Random:       data.Random,
		Key:          data.Key,
		SampleWeight: data.SampleWeight,
		Kwargs:       data.Kwargs,
	})

	if err != nil {
		log.Fatalf("could not lasso: %v", err)
		c.JSON(500, gin.H{"error": "could not lasso"})
		return
	}

	c.JSON(200, gin.H{"dataFrame": r})

}

func lassoLarsReponse(c *gin.Context) {

	bodyAsByteArray, err := io.ReadAll(c.Request.Body)

	if err != nil {
		log.Fatalf("could not lasso Lars: %v", err)
		c.JSON(400, gin.H{"error": "Bad Request"})
		return
	}

	var data pb.LassoLarsLassoExpressionRequest

	err = json.Unmarshal(bodyAsByteArray, &data)

	if err != nil {
		log.Fatalf("could not lasso Lars: %v", err)
		c.JSON(404, gin.H{"error": "missing parameter"})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := linearClient.LassoLarsLassoExpressionEvent(ctx, &pb.LassoLarsLassoExpressionRequest{
		Alpha:        data.Alpha,
		XDropData:    data.XDropData,
		YDropData:    data.YDropData,
		Size:         data.Size,
		Random:       data.Random,
		Key:          data.Key,
		SampleWeight: data.SampleWeight,
		Kwargs:       data.Kwargs,
	})

	if err != nil {
		log.Fatalf("could not lasso Lars: %v", err)
		c.JSON(500, gin.H{"error": "could not lasso Lars"})
		return
	}

	c.JSON(200, gin.H{"dataFrame": r})

}

func bayesianRidgeResponse(c *gin.Context) {

	bodyAsByteArray, err := io.ReadAll(c.Request.Body)

	if err != nil {
		log.Fatalf("could not bayesian Ridge: %v", err)
		c.JSON(400, gin.H{"error": "Bad Request"})
		return
	}

	var data pb.BayesianRidgeRequest

	err = json.Unmarshal(bodyAsByteArray, &data)

	if err != nil {
		log.Fatalf("could not bayesian Ridge: %v", err)
		c.JSON(404, gin.H{"error": "missing parameter"})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := linearClient.BayesianRidgeEvent(ctx, &pb.BayesianRidgeRequest{
		XDropData:    data.XDropData,
		YDropData:    data.YDropData,
		Size:         data.Size,
		Random:       data.Random,
		Key:          data.Key,
		SampleWeight: data.SampleWeight,
		Kwargs:       data.Kwargs,
	})

	if err != nil {
		log.Fatalf("could not bayesian Ridge: %v", err)
		c.JSON(500, gin.H{"error": "could not bayesian Ridge"})
		return
	}

	c.JSON(200, gin.H{"dataFrame": r})

}

func tweedieRegressorResponse(c *gin.Context) {

	bodyAsByteArray, err := io.ReadAll(c.Request.Body)

	if err != nil {
		log.Fatalf("could not tweedie Regressor: %v", err)
		c.JSON(400, gin.H{"error": "Bad Request"})
		return
	}

	var data pb.TweedieRegressorRequest

	err = json.Unmarshal(bodyAsByteArray, &data)

	if err != nil {
		log.Fatalf("could not tweedie Regressor: %v", err)
		c.JSON(404, gin.H{"error": "missing parameter"})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := linearClient.TweedieRegressorEvent(ctx, &pb.TweedieRegressorRequest{
		XDropData:    data.XDropData,
		YDropData:    data.YDropData,
		Size:         data.Size,
		Random:       data.Random,
		Key:          data.Key,
		SampleWeight: data.SampleWeight,
		Kwargs:       data.Kwargs,
	})

	if err != nil {
		log.Fatalf("could not tweedie Regressor: %v", err)
		c.JSON(500, gin.H{"error": "could not tweedie Regressor"})
		return
	}

	c.JSON(200, gin.H{"dataFrame": r})

}

func sgdClassifierResponse(c *gin.Context) {

	bodyAsByteArray, err := io.ReadAll(c.Request.Body)

	if err != nil {
		log.Fatalf("could not sgd Classifier: %v", err)
		c.JSON(400, gin.H{"error": "Bad Request"})
		return
	}

	var data pb.SGDClassifierRequest

	err = json.Unmarshal(bodyAsByteArray, &data)

	if err != nil {
		log.Fatalf("could not sgd Classifier: %v", err)
		c.JSON(404, gin.H{"error": "missing parameter"})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := linearClient.SGDClassifierEvent(ctx, &pb.SGDClassifierRequest{
		XDropData:    data.XDropData,
		YDropData:    data.YDropData,
		Size:         data.Size,
		Random:       data.Random,
		Key:          data.Key,
		SampleWeight: data.SampleWeight,
		Kwargs:       data.Kwargs,
	})

	if err != nil {
		log.Fatalf("could not sgd Classifier: %v", err)
		c.JSON(500, gin.H{"error": "could not sgd Classifier"})
		return
	}

	c.JSON(200, gin.H{"dataFrame": r})

}

func elasticNetResponse(c *gin.Context) {

	bodyAsByteArray, err := io.ReadAll(c.Request.Body)

	if err != nil {
		log.Fatalf("could not elastic Net: %v", err)
		c.JSON(400, gin.H{"error": "Bad Request"})
		return
	}

	var data pb.ElasticNetRequest

	err = json.Unmarshal(bodyAsByteArray, &data)

	if err != nil {
		log.Fatalf("could not elastic Net: %v", err)
		c.JSON(404, gin.H{"error": "missing parameter"})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := linearClient.ElasticNetEvent(ctx, &pb.ElasticNetRequest{
		Alpha:        data.Alpha,
		XDropData:    data.XDropData,
		YDropData:    data.YDropData,
		Size:         data.Size,
		Random:       data.Random,
		Key:          data.Key,
		SampleWeight: data.SampleWeight,
		Kwargs:       data.Kwargs,
	})

	if err != nil {
		log.Fatalf("could not elastic Net: %v", err)
		c.JSON(500, gin.H{"error": "could not elastic Net"})
		return
	}

	c.JSON(200, gin.H{"dataFrame": r})

}
