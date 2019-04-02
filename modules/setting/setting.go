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

package setting

type Error struct {
	HasError bool
	Fatal    error
	Errors   []error
}

type Option struct {
}

const (
	VERSION = 20190325
)

var (
	// Global variables.
	InstallGopath string

	// Global settings.
	RuntimeError = new(Error)

	// System settings.
	IsWindows        bool
	IsWindowsXP      bool
	HasGOPATHSetting bool

	// Configuration settings.

	CommonRes = []string{"views", "templates", "static", "public", "conf"}
)
