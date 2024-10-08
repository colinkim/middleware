package params

import (
	serverconfig "github.com/reapchain/cosmos-sdk/server/config"
)

var (
	// BypassMinFeeMsgTypesKey defines the configuration key for the
	// BypassMinFeeMsgTypes value.
	// nolint: gosec
	BypassMinFeeMsgTypesKey = "bypass-min-fee-msg-types"

	// CustomConfigTemplate defines custom application configuration TOML
	// template. It extends the core SDK template.
	CustomConfigTemplate = serverconfig.DefaultConfigTemplate + `
###############################################################################
###                        Custom MiddleWare Configuration                        ###
###############################################################################
# bypass-min-fee-msg-types defines custom message types the operator may set that
# will bypass minimum fee checks during CheckTx.
#
# Example:
# ["/ibc.core.channel.v1.MsgRecvPacket", "/ibc.core.channel.v1.MsgAcknowledgement", ...]
bypass-min-fee-msg-types = [{{ range .BypassMinFeeMsgTypes }}{{ printf "%q, " . }}{{end}}]
`
)

// CustomAppConfig defines custom application configuration.
type CustomAppConfig struct {
	serverconfig.Config

	// BypassMinFeeMsgTypes defines custom message types the operator may set that
	// will bypass minimum fee checks during CheckTx.
	BypassMinFeeMsgTypes []string `mapstructure:"bypass-min-fee-msg-types"`
}
