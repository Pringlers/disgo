package voice

import (
	"github.com/disgoorg/log"
)

func DefaultConnConfig() *ConnConfig {
	return &ConnConfig{
		Logger:            log.Default(),
		GatewayCreateFunc: NewGateway,
		UDPConnCreateFunc: NewUDPConn,
	}
}

type ConnConfig struct {
	Logger log.Logger

	GatewayCreateFunc GatewayCreateFunc
	GatewayConfigOpts []GatewayConfigOpt

	UDPConnCreateFunc UDPConnCreateFunc
	UDPConnConfigOpts []UDPConnConfigOpt

	EventHandlerFunc EventHandlerFunc
}

type ConnConfigOpt func(connConfig *ConnConfig)

func (c *ConnConfig) Apply(opts []ConnConfigOpt) {
	for _, opt := range opts {
		opt(c)
	}
}

func WithConnLogger(logger log.Logger) ConnConfigOpt {
	return func(config *ConnConfig) {
		config.Logger = logger
	}
}

func WithConnGatewayCreateFunc(gatewayCreateFunc GatewayCreateFunc) ConnConfigOpt {
	return func(config *ConnConfig) {
		config.GatewayCreateFunc = gatewayCreateFunc
	}
}

func WithConnGatewayConfigOpts(opts ...GatewayConfigOpt) ConnConfigOpt {
	return func(config *ConnConfig) {
		config.GatewayConfigOpts = append(config.GatewayConfigOpts, opts...)
	}
}

func WithUDPConnCreateFunc(udpConnCreateFunc UDPConnCreateFunc) ConnConfigOpt {
	return func(config *ConnConfig) {
		config.UDPConnCreateFunc = udpConnCreateFunc
	}
}

func WithUDPConnConfigOpts(opts ...UDPConnConfigOpt) ConnConfigOpt {
	return func(config *ConnConfig) {
		config.UDPConnConfigOpts = append(config.UDPConnConfigOpts, opts...)
	}
}

func WithConnEventHandlerFunc(eventHandlerFunc EventHandlerFunc) ConnConfigOpt {
	return func(config *ConnConfig) {
		config.EventHandlerFunc = eventHandlerFunc
	}
}
