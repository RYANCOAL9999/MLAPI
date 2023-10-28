package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ReadRequestData(c *gin.Context, dataPointer interface{}) error {
	bodyAsByteArray, err := io.ReadAll(c.Request.Body)
	if err != nil {
		return err
	}
	err = json.Unmarshal(bodyAsByteArray, dataPointer)
	return err
}

func HandleGRPCError(c *gin.Context, err error, errMsg string) {
	log.Printf("could not %s: %v", errMsg, err)
	c.JSON(http.StatusInternalServerError, gin.H{"error": errMsg})
}

func HandleBadRequestError(c *gin.Context, err error, errMsg string) {
	log.Printf("could not %s: %v", errMsg, err)
	c.JSON(http.StatusBadRequest, gin.H{"error": "Bad Request"})
}

func SendResponse(c *gin.Context, response interface{}, errMsg string) {
	responseData, err := json.Marshal(response)
	if err != nil {
		log.Printf("%s encoding error: %v", errMsg, err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not encode response"})
		return
	}

	c.JSON(http.StatusOK, responseData)
}
