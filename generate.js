//str = str.replace(str[0],str[0].toUpperCase());
//str = str.replace(str[0],str[0].toLowerCase());

const MongoClient = require('mongodb').MongoClient;
const assert = require('assert');
const request = require('request');
const url = "mongodb://localhost:27017";
const dbName = "poe"

var poeDb

const GetDb = async function() {
    let client;
    try {
        // Use connect method to connect to the Server
        client = await MongoClient.connect(url, {useNewUrlParser: true});

        return client.db(dbName);
    } catch (err) {
        console.log(err.stack);
    }
};

const Insert = async function(data) {
    let client;
    try {
       var col = poeDb.collection("ladder")
        col.insertOne(data)
    } catch (err) {
        console.log(err.stack);
    }
};


const Requst = async function(url) {
    await new Promise((resolve,reject) => {
        request(url, function (error, response, body) {
            if (!error && response.statusCode == 200) {
                console.log(body)
                console.log(JSON.parse(body))
            }else {
                console.log(error)
                console.log(response)
                throw new Error()
            }

        })
    })
};

(async function(){
    poeDb = await GetDb()
    var data = await  Requst("http://api.pathofexile.com/ladders/Standard?offset=1&limit=2")
    console.log(data)
})()
