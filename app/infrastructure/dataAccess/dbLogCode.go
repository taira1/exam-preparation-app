package dataAccess

type dbLogCode struct{ value string }

var failedToGetData = dbLogCode{"データの取得に失敗しました: %#v\n"}
var failedToUpdateData = dbLogCode{"データの更新に失敗しました: %#v\n"}
var failedToDeleteData = dbLogCode{"データの削除に失敗しました: %#v\n"}
var failedToInsertData = dbLogCode{"データの登録に失敗しました: %#v\n"}
var failedToGenerateDbAgent = dbLogCode{"DbAgentを生成できませんでした %#v\n"}
var generatedDbAgent = dbLogCode{"DbAgentを生成しました"}

func (c dbLogCode) string() string {
	if c.value == "" {
		return "未定義"
	}
	return c.value
}
