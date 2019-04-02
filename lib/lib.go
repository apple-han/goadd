// +build go1.10

// Copyright 2014 Unknwon
//
// Licensed under the Apache License, Version 2.0 (the "License"): you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
// WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
// License for the specific language governing permissions and limitations
// under the License.

// Package lib is a library version of Goadd(Add Folder To GoPath).
package lib

import (
	"runtime"

	"github.com/apple-han/goadd/cmd"
	"github.com/apple-han/goadd/modules/cli"
	"github.com/apple-han/goadd/modules/setting"
)

const APP_VER = "0.0.1 Beta"

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func Run(args []string) *setting.Error {
	app := cli.NewApp()
	app.Name = "Goadd"
	app.Usage = "Add Folder To GoPath"
	app.Version = APP_VER
	app.Commands = []cli.Command{
		cmd.CmdRun,
		cmd.CmdList,
	}
	app.Run(args)
	return setting.RuntimeError
}
