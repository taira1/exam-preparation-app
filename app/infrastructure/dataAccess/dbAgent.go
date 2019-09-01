package dataAccess

import (
	"database/sql"
	"log"
)

// DBAgent DBへのコネクションを保持します。
type DBAgent struct {
	conn *sql.DB
}

// NewDbAgent コンストラクタです
func NewDbAgent() *DBAgent {
	if c, err := getDBConnection(); err == nil {
		log.Println("DbAgentを生成しました")
		return &DBAgent{conn: c}
	}
	log.Println("DbAgentを生成できませんでした")
	return nil
}

func getDBConnection() (*sql.DB, error) {
	return sql.Open("mysql", "root:@/exam_preparation")
	//TODO:データベースのハードコードをどうにかする。
}