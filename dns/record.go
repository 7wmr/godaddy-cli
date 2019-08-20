package dns

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/7wmr/godaddy-cli/conf"
	"io/ioutil"
	"net/http"
)

// Record returned from GoDaddy.
type Record struct {
	Name  string `json:"name"`
	TTL   int    `json:"ttl"`
	Value string `json:"data"`
	Type  string `json:"type"`
}

// SetValue update the record value.
func (r *Record) SetValue(value string) {
	r.Value = value
}

// Records returned from GoDaddy.
type Records struct {
	Records []Record    `json:"records"`
	Domain  string      `json:"domain"`
	Config  conf.Config `json:"-"`
}

// GetRecords selected records.
func (r *Records) GetRecords(name string, rtype string) error {
	url := fmt.Sprintf("%s/v1/domains/%s/records/%s/%s", r.Config.GetAPI(), r.Domain, rtype, name)
	client := &http.Client{}

	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set(r.Config.GetAuth())

	res, err := client.Do(req)
	if err != nil {
		return err
	}
	if res.StatusCode != 200 {
		return errors.New(string(res.StatusCode))
	}

	body, _ := ioutil.ReadAll(res.Body)
	err = json.Unmarshal(body, &r.Records)
	if err != nil {
		return err
	}

	return nil
}

// SetRecord to be updated
func (r *Records) SetRecords() error {
	url := fmt.Sprintf("%s/v1/domains/%s/records/%s/%s", r.Config.GetAPI(), r.Domain, r.Records[0].Type, r.Records[0].Name)
	client := &http.Client{}

	data, err := json.Marshal(r.Records)
	req, _ := http.NewRequest("PUT", url, bytes.NewBuffer(data))
	req.Header.Set(r.Config.GetAuth())
	req.Header.Set("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		return err
	}

	if res.StatusCode != 200 {
		return errors.New(string(res.StatusCode))
	}

	return nil
}

// PublicAddress is details about the current public IP address
type PublicAddress struct {
	IP       string `json:"ip"`
	Hostname string `json:"hostname"`
	City     string `json:"city"`
	Region   string `json:"region"`
	Country  string `json:"country"`
	Location string `json:"loc"`
	Org      string `json:"org"`
	Postal   string `json:"postal"`
}

// GetPublicAddress get the current public IP address
func GetPublicAddress() (PublicAddress, error) {
	var ip PublicAddress
	res, err := http.Get("http://ipinfo.io/json")
	if err != nil {
		return ip, err
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return ip, err
	}

	err = json.Unmarshal(body, &ip)
	if err != nil {
		return ip, nil
	}

	return ip, nil
}
