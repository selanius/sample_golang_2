## 概要
- Golangのスタートアップ用プロジェクト第2弾
- 超初学者向け

## prj詳細
基本的なwebアプリの基本
DB接続

## 前提
- host pcにgoをインストール
- 動作確認端末のバージョン  
> go version 1.10.3  
> go version 1.9

- host pcにgo go-sql-driverをインストール
`go get -u github.com/go-sql-driver/mysql`

- host pcにmysqlをインストールしてあること
参考 : `https://qiita.com/hkusu/items/cda3e8461e7a46ecf25d`
localhostのmysqlにはパスワードを設定していないこと
パスワード設定している場合は、db接続のクラスを修正してください

- あらかじめhost pcのmysqlにデータベース追加する
`create database gosample;`

## 起動方法
### for local
- sample_golang_second/run
ターミナルにてsample_golang/run配下でコマンド `go run main.go` で起動

## 検証方法
webブラウザーで下記をURLに入力  
`http://localhost:8080/index`

## 注意事項
- ログイン機能が未実装なので、CSRF対策が施されていないので、この知識のままwebアプリケーションを実装しないようにしてください
- 次章でログイン機能を実装し、様々な脆弱性の対策をしていきます