package main

import "time"

type Call struct {
	Id        		int       	`json:"id"`
	Caller    		string    	`json:"name"`
	Answered 		bool     	`json:"answered"`
	AnswerTime	 	time.Time 	`json:"timeAnswered"`
}

type Calls []Call
