package MongoDto

import (
	"encoding/json"
	"fmt"
	"github.com/asaskevich/govalidator"
	"logs"
)

//type jsonField struct {
//	Name string ``
//	Email string `valid:"email"`
//}

type example struct {
}

type jsonField struct {
	Level  string `valid:"required"`     //日志等级
	Source string `valid:"required"`     //来源 RPC/restful api
	Code   string `valid:"required,int"` // 记录代码
	Msg    string `valid:"required"`     //日志消息内容
}

//type jsonField struct {
//	Level  string `valid:"required"`  //日志等级
//	Source string `valid:"required"` //来源 RPC/restful api
//	Code   string    `valid:"required,int"`   // 记录代码
//	Msg    string `valid:"required"`    //日志消息内容
//}

var jsonMap map[string]interface{}

var MapTemplate = map[string]interface{}{
	"Level":   "required~Miss:日志等级",
	"Source":  "required~Miss:日志来源",
	"Code":    "required~Miss:日志代码,int~日志代码必须是整形",
	"Message": "required~Miss:日志消息",
}

//日志格式验证
func InspectionFormat(jsonData string) (interface{}, bool) {
	err := json.Unmarshal([]byte(jsonData), &jsonMap)
	if err != nil {
		logs.ErrorLog(err)
	}
	fmt.Println(jsonMap)
	result, err := govalidator.ValidateMap(jsonMap, MapTemplate)
	if err != nil {
		fmt.Println("error:" + err.Error())
		return err.Error(), result
	}

	return jsonMap, result
}
