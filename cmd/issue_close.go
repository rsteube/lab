package cmd

import (
	"fmt"
	"log"

	zsh "github.com/rsteube/cobra-zsh-gen"
	"github.com/spf13/cobra"
	"github.com/zaquestion/lab/cmd/action"
	lab "github.com/zaquestion/lab/internal/gitlab"
)

var issueCloseCmd = &cobra.Command{
	Use:     "close [remote] <id>",
	Aliases: []string{"delete"},
	Short:   "Close issue by id",
	Long:    ``,
	Args:    cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		rn, id, err := parseArgs(args)
		if err != nil {
			log.Fatal(err)
		}

		p, err := lab.FindProject(rn)
		if err != nil {
			log.Fatal(err)
		}

		err = lab.IssueClose(p.ID, int(id))
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Issue #%d closed\n", id)
	},
}

func init() {
	issueCmd.AddCommand(issueCloseCmd)
	zsh.Gen(issueCloseCmd).PositionalCompletion(
		action.Remotes(),
		action.Issues(issueList),
	)
}
