package parseconf

var ConfigInstance map[string]interface{}

type redis struct {
	Host string `ini:"host"`
	Port string `ini:"port"`
	Auth string `ini:"auth"`
}

type mysql struct {
	Host   string `ini:"host"`
	Port   string `ini:"port"`
	Passwd string `ini:"passwd"`
	Db     string `ini:"db"`
}

var Redis *redis

var Mysql *mysql

func init() {

	ConfigInstance = make(map[string]interface{})

	ConfigInstance["Global"] = make(map[string]string)

	ConfigInstance["Redis"] = Redis
	ConfigInstance["Mysql"] = Mysql

}
