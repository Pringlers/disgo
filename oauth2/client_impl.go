package oauth2

import (
	"time"

	"github.com/disgoorg/disgo/discord"
	"github.com/disgoorg/disgo/rest"
	"github.com/disgoorg/snowflake/v2"
)

// New returns a new OAuth2 client with the given ID, secret and ConfigOpt(s).
func New(id snowflake.ID, secret string, opts ...ConfigOpt) Client {
	config := DefaultConfig()
	config.Apply(opts)

	return &clientImpl{id: id, secret: secret, config: *config}
}

type clientImpl struct {
	id     snowflake.ID
	secret string
	config Config
}

func (c *clientImpl) ID() snowflake.ID {
	return c.id
}

func (c *clientImpl) Secret() string {
	return c.secret
}

func (c *clientImpl) Rest() rest.OAuth2 {
	return c.config.OAuth2
}

func (c *clientImpl) SessionController() SessionController {
	return c.config.SessionController
}

func (c *clientImpl) StateController() StateController {
	return c.config.StateController
}

func (c *clientImpl) GenerateAuthorizationURL(redirectURI string, permissions discord.Permissions, guildID snowflake.ID, disableGuildSelect bool, scopes ...discord.OAuth2Scope) string {
	authURL, _ := c.GenerateAuthorizationURLState(redirectURI, permissions, guildID, disableGuildSelect, scopes...)
	return authURL
}

func (c *clientImpl) GenerateAuthorizationURLState(redirectURI string, permissions discord.Permissions, guildID snowflake.ID, disableGuildSelect bool, scopes ...discord.OAuth2Scope) (string, string) {
	state := c.StateController().GenerateNewState(redirectURI)
	values := discord.QueryValues{
		"client_id":     c.id,
		"redirect_uri":  redirectURI,
		"response_type": "code",
		"scope":         discord.JoinScopes(scopes),
		"state":         state,
	}

	if permissions != discord.PermissionsNone {
		values["permissions"] = permissions
	}
	if guildID != 0 {
		values["guild_id"] = guildID
	}
	if disableGuildSelect {
		values["disable_guild_select"] = true
	}
	return discord.AuthorizeURL(values), state
}

func (c *clientImpl) StartSession(code string, state string, identifier string, opts ...rest.RequestOpt) (Session, error) {
	redirectURI := c.StateController().ConsumeState(state)
	if redirectURI == "" {
		return nil, ErrStateNotFound
	}
	exchange, err := c.Rest().GetAccessToken(c.id, c.secret, code, redirectURI, opts...)
	if err != nil {
		return nil, err
	}
	return c.SessionController().CreateSessionFromResponse(identifier, *exchange), nil
}

func (c *clientImpl) RefreshSession(identifier string, session Session, opts ...rest.RequestOpt) (Session, error) {
	exchange, err := c.Rest().RefreshAccessToken(c.id, c.secret, session.RefreshToken(), opts...)
	if err != nil {
		return nil, err
	}
	return c.SessionController().CreateSessionFromResponse(identifier, *exchange), nil
}

func (c *clientImpl) GetUser(session Session, opts ...rest.RequestOpt) (*discord.OAuth2User, error) {
	if session.Expiration().Before(time.Now()) {
		return nil, ErrAccessTokenExpired
	}
	if !discord.HasScope(discord.OAuth2ScopeIdentify, session.Scopes()...) {
		return nil, ErrMissingOAuth2Scope(discord.OAuth2ScopeIdentify)
	}
	return c.Rest().GetCurrentUser(session.AccessToken(), opts...)
}

func (c *clientImpl) GetMember(session Session, guildID snowflake.ID, opts ...rest.RequestOpt) (*discord.Member, error) {
	if session.Expiration().Before(time.Now()) {
		return nil, ErrAccessTokenExpired
	}
	if !discord.HasScope(discord.OAuth2ScopeGuildsMembersRead, session.Scopes()...) {
		return nil, ErrMissingOAuth2Scope(discord.OAuth2ScopeGuildsMembersRead)
	}
	return c.Rest().GetCurrentMember(session.AccessToken(), guildID, opts...)
}

func (c *clientImpl) GetGuilds(session Session, opts ...rest.RequestOpt) ([]discord.OAuth2Guild, error) {
	if session.Expiration().Before(time.Now()) {
		return nil, ErrAccessTokenExpired
	}
	if !discord.HasScope(discord.OAuth2ScopeGuilds, session.Scopes()...) {
		return nil, ErrMissingOAuth2Scope(discord.OAuth2ScopeGuilds)
	}
	return c.Rest().GetCurrentUserGuilds(session.AccessToken(), 0, 0, 0, opts...)
}

func (c *clientImpl) GetConnections(session Session, opts ...rest.RequestOpt) ([]discord.Connection, error) {
	if session.Expiration().Before(time.Now()) {
		return nil, ErrAccessTokenExpired
	}
	if !discord.HasScope(discord.OAuth2ScopeConnections, session.Scopes()...) {
		return nil, ErrMissingOAuth2Scope(discord.OAuth2ScopeConnections)
	}
	return c.Rest().GetCurrentUserConnections(session.AccessToken(), opts...)
}
