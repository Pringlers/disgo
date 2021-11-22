package handlers

import (
	"github.com/DisgoOrg/disgo/core"
	events2 "github.com/DisgoOrg/disgo/core/events"
	"github.com/DisgoOrg/disgo/discord"
)

// gatewayHandlerStageInstanceCreate handles core.GatewayEventMessageCreate
type gatewayHandlerStageInstanceCreate struct{}

// EventType returns the discord.GatewayEventTypeStageInstanceCreate
func (h *gatewayHandlerStageInstanceCreate) EventType() discord.GatewayEventType {
	return discord.GatewayEventTypeStageInstanceCreate
}

// New constructs a new payload receiver for the raw gateway event
func (h *gatewayHandlerStageInstanceCreate) New() interface{} {
	return &discord.StageInstance{}
}

// HandleGatewayEvent handles the specific raw gateway event
func (h *gatewayHandlerStageInstanceCreate) HandleGatewayEvent(bot *core.Bot, sequenceNumber int, v interface{}) {
	payload := *v.(*discord.StageInstance)

	bot.EventManager.Dispatch(&events2.StageInstanceCreateEvent{
		GenericStageInstanceEvent: &events2.GenericStageInstanceEvent{
			GenericEvent:    events2.NewGenericEvent(bot, sequenceNumber),
			StageInstanceID: payload.ID,
			StageInstance:   bot.EntityBuilder.CreateStageInstance(payload, core.CacheStrategyYes),
		},
	})
}
