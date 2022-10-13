package main

import "github.com/gin-gonic/gin"

import "net/http"

//https://gin-gonic.com/docs/quickstart/のサンプルコード

func main() {
    //default()とは、デフォルトの設定でginのインスタンスを作成する関数
    engine:= gin.Default()
    //GETメソッドで"/"にアクセスがあった場合の処理
    engine.GET("/", func(c *gin.Context) {
        //jsonメソッドとhttp.StatusOKを返す
        c.JSON(http.StatusOK, gin.H{
            "message": "hello world",
        })
    })
    //runメソッドでポート番号を指定してサーバーを起動する
    engine.Run(":3000")

    //
   // engine.Post("/post", func(c *gin.Context) {
     //   c.JSON(http.StatusOK, gin.H{
       //     "message": "post",
        //})}

    //
    //engine.PUT("/put", func(c *gin.Context) {
    //    c.JSON(http.StatusOK, gin.H{
    //        "message": "put",
    //    })

    //engine.DELETE("/delete", func(c *gin.Context) {
    //    c.JSON(http.StatusOK, gin.H{
    //        "message": "delete",
    //    })
    //})

}

//export GIN_MODE=release