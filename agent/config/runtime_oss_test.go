// +build !consulent

package config

var authMethodEntFields = `{}`

var entMetaJSON = `{}`

var entRuntimeConfigSanitize = `{}`

var entTokenConfigSanitize = `"EnterpriseConfig": {},`

func entFullRuntimeConfig(rt *RuntimeConfig) {}

var enterpriseReadReplicaWarnings = []string{enterpriseConfigKeyError{key: "read_replica"}.Error()}

var enterpriseConfigKeyWarnings = []string{
	enterpriseConfigKeyError{key: "read_replica"}.Error(),
	enterpriseConfigKeyError{key: "segment"}.Error(),
	enterpriseConfigKeyError{key: "segments"}.Error(),
	enterpriseConfigKeyError{key: "autopilot.redundancy_zone_tag"}.Error(),
	enterpriseConfigKeyError{key: "autopilot.upgrade_version_tag"}.Error(),
	enterpriseConfigKeyError{key: "autopilot.disable_upgrade_migration"}.Error(),
	enterpriseConfigKeyError{key: "dns_config.prefer_namespace"}.Error(),
	enterpriseConfigKeyError{key: "acl.msp_disable_bootstrap"}.Error(),
	enterpriseConfigKeyError{key: "acl.tokens.managed_service_provider"}.Error(),
	enterpriseConfigKeyError{key: "audit"}.Error(),
}
