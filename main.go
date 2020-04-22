// web_access_local project main.go
package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"strings"

	"golang.org/x/net/websocket"

	"web_access_local/actions"
	"web_access_local/configs"
	"web_access_local/protocols"
)

func main() {

	http.Handle("/access", websocket.Handler(access))
	http.HandleFunc("/index", index)
	http.HandleFunc("/", index)
	if err := http.ListenAndServe(fmt.Sprintf("127.0.0.1:%d", configs.MyConfig.Port), nil); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func index(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	w.Write([]byte("正常"))
}

func access(ws *websocket.Conn) {
	var err error
	var message, reply string
	var authed = false

	defer ws.Close()

	for {
		if err = websocket.Message.Receive(ws, &message); err != nil {
			if err == io.EOF {
				fmt.Println("client close")
				break
			}
			fmt.Println(fmt.Sprintf("err: %v", err))
			continue
		}

		reply = processMessage(message, &authed)
		err = websocket.Message.Send(ws, reply)
		if err != nil {
			fmt.Println(fmt.Sprintf("err: %v", err))
		}
		if !authed {
			return
		}
	}
}

func processMessage(msgContent string, authed *bool) (replyContent string) {
	var message protocols.Message
	err := json.Unmarshal([]byte(msgContent), &message)
	if err != nil {
		println(fmt.Sprintf("err: %v", err))
	}

	reply := protocols.Reply{
		Sn:     message.Sn,
		Result: false,
	}

	if !*authed {
		if protocols.CmdEnum(message.Cmd) == protocols.Cmd_Auth {
			*authed = auth(message.Args)
			reply.Result = true
		} else {
			reply.Result = false
			reply.Remark = "No authentication"
		}
	} else {
		switch protocols.CmdEnum(message.Cmd) {
		case protocols.Cmd_AppRun:
			if len(message.Args) > 0 {
				cmd := exec.Command(message.Args)
				err = cmd.Start()
				if err == nil {
					reply.Result = true
				} else {
					reply.Result = false
					reply.Remark = fmt.Sprintf("%v", err)
				}
			}
		case protocols.Cmd_CpuReadInfo:
			cpuArr, err := actions.GetCpuInfo()
			reply.Result = cpuArr
			if err != nil {
				reply.Remark = err.Error()
			}
		case protocols.Cmd_Custom:
			reply.Remark = "Operation was not implemented"
		case protocols.Cmd_CustomBinary:
			reply.Remark = "Operation was not implemented"
		case protocols.Cmd_FileReadData:
			filePath := message.Args
			byteArr, err := ioutil.ReadFile(filePath)
			if err != nil {
				reply.Remark = err.Error()
			} else {
				reply.Result = string(byteArr)
			}
		case protocols.Cmd_FileReadDataBinary:
			filePath := message.Args
			byteArr, err := ioutil.ReadFile(filePath)
			if err != nil {
				reply.Remark = err.Error()
			} else {
				reply.Result = byteArr
			}
		case protocols.Cmd_FileReadMeta:
			filePath := message.Args
			file, err := os.Open(filePath)
			defer file.Close()
			if err != nil {
				reply.Remark = err.Error()
			} else {
				var fileInfo os.FileInfo
				fileInfo, err = file.Stat()
				if err != nil {
					reply.Remark = err.Error()
				} else {
					reply.Result = fileInfo
				}
			}
		case protocols.Cmd_MacReadInfo:
			macArr, err := actions.GetMacAddress()
			reply.Result = macArr
			if err != nil {
				reply.Remark = err.Error()
			}
		case protocols.Cmd_MachineReadGuid:
			machineGuid, err := actions.GetMachineGuid()
			reply.Result = machineGuid
			if err != nil {
				reply.Remark = err.Error()
			}
		case protocols.Cmd_RegReadValue:
			argArr := strings.Split(message.Args, ",")
			if len(argArr) == 2 {
				regValue, err := actions.QueryRegValue(argArr[0], argArr[1])
				reply.Result = regValue
				if err != nil {
					reply.Remark = err.Error()
				}
			} else {
				reply.Remark = "Parameter format is incorrect, registry key and key are separated by ,"
			}
		case protocols.Cmd_UDiskReadList:
			udiskArr, err := actions.QueryUDisks()
			reply.Result = udiskArr
			if err != nil {
				reply.Remark = err.Error()
			}
		default:
		}
	}

	byteArr, _ := json.Marshal(reply)
	replyContent = string(byteArr)
	return
}

func auth(accessCode string) bool {
	authUrl := strings.ToLower(strings.TrimSpace(configs.MyConfig.AuthUrl))
	if len(authUrl) == 0 {
		return true
	} else {
		if len(accessCode) == 0 {
			return false
		}

		if !strings.HasPrefix(authUrl, "http://") || strings.HasPrefix(authUrl, "https://") {
			authUrl = "http://" + authUrl
		}
		if !strings.HasSuffix(authUrl, "=") {
			if strings.Index(authUrl, "?") > 0 {
				authUrl = authUrl + "&access="
			} else {
				authUrl = authUrl + "?access="
			}
		}
		rsp, err := http.Get(fmt.Sprintf("%s%s", authUrl, accessCode))
		if rsp != nil && rsp.Body != nil {
			defer rsp.Body.Close()
		}
		if err != nil {
			return false
		}
		body, err := ioutil.ReadAll(rsp.Body)
		if err != nil {
			return false
		}

		if strings.HasPrefix(strings.TrimSpace(strings.ToLower(string(body))), "ok") {
			return true
		}
	}
	return false
}
