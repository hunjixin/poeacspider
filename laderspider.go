package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"
)

func spiderLader(){
	var limit = 200
	offset, err  := ladderCol.Count()
	if err != nil {
		log.Fatal(err)
	}

	for {
		url := "http://api.pathofexile.com/ladders/Synthesis?offset="+ strconv.FormatInt(int64(offset), 10) +"&limit=" +strconv.FormatInt(int64(limit), 10)
		request := &SendRequest{
			Url: url,
			Err: make(chan error),
		}
		poolRequest(request)
		err := <- request.Err
		if err != nil {
			panic(err)
		}
		ladder := &Ladder{}
		err = json.Unmarshal(request.Result, ladder)
		if err != nil {
			log.Fatal(err)
		}

		insertLader(ladder)
		if len(ladder.Entries) >= limit {
			offset += limit
		}else{
			fmt.Println("success to spider all ladder data")
			break
		}
		//spider items in character
		for _, entry := range ladder.Entries {
			spidercharacter(entry.Character.Name, entry.Account.Name)
		}
		fmt.Println("xxxx")
	}


}

func insertLader(stash *Ladder){
	err := ladderCol.Insert(stash)
	if err != nil {
		log.Fatal(err)
	}
}
func getOffset(){

}