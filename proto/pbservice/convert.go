package pbservice

import (
	"github.com/hashicorp/consul/agent/structs"
	"github.com/hashicorp/consul/proto/pbcommon"
)

func RaftIndexToStructs(s pbcommon.RaftIndex) structs.RaftIndex {
	return structs.RaftIndex{
		CreateIndex: s.CreateIndex,
		ModifyIndex: s.ModifyIndex,
	}
}

func NewRaftIndexFromStructs(s structs.RaftIndex) pbcommon.RaftIndex {
	return pbcommon.RaftIndex{
		CreateIndex: s.CreateIndex,
		ModifyIndex: s.ModifyIndex,
	}
}

func MapHeadersToStructs(s map[string]HeaderValue) map[string][]string {
	t := make(map[string][]string, len(s))
	for k, v := range s {
		t[k] = v.Value
	}
	return t
}

func NewMapHeadersFromStructs(t map[string][]string) map[string]HeaderValue {
	s := make(map[string]HeaderValue, len(t))
	for k, v := range t {
		s[k] = HeaderValue{Value: v}
	}
	return s
}
