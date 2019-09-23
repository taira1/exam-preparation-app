package dataAccess

import (
	"exam-preparation-app/app/domain/model"
	"fmt"
	"log"
)

// UniversityAccesser universityテーブルへのアクセッサーです。
type UniversityAccesser struct {
	DBAgent *DBAgent
}

// FindAll すべての大学を取得します。
func (a *UniversityAccesser) FindAll() []model.University {
	rows, err := a.DBAgent.Conn.Query("SELECT * FROM university;")
	if err != nil {
		log.Printf(failedToGetData.value, err)
		return nil
	}
	defer rows.Close()
	var universitiesResult []model.University
	for rows.Next() {
		university := model.University{}
		if err := rows.Scan(&university.ID, &university.Name); err != nil {
			log.Printf(failedToGetData.value, err)
		}
		universitiesResult = append(universitiesResult, university)
	}
	return universitiesResult
}

// FindByID 指定したIDの大学を取得します。
func (a *UniversityAccesser) FindByID(ID int) *model.University {
	rows, err := a.DBAgent.Conn.Query(fmt.Sprintf("SELECT * FROM university WHERE id = %d;", ID))
	if err != nil {
		log.Printf(failedToGetData.value, err)
		return nil
	}
	defer rows.Close()
	university := model.University{}
	for rows.Next() {
		if err := rows.Scan(&university.ID, &university.Name); err != nil {
			log.Printf(failedToGetData.value, err)
		}
	}
	return &university

}
