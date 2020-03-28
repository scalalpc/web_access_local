// Copyright 2020 The web_access_local Authors. All rights reserved.
// Use of this source code is governed by a GPL-style license that can be found in the LICENSE file.

package actions

import (
	"errors"
	"net"
)

func GetMacAddress() (macArr []string, err error) {
	var niArr []net.Interface
	niArr, err = net.Interfaces()
	if err != nil {
		return
	}
	for i := 0; i < len(niArr); i++ {
		if (niArr[i].Flags&net.FlagUp) != 0 && (niArr[i].Flags&net.FlagLoopback) == 0 {
			addrArr, _ := niArr[i].Addrs()
			for _, address := range addrArr {
				ipnet, ok := address.(*net.IPNet)
				if ok && ipnet.IP.IsGlobalUnicast() {
					mac := niArr[i].HardwareAddr.String()
					macArr = append(macArr, mac)
				}
			}
		}
	}
	if len(macArr) == 0 {
		err = errors.New("Unable to get the correct MAC address.")
	}
	return
}
