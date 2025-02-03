package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rcsolis/endoflife_client/internal/rpc"
)

/**
 * GetAll is a handler function to get all the technologies
 * @param c *gin.Context
 */
func GetAll(c *gin.Context) {
	// Get All from RPC server
	response, err := rpc.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"data": response,
	})
}
