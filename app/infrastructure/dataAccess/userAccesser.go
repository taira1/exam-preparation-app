package dataAccess

import (
	"exam-preparation-app/app/domain/model"
	"fmt"
	"log"
)

// UserAccesser userテーブルへのアクセッサーです。
type UserAccesser struct {
	DBAgent         *DBAgent
	SubjectAccesser *SubjectAccesser
}

// FindBySubjectID 指定したSubjectIDのユーザ一覧を取得します。
func (a *UserAccesser) FindBySubjectID(subjectID int) []model.User {
	rows, err := a.DBAgent.conn.Query(fmt.Sprintf("SELECT * FROM user WHERE subject_id = %d;", subjectID))
	if err != nil {
		log.Println("データの取得に失敗しました。")
		return nil
	}
	defer rows.Close()
	var userResult []model.User
	for rows.Next() {
		user := model.User{}
		if err := rows.Scan(&user.ID, &user.Name); err != nil {
			log.Println("クエリの発行に失敗しました。")
		}
		*user.Education = *a.SubjectAccesser.FindByID(subjectID)
		userResult = append(userResult, user)
	}
	return userResult
}

// FindByID 指定したIDのユーザを取得します。
func (a *UserAccesser) FindByID(ID int) *model.User {
	rows, err := a.DBAgent.conn.Query(fmt.Sprintf("SELECT * FROM user WHERE id = %d;", ID))
	if err != nil {
		log.Println("データの取得に失敗しました。")
		return nil
	}
	defer rows.Close()
	user := model.User{}
	var subjectID int
	for rows.Next() {
		if err := rows.Scan(&user.ID, &user.Name, &subjectID); err != nil {
			log.Println("クエリの発行に失敗しました。")
		}
		*user.Education = *a.SubjectAccesser.FindByID(subjectID)
	}
	return &user
}
