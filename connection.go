// Copyright 2015 Matthew Setter. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package callfire provides structs and functions for
// interacting with the CallFire API

package callfire

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// A struct which stores options for making a generic search request.
type SearchOptions struct {
	MaxResults  int `url:"MaxResults,omitempty"`
	FirstResult int `url:"FirstResult,omitempty"`
}

// Query options for the numbers search request.
// For more information checkout
//   https://github.com/google/go-querystring/blob/master/query/encode.go
type RequestOptions struct {
	Count      int    `url:"count"` // How many records to return
	City       string `url:"city,omitempty"`
	State      string `url:"state,omitempty"`
	Country    string `url:"country,omitempty"`
	Prefix     string `url:"prefix,omitempty"`
	Zipcode    string `url:"zipcode,omitempty"`
	Lata       string `url:"lata,omitempty"`
	RateCenter string `url:"ratecenter,omitempty"` // The rate center to use
	Latitude   string `url:"latitude,omitempty"`
	Longitude  string `url:"longitude,omitempty"`
	TimeZone   string `url:"timezone,omitempty"`
	TollFree   bool   `url:"tollfree,omitempty"`
	MaxResults int    `url:"maxresults,omitempty"`
}

// A simple struct to store the core request parameters.
type CallFireRequestOptions struct {
	Url            string
	Login          string
	Secret         string
	RequestOptions string
	ReqType        string
}

// Utility function to initialise a http Client object.
// It takes an http Client and Request object, makes
// a request and returns both objects.
func InitHttpClient(options CallFireRequestOptions) (*http.Client, *http.Request) {
	// Initialise a new http client
	client := &http.Client{}

	// Setup the request object
	request, err := http.NewRequest(
		options.ReqType, options.Url+options.RequestOptions, nil,
	)

	if err != nil {
		log.Fatal(err)
	}

	// Set the authentication credentials
	request.SetBasicAuth(options.Login, options.Secret)

	return client, request
}

// RunRequest simplifies making a request to the CallFire REST API.
// A set of options are supplied in requestOptions, which are then
// used to initialise the request.
func RunRequest(requestOptions CallFireRequestOptions) {
	client, req := InitHttpClient(requestOptions)
	resp, err := client.Do(req)

	if err != nil {
		log.Fatal(err)
	}

	// Retrieve the body of the response
	body, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()

	if err != nil {
		log.Fatal(err)
	}

	// Dump the response
	fmt.Printf("%s", body)
}
