// Package techconferences allows you to find your next tech conference using the Open-source and crowd-sourced conference website https://confs.tech/
package techconferences

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// ConferenceType is the type of conference you want to get the data for.
type ConferenceType int

const (
	// Android conference
	Android ConferenceType = iota
	// Clojure conference
	Clojure
	// CC++ conference
	CPP
	// CSS conference
	CSS
	// Datascience conference
	Data
	// DevOps conference
	DevOps
	// .Net conference
	DotNet
	// Elixer conference
	Elixer
	// General conference
	General
	// Golang conference
	Golang
	// GraphQL conference
	GraphQL
	// Groovy conference
	Groovy
	// iOS conference
	IOS
	// Java conference
	Java
	// JavaScript / Node.js conference
	JavaScript
	// Networking conference
	Networking
	// PHP conference
	PHP
	// Python conference
	Python
	// Ruby conference
	Ruby
	// Rust conference
	Rust
	// Scala conference
	Scala
	// Security conference
	Security
	// Technical Communications / Doc conference
	TechComm
	// UX conference
	UX
	// baseURL is the baseURL of the confs.tech GitHub repository
	baseURL = "https://raw.githubusercontent.com/tech-conferences/conference-data/master/conferences/%d/%s.json"
)

// types are the actual names of the files
var types = [...]string{
	"android",
	"clojure",
	"cpp",
	"css",
	"data",
	"devops",
	"dotnet",
	"elixer",
	"general",
	"golang",
	"graphql",
	"groovy",
	"ios",
	"java",
	"javascript",
	"networking",
	"php",
	"python",
	"ruby",
	"rust",
	"scala",
	"security",
	"tech-comm",
	"ux",
}

// String translates the conference type to a valid string
func (conf ConferenceType) String() string {
	return types[conf]
}

// Conferences is an array of conferences
type Conferences []Conference

// Conference is the conference you might want to go to
type Conference struct {
	Name       string  `json:"name"`
	City       string  `json:"city"`
	Country    string  `json:"country"`
	URL        string  `json:"url"`
	StartDate  string  `json:"startDate"`
	EndDate    string  `json:"endDate"`
	Twitter    *string `json:"twitter,omitempty"`
	CfpURL     *string `json:"cfpUrl,omitempty"`
	CfpEndDate *string `json:"cfpEndDate,omitempty"`
}

// GetConferences gets the conferences for a particular type and year (like DevOps 2019) and will return a list of conferences or an error
func GetConferences(confType ConferenceType, year int) (Conferences, error) {
	resp, err := http.Get(fmt.Sprintf(baseURL, year, confType))
	if err != nil {
		return nil, fmt.Errorf("error building HTTP request: %s", err.Error())
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("error getting conferences: HTTP Status Code %d", resp.StatusCode)
	}

	confData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading HTTP response: %s", err.Error())
	}

	conferences, err := unmarshalConferences(confData)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling JSON data: %s", err.Error())
	}

	return conferences, nil
}

func unmarshalConferences(data []byte) (Conferences, error) {
	var r Conferences
	err := json.Unmarshal(data, &r)
	return r, err
}
