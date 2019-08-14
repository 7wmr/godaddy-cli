package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/7wmr/godaddy-cli/dns"
	"github.com/spf13/cobra"
	"io/ioutil"
	"net/http"
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
		url := fmt.Sprintf("%s/v1/domains/%s/records/%s/%s", config.GetAPI(), recordDomain, recordType, recordName)
		client := &http.Client{}

		req, _ := http.NewRequest("GET", url, nil)
		req.Header.Set(config.GetAuth())

		res, err := client.Do(req)
		if err != nil {
			fmt.Println(err)
			return
		}
		if res.StatusCode != 200 {
			fmt.Println(res.StatusCode)
			return
		}

		body, _ := ioutil.ReadAll(res.Body)
		var records dns.Records
		err = json.Unmarshal(body, &records.Records)
		if err != nil {
			fmt.Println(err)
			return
		}
		records.SetDomains(recordDomain)

		fmt.Println(records)
	},
}

func init() {
	domainCmd.AddCommand(setRecordCmd)

	setRecordCmd.Flags().StringVarP(&recordDomain, "domain", "d", "", "DNS Record Domain")
	setRecordCmd.MarkFlagRequired("domain")

	setRecordCmd.Flags().StringVarP(&recordType, "type", "t", "A", "DNS Record Type [A, CNAME]")
	setRecordCmd.MarkFlagRequired("type")

	setRecordCmd.Flags().StringVarP(&recordName, "name", "n", "", "DNS Record Name")
	setRecordCmd.MarkFlagRequired("name")

	setRecordCmd.Flags().StringVarP(&recordValue, "value", "v", "", "DNS Record Value")
	setRecordCmd.MarkFlagRequired("value")

}
