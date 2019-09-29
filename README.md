# プロジェクト

学生向け過去問共有サイトZuru

Zuruは過去問の掲載を主軸足した大学の定期試験対策のための総合サイトです。

アカウント作成時に大学・学部・学科メールアドレスを登録。

ログイン後はテスト対策用の記事を投稿でき、自分が所属する大学・学部・学科の他の生徒が書いた記事を閲覧できる。

それぞれの大学・学部・学科毎にwebsocketを使ったチャットルームがあり、テスト前日などに複数人で情報共有する場として利用できる。(実装予定)


# デモ

![demo](https://raw.github.com/wiki/yuki-inoue-eng/exam-preparation--app/images/exam-preparation-app.gif)

# 使用技術

サーバーサイド:Go

データベース:MySQL

# 機能説明

・アカウント作成機能

・ログイン・ログアウト機能

・プロフィール登録機能

・記事投稿、編集、削除機能

・記事一覧表示機能

・記事詳細表示機能

・websoketを使用したチャット機能(未実装)


# Requirement

下記の定義に基づいたDBテーブルが必要です。

データベース名  
mysql> select database();    
| database()       |
|---|
| exam_preparation |


大学テーブル  
mysql> describe university;    
| Field | Type         | Null | Key | Default | Extra          |
|---|---|---|---|---|---|
| id    | int(11)      | NO   | PRI | NULL    | auto_increment |
| name  | varchar(100) | NO   |     | NULL    |                |


学部テーブル  
mysql> describe faculty;   
| Field         | Type         | Null | Key | Default | Extra          |
|---|---|---|---|---|---|
| id            | int(11)      | NO   | PRI | NULL    | auto_increment |
| name          | varchar(100) | NO   |     | NULL    |                |
| university_id | int(11)      | NO   |     | NULL    |                |


学科テーブル  
mysql> describe subject;  
| Field      | Type         | Null | Key | Default | Extra          |
|---|---|---|---|---|---|
| id         | int(11)      | NO   | PRI | NULL    | auto_increment |
| name       | varchar(100) | NO   |     | NULL    |                |
| faculty_id | int(11)      | NO   |     | NULL    |                |


ユーザテーブル  
mysql> describe user;  
| Field        | Type         | Null | Key | Default | Extra          |
|---|---|---|---|---|---|
| id           | int(11)      | NO   | PRI | NULL    | auto_increment |
| name         | varchar(100) | NO   |     | NULL    |                |
| comment      | varchar(200) | NO   |     | NULL    |                |
| education_id | int(11)      | NO   |     | NULL    |                |


ユーザ認証情報テーブル  
mysql> describe auth;  
| Field    | Type         | Null | Key | Default | Extra          |
|---|---|---|---|---|---|
| id       | int(11)      | NO   | PRI | NULL    | auto_increment |
| email    | varchar(100) | NO   | UNI | NULL    |                |
| password | varchar(60)  | NO   |     | NULL    |                |
| user_id  | int(11)      | NO   |     | NULL    |                |


記事テーブル  
mysql> describe article;  
| Field      | Type        | Null | Key | Default           | Extra                                         |
|---|---|---|---|---|---|
| id         | int(11)     | NO   | PRI | NULL              | auto_increment                                |
| user_id    | int(11)     | NO   |     | NULL              |                                               |
| lastupdate | datetime    | NO   |     | CURRENT_TIMESTAMP | DEFAULT_GENERATED on update CURRENT_TIMESTAMP |
| title      | text        | YES  |     | NULL              |                                               |
| class      | text        | YES  |     | NULL              |                                               |
| teacher    | text        | YES  |     | NULL              |                                               |
| content    | longtext    | YES  |     | NULL              |                                               |
| status     | varchar(20) | YES  |     | public            |                                               |


※ 上記のテーブル定義を行なった上で、適切に動作しない場合は下記ファイルの関数getDBConnection()内でDBコネクションを取得する際のURLが間違っている可能性があります。適切なURLに修正後コンパイルし直してください。
app/infrastructure/dataAccess/dbAgent.go


※ DBにはMySQLを使用しています。他のDBを使用する場合はソースコードに変更を加える必要があります。
app/infrastructure/dataAccess/dbAgent.go でMySQLドライバをインポートしている文を削除し、適切なドライバをインポートしてください。


※ アカウント作成時に必要な所属大学情報(大学・学部・学科)はマスターデータです。事前にuniversity、faculty、subjectの各テーブルにサービスの対象となる大学のレコードを作成しておく必要があります。(デモでは北里大学のデータを使用)


# 作成者情報

* yuki.inoue
* E-mail yuki.inoue.eng@gmail.com