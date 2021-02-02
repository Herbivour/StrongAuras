package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type config struct {
	EqFolder   string       `json:"eq_folder"`
	Character  string       `json:"character"`
	Server     string       `json:"server"`
	Indicators []*indicator `json:"indicators,omitempty"`
}

var cfg config

func loadConfig() {
	f, err := os.Open("./config.json")
	check(err)

	dat, err := ioutil.ReadAll(f)
	check(err)

	err = json.Unmarshal(dat, &cfg)
	check(err)

	fmt.Println("Cfg loaded: ", cfg)
}
