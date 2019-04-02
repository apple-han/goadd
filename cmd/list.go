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

package cmd

import (
	"fmt"
	"sort"

	"github.com/apple-han/goadd/modules/doc"

	"github.com/apple-han/goadd/modules/cli"
	"github.com/apple-han/goadd/modules/errors"
	"github.com/apple-han/goadd/modules/log"
	"github.com/apple-han/goadd/modules/setting"
)

var CmdList = cli.Command{
	Name:  "list",
	Usage: "list all files of gopath or list . files in the current directory ",
	Description: `Command list lists all files under Gopath

goadd list
`,
	Action: runList,
}

// getDepList gets list of dependencies in root path format and nature order.
func getDepList(ctx *cli.Context, path string) ([]string, error) {

	//If the user passes a point after the list, the file in the current path is returned.

	if args := ctx.Args(); len(args) > 0 && args[0] == "." {
		path = "./"
	}

	imports, err := doc.ListImports(path)
	if err != nil {
		return nil, err
	}
	list := make([]string, 0, len(imports))
	for _, name := range imports {
		list = append(list, name)
	}
	sort.Strings(list)
	return list, nil
}

func runList(ctx *cli.Context) {
	if err := setup(ctx); err != nil {
		errors.SetError(err)
		return
	}

	if !setting.HasGOPATHSetting {
		log.Error("gopath must have")
	}

	list, err := getDepList(ctx, setting.InstallGopath)
	if err != nil {
		errors.SetError(err)
		return
	}

	fmt.Printf("Dependency list (%d):\n", len(list))
	for _, name := range list {
		fmt.Printf("-> %s\n", name)
	}
}
