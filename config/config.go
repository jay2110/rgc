package config

import (
	"errors"
)

type ConfigStruct struct {
	Server struct {
		Port int `yaml:"port"`
	} `yaml:"server"`

	Apikey string `yaml:"apikey"`
	Url    string `yaml:"url"`
}

func (a *ConfigStruct) Validate() (err error) {
	if len(a.Url) <= 0 {
		err = errors.New("Empty url")
	} else if len(a.Apikey) <= 0 {
		err = errors.New("Empty Apikey")
	}
	return err
}
