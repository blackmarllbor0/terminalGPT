package flags

import "flag"

func Parse() (query, apiKey string) {
	flag.StringVar(&query, "q", "", "")
	flag.StringVar(&apiKey, "api-key", "", "")

	flag.Parse()

	return
}
