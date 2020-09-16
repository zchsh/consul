package pbservice

import structs "github.com/hashicorp/consul/agent/structs"

func HealthCheckToStructs(s HealthCheck) structs.HealthCheck {
	var t structs.HealthCheck
	t.Node = s.Node
	t.CheckID = s.CheckID
	t.Name = s.Name
	t.Status = s.Status
	t.Notes = s.Notes
	t.Output = s.Output
	t.ServiceID = s.ServiceID
	t.ServiceName = s.ServiceName
	t.ServiceTags = s.ServiceTags
	t.Type = s.Type
	t.Definition = HealthCheckDefinitionToStructs(s.Definition)
	t.EnterpriseMeta = EnterpriseMetaToStructs(s.EnterpriseMeta)
	t.RaftIndex = RaftIndexToStructs(s.RaftIndex)
	return t
}
func NewHealthCheckFromStructs(t structs.HealthCheck) HealthCheck {
	var s HealthCheck
	s.Node = t.Node
	s.CheckID = t.CheckID
	s.Name = t.Name
	s.Status = t.Status
	s.Notes = t.Notes
	s.Output = t.Output
	s.ServiceID = t.ServiceID
	s.ServiceName = t.ServiceName
	s.ServiceTags = t.ServiceTags
	s.Type = t.Type
	s.Definition = NewHealthCheckDefinitionFromStructs(t.Definition)
	s.EnterpriseMeta = NewEnterpriseMetaFromStructs(t.EnterpriseMeta)
	s.RaftIndex = NewRaftIndexFromStructs(t.RaftIndex)
	return s
}
func HealthCheckDefinitionToStructs(s HealthCheckDefinition) structs.HealthCheckDefinition {
	var t structs.HealthCheckDefinition
	t.HTTP = s.HTTP
	t.TLSSkipVerify = s.TLSSkipVerify
	t.Header = MapHeadersToStructs(s.Header)
	t.Method = s.Method
	t.Body = s.Body
	t.TCP = s.TCP
	t.Interval = s.Interval
	t.OutputMaxSize = uint(s.OutputMaxSize)
	t.Timeout = s.Timeout
	t.DeregisterCriticalServiceAfter = s.DeregisterCriticalServiceAfter
	t.ScriptArgs = s.ScriptArgs
	t.DockerContainerID = s.DockerContainerID
	t.Shell = s.Shell
	t.GRPC = s.GRPC
	t.GRPCUseTLS = s.GRPCUseTLS
	t.AliasNode = s.AliasNode
	t.AliasService = s.AliasService
	t.TTL = s.TTL
	return t
}
func NewHealthCheckDefinitionFromStructs(t structs.HealthCheckDefinition) HealthCheckDefinition {
	var s HealthCheckDefinition
	s.HTTP = t.HTTP
	s.TLSSkipVerify = t.TLSSkipVerify
	s.Header = NewMapHeadersFromStructs(t.Header)
	s.Method = t.Method
	s.Body = t.Body
	s.TCP = t.TCP
	s.Interval = t.Interval
	s.OutputMaxSize = uint32(t.OutputMaxSize)
	s.Timeout = t.Timeout
	s.DeregisterCriticalServiceAfter = t.DeregisterCriticalServiceAfter
	s.ScriptArgs = t.ScriptArgs
	s.DockerContainerID = t.DockerContainerID
	s.Shell = t.Shell
	s.GRPC = t.GRPC
	s.GRPCUseTLS = t.GRPCUseTLS
	s.AliasNode = t.AliasNode
	s.AliasService = t.AliasService
	s.TTL = t.TTL
	return s
}
func NodeToStructs(s Node) structs.Node {
	var t structs.Node
	t.ID = s.ID
	t.Node = s.Node
	t.Address = s.Address
	t.Datacenter = s.Datacenter
	t.TaggedAddresses = s.TaggedAddresses
	t.Meta = s.Meta
	t.RaftIndex = RaftIndexToStructs(s.RaftIndex)
	return t
}
func NewNodeFromStructs(t structs.Node) Node {
	var s Node
	s.ID = t.ID
	s.Node = t.Node
	s.Address = t.Address
	s.Datacenter = t.Datacenter
	s.TaggedAddresses = t.TaggedAddresses
	s.Meta = t.Meta
	s.RaftIndex = NewRaftIndexFromStructs(t.RaftIndex)
	return s
}
func NodeServiceToStructs(s NodeService) structs.NodeService {
	var t structs.NodeService
	t.Kind = s.Kind
	t.ID = s.ID
	t.Service = s.Service
	t.Tags = s.Tags
	t.Address = s.Address
	t.TaggedAddresses = MapStringServiceAddressToStructs(s.TaggedAddresses)
	t.Meta = s.Meta
	t.Port = s.Port
	t.Weights = WeightsPtrToStructs(s.Weights)
	t.EnableTagOverride = s.EnableTagOverride
	t.Proxy = ConnectProxyConfigToStructs(s.Proxy)
	t.Connect = ServiceConnectToStructs(s.Connect)
	t.LocallyRegisteredAsSidecar = s.LocallyRegisteredAsSidecar
	t.EnterpriseMeta = EnterpriseMetaToStructs(s.EnterpriseMeta)
	t.RaftIndex = RaftIndexToStructs(s.RaftIndex)
	return t
}
func NewNodeServiceFromStructs(t structs.NodeService) NodeService {
	var s NodeService
	s.Kind = t.Kind
	s.ID = t.ID
	s.Service = t.Service
	s.Tags = t.Tags
	s.Address = t.Address
	s.TaggedAddresses = NewMapStringServiceAddressFromStructs(t.TaggedAddresses)
	s.Meta = t.Meta
	s.Port = t.Port
	s.Weights = NewWeightsPtrFromStructs(t.Weights)
	s.EnableTagOverride = t.EnableTagOverride
	s.Proxy = NewConnectProxyConfigFromStructs(t.Proxy)
	s.Connect = NewServiceConnectFromStructs(t.Connect)
	s.LocallyRegisteredAsSidecar = t.LocallyRegisteredAsSidecar
	s.EnterpriseMeta = NewEnterpriseMetaFromStructs(t.EnterpriseMeta)
	s.RaftIndex = NewRaftIndexFromStructs(t.RaftIndex)
	return s
}
func ServiceDefinitionToStructs(s ServiceDefinition) structs.ServiceDefinition {
	var t structs.ServiceDefinition
	t.Kind = s.Kind
	t.ID = s.ID
	t.Name = s.Name
	t.Tags = s.Tags
	t.Address = s.Address
	t.TaggedAddresses = MapStringServiceAddressToStructs(s.TaggedAddresses)
	t.Meta = s.Meta
	t.Port = s.Port
	t.Check = CheckTypeToStructs(s.Check)
	t.Checks = CheckTypesToStructs(s.Checks)
	t.Weights = WeightsPtrToStructs(s.Weights)
	t.Token = s.Token
	t.EnableTagOverride = s.EnableTagOverride
	t.Proxy = ConnectProxyConfigPtrToStructs(s.Proxy)
	t.EnterpriseMeta = EnterpriseMetaToStructs(s.EnterpriseMeta)
	t.Connect = ServiceConnectPtrToStructs(s.Connect)
	return t
}
func NewServiceDefinitionFromStructs(t structs.ServiceDefinition) ServiceDefinition {
	var s ServiceDefinition
	s.Kind = t.Kind
	s.ID = t.ID
	s.Name = t.Name
	s.Tags = t.Tags
	s.Address = t.Address
	s.TaggedAddresses = NewMapStringServiceAddressFromStructs(t.TaggedAddresses)
	s.Meta = t.Meta
	s.Port = t.Port
	s.Check = NewCheckTypeFromStructs(t.Check)
	s.Checks = NewCheckTypesFromStructs(t.Checks)
	s.Weights = NewWeightsPtrFromStructs(t.Weights)
	s.Token = t.Token
	s.EnableTagOverride = t.EnableTagOverride
	s.Proxy = NewConnectProxyConfigPtrFromStructs(t.Proxy)
	s.EnterpriseMeta = NewEnterpriseMetaFromStructs(t.EnterpriseMeta)
	s.Connect = NewServiceConnectPtrFromStructs(t.Connect)
	return s
}
