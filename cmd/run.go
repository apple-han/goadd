package cmd

import (
	"os"
	"path"
	"regexp"

	b "github.com/apple-han/goadd/modules/base"
	"github.com/apple-han/goadd/modules/cli"
	"github.com/apple-han/goadd/modules/errors"
	"github.com/apple-han/goadd/modules/log"
	"github.com/apple-han/goadd/modules/setting"
)

var CmdRun = cli.Command{
	Name:        "run",
	Usage:       "add a new folder, and add to gopath",
	Description: `add a new folder, and add to gopath,and execute 'goadd run'.`,
	Action:      runRun,
}

func runRun(ctx *cli.Context) {
	if err := setup(ctx); err != nil {
		errors.SetError(err)
		return
	}

	// 如果是已存在的文件或者是文件夹, 就直接移到GoPath就可以了
	// 如果不是的就要判断:
	// 	1. 不能含有特殊字符 [\/\\\:\*\?\"\<\>\&\#\@\%\|]
	log.Info("Running...")
	if args := ctx.Args(); len(args) > 0 {
		folder := args[0]
		var validFolder = regexp.MustCompile(`[\/\\\:\*\?\"\<\>\&\#\@\%\|]`)

		if validFolder.MatchString(folder) {
			log.Error("file format error")
		}
		InstallPath := path.Join(setting.InstallGopath, folder)

		if b.IsFile(folder) || b.IsDir(folder) {
			if err := os.Rename(folder, InstallPath); err != nil {
				log.Error("fail to move binary: %v", err)
			}
		}

		if err := os.MkdirAll(InstallPath, os.ModePerm); err != nil {
			log.Error("fail to install binary: %v", err)
		}
		log.Info("Local repository path: %s", InstallPath)

	}

	log.Info("Command executed successfully!")
}
