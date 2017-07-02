package main

import(
	pb "github.com/sorashido/mercari/api"
	"log"
	"google.golang.org/grpc"
	"flag"
	"context"
	"math/rand"
	"time"
	"strings"
)

var (
	port = flag.String("port", ":3000", "server address host:post")
	action = flag.String("action", "", "[list|get|add|update|delete]を指定")
	id = flag.String("id", "", "[get|update|delete]を実行時にidを指定")
	page = flag.Int("page", 1, "[list]を実行時にpage指定")
	limit = flag.Int("limit", 5, "[list]を実行時にlimit指定")
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ             "
func randomString(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return strings.TrimSpace(string(b))
}

func listRequest(c *pb.APIClient, p int32, l int32){
	resp, err := (*c).ListItem(context.Background(), &pb.ListItemRequest{
		Page:p,
		Limit:l,
	})

	if err != nil {
		log.Println("cannot get item")
		log.Fatalf("%v", err)
	}
	log.Print("List:")
	for i := range resp.Items{
		log.Println(resp.Items[i].String())
	}
}

func getRequest(c *pb.APIClient, id *string){
	resp, err := (*c).GetItem(context.Background(), &pb.GetItemRequest{
		Id:*id,
	})
	if err != nil {
		log.Println("cannot get item")
		log.Fatalf("%v", err)
	}
	log.Printf("GET: %s", resp.String())
}

func addRequest(c *pb.APIClient, item *pb.Item){
	//ランダムな文字列生成
	resp, err := (*c).AddItem(context.Background(), item)
	if err != nil {
		log.Println("cannot add item")
		log.Fatalf("%v", err)
	}
	log.Printf("ADD: %s", resp.Item.String())
}

func updateRequest(c *pb.APIClient, item *pb.Item){
	resp, err := (*c).UpdateItem(context.Background(), &pb.UpdateItemRequest{
		Item:item,
	})
	if err != nil {
		log.Println("cannot update item")
		log.Fatalf("%v", err)
	}
	log.Printf("UPDATE: %s", resp.Item.String())
}

func deleteRequest(c *pb.APIClient, id *string){
	_, err := (*c).DeleteItem(context.Background(), &pb.DeleteItemRequest{
		Id:*id,
	})
	if err != nil {
		log.Println("cannot delete item")
		log.Fatalf("%v", err)
	}
	log.Printf("DELETE: id %s", *id)
}

func main(){
	//引数
	flag.Parse()

	//IPアドレスとポート指定し, サーバーと接続
  conn, err := grpc.Dial(*port, grpc.WithInsecure())
  if err != nil {
		log.Fatalf("failed to listen: %v", err)
  }
  defer conn.Close()

	//grpcクライアントのインスタンス作成
  c := pb.NewAPIClient(conn)

	//ramdom(add, updateで使用)
	rand.Seed(time.Now().UnixNano())

	switch *action{
	case "list":
		listRequest(&c, int32(*page), int32(*limit))
	case "get":
		getRequest(&c, id)
	case "add":
		item := &pb.Item{
			Name:        randomString(6),
			Title:       randomString(40),
			Description: randomString(400),
			Price:       rand.Int31n(4000),
			Status:      true,
		}
		addRequest(&c, item)
	case "update":
		item := &pb.Item{
			Id:*id,
			Name:        randomString(6),
			Title:       randomString(40),
			Description: randomString(400),
			Price:       rand.Int31n(4000),
			Status:      true,
		}
		updateRequest(&c, item)
	case "delete":
		deleteRequest(&c, id)
	}
}