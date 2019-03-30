package spec

const (
	SecurityHTTP          = "http"
	SecurityAPIKey        = "apiKey"
	SecurityOAuth2        = "oauth2"
	SecurityOpenIDConnect = "openIdConnect"
)

// BasicAuth creates a basic auth security scheme
func BasicAuth() *SecurityScheme {
	ss := &SecurityScheme{}
	ss.Type = SecurityHTTP
	ss.Scheme = "basic"
	return ss
}

// BearerAuth creates a bearer auth security scheme
func BearerAuth(bearerFormat string) *SecurityScheme {
	ss := &SecurityScheme{}
	ss.Type = SecurityHTTP
	ss.Scheme = "bearer"
	ss.BearerFormat = bearerFormat
	return ss
}

// APIKeyAuth creates an api key auth security scheme
func APIKeyAuth(fieldName, valueSource string) *SecurityScheme {
	ss := &SecurityScheme{}
	ss.Type = SecurityAPIKey
	ss.Name = fieldName
	ss.In = valueSource
	return ss
}

// OAuth2Implicit creates an implicit flow oauth2 security scheme
func OAuth2Implicit(authorizationURL string) *SecurityScheme {
	ss := &SecurityScheme{}
	ss.Flows = &OAuthFlows{}
	ss.Type = SecurityOAuth2
	ss.Flows.Implicit.AuthorizationURL = authorizationURL
	return ss
}

// OAuth2Password creates a password flow oauth2 security scheme
func OAuth2Password(tokenURL string) *SecurityScheme {
	ss := &SecurityScheme{}
	ss.Flows = &OAuthFlows{}
	ss.Type = SecurityOAuth2
	ss.Flows.Password.TokenURL = tokenURL
	return ss
}

// OAuth2AuthorizationCode creates an application flow oauth2 security scheme
func OAuth2AuthorizationCode(tokenURL string) *SecurityScheme {
	ss := &SecurityScheme{}
	ss.Flows = &OAuthFlows{}
	ss.Type = SecurityOAuth2
	ss.Flows.AuthorizationCode.TokenURL = tokenURL
	return ss
}

// OAuth2ClientCredentials creates an access token flow oauth2 security scheme
func OAuth2ClientCredentials(authorizationURL, tokenURL string) *SecurityScheme {
	ss := &SecurityScheme{}
	ss.Flows = &OAuthFlows{}
	ss.Type = SecurityOAuth2
	ss.Flows.ClientCredentials.AuthorizationURL = authorizationURL
	ss.Flows.ClientCredentials.TokenURL = tokenURL
	return ss
}
