package wahoo

import "reflect"

const (
	ProposalTag uint8 = iota
	VoteTag
	ElectTag
	ReadyTag
	DoneTag
	ReVoteTag
)

var proposal Block
var vote Vote
var elect Elect
var ready Ready
var done Done
var revote ReVote

var reflectedTypesMap = map[uint8]reflect.Type{
	ProposalTag: reflect.TypeOf(proposal),
	VoteTag:     reflect.TypeOf(vote),
	ElectTag:    reflect.TypeOf(elect),
	ReadyTag:    reflect.TypeOf(ready),
	DoneTag:     reflect.TypeOf(done),
	ReVoteTag:   reflect.TypeOf(revote),
}
