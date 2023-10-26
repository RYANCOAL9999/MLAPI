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
		log.Fatalf("could not linear regression: %v", err)
		c.JSON(400, gin.H{"error": "Bad Request"})
		return
	}

	var data pb.LinearRegressionRequest

	err = json.Unmarshal(bodyAsByteArray, &data)

	if err != nil {
		log.Fatalf("could not linear regression: %v", err)
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
		log.Fatalf("could not linear regression: %v", err)
		c.JSON(500, gin.H{"error": "could not linear regression"})
		return
	}

	responseData, err := json.Marshal(r)

	if err != nil {
		log.Fatalf("response encoding linear regression: %v", err)
		c.JSON(505, gin.H{"error": "could not linear regression"})
		return
	}

	c.JSON(200, responseData)
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
		log.Fatalf("could not ridge: %v", err)
		c.JSON(500, gin.H{"error": "could not ridge"})
		return
	}

	responseData, err := json.Marshal(r)

	if err != nil {
		log.Fatalf("response encoding ridge: %v", err)
		c.JSON(505, gin.H{"error": "could not ridge"})
		return
	}

	c.JSON(200, responseData)

}

func ridgeCVReponse(c *gin.Context) {

	bodyAsByteArray, err := io.ReadAll(c.Request.Body)

	if err != nil {
		log.Fatalf("could not ridge cv: %v", err)
		c.JSON(400, gin.H{"error": "Bad Request"})
		return
	}

	var data pb.LinearRidgeCVRequest

	err = json.Unmarshal(bodyAsByteArray, &data)

	if err != nil {
		log.Fatalf("could not ridge cv: %v", err)
		c.JSON(404, gin.H{"error": "missing parameter"})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := linearClient.LinearRidgeCVEvent(ctx, &pb.LinearRidgeCVRequest{})

	if err != nil {
		log.Fatalf("could not ridge cv: %v", err)
		c.JSON(500, gin.H{"error": "could not ridge cv"})
		return
	}

	responseData, err := json.Marshal(r)

	if err != nil {
		log.Fatalf("response encoding ridge cv %v", err)
		c.JSON(505, gin.H{"error": "could not ridge cv"})
		return
	}

	c.JSON(200, responseData)

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

	responseData, err := json.Marshal(r)

	if err != nil {
		log.Fatalf("response encoding lasso: %v", err)
		c.JSON(505, gin.H{"error": "could not lasso"})
		return
	}

	c.JSON(200, responseData)

}

func lassoLarsReponse(c *gin.Context) {

	bodyAsByteArray, err := io.ReadAll(c.Request.Body)

	if err != nil {
		log.Fatalf("could not lasso lars: %v", err)
		c.JSON(400, gin.H{"error": "Bad Request"})
		return
	}

	var data pb.LassoLarsLassoExpressionRequest

	err = json.Unmarshal(bodyAsByteArray, &data)

	if err != nil {
		log.Fatalf("could not lasso lars: %v", err)
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
		log.Fatalf("could not lasso lars: %v", err)
		c.JSON(500, gin.H{"error": "could not lasso lars"})
		return
	}

	responseData, err := json.Marshal(r)

	if err != nil {
		log.Fatalf("response encoding lasso lars: %v", err)
		c.JSON(505, gin.H{"error": "could not lasso lars"})
		return
	}

	c.JSON(200, responseData)

}

func bayesianRidgeResponse(c *gin.Context) {

	bodyAsByteArray, err := io.ReadAll(c.Request.Body)

	if err != nil {
		log.Fatalf("could not bayesian ridge: %v", err)
		c.JSON(400, gin.H{"error": "Bad Request"})
		return
	}

	var data pb.BayesianRidgeRequest

	err = json.Unmarshal(bodyAsByteArray, &data)

	if err != nil {
		log.Fatalf("could not bayesian ridge: %v", err)
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
		log.Fatalf("could not bayesian ridge: %v", err)
		c.JSON(500, gin.H{"error": "could not bayesian ridge"})
		return
	}

	responseData, err := json.Marshal(r)

	if err != nil {
		log.Fatalf("response encoding bayesian ridge: %v", err)
		c.JSON(505, gin.H{"error": "could not bayesian ridge"})
		return
	}

	c.JSON(200, responseData)

}

func tweedieRegressorResponse(c *gin.Context) {

	bodyAsByteArray, err := io.ReadAll(c.Request.Body)

	if err != nil {
		log.Fatalf("could not tweedie regressor: %v", err)
		c.JSON(400, gin.H{"error": "Bad Request"})
		return
	}

	var data pb.TweedieRegressorRequest

	err = json.Unmarshal(bodyAsByteArray, &data)

	if err != nil {
		log.Fatalf("could not tweedie regressor: %v", err)
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
		log.Fatalf("could not tweedie regressor: %v", err)
		c.JSON(500, gin.H{"error": "could not tweedie regressor"})
		return
	}

	responseData, err := json.Marshal(r)

	if err != nil {
		log.Fatalf("response encoding tweedie regressor: %v", err)
		c.JSON(505, gin.H{"error": "could not tweedie regressor"})
		return
	}

	c.JSON(200, responseData)

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
		log.Fatalf("could not sgd classifier: %v", err)
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
		log.Fatalf("could not sgd classifier: %v", err)
		c.JSON(500, gin.H{"error": "could not sgd classifier"})
		return
	}

	responseData, err := json.Marshal(r)

	if err != nil {
		log.Fatalf("response encoding sgd classifier: %v", err)
		c.JSON(505, gin.H{"error": "could not sgd classifier"})
		return
	}

	c.JSON(200, responseData)

}

func elasticNetResponse(c *gin.Context) {

	bodyAsByteArray, err := io.ReadAll(c.Request.Body)

	if err != nil {
		log.Fatalf("could not elastic net: %v", err)
		c.JSON(400, gin.H{"error": "Bad Request"})
		return
	}

	var data pb.ElasticNetRequest

	err = json.Unmarshal(bodyAsByteArray, &data)

	if err != nil {
		log.Fatalf("could not elastic net: %v", err)
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
		log.Fatalf("could not elastic net: %v", err)
		c.JSON(500, gin.H{"error": "could not elastic net"})
		return
	}

	responseData, err := json.Marshal(r)

	if err != nil {
		log.Fatalf("response encoding elastic net: %v", err)
		c.JSON(505, gin.H{"error": "could not elastic net"})
		return
	}

	c.JSON(200, responseData)

}
