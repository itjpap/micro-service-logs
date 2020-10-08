package mgodb

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"logs"
)

//type mgo struct {
//	uri string //数据库网络地址
//	database string //要连接的数据库
//	collection string //要连接的集合
//}

//type Collection struct {
//	Database *Database
//	Name string // "collection"
//	FullName string  // "db.collection"
//}

var (
	mgoSession *mgo.Session
	dataBase = "myDb"
	collection = "crm-log"
)

//type result struct {
//	level string
//	msg string
//	time string
//}

const URL = "127.0.0.1:27017"

func GetSession(){
	var result interface{}
	session,err := mgo.Dial(URL)
	logs.Record(err)
	c := session.DB(dataBase).C(collection)
	err2 := c.Find(bson.M{"_id":1}).One(&result)
	logs.Record(err2)
	fmt.Println(result)
}

//func (m *mgo)Connect() *mongo.Clooection{
//	ctx,cancel := context.WithTimeout(context.Background(),10*time.Second)
//	defer cancel()
//	client,err := mongo.Contect(ctx,options.Client().AppluURI(m.uri))
//	if err != nil{
//		log.Print(err)
//	}
//	collection := client.Database(m.database).Collection(m.collection)
//	return collection
//}
