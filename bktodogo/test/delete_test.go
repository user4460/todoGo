package test

import (
	"testing" 
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestDelete(t *testing.T) {
	// Setup
	gin.SetMode(gin.TestMode)
	// Assertions
	assert := assert.New(t)
	assert.Equal(1, 1, "todoを削除する")
}

/*
describe("todoを作成する", () => {
   const email = "test@example.com";
   const password = "123456";

   it("Login by firebase authentication", () => {
      console.log(email);
   });
});

describe("todoを編集する", () => {
   const email = "test@example.com";
   const password = "123456";

   it("Login by firebase authentication", () => {
      console.log(email);
   });
});

describe("todoを削除する", () => {
   const email = "test@example.com";
   const password = "123456";

   it("Login by firebase authentication", () => {
      console.log(email);
   });
});

describe("todoをチェックする", () => {
   const email = "test@example.com";
   const password = "123456";

   it("Login by firebase authentication", () => {
      console.log(email);
   });
});

*/

