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
		case "fish":
			cmd, _, _ := RootCmd.Find([]string{"_carapace"})
			cmd.Run(cmd, []string{"fish"})
		case "zsh":
			cmd, _, _ := RootCmd.Find([]string{"_carapace"})
			cmd.Run(cmd, []string{"zsh"})
		default:
			println("only 'bash', 'fish' or 'zsh' allowed")
		}
	},
}

func init() {
	RootCmd.AddCommand(completionCmd)
}
