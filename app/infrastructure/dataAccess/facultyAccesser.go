package dataAccess

import (
	"exam-preparation-app/app/domain/model"
	"fmt"
	"log"
)

// FacultyAccesser faculityテーブルへのアクセッサーです。
type FacultyAccesser struct {
	DBAgent            *DBAgent
	UniversityAccesser *UniversityAccesser
}

// FindAll 学部一覧を取得します。
func (a *FacultyAccesser) FindAll() []model.Faculty {
	rows, err := a.DBAgent.Conn.Query("SELECT * FROM faculity")
	if err != nil {
		log.Println("データの取得に失敗しました。")
		return nil
	}
	defer rows.Close()
	var facultiesResult []model.Faculty
	var universityID int
	for rows.Next() {
		faculty := model.Faculty{}
		if err := rows.Scan(&faculty.ID, &faculty.Name, &universityID); err != nil {
			log.Println("クエリの発行に失敗しました。")
		}
		*faculty.University = *a.UniversityAccesser.FindByID(universityID)
		facultiesResult = append(facultiesResult, faculty)
	}
	return facultiesResult
}

// FindByUniversityID 指定したuniversityIDの学部一覧を取得します。
func (a *FacultyAccesser) FindByUniversityID(universityID int) []model.Faculty {
	rows, err := a.DBAgent.Conn.Query(fmt.Sprintf("SELECT * FROM faculity WHERE university_id = %d;", universityID))
	if err != nil {
		log.Println("データの取得に失敗しました。")
		return nil
	}
	defer rows.Close()
	var facultiesResult []model.Faculty
	for rows.Next() {
		faculty := model.Faculty{}
		if err := rows.Scan(&faculty.ID, &faculty.Name); err != nil {
			log.Println("クエリの発行に失敗しました。")
		}
		*faculty.University = *a.UniversityAccesser.FindByID(universityID)
		facultiesResult = append(facultiesResult, faculty)
	}
	return facultiesResult
}

// FindByID 指定したIDの学部を取得します。
func (a *FacultyAccesser) FindByID(ID int) *model.Faculty {
	rows, err := a.DBAgent.Conn.Query(fmt.Sprintf("SELECT * FROM faculity WHERE id = %d;", ID))
	if err != nil {
		log.Println("データの取得に失敗しました。")
		return nil
	}
	defer rows.Close()
	var universityID int
	faculty := model.Faculty{}
	for rows.Next() {
		if err := rows.Scan(&faculty.ID, &faculty.Name, &universityID); err != nil {
			log.Println("クエリの発行に失敗しました。")
		}
		*faculty.University = *a.UniversityAccesser.FindByID(universityID)
	}
	return &faculty
}
