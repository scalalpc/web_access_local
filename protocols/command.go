// Copyright 2020 The web_access_local Authors. All rights reserved.
// Use of this source code is governed by a GPL-style license that can be found in the LICENSE file.

package protocols

type CmdEnum string

const (
	Cmd_AppRun             CmdEnum = "app_run"
	Cmd_Auth               CmdEnum = "auth"
	Cmd_CpuReadInfo        CmdEnum = "cpu_read_info"
	Cmd_Custom             CmdEnum = "custom"
	Cmd_CustomBinary       CmdEnum = "custom_binary"
	Cmd_FileReadData       CmdEnum = "file_read_data"
	Cmd_FileReadDataBinary CmdEnum = "file_read_data_binary"
	Cmd_FileReadMeta       CmdEnum = "file_read_meta"
	Cmd_MacReadInfo        CmdEnum = "mac_read_info"
	Cmd_MachineReadGuid    CmdEnum = "machine_read_guid"
	Cmd_RegReadValue       CmdEnum = "reg_read_value"
	Cmd_UDiskReadList      CmdEnum = "udisk_read_list"
)
