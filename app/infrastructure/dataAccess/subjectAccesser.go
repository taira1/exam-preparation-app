package dataAccess

import (
	"exam-preparation-app/app/domain/model"
	"fmt"
	"log"
)

// SubjectAccesser subjectテーブルへのアクセッサーです。
type SubjectAccesser struct {
	DBAgent         *DBAgent
	FacultyAccesser *FacultyAccesser
}

// FindByFacultyID 指定したFacultyIDの学科一覧を取得します。
func (a *SubjectAccesser) FindByFacultyID(facultyID int) []model.Subject {
	rows, err := a.DBAgent.conn.Query(fmt.Sprintf("SELECT * FROM subject WHERE faculty_id = %d;", facultyID))
	if err != nil {
		log.Println("データの取得に失敗しました。")
		return nil
	}
	defer rows.Close()
	var subjectResult []model.Subject
	for rows.Next() {
		subject := model.Subject{}
		if err := rows.Scan(&subject.ID, &subject.Name); err != nil {
			log.Println("クエリの発行に失敗しました。")
		}
		*subject.Faculty = *a.FacultyAccesser.FindByID(facultyID)
		subjectResult = append(subjectResult, subject)
	}
	return subjectResult
}

// FindByID 指定したIDの学科を取得します。
func (a *SubjectAccesser) FindByID(ID int) *model.Subject {
	rows, err := a.DBAgent.conn.Query(fmt.Sprintf("SELECT * FROM subject WHERE id = %d;", ID))
	if err != nil {
		log.Println("データの取得に失敗しました。")
		return nil
	}
	defer rows.Close()
	var facultyID int
	subject := model.Subject{}

	for rows.Next() {
		if err := rows.Scan(&subject.ID, &subject.Name, facultyID); err != nil {
			log.Println("クエリの発行に失敗しました。")
		}
		*subject.Faculty = *a.FacultyAccesser.FindByID(facultyID)
	}
	return &subject
}
