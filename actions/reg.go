// Copyright 2020 The web_access_local Authors. All rights reserved.
// Use of this source code is governed by a GPL-style license that can be found in the LICENSE file.

package actions

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"syscall"
	"unsafe"

	"golang.org/x/sys/windows"
)

func GetMachineGuid() (machineGuid string, err error) {
	var handle windows.Handle
	err = windows.RegOpenKeyEx(windows.HKEY_LOCAL_MACHINE, windows.StringToUTF16Ptr(`SOFTWARE\Microsoft\Cryptography`), 0,
		windows.KEY_READ|windows.KEY_WOW64_64KEY, &handle)
	if err != nil {
		return
	}
	defer windows.RegCloseKey(handle)
	const windowsRegBufferLen = 74
	const uuidLen = 36
	var regBuffer [windowsRegBufferLen]uint16
	bufferLen := uint32(windowsRegBufferLen)
	var valueType uint32
	err = windows.RegQueryValueEx(handle, windows.StringToUTF16Ptr("MachineGuid"), nil, &valueType,
		(*byte)(unsafe.Pointer(&regBuffer[0])), &bufferLen)
	if err != nil {
		return
	}
	hostId := windows.UTF16ToString(regBuffer[:])
	hostIdLen := len(hostId)
	if hostIdLen != uuidLen {
		err = fmt.Errorf("Host ID is incorrect: %q\n", hostId)
		return
	}
	machineGuid = hostId
	return
}

func QueryRegValue(regItem, key string) (value interface{}, err error) {
	regItem = strings.ReplaceAll(regItem, "/", `\`)
	regItem = strings.TrimLeft(regItem, `/\`)
	if strings.Index(regItem, `\`) > 0 {
		pathArr := strings.SplitN(regItem, `\`, 2)

		var rootKeyHandle windows.Handle
		rootKeyHandle = windows.HKEY_LOCAL_MACHINE
		switch strings.ToUpper(pathArr[0]) {
		case "HKEY_CLASSES_ROOT":
			rootKeyHandle = windows.HKEY_CLASSES_ROOT
		case "HKEY_CURRENT_USER":
			rootKeyHandle = windows.HKEY_CURRENT_USER
		case "HKEY_LOCAL_MACHINE":
			rootKeyHandle = windows.HKEY_LOCAL_MACHINE
		case "HKEY_USERS":
			rootKeyHandle = windows.HKEY_USERS
		case "HKEY_CURRENT_CONFIG":
			rootKeyHandle = windows.HKEY_CURRENT_CONFIG
		}

		regItemExcludeRoot := pathArr[1]
		var handle windows.Handle
		err = windows.RegOpenKeyEx(rootKeyHandle, windows.StringToUTF16Ptr(regItemExcludeRoot), 0,
			windows.KEY_READ|windows.KEY_WOW64_64KEY, &handle)
		if err != nil {
			return
		}
		defer windows.RegCloseKey(handle)
		const windowsRegBufferLen = 2000
		var regBuffer [windowsRegBufferLen]uint16
		bufferLen := uint32(windowsRegBufferLen)
		var valType uint32
		err = windows.RegQueryValueEx(handle, windows.StringToUTF16Ptr(key), nil, &valType,
			(*byte)(unsafe.Pointer(&regBuffer[0])), &bufferLen)
		if err != nil {
			return
		}

		switch valType {
		case 1, 2, 7:
			value = windows.UTF16ToString(regBuffer[:bufferLen])
		case 3:
			value = regBuffer[:]
		case 4:
			intValue := uint32(0)
			var byteArr = make([]byte, 0)
			for i := 0; i < int(bufferLen); i++ {
				byteArr = append(byteArr, byte((regBuffer[int(bufferLen)-1-i]&0xff00)>>8))
				byteArr = append(byteArr, byte(regBuffer[int(bufferLen)-1-i]&0xff))
			}
			for i := 0; i < len(byteArr); i++ {
				intValue += uint32(byteArr[i]) << (8 * (len(byteArr) - 1 - i))
			}
			value = interface{}(intValue)
		case 11:
			longValue := uint64(0)
			var byteArr = make([]byte, 0)
			for i := 0; i < int(bufferLen); i++ {
				byteArr = append(byteArr, byte((regBuffer[int(bufferLen)-1-i]&0xff00)>>8))
				byteArr = append(byteArr, byte(regBuffer[int(bufferLen)-1-i]&0xff))
			}
			for i := 0; i < len(byteArr); i++ {
				longValue += uint64(byteArr[i]) << (8 * (len(byteArr) - 1 - i))
			}
			value = interface{}(longValue)
		}

	} else {
		err = errors.New("The registry key is incomplete")
	}
	return
}

func QueryUDisks() (udiskArr []string, err error) {
	var handle windows.Handle
	err = windows.RegOpenKeyEx(windows.HKEY_LOCAL_MACHINE, windows.StringToUTF16Ptr(`SYSTEM\CurrentControlSet\Services\USBSTOR\Enum`), 0,
		windows.KEY_READ|windows.KEY_WOW64_64KEY, &handle)
	if err != nil {
		err = nil
		udiskArr = []string{}
		return
	}

	defer windows.RegCloseKey(handle)

	const windowsRegBufferLen = 4
	var regBuffer [windowsRegBufferLen]byte
	bufferLen := uint32(windowsRegBufferLen)
	var valueType uint32

	err = windows.RegQueryValueEx(handle, windows.StringToUTF16Ptr("Count"), nil, &valueType,
		(*byte)(unsafe.Pointer(&regBuffer[0])), &bufferLen)
	if err != nil {
		err = nil
		udiskArr = []string{}
		return
	}

	count := regBuffer[0]
	allArr := getSystemDisks()
	udiskArr = allArr[len(allArr)-int(count):]
	return
}

func getSystemDisks() (distCharArr []string) {
	kernel32 := syscall.MustLoadDLL("kernel32.dll")
	GetLogicalDrives := kernel32.MustFindProc("GetLogicalDrives")
	n, _, _ := GetLogicalDrives.Call()
	s := strconv.FormatInt(int64(n), 2)
	var allDrives = []string{
		"A:", "B:", "C:", "D:", "E:", "F:", "G:", "H:", "I:", "J:", "K:", "L:", "M:", "N:", "O:", "P:",
		"Q:", "R:", "S:", "T:", "U:", "V:", "W:", "X:", "Y:", "Z",
	}
	temp := allDrives[:len(s)]
	var d []string
	for i, v := range s {
		if v == 49 {
			l := len(s) - i - 1
			d = append(d, temp[l])
		}
	}
	var drives []string
	for i, v := range d {
		drives = append(drives[i:], append([]string{v}, drives[:i]...)...)
	}
	return drives
}
