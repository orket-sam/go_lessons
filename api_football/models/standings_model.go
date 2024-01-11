package model

// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    apidata, err := UnmarshalApidata(bytes)
//    bytes, err = apidata.Marshal()

import "encoding/json"

func UnmarshalApidata(data []byte) (Apidata, error) {
	var r Apidata
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Apidata) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type Apidata struct {
	Get        string        `json:"get"`
	Parameters Parameters    `json:"parameters"`
	Errors     []interface{} `json:"errors"`
	Results    int64         `json:"results"`
	Paging     Paging        `json:"paging"`
	Response   []Response    `json:"response"`
}

type Paging struct {
	Current int64 `json:"current"`
	Total   int64 `json:"total"`
}

type Parameters struct {
	League string `json:"league"`
	Season string `json:"season"`
}

type Response struct {
	League League `json:"league"`
}

type League struct {
	ID        int64        `json:"id"`
	Name      Name         `json:"name"`
	Country   string       `json:"country"`
	Logo      string       `json:"logo"`
	Flag      string       `json:"flag"`
	Season    int64        `json:"season"`
	Standings [][]Standing `json:"standings"`
}

type Standing struct {
	Rank        int64   `json:"rank"`
	Team        Team    `json:"team"`
	Points      int64   `json:"points"`
	GoalsDiff   int64   `json:"goalsDiff"`
	Group       Name    `json:"group"`
	Form        string  `json:"form"`
	Status      Status  `json:"status"`
	Description *string `json:"description"`
	All         All     `json:"all"`
	Home        All     `json:"home"`
	Away        All     `json:"away"`
	Update      string  `json:"update"`
}

type All struct {
	Played int64 `json:"played"`
	Win    int64 `json:"win"`
	Draw   int64 `json:"draw"`
	Lose   int64 `json:"lose"`
	Goals  Goals `json:"goals"`
}

type Goals struct {
	For     int64 `json:"for"`
	Against int64 `json:"against"`
}

type Team struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
	Logo string `json:"logo"`
}

type Name string

const (
	PremierLeague Name = "Premier League"
)

type Status string

const (
	Same Status = "same"
)
