package main

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"net/http"
	"path/filepath"
)

type Config struct {
	Username string
	Password string
}

func main() {
	filename, _ := filepath.Abs("./config.yml")
	yamlFile, err := ioutil.ReadFile(filename)

	if err != nil {
		panic(err)
	}

	var config Config

	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		panic(err)
	}

	// fmt.Printf("Value: %#v\n", string(config.Password))

	resp, err := http.Get("http://" + config.Username + ":" + config.Password + "@cp.acs.li/nic/update?hostname=casa.genar.me")

	if err != nil {
		fmt.Printf("Some error ocurred")
	} else {
		defer resp.Body.Close()
		contents, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Printf("%s", err)
		}
		fmt.Printf("%s\n", string(contents))
	}

	defer resp.Body.Close()
	// body, err := ioutil.ReadAll(resp.Body)
}
