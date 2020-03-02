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
		if snips, err := snippetList(args[:0]); err != nil {
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

func Issues(issueList func(args []string) ([]*gitlab.Issue, error)) zsh.Action {
	return zsh.ActionCallback(func(args []string) zsh.Action {
		if issues, err := issueList(args[:0]); err != nil {
			return zsh.ActionMessage(err.Error())
		} else {
			values := make([]string, len(issues)*2)
			for index, issue := range issues {
				values[index*2] = strconv.Itoa(issue.IID)
				values[index*2+1] = issue.Title
			}
			return zsh.ActionValuesDescribed(values...)
		}
	})
}

func MergeRequests(mrList func(args []string) ([]*gitlab.MergeRequest, error)) zsh.Action {
	return zsh.ActionCallback(func(args []string) zsh.Action {
		if mergeRequests, err := mrList(args[:0]); err != nil {
			return zsh.ActionMessage(err.Error())
		} else {
			values := make([]string, len(mergeRequests)*2)
			for index, mergeRequest := range mergeRequests {
				values[index*2] = strconv.Itoa(mergeRequest.IID)
				values[index*2+1] = mergeRequest.Title
			}
			return zsh.ActionValuesDescribed(values...)
		}
	})
}
