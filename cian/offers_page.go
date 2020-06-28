package cian

import (
	"bytes"
	"encoding/json"
	"net/http"
)

func SearchOffersDesktop(req []byte) (*OffersResponse, error) {
	url := "https://api.cian.ru/search-offers/v2/search-offers-desktop/"
	resp, err := http.Post(url, "application/json", bytes.NewReader(req))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var obj OffersResponse
	err = json.NewDecoder(resp.Body).Decode(&obj)
	if err != nil {
		return nil, err
	}

	return &obj, nil
}