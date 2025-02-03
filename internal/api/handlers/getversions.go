package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rcsolis/endoflife_client/internal/model"
	"github.com/rcsolis/endoflife_client/internal/rpc"
)

/**
 * GetVersions is a handler for the GET /api/versions/:name endpoint
 * @param c *gin.Context
 */
func GetVersions(c *gin.Context) {
	name := c.Param("name")
	// Clear the model before fetching new data
	model.TechnologiesCycle = nil
	// Get from server
	err := rpc.GetAllVersions(name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": model.TechnologiesCycle,
	})
}
