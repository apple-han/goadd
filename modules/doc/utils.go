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

package doc

import (
	"io/ioutil"
)

// ListImports checks and returns a list of imports of given import path and options.
func ListImports(target string) (r []string, err error) {
	// 获取目标目录下的文件名称
	files, _ := ioutil.ReadDir(target)
	for _, f := range files {
		r = append(r, f.Name())
	}
	return r, nil
}
