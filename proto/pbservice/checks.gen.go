package pbservice

import structs "github.com/hashicorp/consul/agent/structs"

func CheckTypeToStructs(s CheckType) structs.CheckType {
	var t structs.CheckType
	t.CheckID = s.CheckID
	t.Name = s.Name
	t.Status = s.Status
	t.Notes = s.Notes
	t.ScriptArgs = s.ScriptArgs
	t.HTTP = s.HTTP
	t.Header = MapHeadersToStructs(s.Header)
	t.Method = s.Method
	t.Body = s.Body
	t.TCP = s.TCP
	t.Interval = s.Interval
	t.AliasNode = s.AliasNode
	t.AliasService = s.AliasService
	t.DockerContainerID = s.DockerContainerID
	t.Shell = s.Shell
	t.GRPC = s.GRPC
	t.GRPCUseTLS = s.GRPCUseTLS
	t.TLSSkipVerify = s.TLSSkipVerify
	t.Timeout = s.Timeout
	t.TTL = s.TTL
	t.SuccessBeforePassing = int(s.SuccessBeforePassing)
	t.FailuresBeforeCritical = int(s.FailuresBeforeCritical)
	t.ProxyHTTP = s.ProxyHTTP
	t.ProxyGRPC = s.ProxyGRPC
	t.DeregisterCriticalServiceAfter = s.DeregisterCriticalServiceAfter
	t.OutputMaxSize = int(s.OutputMaxSize)
	return t
}
func NewCheckTypeFromStructs(t structs.CheckType) CheckType {
	var s CheckType
	s.CheckID = t.CheckID
	s.Name = t.Name
	s.Status = t.Status
	s.Notes = t.Notes
	s.ScriptArgs = t.ScriptArgs
	s.HTTP = t.HTTP
	s.Header = NewMapHeadersFromStructs(t.Header)
	s.Method = t.Method
	s.Body = t.Body
	s.TCP = t.TCP
	s.Interval = t.Interval
	s.AliasNode = t.AliasNode
	s.AliasService = t.AliasService
	s.DockerContainerID = t.DockerContainerID
	s.Shell = t.Shell
	s.GRPC = t.GRPC
	s.GRPCUseTLS = t.GRPCUseTLS
	s.TLSSkipVerify = t.TLSSkipVerify
	s.Timeout = t.Timeout
	s.TTL = t.TTL
	s.SuccessBeforePassing = int32(t.SuccessBeforePassing)
	s.FailuresBeforeCritical = int32(t.FailuresBeforeCritical)
	s.ProxyHTTP = t.ProxyHTTP
	s.ProxyGRPC = t.ProxyGRPC
	s.DeregisterCriticalServiceAfter = t.DeregisterCriticalServiceAfter
	s.OutputMaxSize = int32(t.OutputMaxSize)
	return s
}
