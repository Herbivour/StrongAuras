package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Config struct {
	EqFolder       string       `json:"eq_folder"`
	Character      string       `json:"character"`
	Server         string       `json:"server"`
	WindowPosition Position     `json:"window_position"`
	Indicators     []*indicator `json:"indicators,omitempty"`
}

func LoadConfig(path string) Config {
	var cfg Config
	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}

	dat, err := ioutil.ReadAll(f)
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(dat, &cfg)
	if err != nil {
		panic(err)
	}

	fmt.Println("Cfg loaded: ", cfg)
	return cfg
}
