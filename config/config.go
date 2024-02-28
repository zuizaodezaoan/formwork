package config

type Nacos struct {
	Host        string `mapstructure:"host" json:"host"`
	Port        int    `mapstructure:"port" json:"port"`
	NamespaceId string `mapstructure:"namespace_id" json:"namespace_id"`
	LogDir      string `mapstructure:"log_dir" json:"log_dir"`
	CacheDir    string `mapstructure:"cache_dir" json:"cache_dir"`
	LogLevel    string `mapstructure:"log_level" json:"log_level"`
	DataId      string `mapstructure:"data_id" json:"data_id"`
	Group       string `mapstructure:"group" json:"group"`
}

type Mysql struct {
	Number   string `mapstructure:"number" json:"number"`
	Password string `mapstructure:"password" json:"password"`
	Host     string `mapstructure:"host" json:"host"`
	Port     int    `mapstructure:"port" json:"port"`
	DbName   string `mapstructure:"db_name" json:"db_name"`
}
