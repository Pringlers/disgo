package cache

import (
	"github.com/disgoorg/disgo/discord"
	"github.com/disgoorg/snowflake/v2"
)

// ChannelCache is a Cache for all channel types
type ChannelCache interface {
	Cache[discord.Channel]

	// GuildChannels returns all discord.GuildChannel in a guild and a bool indicating if it exists.
	GuildChannels(guildID snowflake.ID) []discord.GuildChannel

	// GuildThreadsInChannel returns all discord.GuildThread from the ChannelCache and a bool indicating if it exists.
	GuildThreadsInChannel(channelID snowflake.ID) []discord.GuildThread

	// GetGuildChannel returns a discord.GuildChannel from the ChannelCache and a bool indicating if it exists.
	GetGuildChannel(channelID snowflake.ID) (discord.GuildChannel, bool)

	// GetMessageChannel returns a discord.MessageChannel from the ChannelCache and a bool indicating if it exists.
	GetMessageChannel(channelID snowflake.ID) (discord.MessageChannel, bool)

	// GetGuildMessageChannel returns a discord.GuildMessageChannel from the ChannelCache and a bool indicating if it exists.
	GetGuildMessageChannel(channelID snowflake.ID) (discord.GuildMessageChannel, bool)

	// GetGuildThread returns a discord.GuildThread from the ChannelCache and a bool indicating if it exists.
	GetGuildThread(channelID snowflake.ID) (discord.GuildThread, bool)

	// GetGuildAudioChannel returns a discord.GetGuildAudioChannel from the ChannelCache and a bool indicating if it exists.
	GetGuildAudioChannel(channelID snowflake.ID) (discord.GuildAudioChannel, bool)

	// GetGuildTextChannel returns a discord.GuildTextChannel from the ChannelCache and a bool indicating if it exists.
	GetGuildTextChannel(channelID snowflake.ID) (discord.GuildTextChannel, bool)

	// GetDMChannel returns a discord.DMChannel from the ChannelCache and a bool indicating if it exists.
	GetDMChannel(channelID snowflake.ID) (discord.DMChannel, bool)

	// GetGuildVoiceChannel returns a discord.GuildVoiceChannel from the ChannelCache and a bool indicating if it exists.
	GetGuildVoiceChannel(channelID snowflake.ID) (discord.GuildVoiceChannel, bool)

	// GetGuildCategoryChannel returns a discord.GuildCategoryChannel from the ChannelCache and a bool indicating if it exists.
	GetGuildCategoryChannel(channelID snowflake.ID) (discord.GuildCategoryChannel, bool)

	// GetGuildNewsChannel returns a discord.GuildNewsChannel from the ChannelCache and a bool indicating if it exists.
	GetGuildNewsChannel(channelID snowflake.ID) (discord.GuildNewsChannel, bool)

	// GetGuildNewsThread returns a discord.GuildThread from the ChannelCache and a bool indicating if it exists.
	GetGuildNewsThread(channelID snowflake.ID) (discord.GuildThread, bool)

	// GetGuildPublicThread returns a discord.GuildThread from the ChannelCache and a bool indicating if it exists.
	GetGuildPublicThread(channelID snowflake.ID) (discord.GuildThread, bool)

	// GetGuildPrivateThread returns a discord.GuildThread from the ChannelCache and a bool indicating if it exists.
	GetGuildPrivateThread(channelID snowflake.ID) (discord.GuildThread, bool)

	// GetGuildStageVoiceChannel returns a discord.GuildStageVoiceChannel from the ChannelCache and a bool indicating if it exists.
	GetGuildStageVoiceChannel(channelID snowflake.ID) (discord.GuildStageVoiceChannel, bool)

	// GetGuildForumChannel returns a discord.GuildForumChannel from the ChannelCache and a bool indicating if it exists.
	GetGuildForumChannel(channelID snowflake.ID) (discord.GuildForumChannel, bool)
}

// NewChannelCache returns a new channelCacheImpl with the given flags and policy.
// channelCacheImpl is thread safe and can be used in multiple goroutines.
func NewChannelCache(flags Flags, policy Policy[discord.Channel]) ChannelCache {
	return &channelCacheImpl{
		Cache: NewCache[discord.Channel](flags, FlagChannels, policy),
	}
}

type channelCacheImpl struct {
	Cache[discord.Channel]
}

func (c *channelCacheImpl) GuildChannels(guildID snowflake.ID) []discord.GuildChannel {
	channels := c.FindAll(func(channel discord.Channel) bool {
		if ch, ok := channel.(discord.GuildChannel); ok {
			return ch.GuildID() == guildID
		}
		return false
	})
	guildChannels := make([]discord.GuildChannel, len(channels))
	for i, channel := range channels {
		guildChannels[i] = channel.(discord.GuildChannel)
	}
	return guildChannels
}

func (c *channelCacheImpl) GuildThreadsInChannel(channelID snowflake.ID) []discord.GuildThread {
	channels := c.FindAll(func(channel discord.Channel) bool {
		if thread, ok := channel.(discord.GuildThread); ok {
			return *thread.ParentID() == channelID
		}
		return false
	})
	threads := make([]discord.GuildThread, len(channels))
	for i, channel := range channels {
		threads[i] = channel.(discord.GuildThread)
	}
	return threads
}

func (c *channelCacheImpl) GetGuildChannel(channelID snowflake.ID) (discord.GuildChannel, bool) {
	if ch, ok := c.Get(channelID); ok {
		if cCh, ok := ch.(discord.GuildChannel); ok {
			return cCh, true
		}
	}
	return nil, false
}

func (c *channelCacheImpl) GetMessageChannel(channelID snowflake.ID) (discord.MessageChannel, bool) {
	if ch, ok := c.Get(channelID); ok {
		if cCh, ok := ch.(discord.MessageChannel); ok {
			return cCh, true
		}
	}
	return nil, false
}

func (c *channelCacheImpl) GetGuildMessageChannel(channelID snowflake.ID) (discord.GuildMessageChannel, bool) {
	if ch, ok := c.Get(channelID); ok {
		if chM, ok := ch.(discord.GuildMessageChannel); ok {
			return chM, true
		}
	}
	return nil, false
}

func (c *channelCacheImpl) GetGuildThread(channelID snowflake.ID) (discord.GuildThread, bool) {
	if ch, ok := c.Get(channelID); ok {
		if cCh, ok := ch.(discord.GuildThread); ok {
			return cCh, true
		}
	}
	return discord.GuildThread{}, false
}

func (c *channelCacheImpl) GetGuildAudioChannel(channelID snowflake.ID) (discord.GuildAudioChannel, bool) {
	if ch, ok := c.Get(channelID); ok {
		if cCh, ok := ch.(discord.GuildAudioChannel); ok {
			return cCh, true
		}
	}
	return nil, false
}

func (c *channelCacheImpl) GetGuildTextChannel(channelID snowflake.ID) (discord.GuildTextChannel, bool) {
	if ch, ok := c.Get(channelID); ok {
		if cCh, ok := ch.(discord.GuildTextChannel); ok {
			return cCh, true
		}
	}
	return discord.GuildTextChannel{}, false
}

func (c *channelCacheImpl) GetDMChannel(channelID snowflake.ID) (discord.DMChannel, bool) {
	if ch, ok := c.Get(channelID); ok {
		if cCh, ok := ch.(discord.DMChannel); ok {
			return cCh, true
		}
	}
	return discord.DMChannel{}, false
}

func (c *channelCacheImpl) GetGuildVoiceChannel(channelID snowflake.ID) (discord.GuildVoiceChannel, bool) {
	if ch, ok := c.Get(channelID); ok {
		if cCh, ok := ch.(discord.GuildVoiceChannel); ok {
			return cCh, true
		}
	}
	return discord.GuildVoiceChannel{}, false
}

func (c *channelCacheImpl) GetGuildCategoryChannel(channelID snowflake.ID) (discord.GuildCategoryChannel, bool) {
	if ch, ok := c.Get(channelID); ok {
		if cCh, ok := ch.(discord.GuildCategoryChannel); ok {
			return cCh, true
		}
	}
	return discord.GuildCategoryChannel{}, false
}

func (c *channelCacheImpl) GetGuildNewsChannel(channelID snowflake.ID) (discord.GuildNewsChannel, bool) {
	if ch, ok := c.Get(channelID); ok {
		if cCh, ok := ch.(discord.GuildNewsChannel); ok {
			return cCh, true
		}
	}
	return discord.GuildNewsChannel{}, false
}

func (c *channelCacheImpl) GetGuildNewsThread(channelID snowflake.ID) (discord.GuildThread, bool) {
	if ch, ok := c.GetGuildThread(channelID); ok && ch.Type() == discord.ChannelTypeGuildNewsThread {
		return ch, true
	}
	return discord.GuildThread{}, false
}

func (c *channelCacheImpl) GetGuildPublicThread(channelID snowflake.ID) (discord.GuildThread, bool) {
	if ch, ok := c.GetGuildThread(channelID); ok && ch.Type() == discord.ChannelTypeGuildPublicThread {
		return ch, true
	}
	return discord.GuildThread{}, false
}

func (c *channelCacheImpl) GetGuildPrivateThread(channelID snowflake.ID) (discord.GuildThread, bool) {
	if ch, ok := c.GetGuildThread(channelID); ok && ch.Type() == discord.ChannelTypeGuildPrivateThread {
		return ch, true
	}
	return discord.GuildThread{}, false
}

func (c *channelCacheImpl) GetGuildStageVoiceChannel(channelID snowflake.ID) (discord.GuildStageVoiceChannel, bool) {
	if ch, ok := c.Get(channelID); ok {
		if cCh, ok := ch.(discord.GuildStageVoiceChannel); ok {
			return cCh, true
		}
	}
	return discord.GuildStageVoiceChannel{}, false
}

func (c *channelCacheImpl) GetGuildForumChannel(channelID snowflake.ID) (discord.GuildForumChannel, bool) {
	if ch, ok := c.Get(channelID); ok {
		if cCh, ok := ch.(discord.GuildForumChannel); ok {
			return cCh, true
		}
	}
	return discord.GuildForumChannel{}, false
}
