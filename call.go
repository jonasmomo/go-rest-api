package main

import "time"

type Call struct {
	Id        		int       	`json:"id"`
	Caller    		string    	`json:"name"`
	Answered 		bool     	`json:"completed"`
	AnswerTime	 	time.Time 	`json:"due"`
}

type Calls []Call
