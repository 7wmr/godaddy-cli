package dns

// Records retruned from GoDaddy.
type Records struct {
	Records []Record
}

// SetDomains update all records with domain.
func (r *Records) SetDomains(domain string) {
	for index := range r.Records {
		r.Records[index].Domain = domain
	}
}

// Record returned from GoDaddy.
type Record struct {
	Name   string `json:"name"`
	TTL    int    `json:"ttl"`
	Domain string `json:"-"`
	Value  string `json:"data"`
	Type   string `json:"type"`
}

// SetValue update the record value.
func (r *Record) SetValue(value string) {
	r.Value = value
}

// SetDomain update the domain value.
func (r *Record) SetDomain(domain string) {
	r.Domain = domain
}
