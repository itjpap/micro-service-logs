package main

import (
	MongoDto "dto"
	"fmt"
	log "github.com/sirupsen/logrus"
	"logs"
	"os"
)


func init(){
	//设置日志格是为json格式
	log.SetFormatter(&log.JSONFormatter{})

	//设置将日志输出到标准输出（默认输出为stderr，标准错误）
	//日志消息输出可以是任意的io.writer类型
	log.SetOutput(os.Stdout)

	//设置日志级别为warn以上
	log.SetLevel(log.WarnLevel)
}

func main()  {
	////设置日志输出到控制台
	//log.Out = os.Stdout
	//
	//fiel,err := os.OpenFile("logrus.log",os.O_CREATE|os.O)
	//mgodb.GetSession()
	logHandle()


}

func logHandle() interface{}{
	//fmt.Println("run log test")
	//jsonData := `{"Level":"warn","Source":"RPC","Code":"10001","Message":"this is rpc error msg"}`
	result,ok := MongoDto.InspectionFormat(jsonData)

	if !ok {
		fmt.Println(result)
		return result.(string)
	}
	fmt.Println("break point 10001 ")
	logs.Record(result.(map[string]interface{}))
	return nil
}



