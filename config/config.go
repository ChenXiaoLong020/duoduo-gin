package config

type Server struct {
	App   App   `yaml:"app"`
	Mysql Mysql `yaml:"mysql"`
}
