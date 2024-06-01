package model

import (
	"fmt"
)

const (
	TERM_WIDTH_DEFAULT                 = 150
	DEFAULT_TOKEN_ENV_KEY              = "DEFAULT_OPAQUE"
	JWT_GO_USE_TABLE                   = "JWT_GO_USE_TABLE"
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
	CartIdFlag                         = "c"
	CartIdInfo                         = "A cartId input for the session call"
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
	TravelerAPIEndPoint                = "ecomm-checkout-traveller-api.rcp.us-west-2.checkout.test.exp-aws.net:443"
	CreateCheckoutSession              = "expediagroup.travelerapi.checkout.CheckoutAPI/CreateCheckoutSession"
)

type TableInput struct {
	Name    string
	Content string
}

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

type Session struct {
	CheckoutSession CheckoutSession `json:"checkout_session"`
}

type CheckoutSession struct {
	CheckoutIdentifier CheckoutIdentifier `json:"checkout_identifier"`
}

type CheckoutIdentifier struct {
	SessionID    string `json:"session_id"`
	SessionToken string `json:"session_token"`
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
	return fmt.Sprintf(`EGToken Principal-JWT=%s`, r.EncodedJwt)
}

func (r *Result) PrintEGJwtWithAuth() string {
	return fmt.Sprintf(`Authorization: %s`, r.PrintEGJwt())
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

type GetSessionReq struct {
	Context Context `json:"context"`
	CartID  string  `json:"cart_id"`
}

type Context struct {
	Locale                    string                     `json:"locale"`
	Currency                  string                     `json:"currency"`
	DeviceContext             DeviceContext              `json:"device_context"`
	Experiments               []Experiment               `json:"experiments"`
	ResourceOwner             ResourceOwner              `json:"resource_owner"`
	PlatformProductID         string                     `json:"platform_product_id"`
	AdditionalPartnerAccounts []AdditionalPartnerAccount `json:"additional_partner_accounts"`
	PersonalizationContext    PersonalizationContext     `json:"personalization_context"`
}

type DeviceContext struct {
	DeviceUserAgentID string           `json:"device_user_agent_id"`
	CrossDomainID     string           `json:"cross_domain_id"`
	IP                string           `json:"ip"`
	UserAgent         string           `json:"user_agent"`
	DeviceType        string           `json:"device_type"`
	UserAgentContext  UserAgentContext `json:"user_agent_context"`
	TrustContext      TrustContext     `json:"trust_context"`
}

type UserAgentContext struct {
	BotnessScore   float64 `json:"botness_score"`
	AlignmentScore float64 `json:"alignment_score"`
	Classification string  `json:"classification"`
}

type TrustContext struct {
	Source  string `json:"source"`
	Payload string `json:"payload"`
}

type Experiment struct {
	ID     int `json:"id"`
	Bucket int `json:"bucket"`
}

type ResourceOwner struct {
	UserID string `json:"user_id"`
}

type AdditionalPartnerAccount struct {
	PartnerAccountID  string `json:"partner_account_id"`
	PlatformProductID string `json:"platform_product_id"`
	Type              int    `json:"type"`
}

type PersonalizationContext struct {
	SearchFilters []SearchFilter `json:"search_filters"`
}

type SearchFilter struct {
	ProductLine int    `json:"product_line"`
	Urn         string `json:"urn"`
	Explanation int    `json:"explanation"`
}

func GetSessionReData(cartId string) GetSessionReq {
	return GetSessionReq{
		Context: Context{
			Locale:   "en_US",
			Currency: "USD",
			DeviceContext: DeviceContext{
				DeviceUserAgentID: "a143bb33-3a40-4731-b730-674f8d3a91c5",
				CrossDomainID:     "ad58c8ce-a2ec-41d7-8865-723bc883cef1",
				IP:                "Hello",
				UserAgent:         "Hello",
				DeviceType:        "DESKTOP",
			},
			Experiments: []Experiment{
				{ID: 47688, Bucket: 1},
				{ID: 46893, Bucket: 1},
				{ID: 48088, Bucket: 0},
			},
			ResourceOwner: ResourceOwner{
				UserID: "6e3cc17f-bc98-4231-9e8f-3601a6cd739d",
			},
			PlatformProductID: "8a1938c7-f88c-4b76-b682-c7f4f2ada5e7",
			AdditionalPartnerAccounts: []AdditionalPartnerAccount{
				{
					PartnerAccountID:  "a2147247-59a1-4541-8e88-5d055c1b16b7",
					PlatformProductID: "dac30150-d377-4cf1-ac78-ca226e91c16f",
					Type:              0,
				},
			},
			PersonalizationContext: PersonalizationContext{
				SearchFilters: []SearchFilter{
					{ProductLine: 0, Urn: "Hello", Explanation: 0},
				},
			},
		},
		CartID: cartId,
	}
}

func GetTableInput(name string, content string) TableInput {
	return TableInput{
		Name:    name,
		Content: content,
	}
}
