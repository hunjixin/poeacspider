package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

var (
	Queue = make(chan *SendRequest, 4)
)

func init (){
	go process()
}
type SendRequest struct {
	Url		 string
	Err 	chan error
	Result 	[]byte
}

func process (){
	for {
		select {
		case request := <- Queue:
			for {
				resp, err := http.Get(request.Url)
				if err != nil {
					fmt.Println(err)
					fmt.Println("request , try again")
					time.Sleep(time.Second)
					continue
				}
				defer resp.Body.Close()
				body, err := ioutil.ReadAll(resp.Body)
				if err != nil {
					fmt.Println(err)
					fmt.Println("reade pool fail")
					time.Sleep(time.Second)
					continue
				}
				fmt.Println(string(body))
				res := &Resp{}
				err = json.Unmarshal(body, res)
				if err != nil {
					fmt.Println(err)
					fmt.Println("Unmarshal result fail")
					time.Sleep(time.Second)
					continue
				}
				if res.Error.Code != 0 {
					if res.Error.Message == "Forbidden" {
						request.Result = body
						request.Err <- errors.New("Forbidden")
						break
					}
					time.Sleep(time.Second)
					continue
				}

				request.Result = body
				request.Err <- nil
				break
			}
		default:
			time.Sleep(time.Second)
		}
	}
}

func poolRequest(request *SendRequest){
	Queue <- request
}

type Resp struct {
	Error 	 ErroDetail `json:"error"`
}

type ErroDetail struct {
	Code 	 int 		`json:"code"`
	Message	 string 	`json:"message"`
}