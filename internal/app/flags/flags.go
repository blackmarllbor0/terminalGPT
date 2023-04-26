package flags

import "flag"

type Flags struct{}

func NewFlags() *Flags {
	return &Flags{}
}

func (f *Flags) Parse() (query, apiKey string) {
	flag.StringVar(&query, "q", "", "")
	flag.StringVar(&apiKey, "api-key", "", "")

	flag.Parse()

	return
}

func (f *Flags) GetQuery() string {
	var query string
	flag.StringVar(&query, "q", "", "")

	flag.Parse()

	return query
}

func (f *Flags) SetApiKey() string {
	var apiKey string
	flag.StringVar(&apiKey, "api-key", "", "")

	flag.Parse()

	return apiKey
}
