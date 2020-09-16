package pbservice

import structs "github.com/hashicorp/consul/agent/structs"

func ConnectProxyConfigToStructs(s ConnectProxyConfig) structs.ConnectProxyConfig {
	var t structs.ConnectProxyConfig
	t.DestinationServiceName = s.DestinationServiceName
	t.DestinationServiceID = s.DestinationServiceID
	t.LocalServiceAddress = s.LocalServiceAddress
	t.LocalServicePort = s.LocalServicePort
	t.Config = MapStringInterfaceToStructs(s.Config)
	t.Upstreams = UpstreamsToStructs(s.Upstreams)
	t.MeshGateway = MeshGatewayConfigToStructs(s.MeshGateway)
	t.Expose = ExposeConfigToStructs(s.Expose)
	return t
}
func NewConnectProxyConfigFromStructs(t structs.ConnectProxyConfig) ConnectProxyConfig {
	var s ConnectProxyConfig
	s.DestinationServiceName = t.DestinationServiceName
	s.DestinationServiceID = t.DestinationServiceID
	s.LocalServiceAddress = t.LocalServiceAddress
	s.LocalServicePort = t.LocalServicePort
	s.Config = NewMapStringInterfaceFromStructs(t.Config)
	s.Upstreams = NewUpstreamsFromStructs(t.Upstreams)
	s.MeshGateway = NewMeshGatewayConfigFromStructs(t.MeshGateway)
	s.Expose = NewExposeConfigFromStructs(t.Expose)
	return s
}
func ExposeConfigToStructs(s ExposeConfig) structs.ExposeConfig {
	var t structs.ExposeConfig
	t.Checks = s.Checks
	t.Paths = ExposePathSliceToStructs(s.Paths)
	return t
}
func NewExposeConfigFromStructs(t structs.ExposeConfig) ExposeConfig {
	var s ExposeConfig
	s.Checks = t.Checks
	s.Paths = NewExposePathSliceFromStructs(t.Paths)
	return s
}
func ExposePathToStructs(s ExposePath) structs.ExposePath {
	var t structs.ExposePath
	t.ListenerPort = int(s.ListenerPort)
	t.Path = s.Path
	t.LocalPathPort = int(s.LocalPathPort)
	t.Protocol = s.Protocol
	t.ParsedFromCheck = s.ParsedFromCheck
	return t
}
func NewExposePathFromStructs(t structs.ExposePath) ExposePath {
	var s ExposePath
	s.ListenerPort = int32(t.ListenerPort)
	s.Path = t.Path
	s.LocalPathPort = int32(t.LocalPathPort)
	s.Protocol = t.Protocol
	s.ParsedFromCheck = t.ParsedFromCheck
	return s
}
func MeshGatewayConfigToStructs(s MeshGatewayConfig) structs.MeshGatewayConfig {
	var t structs.MeshGatewayConfig
	t.Mode = s.Mode
	return t
}
func NewMeshGatewayConfigFromStructs(t structs.MeshGatewayConfig) MeshGatewayConfig {
	var s MeshGatewayConfig
	s.Mode = t.Mode
	return s
}
func ServiceConnectToStructs(s ServiceConnect) structs.ServiceConnect {
	var t structs.ServiceConnect
	t.Native = s.Native
	t.SidecarService = ServiceDefinitionPtrToStructs(s.SidecarService)
	return t
}
func NewServiceConnectFromStructs(t structs.ServiceConnect) ServiceConnect {
	var s ServiceConnect
	s.Native = t.Native
	s.SidecarService = NewServiceDefinitionPtrFromStructs(t.SidecarService)
	return s
}
func UpstreamToStructs(s Upstream) structs.Upstream {
	var t structs.Upstream
	t.DestinationType = s.DestinationType
	t.DestinationNamespace = s.DestinationNamespace
	t.DestinationName = s.DestinationName
	t.Datacenter = s.Datacenter
	t.LocalBindAddress = s.LocalBindAddress
	t.LocalBindPort = s.LocalBindPort
	t.Config = MapStringInterfaceToStructs(s.Config)
	t.MeshGateway = MeshGatewayConfigToStructs(s.MeshGateway)
	return t
}
func NewUpstreamFromStructs(t structs.Upstream) Upstream {
	var s Upstream
	s.DestinationType = t.DestinationType
	s.DestinationNamespace = t.DestinationNamespace
	s.DestinationName = t.DestinationName
	s.Datacenter = t.Datacenter
	s.LocalBindAddress = t.LocalBindAddress
	s.LocalBindPort = t.LocalBindPort
	s.Config = NewMapStringInterfaceFromStructs(t.Config)
	s.MeshGateway = NewMeshGatewayConfigFromStructs(t.MeshGateway)
	return s
}
