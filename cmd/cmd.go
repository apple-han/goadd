package cmd

import (
	"fmt"
	"runtime"

	"github.com/apple-han/goadd/modules/base"
	"github.com/apple-han/goadd/modules/cli"
	"github.com/apple-han/goadd/modules/setting"
)

// setup initializes and checks common environment variables.
func setup(ctx *cli.Context) (err error) {

	if runtime.GOOS == "windows" {
		setting.IsWindows = true
	}

	// Get GOPATH.
	setting.InstallGopath = base.GetGOPATHs()[0]
	if base.IsDir(setting.InstallGopath) {

		setting.InstallGopath += "/src"
		setting.HasGOPATHSetting = true
	} else {
		if ctx.Bool("gopath") {
			return fmt.Errorf("Local GOPATH does not exist or is not a directory: %s",
				setting.InstallGopath)
		} else {
			// It's OK that no GOPATH setting
			// when user does not specify to use.
		}
	}
	return nil
}
