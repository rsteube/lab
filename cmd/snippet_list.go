package cmd

import (
	"fmt"
	"log"

	zsh "github.com/rsteube/cobra-zsh-gen"
	"github.com/spf13/cobra"
	gitlab "github.com/xanzy/go-gitlab"
	"github.com/zaquestion/lab/internal/git"
	lab "github.com/zaquestion/lab/internal/gitlab"
)

var snippetListConfig struct {
	Number int
	All    bool
}

// snippetListCmd represents the snippetList command
var snippetListCmd = &cobra.Command{
	Use:     "list [remote]",
	Aliases: []string{"ls"},
	Short:   "List personal or project snippets",
	Long:    ``,
	Run: func(cmd *cobra.Command, args []string) {
		rn, _, err := parseArgs(args)
		if err != nil {
			log.Fatal(err)
		}
		listOpts := gitlab.ListOptions{
			PerPage: snippetListConfig.Number,
		}

		num := snippetListConfig.Number
		if snippetListConfig.All {
			num = -1
		}
		// See if we're in a git repo or if global is set to determine
		// if this should be a personal snippet
		if global || rn == "" {
			opts := gitlab.ListSnippetsOptions(listOpts)
			snips, err := lab.SnippetList(opts, num)
			if err != nil {
				log.Fatal(err)
			}
			for _, snip := range snips {
				fmt.Printf("#%d %s\n", snip.ID, snip.Title)
			}
			return
		}

		project, err := lab.FindProject(rn)
		if err != nil {
			log.Fatal(err)
		}
		opts := gitlab.ListProjectSnippetsOptions(listOpts)
		snips, err := lab.ProjectSnippetList(project.ID, opts, num)
		if err != nil {
			log.Fatal(err)
		}
		for _, snip := range snips {
			fmt.Printf("#%d %s\n", snip.ID, snip.Title)
		}
	},
}

func init() {
	snippetListCmd.Flags().IntVarP(&snippetListConfig.Number, "number", "n", 10, "Number of snippets to return")
	snippetListCmd.Flags().BoolVarP(&snippetListConfig.All, "all", "a", false, "List all snippets")

	snippetCmd.AddCommand(snippetListCmd)
    zsh.Gen(snippetListCmd).PositionalCompletion(
      zsh.ActionCallback(func(args []string) zsh.Action {
        if remotes, err := git.Remotes(); err != nil {
          return zsh.ActionMessage(err.Error())
        } else {
          return zsh.ActionValues(remotes...)
        }
      }),
    )
}
