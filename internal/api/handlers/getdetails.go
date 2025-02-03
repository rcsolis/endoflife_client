package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
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

	// call the rpc server
	response, err := rpc.GetDetails(name, version)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"data": response,
	})
}
