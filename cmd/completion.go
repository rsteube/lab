package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// completionCmd represents the completion command
var completionCmd = &cobra.Command{
	Use:       "completion",
	Short:     "Generates the shell autocompletion",
	Long:      `'completion bash' generates the bash and 'completion zsh' the zsh autocompletion`,
	Args:      cobra.ExactArgs(1),
	ValidArgs: []string{"bash", "fish", "zsh"},
	Run: func(cmd *cobra.Command, args []string) {
		switch args[0] {
		case "bash":
			RootCmd.GenBashCompletion(os.Stdout)
		case "fish":
			RootCmd.GenFishCompletion(os.Stdout)
		case "zsh":
			RootCmd.GenZshCompletion(os.Stdout)
		default:
			println("expected one of {bash, fish, zsh}")
		}
	},
}

func init() {
	RootCmd.AddCommand(completionCmd)
}
