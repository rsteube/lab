package cmd

import (
	"log"
	"strconv"

	"github.com/spf13/cobra"
	"github.com/zaquestion/lab/internal/browser"
	"github.com/zaquestion/lab/internal/gitlab"
)

var browse = browser.Open

var issueBrowseCmd = &cobra.Command{
	Use:     "browse [remote] <id>",
	Aliases: []string{"b"},
	Short:   "View issue in a browser",
	Long:    ``,
	Run: func(cmd *cobra.Command, args []string) {
		rn, num, err := parseArgs(args)
		if err != nil {
			log.Fatal(err)
		}

		project, err := gitlab.FindProject(rn)
		if err != nil {
			log.Fatal(err)
		}

		// path.Join will remove 1 "/" from "http://" as it's consider that's
		// file system path. So we better use normal string concat
		issueURL := project.WebURL + "/issues"
		if num > 0 {
			issueURL = issueURL + "/" + strconv.FormatInt(num, 10)
		}

		err = browse(issueURL)
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
    // TODO
	//issueBrowseCmd.MarkZshCompPositionalArgumentCustom(1, "__lab_completion_remote")
	//issueBrowseCmd.MarkZshCompPositionalArgumentCustom(2, "__lab_completion_issue $words[2]")
	issueCmd.AddCommand(issueBrowseCmd)
}
