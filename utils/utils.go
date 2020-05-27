package utils

import (
	"encoding/json"
	"github.com/AliyunContainerService/k8s-version-reminder/config"
	"io/ioutil"
)

func Load(filename string) (c *config.Config, err error) {
	var cfg config.Config
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(data, &cfg)
	if err != nil {
		return nil, err
	}
	return &cfg, err
}
