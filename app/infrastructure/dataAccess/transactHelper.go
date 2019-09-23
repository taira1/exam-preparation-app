package dataAccess

import (
	"database/sql"
	"log"
)

// Transact トランザクションのラッパー関数です
func Transact(db *sql.DB, txFunc func(*sql.Tx) error) (err error) {
	tx, err := db.Begin()
	if err != nil {
		return
	}
	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p) // re-throw panic after Rollback
		} else if err != nil {
			tx.Rollback()
		} else {
			err = tx.Commit()
		}
	}()
	err = txFunc(tx)
	return err
}

// GetAutoNumberedID 自動採番されたidを取得します。
func GetAutoNumberedID(dbAgent *DBAgent) int {
	AutoNumberedID := -1
	rows, err := dbAgent.Conn.Query("SELECT LAST_INSERT_ID()")
	if err != nil {
		log.Printf(failedToGetData.value, err)
		return AutoNumberedID
	}
	defer rows.Close()
	for rows.Next() {
		if err := rows.Scan(&AutoNumberedID); err != nil {
			log.Printf(failedToGetData.value, err)
		}
	}
	return AutoNumberedID
}
