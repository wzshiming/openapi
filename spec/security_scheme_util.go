package spec

const (
	basic  = "basic"
	apiKey = "apiKey"
	oauth2 = "oauth2"
)

// BasicAuth creates a basic auth security scheme
func BasicAuth() *SecurityScheme {
	ss := &SecurityScheme{}
	ss.Type = basic
	return ss
}

// APIKeyAuth creates an api key auth security scheme
func APIKeyAuth(fieldName, valueSource string) *SecurityScheme {
	ss := &SecurityScheme{}
	ss.Type = apiKey
	ss.Name = fieldName
	ss.In = valueSource
	return ss
}

// OAuth2Implicit creates an implicit flow oauth2 security scheme
func OAuth2Implicit(authorizationURL string) *SecurityScheme {
	ss := &SecurityScheme{}
	ss.Flows = &OAuthFlows{}
	ss.Type = oauth2
	ss.Flows.Implicit.AuthorizationURL = authorizationURL
	return ss
}

// OAuth2Password creates a password flow oauth2 security scheme
func OAuth2Password(tokenURL string) *SecurityScheme {
	ss := &SecurityScheme{}
	ss.Flows = &OAuthFlows{}
	ss.Type = oauth2
	ss.Flows.Password.TokenURL = tokenURL
	return ss
}

// OAuth2AuthorizationCode creates an application flow oauth2 security scheme
func OAuth2AuthorizationCode(tokenURL string) *SecurityScheme {
	ss := &SecurityScheme{}
	ss.Flows = &OAuthFlows{}
	ss.Type = oauth2
	ss.Flows.AuthorizationCode.TokenURL = tokenURL
	return ss
}

// OAuth2ClientCredentials creates an access token flow oauth2 security scheme
func OAuth2ClientCredentials(authorizationURL, tokenURL string) *SecurityScheme {
	ss := &SecurityScheme{}
	ss.Flows = &OAuthFlows{}
	ss.Type = oauth2
	ss.Flows.ClientCredentials.AuthorizationURL = authorizationURL
	ss.Flows.ClientCredentials.TokenURL = tokenURL
	return ss
}
