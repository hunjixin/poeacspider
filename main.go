package main

import (
	"fmt"
	"gopkg.in/mgo.v2"
)

var db *mgo.Database
var stashCol *mgo.Collection
var ladderCol *mgo.Collection
var characterCol *mgo.Collection

func init(){
	session, err := mgo.Dial("mongodb://localhost:27017")
	if err != nil {
		panic(err)
	}

	db 				= 	session.DB("poedb")
	stashCol 		= 	db.C("stash")
	ladderCol 		= 	db.C("ladder")
	characterCol 	= 	db.C("character")
}

func main(){
	defer func() {
		if err:=recover();err!=nil{
			fmt.Println("app action err", err)
		}
	}()
	spiderLader()
}

