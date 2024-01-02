package config

type Mysql struct {
	Host    string `yaml:"host"`
	Port    string `yaml:"port"`
	User    int    `yaml:"user"`
	Pwd     int    `yaml:"pwd"`
	Charset int    `yaml:"charset"`
}
