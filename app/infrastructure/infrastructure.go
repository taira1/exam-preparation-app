package infrastructure

import (
	"exam-preparation-app/app/infrastructure/dataAccess"
	"log"
)

// InfrastructureOBJ インフラのオブジェクトです。
var InfrastructureOBJ = newInfrastructure()

// Infrastructure インフラ
type Infrastructure struct {
	UniversityAccesser *dataAccess.UniversityAccesser
	FacultyAccesser    *dataAccess.FacultyAccesser
	SubjectAccesser    *dataAccess.SubjectAccesser
	UserAccesser       *dataAccess.UserAccesser
	ArticleAccesser    *dataAccess.ArticleAccesser
	AuthAccesser       *dataAccess.AuthAccesser
	DBAgent            *dataAccess.DBAgent
}

func newUniversityAccesser(dbAgent *dataAccess.DBAgent) *dataAccess.UniversityAccesser {
	return &dataAccess.UniversityAccesser{DBAgent: dbAgent}
}
func newFacultyAccesser(dbAgent *dataAccess.DBAgent, accesser *dataAccess.UniversityAccesser) *dataAccess.FacultyAccesser {
	return &dataAccess.FacultyAccesser{DBAgent: dbAgent, UniversityAccesser: accesser}
}
func newSubjectAccesser(dbAgent *dataAccess.DBAgent, accesser *dataAccess.FacultyAccesser) *dataAccess.SubjectAccesser {
	return &dataAccess.SubjectAccesser{DBAgent: dbAgent, FacultyAccesser: accesser}
}
func newUserAccesser(dbAgent *dataAccess.DBAgent, accesser *dataAccess.SubjectAccesser) *dataAccess.UserAccesser {
	return &dataAccess.UserAccesser{DBAgent: dbAgent, SubjectAccesser: accesser}
}
func newArticleAccesser(dbAgent *dataAccess.DBAgent) *dataAccess.ArticleAccesser {
	return &dataAccess.ArticleAccesser{DBAgent: dbAgent}
}
func newAuthAccesser(dbAgent *dataAccess.DBAgent) *dataAccess.AuthAccesser {
	return &dataAccess.AuthAccesser{DBAgent: dbAgent}
}

// コンストラクタです。
func newInfrastructure() *Infrastructure {
	dbAgent := dataAccess.NewDbAgent()
	universityAccesser := newUniversityAccesser(dbAgent)
	facultyAccesser := newFacultyAccesser(dbAgent, universityAccesser)
	subjectAccesser := newSubjectAccesser(dbAgent, facultyAccesser)
	userAccesser := newUserAccesser(dbAgent, subjectAccesser)
	articleAccesser := newArticleAccesser(dbAgent)
	authAccesser := newAuthAccesser(dbAgent)

	return &Infrastructure{
		DBAgent:            dbAgent,
		UniversityAccesser: universityAccesser,
		FacultyAccesser:    facultyAccesser,
		SubjectAccesser:    subjectAccesser,
		UserAccesser:       userAccesser,
		ArticleAccesser:    articleAccesser,
		AuthAccesser:       authAccesser,
	}
}

// Close dbコネクションをクローズします。
func (i *Infrastructure) Close() {
	err := i.DBAgent.Conn.Close()
	if err != nil {
		log.Fatalf("DBコネクションを正常に終了できませんでした。:%#v", err)
	}
}
