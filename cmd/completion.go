package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// completionCmd represents the completion command
var completionCmd = &cobra.Command{
	Use:       "completion",
	Short:     "Generates the shell autocompletion",
	Long:      `'completion bash' generates the bash and 'completion carapace' the carapace autocompletion`,
	Args:      cobra.ExactArgs(1),
	ValidArgs: []string{"bash", "carapace"},
	Run: func(cmd *cobra.Command, args []string) {
		switch args[0] {
		case "bash":
			RootCmd.GenBashCompletion(os.Stdout)
		case "carapace":
			cmd, _, _ := RootCmd.Find([]string{"_carapace_completion"})
			cmd.Run(cmd, []string{})
		default:
			println("only 'bash' or 'carapace' allowed")
		}
	},
}

func init() {
	RootCmd.AddCommand(completionCmd)
}
