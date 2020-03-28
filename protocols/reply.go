// Copyright 2020 The web_access_local Authors. All rights reserved.
// Use of this source code is governed by a GPL-style license that can be found in the LICENSE file.

package protocols

type Reply struct {
	Sn     uint        `json:"sn"`
	Result interface{} `json:"result"`
	Remark string      `json:"remark"`
}
