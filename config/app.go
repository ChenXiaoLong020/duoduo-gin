package config

type App struct {
	Domain string `yaml:"domain"`
	Host   string `yaml:"host"`
	Port   int    `yaml:"port"`
}
