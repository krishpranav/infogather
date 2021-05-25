package cmd

import (
	"fmt"
	"log"
	"strconv"

	"github.com/krishpranav/infogather/pkg/noauth/discord"
	"github.com/spf13/cobra"
)

// noauthDiscordTokenCmd represents the noauthDiscordToken command
var noauthDiscordTokenCmd = &cobra.Command{
	Use:   "noauthDiscordToken <token>",
	Short: "Returns information regarding the passed token. Auth: NONE",
	Long:  `Returns information regarding the passed token. Auth: NONE`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			cmd.Usage()
		} else {
			x, err := discord.TokenLookup(args[0])
			if err != nil {
				log.Panic(err)
			} else {
				fmt.Println("ID: " + x.ID)
				fmt.Println("Username: " + x.Username)
				fmt.Println("Avatar: " + x.Avatar)
				fmt.Println("Discriminator: " + x.Discriminator)
				fmt.Println("Public Flags: " + strconv.Itoa(x.PublicFlags))
				fmt.Println("Flags: " + strconv.Itoa(x.Flags))
				fmt.Println("Purchased Flags: " + strconv.Itoa(x.PurchasedFlags))
				fmt.Println("Locale: " + x.Locale)
				fmt.Println("NSFW Allowed: " + strconv.FormatBool(x.NsfwAllowed))
				fmt.Println("MFA Enabled: " + strconv.FormatBool(x.MfaEnabled))
				fmt.Println("Email: " + x.Email)
				fmt.Println("Verified: " + strconv.FormatBool(x.Verified))
				fmt.Println("Phone: " + x.Phone)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(noauthDiscordTokenCmd)
}
