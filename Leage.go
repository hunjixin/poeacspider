package main

import "time"

type Leage struct {
	Id 			string		`json:"id"`
	Realm 		string		`json:"realm"`
	Url 		string		`json:"url"`
	StartAt		time.Time	`json:"startAt"`
	EndAt  		time.Time	`json:"endAt"`
	DelveEvent 	string		`json:"delveEvent"`
}