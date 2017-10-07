package main

import "net/http"

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		Index,
	},
	Route{
		"AgentIndex",
		"GET",
		"/agents",
		AgentIndex,
	},
	Route{
		"AgentCreate",
		"POST",
		"/agents",
		AgentCreate,
	},
	Route{
		"AgentShow",
		"GET",
		"/agents/{agentId}",
		AgentShow,
	},

	// calls
	Route{
		"CallIndex",
		"GET",
		"/calls",
		CallIndex,
	},
	/*Route{
		"CallCreate",
		"POST",
		"/calls",
		CallCreate,
	},*/
	Route{
		"CallShow",
		"GET",
		"/calls/{callId}",
		CallShow,
	},
}
