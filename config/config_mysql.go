package config

import "strconv"

type Mysql struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`   //port应该为int类型
	Config   string `yaml:"config"` //高级配置，例如 charset
	DB       string `yaml:"db"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	LogLevel string `yaml:"log_level"` //日志等级，debug为输出全部sql，dev，release
}

func (m Mysql) Dsn() string {
	// m.Port为Int类型，要转成字符串
	return m.User + ":" + m.Password + "@tcp(" + m.Host + ":" + strconv.Itoa(m.Port) + ")/" + m.DB + "?" + m.Config
}
