package handlers

import (
	"github.com/DisgoOrg/disgo/core"
	events2 "github.com/DisgoOrg/disgo/core/events"
	"github.com/DisgoOrg/disgo/discord"
)

// gatewayHandlerGuildBanAdd handles core.GatewayEventGuildBanAdd
type gatewayHandlerGuildBanAdd struct{}

// EventType returns the core.GatewayGatewayEventType
func (h *gatewayHandlerGuildBanAdd) EventType() discord.GatewayEventType {
	return discord.GatewayEventTypeGuildBanAdd
}

// New constructs a new payload receiver for the raw gateway event
func (h *gatewayHandlerGuildBanAdd) New() interface{} {
	return &discord.GuildBanAddGatewayEvent{}
}

// HandleGatewayEvent handles the specific raw gateway event
func (h *gatewayHandlerGuildBanAdd) HandleGatewayEvent(bot *core.Bot, sequenceNumber int, v interface{}) {
	payload := *v.(*discord.GuildBanAddGatewayEvent)

	bot.EventManager.Dispatch(&events2.GuildBanEvent{
		GenericEvent: events2.NewGenericEvent(bot, sequenceNumber),
		GuildID:      payload.GuildID,
		User:         bot.EntityBuilder.CreateUser(payload.User, core.CacheStrategyNo),
	})
}
