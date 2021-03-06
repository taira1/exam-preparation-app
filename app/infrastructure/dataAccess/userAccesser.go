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
	rows, err := a.DBAgent.Conn.Query(fmt.Sprintf("SELECT * FROM user WHERE subject_id = %d;", subjectID))
	if err != nil {
		log.Printf(failedToGetData.value, err)
		return nil
	}
	defer rows.Close()
	var userResult []model.User
	for rows.Next() {
		user := model.User{}
		if err := rows.Scan(&user.ID, &user.Name); err != nil {
			log.Printf(failedToGetData.value, err)
		}
		user.Education = a.SubjectAccesser.FindByID(subjectID)
		userResult = append(userResult, user)
	}
	return userResult
}

// FindByID 指定したIDのユーザを取得します。
func (a *UserAccesser) FindByID(ID int) *model.User {
	rows, err := a.DBAgent.Conn.Query(fmt.Sprintf("SELECT * FROM user WHERE id = %d;", ID))
	if err != nil {
		log.Printf(failedToGetData.value, err)
		return nil
	}
	defer rows.Close()
	user := model.User{}
	var subjectID int
	for rows.Next() {
		if err := rows.Scan(&user.ID, &user.Name, &user.Comment, &subjectID); err != nil {
			log.Printf(failedToGetData.value, err)
		}
		user.Education = a.SubjectAccesser.FindByID(subjectID)
	}
	return &user
}

// Insert 引数で渡したuserをDBに登録,自動採番されたIDを返す。
func (a *UserAccesser) Insert(u *model.User) int {
	ins, err := a.DBAgent.Conn.Prepare("INSERT INTO user(name,comment,education_id) VALUES(?,?,?)")
	if err != nil {
		log.Printf(failedToInsertData.value, err)
		err = nil
		return -1
	}
	if _, e := ins.Exec(u.Name, u.Comment, u.Education.ID); e != nil {
		log.Printf(failedToInsertData.value, err)
		return -1
	}
	return GetAutoNumberedID(a.DBAgent)
}

// Update 引数で渡したuserをDBに登録。成功したらtrue失敗したらfalseを返す
func (a *UserAccesser) Update(u *model.User) bool {
	_, err := a.DBAgent.Conn.Exec("UPDATE user SET name = ?, comment = ? WHERE id = ?;", u.Name, u.Comment, u.ID)
	if err != nil {
		log.Printf(failedToUpdateData.value, err)
		return false
	}
	return true
}
