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
	assert.Equal(1, 1, "todoを作成する")
}

/*

describe("todoを作成する", () => {
   const email = "test@example.com";
   const password = "123456";

   it("Login by firebase authentication", () => {
      console.log(email);
   });
});

*/
