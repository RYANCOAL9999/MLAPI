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

	r.GET("/api/head", headerResponse)

	r.GET("/api/info", infoResponse)

	r.GET("/api/descible", describleResponse)

	r.POST("/api/linearSVC", linearSVCResponse)

	r.POST("/api/linearSVR", linearSVRResponse)

	r.POST("/api/svc", svcResponse)

	r.POST("/api/polynomialFeatures", polynomialFeaturesResponse)

	r.POST("/api/polynomialFeaturesFitTransform", polynomialFeaturesFitTransformResponse)

	r.POST("/api/nearestNeighbors", nearestNeighborsRespone)

	r.POST("/api/kdTree", kdTreeResponse)

	r.POST("/api/nearestCentroid", nearestCentroidResponse)

	r.POST("/api/linearRegression", linearRegressionResponse)

	r.POST("/api/ridge", ridgeRespone)

	r.POST("/api/ridgeCV", ridgeCVReponse)

	r.POST("/api/lasso", lassoReponse)

	r.POST("/api/lassoLars", lassoLarsReponse)

	r.POST("/api/bayesianRidge", bayesianRidgeResponse)

	r.POST("/api/tweedieRegressor", tweedieRegressorResponse)

	r.POST("/api/sgdClassifier", sgdClassifierResponse)

	r.POST("/api/elasticNet", elasticNetResponse)

	ginLambda = ginadapter.New(r)

}

func Handler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return ginLambda.ProxyWithContext(ctx, request)
}

func main() {
	lambda.Start(Handler)
}
