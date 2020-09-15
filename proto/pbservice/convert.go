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
