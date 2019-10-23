// Package techconferences allows you to find your next tech conference using the Open-source and crowd-sourced conference website https://confs.tech/
package techconferences

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

// ConferenceType is the type of conference you want to get the data for.
type ConferenceType int

const (
	// Android conference
	Android ConferenceType = iota
	// Clojure conference
	Clojure
	// C++ conference
	CPP
	// CSS conference
	CSS
	// Data Science conference
	Data
	// DevOps conference
	DevOps
	// DotNet conference
	DotNet
	// Elixer conference
	Elixer
	// Elm conference
	Elm
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
	// Leadership conference
	Leadership
	// Networking conference
	Networking
	// PHP conference
	PHP
	// Product Management conference
	Product
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

const dateLayout = "2006-01-02"

// Time is a custom struct to unmarshal the dates into a proper format
type Time struct {
	time.Time
}

// UnmarshalJSON unmarshals the Time struct
func (ct *Time) UnmarshalJSON(b []byte) (err error) {
	s := strings.Trim(string(b), "\"")
	if s == "null" {
		ct.Time = time.Time{}
		return
	}
	ct.Time, err = time.Parse(dateLayout, s)
	return
}

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
	"elm",
	"general",
	"golang",
	"graphql",
	"groovy",
	"ios",
	"java",
	"javascript",
	"leadership",
	"networking",
	"php",
	"product",
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
	StartDate  Time    `json:"startDate"`
	EndDate    Time    `json:"endDate"`
	Twitter    *string `json:"twitter,omitempty"`
	CfpURL     *string `json:"cfpUrl,omitempty"`
	CfpEndDate *Time   `json:"cfpEndDate,omitempty"`
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
