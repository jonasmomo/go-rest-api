package main

import (
	"fmt"
	//"net/rpc"
)

var currentCallId int

var calls Calls

// Give us some seed data
func init() {
	RepoCreateCall(Call{Caller: "Test Caller 1"})
	RepoCreateCall(Call{Caller: "Test Caller 2"})
}

func RepoFindCall(id int) Call {
	for _, c := range calls {
		if c.Id == id {
			return c
		}
	}

	return Call{}
}

func RepoCreateCall(c Call) Call {
	currentCallId += 1
	c.Id = currentCallId
	calls = append(calls, c)
	return c
}

func RepoDestroyCall(id int) error {
	for i, t := range calls {
		if t.Id == id {
			calls = append(calls[:i], calls[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("Could not find Todo with id of %d to delete", id)
}
