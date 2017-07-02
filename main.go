// main.go
package main

import (
    "log"
    "github.com/garyburd/redigo/redis"
    "golang.org/x/net/context"
    "google.golang.org/grpc"
    pb "github.com/sorashido/mercari/api"
    "net"
    "fmt"
    "strings"
    "strconv"
)

const (
    httpport = ":3000"//localhost
    dbport = "redis:6379"//docker上
)

type server struct {
    db  redis.Conn
}

func NewServer() *server {
    return &server{
    }
}

//Itemをまとめて返す(pagnition機能)
func (s *server) ListItem(ctx context.Context, m *pb.ListItemRequest) (*pb.ListItemResponse, error){
    //データベースからリストを持ってきて返す
    var listnums []string
    for i :=(m.Page-1)*m.Limit+1; i <= m.Page*m.Limit; i++{
        listnums = append(listnums,strconv.Itoa(int(i)))
    }
    //(idの最大値の取得)
    dataNum, _ := redis.Int(s.db.Do("GET", "dataNum"))
    var items []*pb.Item
    var count int = 0
    for i := 0; i< dataNum; i+=1{
        rep, err := redis.String(s.db.Do("GET", i))
        if err == nil{
            count+=1
            if(contains(listnums,strconv.Itoa(count))){
                item := stringToItem(rep, 0)
                items = append(items, item)
            }
        }
    }

    list := pb.ListItemResponse{
        Items:items,
    }
    fmt.Println("list:",items)
    return &list, nil
}
func contains(s []string, e string) bool {
    for _, a := range s {
        if a == e {
            return true
        }
    }
    return false
}

//idと一致するitemを取得(参照回数カウント機能付き)
func (s *server) GetItem(ctx context.Context, m *pb.GetItemRequest) (*pb.Item, error){
    rep, err := redis.String(s.db.Do("GET", m.Id))
    if err != nil{
        fmt.Println(err)//
        return &pb.Item{}, err
    }

    //stringをitemに変換
    item := stringToItem(rep, 1)

    //データベースの更新
    _, err = s.db.Do("GETSET", m.Id, item.String())
    if err != nil{
        fmt.Println(err)
        return item, err
    }
    fmt.Println("get:", item.String())
    return item,nil
}

//stringをitemに変換, nはpvに加える数(0, 1)
func stringToItem(str string, n int32)(*pb.Item){
    result := make(map[string]string)
    strings.Replace(str, "\"", "", -1)
    result["id"] = extStr(&str, "id:","\" ")
    result["name"] = extStr(&str, "name:","\" ")
    result["title"] = extStr(&str, "title:","\" ")
    result["description"] = extStr(&str, "description:","\" ")
    price := extInt(&str, "price:"," ")
    pv := extInt(&str, "pv:"," ")
    status := extBool(&str, "status:")

    item := pb.Item{
        Id:result["id"],
        Name:result["name"],
        Title:result["title"],
        Description:result["description"],
        Price:price,
        Pv:pv+n,
        Status:status,
    }
    return &item
}

//抽出(抽出された文字列は削除)
//文字列中のnameからjudgeまでを抽出 int型
func extStr(str *string, name string, judge string) string{
    if strings.Contains(*str,name){
        f := strings.Index(*str,name)+len(name)+1
        t := strings.Index(*str, judge)
        ans := (*str)[f:t]
        *str = (*str)[t+2:]
        return ans
    }
    return ""
}
//抽出(抽出された文字列は削除) int型
func extInt(str *string, name string, judge string) int32{
    if strings.Contains(*str,name){
        f := strings.Index(*str,name)+len(name)
        t := strings.Index(*str, judge)
        ans := (*str)[f:t]
        *str = (*str)[t+1:]

        pr64, err := strconv.ParseInt(ans, 10, 32)
        if err != nil{pr64 = 0}
        x := int32(pr64)
        return x
    }
    return 0
}
//文字列中にtrueがあればreturn
func extBool(str *string, name string) bool{
    if strings.Contains(*str,name){
        if strings.Contains(*str, "true"){
            return true
        }
    }
    return false
}

//itemをデータベースに追加
func (s *server) AddItem(ctx context.Context, m *pb.Item) (*pb.AddItemResponse, error){
    //idの取得(1から+1ずつ)
    s.db.Do("INCR", "dataNum")
    count, _ := redis.String(s.db.Do("GET", "dataNum"))

    m.Id = count
    _, err := s.db.Do("SET",m.Id, m.String())
    if err != nil{
        return &pb.AddItemResponse{}, err
    }
    fmt.Println("add:",m.String())
    itemResponce := pb.AddItemResponse{
        Item:m,
    }
    return &itemResponce, nil
}

//idと一致するデータをupdate
//idがない場合, エラー
func (s *server) UpdateItem(ctx context.Context, m *pb.UpdateItemRequest) (*pb.UpdateItemResponse, error){
    //itemの更新
    item := pb.Item{
        Id:m.Item.Id,
        Name:m.Item.Name,
        Title:m.Item.Title,
        Description:m.Item.Description,
        Price:m.Item.Price,
        Pv:0,
        Status:m.Item.Status,
    }
    _, err := s.db.Do("GETSET", item.Id, item.String())
    if err != nil{
        return &pb.UpdateItemResponse{}, err
    }
    updateResponce := pb.UpdateItemResponse{
        Item:m.Item,
    }
    fmt.Println("update:",item.String())
    return &updateResponce, nil
}

//idと一致するitemをdelete
func (s *server) DeleteItem(ctx context.Context, m *pb.DeleteItemRequest) (*pb.DeleteItemResponse, error){
    item, _ := redis.String(s.db.Do("GET", m.Id))
    fmt.Println("delete:", item)

    _, err := s.db.Do("DEL", m.Id)
    if err != nil{
        fmt.Println(err)//
        return &pb.DeleteItemResponse{}, err
    }

    deleteResponse := pb.DeleteItemResponse{}
    return &deleteResponse, nil
}

func main() {
    //requestを受け付けるポート指定
    lis, err := net.Listen("tcp", httpport)
    if err != nil {
        log.Fatalf("failed to listen: %v", err)
    }

    //grpcサーバー(DB:redis)のインスタンス作成
    s := grpc.NewServer()

    //DBにredisを登録
    ns := NewServer()
    c, err := redis.Dial("tcp", dbport)
    if err != nil {
        log.Fatal(err)
    }
    defer c.Close()
    ns.db = c

    pb.RegisterAPIServer(s, ns)
    if err := s.Serve(lis); err != nil {
        log.Fatalf("failed to serve: %v", err)
    }
    defer s.Stop()
}