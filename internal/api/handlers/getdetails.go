package handlers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rcsolis/endoflife_client/internal/database"
	"github.com/rcsolis/endoflife_client/internal/rpc"
)

/**
 * GetDetails is a handler for the GET /:name/version/:version route
 * @param c *gin.Context
 */
func GetDetails(c *gin.Context) {
	// Get params
	name := c.Param("name")
	version := c.Param("version")
	// First check if name and version are empty
	if name == "" || version == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Name and version are required",
		})
		return
	}
	data, ok := database.GetDetails(name, version)
	if ok {
		log.Printf("Data found in the database")
		c.JSON(http.StatusOK, gin.H{
			"data": data,
		})
		return
	}
	// call the rpc server
	log.Print("Data requested from the rpc server")
	response, err := rpc.GetDetails(name, version)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	// Save the data to the database
	err = database.SaveDetails(name, version, response)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": response,
	})
}
