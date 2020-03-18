package parseconf

import (
	"bufio"
	"errors"
	"go-simple-parse-ini/util"
	"io"
	"os"
	"reflect"
	"strings"
)

var (
	NOTEXIST   = errors.New("File not exist")
	NOTINITYPE = errors.New("File isn't ini type file")
)

//验证文件

func InitialConf(file *string) error {

	//判断文件是否存在
	_, err := os.Stat(*file)
	if err != nil {
		if os.IsNotExist(err) {
			return NOTEXIST
		}
		panic(err)
	}

	//判断后缀是否为ini
	if !strings.HasSuffix(*file, ".ini") {
		return NOTINITYPE
	}

	//进行文件读取解析
	assignStru(file)
	return nil

}

func assignStru(file *string) {

	//读取配置文件解析
	filehandle, err := os.Open(*file)

	if err != nil {
		panic(err)
	}

	var StrName string

	paraMap := make(map[string]string, 1)

	readhandle := bufio.NewReader(filehandle)

	defer filehandle.Close()

	for {
		data, _, err := readhandle.ReadLine()

		if err == io.EOF {
			break
		}

		if len(data) > 0 {
			switch data[0] {
			case ';':
				continue //注释行
			case '[':
				data = data[1 : len(data)-1]
				strTempName := string(data)
				strTempName = strings.ToLower(strTempName)
				util.FirstToUpper(&strTempName) //切换首字母为大写
				StrName = strTempName

				instanceVal := reflect.ValueOf(ConfigInstance[StrName])

				if instanceVal.IsNil() { //需要通过StrName来映射 Redis和Mysql 没有想到更好的方式 暂用map

					srcPrtType := reflect.TypeOf(ConfigInstance[StrName]) //因为是指针类型
					srcType := srcPrtType.Elem()                          //值类型
					srcValue := reflect.New(srcType)
					instace := srcValue.Interface()

					//创建回map中
					ConfigInstance[StrName] = instace

				}

			default:

				str := string(data)
				if strings.Contains(str, "=") { //包含等于
					strarg := strings.Split(str, "=")

					if len(strarg) > 1 {
						strarg[0] = strings.Trim(strarg[0], " ")
						strarg[1] = strings.Trim(strarg[1], " ")
					}

					if StrName == "" {
						paraMap[strarg[0]] = strarg[1]
					} else if StrName != "" {

						typeInstance := reflect.TypeOf(ConfigInstance[StrName]).Elem()
						valInstance := reflect.ValueOf(ConfigInstance[StrName]).Elem()

						for i := 0; i < typeInstance.NumField(); i++ {
							fieldName := typeInstance.Field(i).Name
							tagName := typeInstance.Field(i).Tag.Get("ini")

							if tagName == strarg[0] {
								valInstance.FieldByName(fieldName).SetString(strarg[1])
								break
							}

						}

					}

				}

			}
		}

	}

	//从开始没有匹配到section 默认归属global
	ConfigInstance["Global"] = paraMap

}
