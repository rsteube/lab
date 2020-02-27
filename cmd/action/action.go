package action

import (
	zsh "github.com/rsteube/cobra-zsh-gen"
	"github.com/zaquestion/lab/internal/git"
)

func Remotes() zsh.Action {
	return zsh.ActionCallback(func(args []string) zsh.Action {
		if remotes, err := git.Remotes(); err != nil {
			return zsh.ActionMessage(err.Error())
		} else {
			return zsh.ActionValues(remotes...)
		}
	})
}

func RemoteBranches(argIndex int) zsh.Action {
	return zsh.ActionCallback(func(args []string) zsh.Action {
        remote := ""
        if argIndex >= 0 {
          remote = args[argIndex]
        }
		if branches, err := git.RemoteBranches(remote); err != nil {
			return zsh.ActionMessage(err.Error())
		} else {
			return zsh.ActionValues(branches...)
		}
	})
}
