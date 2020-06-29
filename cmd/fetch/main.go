package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"time"

	"github.com/petuhovskiy/cian-search/cian"
	"github.com/petuhovskiy/cian-search/nometa"
)

func main() {
	var query string
	var result string

	flag.StringVar(&query, "query", "query.json", "path to json query file")
	flag.StringVar(&result, "result", "result.json", "file to save request output")
	flag.Parse()

	queryBody, err := ioutil.ReadFile(query)
	if err != nil {
		panic(err)
	}

	pageParam := []byte("$page")

	hasPageParam := bytes.Contains(queryBody, pageParam)
	if hasPageParam {
		log.Println("Found page param in query, activating paging strategy")
	}

	var offersList []nometa.Offer
	for page := 1; ; page++ {
		log.Printf("Sleeping, just in case")
		time.Sleep(time.Second)

		currentQuery := queryBody

		if hasPageParam {
			strPage := []byte(fmt.Sprintf("%d", page))
			currentQuery = bytes.Replace(currentQuery, pageParam, strPage, 1)
			log.Printf("Selecting page %d", page)
		}

		offers, err := cian.SearchOffersDesktop(currentQuery)
		if err != nil {
			panic(err)
		}

		hasAny := false
		for _, o := range offers.Data.OffersSerialized {
			offersList = append(offersList, nometa.ConvertOffer(o))
			hasAny = true
		}

		if !hasPageParam || !hasAny {
			break
		}
	}

	pretty, err := json.MarshalIndent(offersList, "", "\t")
	if err != nil {
		panic(err)
	}

	err = ioutil.WriteFile(result, pretty, 0666)
	if err != nil {
		panic(err)
	}
}
