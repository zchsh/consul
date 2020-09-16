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

// TODO: handle this with mog, once mog handles pointers
func WeightsPtrToStructs(s *Weights) *structs.Weights {
	if s == nil {
		return nil
	}
	var t structs.Weights
	t.Passing = s.Passing
	t.Warning = s.Warning
	return &t
}

// TODO: handle this with mog, once mog handles pointers
func NewWeightsPtrFromStructs(t *structs.Weights) *Weights {
	if t == nil {
		return nil
	}
	var s Weights
	s.Passing = t.Passing
	s.Warning = t.Warning
	return &s
}
