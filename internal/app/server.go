package app

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

const (
	minValidNumber = 0
	maxValidNumber = 10000000
)

// IndexFinder is an interface for finding the index of a number
type IndexFinder interface {
	FindIndex(value int) (int, bool)
}

// NewServer creates a new HTTP server for the object storage gateway
func NewServer(indexFinder IndexFinder) http.Handler {
	r := gin.Default()
	addRoutes(
		r,
		indexFinder,
	)
	var handler http.Handler = r
	return handler
}

func addRoutes(
	mux *gin.Engine,
	indexFinder IndexFinder,
) {
	mux.GET("/index/:number", findIndex(indexFinder))
}

func findIndex(indexFinder IndexFinder) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("number")
		value, err := strconv.Atoi(id)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid path parameter. must be a number"})
			return
		}

		if value < minValidNumber || value > maxValidNumber {
			c.JSON(http.StatusBadRequest, gin.H{"error": "number must be between 0 and 10000000"})
			return
		}

		index, found := indexFinder.FindIndex(value)
		if !found {
			c.JSON(http.StatusNotFound, gin.H{"error": "valid index not found"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"index": index})
		return
	}
}
