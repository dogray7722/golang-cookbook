package app

import (
	"os"
	"testing"

	"github.com/gin-gonic/gin"
)

// TestMain initializes the api mock tests
func TestMain(m *testing.M) {
	gin.SetMode((gin.TestMode))

	os.Exit(m.Run())
}