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
	rows, err := a.DBAgent.Conn.Query("SELECT * FROM faculty")
	if err != nil {
		log.Printf(failedToGetData.value, err)
		return nil
	}
	defer rows.Close()
	var facultiesResult []model.Faculty
	var universityID int
	for rows.Next() {
		faculty := model.Faculty{}
		if err := rows.Scan(&faculty.ID, &faculty.Name, &universityID); err != nil {
			log.Printf(failedToGetData.value, err)
		}
		faculty.University = a.UniversityAccesser.FindByID(universityID)
		facultiesResult = append(facultiesResult, faculty)
	}
	return facultiesResult
}

// FindByUniversityID 指定したuniversityIDの学部一覧を取得します。
func (a *FacultyAccesser) FindByUniversityID(universityID int) []model.Faculty {
	rows, err := a.DBAgent.Conn.Query(fmt.Sprintf("SELECT * FROM faculty WHERE university_id = %d;", universityID))
	if err != nil {
		log.Printf(failedToGetData.value, err)
		return nil
	}
	defer rows.Close()
	var facultiesResult []model.Faculty
	var trash int
	for rows.Next() {
		faculty := model.Faculty{}
		if err := rows.Scan(&faculty.ID, &faculty.Name, &trash); err != nil {
			log.Printf(failedToGetData.value, err)
		}
		faculty.University = a.UniversityAccesser.FindByID(universityID)
		facultiesResult = append(facultiesResult, faculty)
	}
	return facultiesResult
}

// FindByID 指定したIDの学部を取得します。
func (a *FacultyAccesser) FindByID(ID int) *model.Faculty {
	rows, err := a.DBAgent.Conn.Query(fmt.Sprintf("SELECT * FROM faculty WHERE id = %d;", ID))
	if err != nil {
		log.Printf(failedToGetData.value, err)
		return nil
	}
	defer rows.Close()
	var universityID int
	faculty := model.Faculty{}
	for rows.Next() {
		if err := rows.Scan(&faculty.ID, &faculty.Name, &universityID); err != nil {
			log.Printf(failedToGetData.value, err)
		}
		faculty.University = a.UniversityAccesser.FindByID(universityID)
	}
	return &faculty
}
