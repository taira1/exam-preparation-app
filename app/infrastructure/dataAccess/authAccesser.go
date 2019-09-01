package dataAccess

import (
	"exam-preparation-app/app/domain/model"
	"fmt"
	"log"
)

// AuthAccesser authテーブルへのアクセッサーです
type AuthAccesser struct {
	DBAgent *DBAgent
}

// FindByEmail 指定したemailのAuthを取得します。
func (a *AuthAccesser) FindByEmail(email string) *model.Auth {
	rows, err := a.DBAgent.Conn.Query(fmt.Sprintf("SELECT * FROM auth WHERE email = '%s';", email))
	if err != nil {
		log.Println("データの取得に失敗しました。")
		return nil
	}
	defer rows.Close()
	auth := model.Auth{}
	for rows.Next() {
		if err := rows.Scan(&auth.ID, &auth.Email, &auth.UserID); err != nil {
			log.Println("クエリの発行に失敗しました。")
		}
	}
	return &auth
}
