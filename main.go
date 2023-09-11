package main

import (
	"log"

	"github.com/elastic/go-elasticsearch/v8"
)

func main() {

	cfg := elasticsearch.Config{
		Addresses: []string{
			"https://10.2.65.10:9200",
		},
		Username:               "elastic",
		Password:               "1+GyOz=CZIqmbK=5c6hl",
		CertificateFingerprint: "a9d7f76153a39c588171ba0ef0da8f1d2125e79a37abe55a944b9e6bd534b3aa",
	}
	es, _ := elasticsearch.NewClient(cfg)

	log.Println(elasticsearch.Version)
	log.Println(es.Info())
}
