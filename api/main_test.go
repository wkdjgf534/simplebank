package api

import (
	"os"
	"testing"

	"github.com/gin-gonic/gin"
)

// Add "github.com/lib/pq" before execituing this test
func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode) // reduce log's amounts during tests
	os.Exit(m.Run())
}
