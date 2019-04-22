package main

import "time"

//https://www.pathofexile.com/character-window/get-characters?accountName=Acemi07
//https://pathofexile.com/character-window/get-items?character=qimmylee&accountName=hunjixin
type Ladder struct {
	total int64 `json:"total"`
	CachedSince time.Time `json:"cached_since"`
	Entries []Entry `json:"entries"`
}

type Entry struct {
	Rank int			`json:"rank"`
	Dead bool			`json:"dead"`
	Online bool			`json:"online"`
	Character Character `json:"character"`
	Account Account 	`json:"account"`
	Depth  Depth		`json:"depth"`
}

type Character struct {
	Name 			string				`json:"name"`
	League 			string				`json:"league"`
	CalssId 		int					`json:"classId"`
	AscendancyClass int					`json:"ascendancyClass"`
	Class 			string				`json:"class"`
	Level 			int					`json:"level"`
	Experience 		int64				`json:"experience"`
	LastActive 		bool				`json:"lastActive"`
}

type Account struct {
	Name string				`json:"name"`
	Realm string				`json:"realm"`
	Challenges Challenge	`json:"challenges"`
	Twitch Twitch			`json:"twitch"`
}

type Challenge struct {
	Total int64			`json:"total"`
}

type Twitch struct {
	Name string			`json:"name"`
}

type Depth struct {
	Default int			`json:"default"`
	Solo int			`json:"solo"`
}