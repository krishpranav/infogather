package cmd

import (
	"fmt"
	"log"
	"strconv"

	"github.com/krishpranav/infogather/pkg/noauth/ip"
	"github.com/spf13/cobra"
)

// noauthIpCmd represents the noauthIp command
var noauthIpCmd = &cobra.Command{
	Use:   "noauthIp <ipv4 address>",
	Short: "Returns public information regarding an IPV4 Address. Auth: NONE",
	Long:  `Returns public information regarding an IPV4 Address. Auth: NONE`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			cmd.Usage()
		} else {
			x, err := ip.IPLookup(args[0])
			if err != nil {
				log.Panic(err)
			} else {
				if x.Status == "fail" {
					fmt.Println("Request failed.")
				} else {
					fmt.Println("Status: " + x.Status)
					fmt.Println("Continent: " + x.Continent)
					fmt.Println("Country: " + x.Country)
					fmt.Println("Region Name: " + x.RegionName)
					fmt.Println("City: " + x.City)
					fmt.Println("District: " + x.District)
					fmt.Println("Zip: " + x.Zip)
					fmt.Println("Latitude: " + strconv.FormatFloat(x.Lat, 'f', 5, 64))
					fmt.Println("Longitude: " + strconv.FormatFloat(x.Lon, 'f', 5, 64))
					fmt.Println("Timezone: " + x.Timezone)
					fmt.Println("Currency: " + x.Currency)
					fmt.Println("ISP: " + x.Isp)
					fmt.Println("Org: " + x.Org)
					fmt.Println("As: " + x.As)
					fmt.Println("Asname: " + x.Asname)
					fmt.Println("Reverse: " + x.Reverse)
					fmt.Println("Mobile: " + strconv.FormatBool(x.Mobile))
					fmt.Println("Proxy: " + strconv.FormatBool(x.Proxy))
					fmt.Println("Hosting: " + strconv.FormatBool(x.Hosting))
					fmt.Println("Query: " + x.Query)
				}
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(noauthIpCmd)
}
