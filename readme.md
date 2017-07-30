# grpcSample
CRUD APIの作成

## フォルダ構成
```
Dockerfile
docker-compose.yml
main.go             #サーバ側のプログラム
api/
  api.pb.go         #gRPCを用いて作成したインターフェース
  api.proto
client/
  test-client.go    #クライアント側のプログラム
```

## 環境
- Go(1.8)
- Redis
- Docker(17.03.1-ce-mac12)

## 仕様
### サーバ
#### 起動
```bash
docker-compose build
docker-compose up
```

### クライアント
#### 使用方法
```bash
Usage of ./test-client:
  -action string
    	[list|get|add|update|delete]を指定
  -id string
    	[get|update|delete]を実行時にidを指定
  -limit int
    	[list]を実行時にlimit指定 (default 5)
  -page int
    	[list]を実行時にpage指定 (default 1)
  -port string
    	server address host:post (default ":3000")
```

### 仕様
[action]によりリクエスト内容を変化
#### list
- 引数としてpage, limitを受け取り, ListItemRequestを送る
- 受け取ったアイテムリストを表示

#### get
- 引数としてidを受け取り, GetItemRequestを送る

#### add
- ランダムに文字列と数値を生成し, Itemを送る

#### update
- ランダムに文字列と数値を生成し, UpdateItemRequestを送る

#### delete
- 引数としてidを受け取り, DeleteItemRequestを送る
