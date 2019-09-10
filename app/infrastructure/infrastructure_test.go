package infrastructure

import "testing"

func TestInfrastructureInstanceWasGenerated(t *testing.T) {
	i := newInfrastructure()
	if i == nil {
		t.Fatalf("InfrastructureOBJの作成に失敗しました")
	}
	defer i.DBAgent.Conn.Close()
}

func TestUniversityFindAll(t *testing.T) {
	i := newInfrastructure()
	result := i.UniversityAccesser.FindAll()
	if result == nil {
		t.Fatal("クエリに問題がある可能性があります。")
	}
}

func TestUniversityFindByID(t *testing.T) {
	i := newInfrastructure()
	result := i.UniversityAccesser.FindByID(1)
	if result == nil {
		t.Fatal("クエリに問題がある可能性があります。")
	}
}

func TestFacultyFindAll(t *testing.T) {
	i := newInfrastructure()
	result := i.FacultyAccesser.FindAll()
	if result == nil {
		t.Fatal("クエリに問題がある可能性があります。")
	}
}

func TestFacultyFindByUniversityID(t *testing.T) {
	i := newInfrastructure()
	result := i.FacultyAccesser.FindByUniversityID(1)
	if result == nil {
		t.Fatal("クエリに問題がある可能性があります。")
	}
}

func TestFacultyFindByID(t *testing.T) {
	i := newInfrastructure()
	result := i.FacultyAccesser.FindByID(1)
	if result == nil {
		t.Fatal("クエリに問題がある可能性があります。")
	}
}

func TestSubjectFindAll(t *testing.T) {
	i := newInfrastructure()
	result := i.SubjectAccesser.FindAll()
	if result == nil {
		t.Fatal("クエリに問題がある可能性があります。")
	}
}

func TestSubjectFindByFacultyID(t *testing.T) {
	i := newInfrastructure()
	result := i.SubjectAccesser.FindByFacultyID(1)
	if result == nil {
		t.Fatal("クエリに問題がある可能性があります。")
	}
}

func TestSubjectFindByID(t *testing.T) {
	i := newInfrastructure()
	result := i.SubjectAccesser.FindByID(1)
	if result == nil {
		t.Fatal("クエリに問題がある可能性があります。")
	}
}
