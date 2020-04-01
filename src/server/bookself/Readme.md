# go-nic.gin を使用してRESTFulなAPIサーバーを開発する
##　作成1:図書の記事を参照、作成、更新、削除を行う
+ Request: POST http://localhost/v1/articl/
+ Response: {id: "id", name: "username"}

+ Request: GET http://localohst/v1/articl/:id
+ Response: {id: "id", articl: "articl"}

+ Request: PUT http://localhost/v1/articl/:id
+ Response: {id: "id", articl: "articl"}

+ Request: DELETE http://localhost/v1/articl/:id
+ Response: {mess: "Message"}