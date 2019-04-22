package main

import (
	"encoding/json"
	"gopkg.in/mgo.v2/bson"
	"log"
)



func spiderStash(){
	id := getLatestId()
	for {
		url := "http://api.pathofexile.com/public-stash-tabs/?id="+id
		request := &SendRequest{
			Url: url,
			Err: make(chan error),
		}
		poolRequest(request)
		err := <- request.Err
		if err != nil {
			panic(err)
		}


		stashList := &StashList{}
		err = json.Unmarshal(request.Result, stashList)
		if err != nil {
			log.Fatal(err)
		}
		insertStash(stashList)
		id = stashList.NextChangeId
	}
}



func insertStash(stash *StashList){
	col := db.C("stash")
	err := col.Insert(stash)
	if err != nil {
		log.Fatal(err)
	}
}
func getLatestId() string{
	count, err  := stashCol.Count()
	if err != nil {
		log.Fatal(err)
	}
	stash := &StashList{}
	stashCol.Find(bson.M{}).Skip(count-1).One(stash)
	return stash.NextChangeId
}


