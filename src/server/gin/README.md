# github.com/gin-gonic/ginを使ってAPIサーバー作成
## 目標
+ json形式で返す
	+ curlコマンド curl -X POST localhost:8800/ を叩くとjsonを返してくれる
+ jsonをPOSTで送ったら、AccessTokenを作成、送信してくれる
	+ curlコマンド curl -X POST -H "Content-Type: application/json" -d '{"Name":"sensuikan1973", "Age":"100"}' localhost:8800/ を叩くとAccessTokenを返してくる
+ 登録機能作成
	+ ユーザー名、email、passwordを送るとユーザー登録をしてくれる
+ ログイン機能作成
	+ ユーザー名、passwprdを送るとログインしてくれる
