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
		log.Printf(failedToGetData.value, err)
		return nil
	}
	defer rows.Close()
	auth := model.Auth{}
	for rows.Next() {
		if err := rows.Scan(&auth.ID, &auth.Email, &auth.Password, &auth.UserID); err != nil {
			log.Printf(failedToGetData.value, err)
			return nil
		}
	}
	if auth.ID == 0 {
		log.Println("指定したemailは未登録です")
		return nil
	}
	return &auth
}

// Insert 引数で受け取ったauth,Userをインサートします。
func (a *AuthAccesser) Insert(auth *model.Auth, userID int) {
	ins, err := a.DBAgent.Conn.Prepare("INSERT INTO auth(email,password,user_id) VALUES(?,?,?)")
	if err != nil {
		log.Printf(failedToInsertData.value, err)
		err = nil
		return
	}
	if _, e := ins.Exec(auth.Email, auth.Password, userID); e != nil {
		log.Printf(failedToInsertData.value, err)
	}
}
