package dataAccess

import (
	"database/sql"
	"log"

	//mysplのドライバを使用
	_ "github.com/go-sql-driver/mysql"
)

// DBAgent DBへのコネクションを保持します。
type DBAgent struct {
	Conn *sql.DB
}

// NewDbAgent コンストラクタです
func NewDbAgent() *DBAgent {
	c, err := getDBConnection()
	if err == nil {
		log.Println("DbAgentを生成しました")
		return &DBAgent{Conn: c}
	}
	log.Fatalf("DbAgentを生成できませんでした %#v", err)
	return nil
}

func getDBConnection() (*sql.DB, error) {
	return sql.Open("mysql", "root:@/exam_preparation")
	//TODO:データベースURLのハードコードをどうにかする。
}
