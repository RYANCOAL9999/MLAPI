package main

import (
	"context"
	"log"

	model "github.com/RYANCOAL9999/MLAPI/GoLang/quotes/model"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	ginadapter "github.com/awslabs/aws-lambda-go-api-proxy/gin"
	"github.com/gin-gonic/gin"
)

var ginLambda *ginadapter.GinLambda

func init() {

	log.Printf("Gin cold start")

	r := gin.Default()

	generalClient, err := model.NewGeneralGRPCClient("localhost:50051")
	if err != nil {
		log.Fatalf(err.Error())
		return
	}

	linearClient, err := model.NewLinearGRPCClient("localhost:50051")
	if err != nil {
		log.Fatalf(err.Error())
		return
	}

	neighborsClient, err := model.NewNeighborsGPRCClient("localhost:50052")
	if err != nil {
		log.Fatalf(err.Error())
		return
	}

	polynomialClient, err := model.NewPolynomialGRPCClient("localhost:50053")
	if err != nil {
		log.Fatalf(err.Error())
		return
	}

	svmClient, err := model.NewSVMGRPCClient("localhost:50054")
	if err != nil {
		log.Fatalf(err.Error())
		return
	}

	r.GET("/api/head", model.HeaderResponse(generalClient))

	r.GET("/api/info", model.InfoResponse(generalClient))

	r.GET("/api/descible", model.DescriblerResponse(generalClient))

	r.POST("/api/linearSVC", model.LinearSVCResponse(svmClient))

	r.POST("/api/linearSVR", model.LinearSVRResponse(svmClient))

	r.POST("/api/svc", model.SVCResponse(svmClient))

	r.POST("/api/polynomialFeatures", model.PolynomialFeaturesResponse(polynomialClient))

	r.POST("/api/polynomialFeaturesFitTransform", model.PolynomialFeaturesFitTransformResponse(polynomialClient))

	r.POST("/api/nearestNeighbors", model.NearestNeighborsResponse(neighborsClient))

	r.POST("/api/kdTree", model.KDTreeResponse(neighborsClient))

	r.POST("/api/nearestCentroid", model.NearestCentroidResponse(neighborsClient))

	r.POST("/api/linearRegression", model.LinearRegressionResponse(linearClient))

	r.POST("/api/ridge", model.RidgeResponse(linearClient))

	r.POST("/api/ridgeCV", model.RidgeCVResponse(linearClient))

	r.POST("/api/lasso", model.LassoResponse(linearClient))

	r.POST("/api/lassoLars", model.LassoLarsResponse(linearClient))

	r.POST("/api/bayesianRidge", model.BayesianRidgeResponse(linearClient))

	r.POST("/api/tweedieRegressor", model.TweedieRegressorResponse(linearClient))

	r.POST("/api/sgdClassifier", model.SGDClassifierResponse(linearClient))

	r.POST("/api/elasticNet", model.ElasticNetResponse(linearClient))

	ginLambda = ginadapter.New(r)

}

func Handler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return ginLambda.ProxyWithContext(ctx, request)
}

func main() {
	lambda.Start(Handler)
}
