package model

import (
	"fmt"
)

const (
	BLANK                              = ""
	Post                               = "POST"
	DefaultToken                       = ""
	Endpoint                           = ""
	AcceptKey                          = "Accept"
	ContentTypeKey                     = "Content-Type"
	AcceptValue                        = "application/json"
	ContentTypeValue                   = "application/json"
	TraceIdKey                         = "Trace-Id"
	ClientInfoValue                    = ""
	DebugTraceValue                    = "true"
	ProgressiveDdeploymentVariantValue = "template"
	OpaqueFlag                         = "o"
	OpaqueInfo                         = "an opaque token input provided by the user"
	TemplateFlag                       = "t"
	TemplateInfo                       = "adding the template field"
	PartnerAccIdFlag                   = "p"
	PartnerAccIdInfo                   = "partnet account id for different sessions"
	AuthorizationKey                   = "Authorization"
	AuthorizationBasicValue            = "Basic _"
	CookieKey                          = "Cookie"
	CookieValue                        = ""
	AppJwtEndpoint                     = ""
	OpaqueEndpoint                     = ""
	MessageIdKey                       = "Message-ID"
	ParentMessageIdKey                 = "Parent-Message-ID"
	DeviceUserAgentIdKey               = "Device-User-Agent-ID"
	XDIdKey                            = "XD-ID"
)

type KeepAliveStrategy struct {
	Lifecycle string `json:"lifecycle"`
	Expiry    string `json:"expiry"`
}

type State struct {
	KeepAliveStrategy   KeepAliveStrategy `json:"keepAliveStrategy"`
	AuthenticationState string            `json:"authenticationState"`
}

type CustomClaims struct {
	IdentityProvider string `json:"IDENTITY_PROVIDER"`
	CredentialsGiven string `json:"credentialsGiven"`
	SoftAccount      string `json:"softAccount"`
	SubjectTuid      string `json:"subjectTuid"`
	PrincipalTuid    string `json:"principalTuid"`
}

type Data struct {
	ActorID          string       `json:"actorId"`
	SubjectID        string       `json:"subjectId"`
	IDP              string       `json:"idp"`
	States           []State      `json:"states"`
	CustomClaims     CustomClaims `json:"customClaims"`
	Issuer           string       `json:"issuer"`
	Audience         string       `json:"audience"`
	ClientID         string       `json:"clientId"`
	AuthTime         string       `json:"authTime"`
	Scope            string       `json:"scope"`
	PartnerAccountID string       `json:"partnerAccountId"`
}

type JSONStringer interface {
	PrintEGJwt() string
	PrintJwtBearer() string
}

type Payload struct {
	OpaqueToken string `json:"opaqueToken"`
}

type Result struct {
	EncodedJwt string `json:"encodedJwt"`
}

type OpaqueResult struct {
	OpaqueToken string `json:"opaqueToken"`
}

type HeaderJson struct {
	ClientInfo    string `json:"Client-Info"`
	Authorization string `json:"Authorization"`
	DebugTrace    string `json:"Debug-Trace"`
}

type TemplatedHeaderJson struct {
	HeaderJson
	ProgressiveDdeploymentVariant string `json:"progressive-deployment-variant"`
}

func (r *Result) PrintEGJwt() string {
	return fmt.Sprintf(`JWT=%s`, r.EncodedJwt)
}

func (r *Result) PrintJwtBearer() string {
	return fmt.Sprintf(`JWT-Bearer  %s`, r.EncodedJwt)
}

func GetData() Data {
	return Data{
		ActorID:   "",
		SubjectID: "",
		IDP:       "",
		States: []State{
			{
				KeepAliveStrategy: KeepAliveStrategy{
					Lifecycle: "",
					Expiry:    "",
				},
				AuthenticationState: "",
			},
		},
		CustomClaims: CustomClaims{
			IdentityProvider: "",
			CredentialsGiven: "",
			SoftAccount:      "",
			SubjectTuid:      "",
			PrincipalTuid:    "",
		},
		Issuer:           "",
		Audience:         "",
		ClientID:         "",
		AuthTime:         "",
		Scope:            "",
		PartnerAccountID: "",
	}
}
