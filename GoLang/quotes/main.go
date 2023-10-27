package main

import (
	"context"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	ginadapter "github.com/awslabs/aws-lambda-go-api-proxy/gin"
	"github.com/gin-gonic/gin"
)

var ginLambda *ginadapter.GinLambda

func init() {

	log.Printf("Gin cold start")

	r := gin.Default()

	generalClient, err := newGeneralGRPCClient("localhost:50051")
	if err != nil {
		log.Fatalf(err.Error())
		return
	}

	linearClient, err := newLinearGRPCClient("localhost:50051")
	if err != nil {
		log.Fatalf(err.Error())
		return
	}

	neighborsClient, err := newNeighborsGPRCClient("localhost:50052")
	if err != nil {
		log.Fatalf(err.Error())
		return
	}

	polynomialClient, err := newPolynomialGRPCClient("localhost:50053")
	if err != nil {
		log.Fatalf(err.Error())
		return
	}

	svmClient, err := newSVMGRPCClient("localhost:50054")
	if err != nil {
		log.Fatalf(err.Error())
		return
	}

	r.GET("/api/head", headResponse(generalClient))

	r.GET("/api/info", infoResponse(generalClient))

	r.GET("/api/descible", describleResponse(generalClient))

	r.POST("/api/linearSVC", linearSVCResponse(svmClient))

	r.POST("/api/linearSVR", linearSVRResponse(svmClient))

	r.POST("/api/svc", svcResponse(svmClient))

	r.POST("/api/polynomialFeatures", polynomialFeaturesResponse(polynomialClient))

	r.POST("/api/polynomialFeaturesFitTransform", polynomialFeaturesFitTransformResponse(polynomialClient))

	r.POST("/api/nearestNeighbors", nearestNeighborsResponse(neighborsClient))

	r.POST("/api/kdTree", kdTreeResponse(neighborsClient))

	r.POST("/api/nearestCentroid", nearestCentroidResponse(neighborsClient))

	r.POST("/api/linearRegression", linearRegressionResponse(linearClient))

	r.POST("/api/ridge", ridgeResponse(linearClient))

	r.POST("/api/ridgeCV", ridgeCVResponse(linearClient))

	r.POST("/api/lasso", lassoResponse(linearClient))

	r.POST("/api/lassoLars", lassoLarsResponse(linearClient))

	r.POST("/api/bayesianRidge", bayesianRidgeResponse(linearClient))

	r.POST("/api/tweedieRegressor", tweedieRegressorResponse(linearClient))

	r.POST("/api/sgdClassifier", sgdClassifierResponse(linearClient))

	r.POST("/api/elasticNet", elasticNetResponse(linearClient))

	ginLambda = ginadapter.New(r)

}

func Handler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return ginLambda.ProxyWithContext(ctx, request)
}

func main() {
	lambda.Start(Handler)
}
