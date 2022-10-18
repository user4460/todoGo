package auth

//auth関数を作成する
import (
	"os"
	"bufio"
	"log"
	"fmt"
)

func Signup() {
	// ファイルを開く
	file, err := os.OpenFile("users.txt", os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// ファイルに書き込む
	writer := bufio.NewWriter(file)
	fmt.Print("ユーザー名を入力してください: ")
	var username string
	fmt.Scan(&username)
	fmt.Print("パスワードを入力してください: ")
	var password string
	fmt.Scan(&password)
	fmt.Fprintln(writer, username + "," + password)
	writer.Flush()
}
