package cmd

import (
	"log"
	"net/url"
	"path"
	"strconv"

	zsh "github.com/rsteube/cobra-zsh-gen"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/zaquestion/lab/cmd/action"
)

var snippetBrowseCmd = &cobra.Command{
	Use:   "browse [remote] <id>",
	Short: "View personal or project snippet in a browser",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		rn, id, err := parseArgs(args)
		if err != nil {
			log.Fatal(err)
		}

		c := viper.AllSettings()["core"]
		var cfg map[string]interface{}
		switch v := c.(type) {
		// Most run this is the type
		case []map[string]interface{}:
			cfg = v[0]
		// On the first run when the cfg is created it comes in as this type
		// for whatever reason
		case map[string]interface{}:
			cfg = v
		}
		host := cfg["host"].(string)
		hostURL, err := url.Parse(host)
		if err != nil {
			log.Fatal(err)
		}

		// See if we're in a git repo or if global is set to determine
		// if this should be a personal snippet
		if global || rn == "" {
			hostURL.Path = path.Join(hostURL.Path, "dashboard", "snippets")
		} else {
			hostURL.Path = path.Join(hostURL.Path, rn, "snippets")
		}

		if id > 0 {
			hostURL.Path = path.Join(hostURL.Path, strconv.FormatInt(id, 10))
		}

		err = browse(hostURL.String())
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	snippetCmd.AddCommand(snippetBrowseCmd)
    zsh.Gen(snippetBrowseCmd).PositionalCompletion(
      action.Remotes(),
      zsh.ActionCallback(func(args []string) zsh.Action {
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
      }),
    )
}
