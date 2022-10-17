//httpコード　テスト
//https://gin-gonic.com/docs/testing/
package test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	 "github.com/gin-gonic/setupRouter"
)

func TestPingRoute(t *testing.T) {
	router := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/ping", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "pong", w.Body.String())
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