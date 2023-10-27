package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func readRequestData(c *gin.Context, dataPointer interface{}) error {
	bodyAsByteArray, err := io.ReadAll(c.Request.Body)
	if err != nil {
		return err
	}
	err = json.Unmarshal(bodyAsByteArray, dataPointer)
	return err
}

func handleGRPCError(c *gin.Context, err error, errMsg string) {
	log.Printf("could not %s: %v", errMsg, err)
	c.JSON(http.StatusInternalServerError, gin.H{"error": errMsg})
}

func handleBadRequestError(c *gin.Context, err error, error_code int, errMsg string) {
	log.Printf("%s could not: %v", errMsg, err)
	c.JSON(error_code, gin.H{"error": "Bad Request"})
}

func sendResponse(c *gin.Context, response interface{}, errMsg string) {
	responseData, err := json.Marshal(response)
	if err != nil {
		log.Printf("%v response encoding: %v", errMsg, err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not encode response"})
		return
	}

	c.JSON(http.StatusOK, responseData)
}
