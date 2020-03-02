package cmd

import (
	zsh "github.com/rsteube/cobra-zsh-gen"
	"github.com/spf13/cobra"
	"github.com/zaquestion/lab/cmd/action"
)

var mergeRequestCmd = &cobra.Command{
	Use:   "merge-request [remote [branch]]",
	Short: mrCreateCmd.Short,
	Long:  mrCreateCmd.Long,
	Args:  mrCreateCmd.Args,
	Run: func(cmd *cobra.Command, args []string) {
		runMRCreate(cmd, args)
	},
}

func init() {
	RootCmd.AddCommand(mergeRequestCmd)
	zsh.Gen(mergeRequestCmd).PositionalCompletion(
		action.Remotes(),
		action.RemoteBranches(0),
	)
}
