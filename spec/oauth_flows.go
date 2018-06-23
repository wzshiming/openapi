package spec

// OAuthFlows Allows configuration of the supported OAuth Flows.
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
