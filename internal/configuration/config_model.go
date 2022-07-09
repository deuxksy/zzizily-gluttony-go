package configuration

var RuntimeConf = RuntimeConfig{}

type RuntimeConfig struct {
	Scenario 		[]Scene      `yaml:"Scenario,mapstructure"`
}

type Scene struct {
	Name string `yaml:"Name"`
	Url string `yaml:"Url"`
	Action string `yaml:"Action"`
}