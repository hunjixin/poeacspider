package main

import (
	"encoding/json"
	"log"
)

type ItemsInCharacter struct {
	Items 		[]Item 		`json:"items"`
	Character 	Character 	`json:"character"`
	Account		string		`json:"account"`
}

func spidercharacter(characterName, accountName string){
	url :=  "https://pathofexile.com/character-window/get-items?character="+characterName+"&accountName="+accountName
	request := &SendRequest{
		Url: url,
		Err: make(chan error),
	}
	poolRequest(request)
	err := <- request.Err
	if err != nil {
		if err.Error() == "Forbidden" {
			return
		}
		panic(err)
	}


	itemsInCharacter := &ItemsInCharacter{}
	err = json.Unmarshal(request.Result, itemsInCharacter)
	if err != nil {
		log.Fatal(err)
	}
	itemsInCharacter.Account = accountName
	insertCharacter(itemsInCharacter)
}

func insertCharacter(items *ItemsInCharacter){
	err := characterCol.Insert(items)
	if err != nil {
		log.Fatal(err)
	}
}