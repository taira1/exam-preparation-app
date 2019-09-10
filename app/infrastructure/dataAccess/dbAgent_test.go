package dataAccess

import "testing"

func TestAccessSuccess(t *testing.T) {
	agent := NewDbAgent()
	if agent == nil {
		t.Fatalf("DBコネクションの取得に失敗しました。")
	}
	defer agent.Conn.Close()
}
