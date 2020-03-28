// Copyright 2020 The web_access_local Authors. All rights reserved.
// Use of this source code is governed by a GPL-style license that can be found in the LICENSE file.

package configs

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"

	"github.com/Unknwon/goconfig"
)

var MyConfig *Config

type Config struct {
	Port    int
	AuthUrl string
}

func init() {
	MyConfig = &Config{
		Port:    9099,
		AuthUrl: "",
	}

	configFilePath := fmt.Sprintf("%s%s", filepath.Base(os.Args[0]), ".ini")
	gocfg, err := goconfig.LoadConfigFile(configFilePath)
	if err == nil {
		portStr, err := gocfg.GetValue("app", "port")
		if err != nil {
			fmt.Println(err.Error())
		} else {
			MyConfig.Port, err = strconv.Atoi(portStr)
			if err != nil {
				fmt.Println(err.Error())
				os.Exit(1)
				return
			}
		}
		authUrl, err := gocfg.GetValue("app", "authUrl")
		if err != nil {
			// fmt.Println(err.Error())
		} else {
			MyConfig.AuthUrl = authUrl
		}
	}
}
