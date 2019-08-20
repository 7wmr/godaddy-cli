package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/7wmr/godaddy-cli/dns"
	"github.com/spf13/cobra"
	"os"
)

var recordDomain string
var recordType string
var recordName string
var recordValue string

// setRecordCmd represents the set DNS record command
var setRecordCmd = &cobra.Command{
	Use:   "set-record",
	Short: "Set a DNS record for a domain.",
	Long:  `TBC`,
	Run: func(cmd *cobra.Command, args []string) {
		records := dns.Records{Domain: recordDomain, Config: config}
		err := records.GetRecords(recordName, recordType)
		if err != nil {
			fmt.Println("Error getting records ", err)
			os.Exit(1)
		}

		if recordValue == "" {
			ip, _ := dns.GetPublicAddress()
			records.Records[0].SetValue(ip.IP)
		}

		err = records.SetRecords()
		if err != nil {
			fmt.Println("Error updating records ", err)
			os.Exit(1)
		}

		data, _ := json.MarshalIndent(records, "", "\t")
		fmt.Println(string(data))
	},
}

func init() {
	domainCmd.AddCommand(setRecordCmd)

	setRecordCmd.Flags().StringVarP(&recordDomain, "domain", "d", "", "DNS Record Domain")
	setRecordCmd.MarkFlagRequired("domain")

	setRecordCmd.Flags().StringVarP(&recordType, "type", "t", "A", "DNS Record Type [A, CNAME]")

	setRecordCmd.Flags().StringVarP(&recordName, "name", "n", "", "DNS Record Name")
	setRecordCmd.MarkFlagRequired("name")

	setRecordCmd.Flags().StringVarP(&recordValue, "value", "v", "", "DNS Record Value")
}
