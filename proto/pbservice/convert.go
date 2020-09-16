package pbservice

import (
	"github.com/gogo/protobuf/types"
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

// TODO: handle this with mog
func MapStringServiceAddressToStructs(s map[string]ServiceAddress) map[string]structs.ServiceAddress {
	t := make(map[string]structs.ServiceAddress, len(s))
	for k, v := range s {
		t[k] = structs.ServiceAddress{Address: v.Address, Port: v.Port}
	}
	return t
}

// TODO: handle this with mog
func NewMapStringServiceAddressFromStructs(t map[string]structs.ServiceAddress) map[string]ServiceAddress {
	s := make(map[string]ServiceAddress, len(t))
	for k, v := range t {
		s[k] = ServiceAddress{Address: v.Address, Port: v.Port}
	}
	return s
}

// TODO: handle this with mog
func ExposePathSliceToStructs(s []ExposePath) []structs.ExposePath {
	t := make([]structs.ExposePath, len(s))
	for i, v := range s {
		t[i] = ExposePathToStructs(v)
	}
	return t
}

// TODO: handle this with mog
func NewExposePathSliceFromStructs(t []structs.ExposePath) []ExposePath {
	s := make([]ExposePath, len(t))
	for i, v := range t {
		s[i] = NewExposePathFromStructs(v)
	}
	return s
}

func MapStringInterfaceToStructs(s *types.Struct) map[string]interface{} {
	// TODO: how to convert these?
	return nil
}

func NewMapStringInterfaceFromStructs(t map[string]interface{}) *types.Struct {
	// TODO: how to convert these?
	return nil
}

// TODO: handle this with mog
func UpstreamsToStructs(s []Upstream) structs.Upstreams {
	t := make(structs.Upstreams, len(s))
	for i, v := range s {
		t[i] = UpstreamToStructs(v)
	}
	return t
}

// TODO: handle this with mog
func NewUpstreamsFromStructs(t structs.Upstreams) []Upstream {
	s := make([]Upstream, len(t))
	for i, v := range t {
		s[i] = NewUpstreamFromStructs(v)
	}
	return s
}

// TODO: handle this with mog
func CheckTypesToStructs(s []*CheckType) structs.CheckTypes {
	t := make(structs.CheckTypes, len(s))
	for i, v := range s {
		if v == nil {
			continue
		}
		newV := CheckTypeToStructs(*v)
		t[i] = &newV
	}
	return t
}

// TODO: handle this with mog
func NewCheckTypesFromStructs(t structs.CheckTypes) []*CheckType {
	s := make([]*CheckType, len(t))
	for i, v := range t {
		if v == nil {
			continue
		}
		newV := NewCheckTypeFromStructs(*v)
		s[i] = &newV
	}
	return s
}
