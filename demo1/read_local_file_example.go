package demo1

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type Config struct {
	Service Service `yaml:"Service"`
}

type Service struct {
	Name string `yaml:"Name"`
	Port string `yaml:"Port"`
}

func ReadFile() {
	var config = &Config{}
	fmt.Printf("conf --->>>%s\n", config)
	c, err := ioutil.ReadFile("demo1/go.yaml")
	if err != nil {
		fmt.Println("read file failed ,error ", err.Error())
	} else {
		fmt.Println("read file success ,res ", c)
	}
	err = yaml.Unmarshal(c, config)
	if err != nil {
		fmt.Println("read file failed ,error ", err.Error())
	}
	fmt.Printf("conf --->>>%s\n", config)
}
