package spec

// OAuthFlow Allows configuration of the supported OAuth Flows.
type OAuthFlows struct {

	// Configuration for the OAuth Implicit flow
	Implicit *OAuthFlow `json:"implicit,omitempty"`

	// Configuration for the OAuth Resource Owner Password flow
	Password *OAuthFlow `json:"password,omitempty"`

	// Configuration for the OAuth Client Credentials flow.
	// Previously called application in OpenAPI 2.0.
	ClientCredentials *OAuthFlow `json:"clientCredentials,omitempty"`

	// Configuration for the OAuth Authorization Code flow.
	// Previously called accessCode in OpenAPI 2.0.
	AuthorizationCode *OAuthFlow `json:"authorizationCode,omitempty"`
}

// OAuthFlow Configuration details for a supported OAuth Flow
type OAuthFlow struct {

	// oauth2 ("implicit", "authorizationCode")
	// REQUIRED.
	// The authorization URL to be used for this flow.
	// This MUST be in the form of a URL.
	AuthorizationURL string `json:"authorizationUrl,omitempty"`

	// oauth2 ("password", "clientCredentials", "authorizationCode")
	// REQUIRED.
	// The token URL to be used for this flow.
	// This MUST be in the form of a URL.
	TokenURL string `json:"tokenUrl,omitempty"`

	// oauth2	The URL to be used for obtaining refresh tokens.
	// This MUST be in the form of a URL.
	RefreshURL string `json:"refreshUrl,omitempty"`

	// oauth2
	// REQUIRED.
	// The available scopes for the OAuth2 security scheme.
	// A map between the scope name and a short description for it.
	Acopes map[string]string `json:"acopes,omitempty"`
}
