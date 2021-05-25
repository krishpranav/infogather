package cmd

import (
	"fmt"
	"strconv"

	"github.com/krishpranav/infogather/pkg/noauth/userlookup"
	"github.com/spf13/cobra"
)

// noauthUserCmd represents the noauthUser command
var noauthUserCmd = &cobra.Command{
	Use:   "noauthUser <username>",
	Short: "Returns websites that have accounts registered with the specified username. Auth: NONE",
	Long:  `Returns websites that have accounts registered with the specified username. Auth: NONE`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			cmd.Usage()
		} else {
			x := userlookup.UserLookup(args[0])
			for _, v := range x {
				fmt.Println("[" + strconv.FormatBool(v.Valid) + "] - " + v.Domain)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(noauthUserCmd)
}
