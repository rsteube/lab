package cmd

import (
	"log"

	zsh "github.com/rsteube/cobra-zsh-gen"
	"github.com/spf13/cobra"
	"github.com/zaquestion/lab/cmd/action"
	lab "github.com/zaquestion/lab/internal/gitlab"
)

var mrRebaseCmd = &cobra.Command{
	Use:     "rebase [remote] <id>",
	Aliases: []string{"delete"},
	Short:   "Rebase an open merge request",
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

		err = lab.MRRebase(p.ID, int(id))
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	mrCmd.AddCommand(mrRebaseCmd)
    zsh.Gen(mrRebaseCmd).PositionalCompletion(
      action.Remotes(),
      action.MergeRequests(mrList),
    )
}
