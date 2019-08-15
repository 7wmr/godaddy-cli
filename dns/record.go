package dns

import (
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

// Get selected records.
func (r *Records) Get(name string, rtype string) error {
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
