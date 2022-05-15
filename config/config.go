package config

type Revgeo struct {
	Server struct {
		Port int `yaml:"port"`
	} `yaml:"server"`

	Apikey string `yaml:"apikey"`
	Url    string `yaml:"url"`
}
