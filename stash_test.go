package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"testing"
)

func TestStash(t *testing.T) {
	bytes, _ := ioutil.ReadFile("stash.json")
	stashList := &StashList{}
	err := json.Unmarshal(bytes, stashList)
	if err != nil {
		log.Fatal(err)
	}
}