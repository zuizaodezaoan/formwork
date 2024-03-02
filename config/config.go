package config

var Usersrv SrvConfig

type SrvConfig struct {
	UserName string `json:"user_name"`
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Mysql    Mysqls `json:"mysql"`
	Nacos    Nacoss `json:"Nacos"`
}

type Nacoss struct {
	Host        string `json:"host"`
	Port        int    `json:"port"`
	NamespaceId string `json:"namespaceId"`
	LogDir      string `json:"logDir"`
	CacheDir    string `json:"cacheDir"`
	LogLevel    string `json:"logLevel"`
	DataId      string `json:"dataId"`
	Group       string `json:"group" mapstructure:"group" yaml:"group"`
}

type Mysqls struct {
	Number   string `json:"number"`
	Password string `json:"password"`
	Host     string `json:"host"`
	Port     int    `json:"port"`
	DbName   string `json:"db_name"`
}
