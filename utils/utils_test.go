package utils

import (
	"fmt"
	"testing"
)

func TestJsonStruct_Load(t *testing.T) {
	t.Run("readConfig", func(t *testing.T) {
		c,err := Load("../config/version.json")
		if err != nil {
			t.Error(err)
		}
		fmt.Println(c.Deployment)
	})

}
