package auth

import (
	"http"
	"store"
)
//authのマイページ
func MyPage(w http.ResponseWriter, r *http.Request) {	
	//セッションを取得
	session, err := store.Get(r, "session")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	//セッションにユーザーIDがあるか確認
	if session.Values["user_id"] == nil {
		http.Error(w, "ログインしてください", http.StatusUnauthorized)
		return
	}
	//ユーザーIDを取得
	userID := session.Values["user_id"].(string)
	//ユーザーIDからユーザー情報を取得
	user, err := model.GetUserByID(userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	//ユーザー情報を表示
	fmt.Fprintf(w, "ユーザー名: %s", user.Name)
}
