package main

import (
	"fmt"
	//"net/rpc"
)

var currentAgentId int

var agents Agents

// Give us some seed data
func init() {
	RepoCreateTodo(Agent{Name: "Jonas"})
	RepoCreateTodo(Agent{Name: "Kenan"})
}

func RepoFindAgent(id int) Agent {
	for _, a := range agents {
		if a.Id == id {
			return a
		}
	}
	// return empty Todo if not found
	return Agent{}
}

//this is bad, I don't think it passes race condtions
func RepoCreateTodo(a Agent) Agent {
	currentAgentId += 1
	a.Id = currentAgentId
	agents = append(agents, a)
	return a
}

func RepoDestroyAgent(id int) error {
	for i, a := range agents {
		if a.Id == id {
			agents = append(agents[:i], agents[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("Could not find Agent with id of %d to delete", id)
}
