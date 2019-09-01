package infrastructure

import "exam-preparation-app/app/infrastructure/dataAccess"

// Infrastructure インフラ
type Infrastructure struct {
	UniversityAccesser *dataAccess.UniversityAccesser
	FacultyAccesser    *dataAccess.FacultyAccesser
	SubjectAccesser    *dataAccess.SubjectAccesser
	UserAccesser       *dataAccess.UserAccesser
	AuthAccesser       *dataAccess.AuthAccesser
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

func newAuthAccesser(dbAgent *dataAccess.DBAgent) *dataAccess.AuthAccesser {
	return &dataAccess.AuthAccesser{DBAgent: dbAgent}
}

// NewInfrastructure コンストラクタです。
func NewInfrastructure() *Infrastructure {
	dbAgent := dataAccess.NewDbAgent()
	universityAccesser := newUniversityAccesser(dbAgent)
	facultyAccesser := newFacultyAccesser(dbAgent, universityAccesser)
	subjectAccesser := newSubjectAccesser(dbAgent, facultyAccesser)
	userAccesser := newUserAccesser(dbAgent, subjectAccesser)
	authAccesser := newAuthAccesser(dbAgent)

	return &Infrastructure{
		UniversityAccesser: universityAccesser,
		FacultyAccesser:    facultyAccesser,
		SubjectAccesser:    subjectAccesser,
		UserAccesser:       userAccesser,
		AuthAccesser:       authAccesser,
	}
}
