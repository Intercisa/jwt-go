package call

import (
	"bsipiczki.com/jwt-go/model"
	"github.com/go-zoox/fetch"
	"github.com/google/uuid"
)

func call[D interface{}, R interface{}](data *D, result R, header map[string]string, endpoint string) R {
	var body interface{}
	if data != nil {
		body = *data
	}

	response, err := fetch.Post(endpoint, &fetch.Config{
		Body:    body,
		Headers: header,
	})

	if err != nil {
		panic(err)
	}

	response.UnmarshalJSON(&result)
	return result
}

func PrincipalTokenCall(opaqueToken string) model.Result {
	data := model.Payload{
		OpaqueToken: opaqueToken,
	}

	header := make(map[string]string)
	header[model.AuthorizationKey] = model.AuthorizationBasicValue
	header[model.CookieKey] = model.CookieValue

	var result model.Result
	return call(&data, result, header, model.Endpoint)
}

func AppJwtCall() model.Result {
	header := make(map[string]string)
	header[model.AuthorizationKey] = model.AuthorizationBasicValue
	header[model.CookieKey] = model.CookieValue

	var result model.Result
	var nilData *map[string]interface{}
	return call(nilData, result, header, model.AppJwtEndpoint)
}

func OpaqueTokenCall(bearerResult model.Result, partnerAccId string) model.OpaqueResult {
	header := make(map[string]string)
	header[model.ContentTypeKey] = model.ContentTypeValue
	header[model.AcceptKey] = model.AcceptValue
	header[model.AuthorizationKey] = bearerResult.PrintJwtBearer()
	header[model.TraceIdKey] = uuid.New().String()
	header[model.MessageIdKey] = uuid.New().String()
	header[model.ParentMessageIdKey] = uuid.New().String()
	header[model.DeviceUserAgentIdKey] = uuid.New().String()
	header[model.XDIdKey] = uuid.New().String()

	data := model.GetData()
	data.PartnerAccountID = partnerAccId

	var result model.OpaqueResult
	return call(&data, result, header, model.OpaqueEndpoint)
}
