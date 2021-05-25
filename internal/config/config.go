package config

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type Conf struct {
	MelissaKey string `json:"melissaKeyCred"`
	HibpKey    string `json:"hibpKey"`
	DataGovKey string `json:"dataGovKey"`
}

func LoadConfig(path string) (*Conf, error) {
	jsonFile, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer jsonFile.Close()
	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return nil, err
	}
	conf := new(Conf)
	err = json.Unmarshal(byteValue, &conf)
	if err != nil {
		return nil, err
	}
	return conf, nil
}
