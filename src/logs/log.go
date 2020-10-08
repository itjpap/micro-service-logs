package logs

import (
	BusinessCode "enum"
	"fmt"
	rotatelogs "github.com/lestrrat/go-file-rotatelogs"
	"github.com/pkg/errors"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"github.com/weekface/mgorus"
	"path"
	"reflect"
)


var log = logrus.New()

func init(){
	hooker,errConn := mgorus.NewHooker("127.0.0.1:27017","crm-logs","collection1")
	log.Hooks.Add(hooker)
	if errConn != nil{
		log.Fatal(errConn)
	}

}



////记录错误日志
//func Record(err error){
//	log.Formatter = &logrus.JSONFormatter{}
//	log.WithFields(logrus.Fields{
//		"omg":true,
//		"number":100,
//	}).Error(err)
//}

type jsonField struct {
	Level  string `json:"level"`  //日志等级
	Source string `json:"source"` //来源 RPC/restful api
	Code   int    `json:"code"`   // 记录代码
	Msg    string `json:"msg"`    //日志消息内容
}


type logField struct {
	Level string //日志等级
	Source string  //来源 RPC/restful api
	Code   int      // 记录代码
	Msg    string     //日志消息内容
}

//type WriteField map[string]interface{}  {
//	"aaa":a,
//}




//将json格式日志通过钩子写入mongo
func WriteMongo(content string){
	//jsonData := `{"Level":"warn","Source":"RPC","Code" : 10001,"Message":"this is rpc error msg"}`
	//var fields jsonField
	//errJson := json.Unmarshal([]byte(jsonData),&fields)
	//if errJson != nil{
	//	ErrorLog(errJson)
	//}
	//logStruct := TurnStructLog(fields)
	//record(logStruct,log)
}



//将json结构体转成日志结构体
func TurnStructLog(fields jsonField)  logField{
	//定义一个新的日志结构体
	var logField logField
	ftObj := reflect.TypeOf(fields)
	fvObj := reflect.ValueOf(fields)
	lfObj := reflect.ValueOf(&logField).Elem()
	for i:=0;i < ftObj.NumField();i++{
		name := ftObj.Field(i).Name
		val := fvObj.Field(i)
		lfObj.FieldByName(name).Set(val)
	}
	return logField
}



//记录日志，日志等级解析
func Record(fields map[string]interface{}){
	//fields := make(map[string]interface{})
	//err := mapstructure.Decode(fields,fieldsStruct)
	//ErrorLog(err)

	fmt.Println("break point 10002")
	fmt.Println(fields)
	switch fields["Level"] {
	case "debug":
		log.WithFields(fields).Debug()
		break
	case "info":
		log.WithFields(fields).Info()
		break
	case "warn":
		log.WithFields(fields).Warn()
		break
	case "error":
		log.WithFields(fields).Error()
		break
	case "fatal":
		log.WithFields(fields).Fatal()
		break

	case "panic":
		log.WithFields(fields).Panic()
		break
	default:
		log.WithFields(logrus.Fields{
			"Source":"LOCAL",
			"Code":BusinessCode.LEVEL_NOT_ALLOW,
			"Message":"level type not match",
		})
		break
	}
	fmt.Println("break point 10003")
}



//记录运行或解析变量时产生的错误
func ErrorLog(err error){
	log.WithFields(logrus.Fields{
		"Source":"LOCAL",
		"Code":BusinessCode.RUNTIME_ERROR,
		"Message":err.Error(),
	})
}



/**
  通过钩子写入日志到本地文件
  通过文件形式写入日志效率会比较慢
 */
func WriterLocalFile(level string,content string){
	baseLogPath := path.Join("./app_logs")
	writer,err := rotatelogs.New(
		baseLogPath+".%Y%m&d%H%M",
		rotatelogs.WithLinkName(baseLogPath),  //生成软链，指向最新日志
		rotatelogs.WithMaxAge(10000), //文件最大保存时间
		rotatelogs.WithRotationTime(3600),
	)

	if err != nil {
		log.Errorf("config local file system logger error .%+v",errors.WithStack(err))
	}
	lfHook := lfshook.NewHook(lfshook.WriterMap{
		logrus.DebugLevel: writer, //为不同级别设置不同的输出目的
		logrus.InfoLevel:writer,
		logrus.WarnLevel:writer,
		logrus.ErrorLevel:writer,
		logrus.FatalLevel:writer,
		logrus.PanicLevel: writer,
	},&logrus.JSONFormatter{})
	log.AddHook(lfHook)


}


