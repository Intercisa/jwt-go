package call

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	"bsipiczki.com/jwt-go/model"
	"github.com/google/uuid"
)

func PrincipalTokenCall(opaqueToken string) model.Result {
	data := model.Payload{
		OpaqueToken: opaqueToken,
	}

	payloadBytes, err := json.Marshal(data)
	if err != nil {
		log.Fatal(err)
	}
	body := bytes.NewReader(payloadBytes)

	req, err := http.NewRequest(model.Post, model.Endpoint, body)
	if err != nil {
		log.Fatal(err)
	}

	headerProps := make(map[string]string)
	headerProps[model.AcceptKey] = model.AcceptValue
	headerProps[model.ContentTypeKey] = model.ContentTypeValue
	headerProps[model.TraceIdKey] = uuid.New().String()

	setHeader(req.Header, headerProps)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	var result model.Result
	err = json.Unmarshal(bodyBytes, &result)
	if err != nil {
		panic(err)
	}

	return result
}

func AppJwtCall() model.Result {
	client := &http.Client{}
	req, err := http.NewRequest(model.Post, model.AppJwtEndpoint, nil)
	if err != nil {
		log.Fatal(err)
	}

	headerProps := make(map[string]string)
	headerProps[model.AuthorizationKey] = model.AuthorizationBasicValue
	headerProps[model.CookieKey] = model.CookieValue

	setHeader(req.Header, headerProps)

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var result model.Result
	err = json.Unmarshal(bodyBytes, &result)
	if err != nil {
		panic(err)
	}

	return result
}

func OpaqueTokenCall(bearerResult model.Result, partnerAccId string) model.OpaqueResult {
	client := &http.Client{}
	defaultData := model.GetData()
	defaultData.PartnerAccountID = partnerAccId

	jsonData, err := json.MarshalIndent(defaultData, "", "  ")
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
	}

	req, err := http.NewRequest(model.Post, model.OpaqueEndpoint, strings.NewReader(string(jsonData)))
	if err != nil {
		log.Fatal(err)
	}

	headerProps := make(map[string]string)
	headerProps[model.ContentTypeKey] = model.ContentTypeValue
	headerProps[model.AcceptKey] = model.AcceptValue
	headerProps[model.AuthorizationKey] = bearerResult.PrintJwtBearer()
	headerProps[model.TraceIdKey] = uuid.New().String()
	headerProps[model.MessageIdKey] = uuid.New().String()
	headerProps[model.ParentMessageIdKey] = uuid.New().String()
	headerProps[model.DeviceUserAgentIdKey] = uuid.New().String()
	headerProps[model.XDIdKey] = uuid.New().String()

	setHeader(req.Header, headerProps)

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	var result model.OpaqueResult
	err = json.Unmarshal(bodyBytes, &result)
	if err != nil {
		panic(err)
	}

	return result
}

func setHeader(header http.Header, headerProps map[string]string) {
	for key, value := range headerProps {
		header.Set(key, value)
	}
}
