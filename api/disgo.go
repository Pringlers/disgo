package api

import (
	"encoding/json"
	"runtime"
	"strings"
	"time"
)

// Disgo is the main discord interface
type Disgo interface {
	Connect() error
	Close()
	Token() string
	Gateway() Gateway
	RestClient() RestClient
	Cache() Cache
	Intents() Intent
	ApplicationID() Snowflake
	SelfUser() *User
	SetSelfUser(User)
	EventManager() EventManager
	CreateCommand(string, string) GlobalCommandBuilder
	HeartbeatLatency() time.Duration
}

// GatewayEventProvider is used to add new raw gateway events
type GatewayEventProvider interface {
	New() interface{}
	Handle(Disgo, EventManager, interface{})
}

// EventListener is used to create new EventListener to listen to events
type EventListener interface {
	OnEvent(interface{})
}

// EventManager lets you listen for specific events triggered by raw gateway events
type EventManager interface {
	AddEventListeners(...EventListener)
	Handle(string, json.RawMessage)
	Dispatch(GenericEvent)
}

// GetOS returns the simplified version of the operating system for sending to Discord in the IdentifyCommandDataProperties.OS payload
func GetOS() string {
	OS := runtime.GOOS
	if strings.HasPrefix(OS, "windows") {
		return "windows"
	}
	if strings.HasPrefix(OS, "darwin") {
		return "darwin"
	}
	return "linux"
}
