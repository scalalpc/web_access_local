// Copyright 2020 The web_access_local Authors. All rights reserved.
// Use of this source code is governed by a GPL-style license that can be found in the LICENSE file.

package actions

import (
	"context"

	"github.com/StackExchange/wmi"
)

type CpuInfo struct {
	Cpu          int    `json:"cpu"`
	Manufacturer string `json:"Manufacturer"`
	ProcessorId  string `json:"ProcessorID"`
}

type win32_Processor struct {
	Manufacturer string
	ProcessorId  *string
}

func GetCpuInfo() (cpuArr []CpuInfo, err error) {
	var procArr []win32_Processor
	q := wmi.CreateQuery(&procArr, "")
	if err = wmiQuery(q, &procArr); err != nil {
		return
	}

	var procId string
	for idx, proc := range procArr {
		procId = ""
		if proc.ProcessorId != nil {
			procId = *proc.ProcessorId
		}
		cpu := CpuInfo{
			Cpu:          idx,
			Manufacturer: proc.Manufacturer,
			ProcessorId:  procId,
		}
		cpuArr = append(cpuArr, cpu)
	}
	return
}

func wmiQuery(query string, dst interface{}, connectServerArgs ...interface{}) (err error) {
	ctx := context.Background()
	if _, ok := ctx.Deadline(); !ok {
		timeoutCtx, cancel := context.WithTimeout(ctx, 3000000000)
		defer cancel()
		ctx = timeoutCtx
	}

	errChan := make(chan error, 1)
	go func() {
		errChan <- wmi.Query(query, dst, connectServerArgs...)
	}()

	select {
	case <-ctx.Done():
		err = ctx.Err()
		return
	case err = <-errChan:
		return err
	}
}
