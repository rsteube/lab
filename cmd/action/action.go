package action

import (
	"strconv"

	zsh "github.com/rsteube/cobra-zsh-gen"
	"github.com/xanzy/go-gitlab"
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

func Snippets(snippetList func(args []string) ([]*gitlab.Snippet, error)) zsh.Action {
      return zsh.ActionCallback(func(args []string) zsh.Action {
        if snips, err := snippetList(args); err != nil {
          return zsh.ActionMessage(err.Error())
        } else {
          values := make([]string, len(snips)*2)
          for index, snip := range snips {
            values[index*2] = strconv.Itoa(snip.ID)
            values[index*2+1] = snip.Title
          }
          return zsh.ActionValuesDescribed(values...)
        }
      })
}
