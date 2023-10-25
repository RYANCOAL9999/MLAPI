package main

import {
	"github.com/gin-gonic/gin"
	"github.com/go-openapi/loads"
	"github.com/go-openapi/runtime/middleware"
	"mlAPI/restapi"
    "mlAPI/restapi/operations"
	"github.com/go-openapi/swag"
	"google.golang.org/grpc"
}

func main() {
	r := gin.Default()

	doc, _ := loads.Analyzed(restapi.SwaggerJSON, "")
	api := operations.NewMyServerlessAppAPI(doc)

	// Load Swagger documentation
	r.Static("/swaggerui/", "./swaggerui/")

	// Serve Swagger UI
	r.GET("/swagger.json", func(c *gin.Context) {
		c.JSON(200, doc.Spec())
	})

	// Set up gRPC client connection
    conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
    if err != nil {
        log.Fatalf("Could not connect to gRPC server: %v", err)
    }
    defer conn.Close()

    // Create a gRPC client
    client := your_grpc_package.NewYourGRPCServiceClient(conn)

	r.GET("/", func(c *gin.Context) {
        c.JSON(204, gin.H{"message": "Do not wanna attack my server!"})
    })

	r.GET("/head", func(c *gin.Context) {
        // Use the gRPC client to make requests to your gRPC service
        response, err := client.YourGRPCMethod(context.Background(), &your_grpc_package.YourRequest{})
        if err != nil {
            c.JSON(500, gin.H{"error": err.Error()})
            return
        }
        c.JSON(200, gin.H{"message": response.Message})
    })



	r.Use(middleware.Recover())
    r.Use(middleware.Spec(api.Context, doc))

	// Define your API routes
    restapi.RegisterHandlers(r, api)

	r.Run()

}