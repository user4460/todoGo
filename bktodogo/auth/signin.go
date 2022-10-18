package auth
import(
	"net/http"
	"html/template"
	"log"
	"os"
	"bufio"
	"strings"
	"store"
)
//singin関数を作成する
func signin(w http.ResponseWriter, r *http.Request) {
	// ファイルを開く
	file, err := os.OpenFile("users.txt", os.O_RDONLY, 0)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// ファイルの内容を読み込む
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// ユーザー名とパスワードを取得する
		user := strings.Split(scanner.Text(), ",")
		username := user[0]
		password := user[1]

		// ユーザー名とパスワードが正しいか確認する
		if r.FormValue("username") == username && r.FormValue("password") == password {
			// セッションを作成する
			session, _ := store.Get(r, "session")
			session.Values["authenticated"] = true
			session.Save(r, w)

			// リダイレクトする
			http.Redirect(w, r, "/todos", 302)
			return
		}
	}

	// ログインに失敗した場合
	http.Redirect(w, r, "/signin", 302)
}