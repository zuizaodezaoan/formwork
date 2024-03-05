package config

var Usersrv SrvConfig

type SrvConfig struct {
	UserName string   `json:"user_name"`
	Host     string   `json:"host"`
	Port     int      `json:"port"`
	Tags     []string `json:"tags"`
	Mysql    Mysqls   `json:"mysql"`
	Redis    Rediss   `json:"redis"`
	Nacos    Nacoss   `json:"Nacos"`
	Consul   Consuls  `json:"consul"`
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

type Rediss struct {
	Host string `json:"host"`
	Port int    `json:"port"`
}

type Consuls struct {
	Host string `json:"host"`
	Port int    `json:"port"`
}
