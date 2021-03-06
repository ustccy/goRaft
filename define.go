package raft

// RunningStates if server is in these states, the server is considered as running
var RunningStates = map[string]bool{
	Leader:       true,
	Follower:     true,
	Candidate:    true,
	Snapshotting: true,
}

// DefaultConfig default configuration
var DefaultConfig = `{"logprefix":"raft-log-","bExpect":2,"name":"server0"}`
