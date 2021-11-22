package handlers

import (
	"github.com/DisgoOrg/disgo/core"
	events2 "github.com/DisgoOrg/disgo/core/events"
	"github.com/DisgoOrg/disgo/discord"
)

// gatewayHandlerGuildRoleCreate handles core.GuildRoleCreateGatewayEvent
type gatewayHandlerGuildRoleCreate struct{}

// EventType returns the core.GatewayGatewayEventType
func (h *gatewayHandlerGuildRoleCreate) EventType() discord.GatewayEventType {
	return discord.GatewayEventTypeGuildRoleCreate
}

// New constructs a new payload receiver for the raw gateway event
func (h *gatewayHandlerGuildRoleCreate) New() interface{} {
	return &discord.GuildRoleCreateGatewayEvent{}
}

// HandleGatewayEvent handles the specific raw gateway event
func (h *gatewayHandlerGuildRoleCreate) HandleGatewayEvent(bot *core.Bot, sequenceNumber int, v interface{}) {
	payload := *v.(*discord.GuildRoleCreateGatewayEvent)

	bot.EventManager.Dispatch(&events2.RoleCreateEvent{
		GenericRoleEvent: &events2.GenericRoleEvent{
			GenericEvent: events2.NewGenericEvent(bot, sequenceNumber),
			GuildID:      payload.GuildID,
			RoleID:       payload.Role.ID,
			Role:         bot.EntityBuilder.CreateRole(payload.GuildID, payload.Role, core.CacheStrategyYes),
		},
	})
}
