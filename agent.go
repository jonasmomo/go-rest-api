package main

import "time"

type Agent struct {
	Id        	int       `json:"id"`
	Name      	string    `json:"name"`
	Active 		bool      `json:"loggedIn"`
	Login       time.Time `json:"login"`
}

type Agents []Agent
