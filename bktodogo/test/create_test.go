package test

import (
	"testing"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestCreate(t *testing.T) {
	// Setup
	gin.SetMode(gin.TestMode)
	// Assertions
	assert := assert.New(t)
	assert.Equal(1, 1, "they should be equal")
}
