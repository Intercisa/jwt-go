package main

import (
	"flag"

	"bsipiczki.com/jwt-go/call"
	"bsipiczki.com/jwt-go/grpc"
	"bsipiczki.com/jwt-go/model"
	"bsipiczki.com/jwt-go/print"
	"bsipiczki.com/jwt-go/table"
	"bsipiczki.com/jwt-go/util"
)

func main() {
	defaultToken := util.GetEnv(model.DEFAULT_TOKEN_ENV_KEY)
	useTable := util.CheckBoolEnv(model.JWT_GO_USE_TABLE)

	opaqueToken := flag.String(model.OpaqueFlag, defaultToken, model.OpaqueInfo)
	templated := flag.Bool(model.TemplateFlag, false, model.TemplateInfo)
	partnerAccId := flag.String(model.PartnerAccIdFlag, model.BLANK, model.PartnerAccIdInfo)
	cartId := flag.String(model.CartIdFlag, model.BLANK, model.CartIdInfo)
	flag.Parse()

	if useTable {
		printWithTable(*cartId, *templated, partnerAccId, opaqueToken)
	} else {
		printWithoutTable(*cartId, *templated, partnerAccId, opaqueToken)
	}
}

func printWithTable(cartId string, templated bool, partnerAccId *string, opaqueToken *string) {
	var inputs []model.TableInput
	jwtResult := getJwt(partnerAccId, opaqueToken)

	jwtInput := model.GetTableInput("JWT", jwtResult.PrintEGJwt())
	jwtJSONInput := model.GetTableInput("JWT-JSON", string(print.GetJWTJson(jwtResult, templated)))

	inputs = append(inputs, jwtInput, jwtJSONInput)

	if cartId != "" {
		session, err := grpc.CallSessionV2(jwtResult, cartId)
		if err == nil {
			sessionIdInput := model.GetTableInput("SessionId", session.CheckoutSession.CheckoutIdentifier.SessionID)
			sessionToken := model.GetTableInput("SessionToken", session.CheckoutSession.CheckoutIdentifier.SessionToken)

			inputs = append(inputs, sessionIdInput, sessionToken)
		}
	}
	table.Render(inputs)
}

func printWithoutTable(cartId string, templated bool, partnerAccId *string, opaqueToken *string) {
	jwtResult := getJwt(partnerAccId, opaqueToken)
	print.PrintJWTAndAddToClipboard(jwtResult, templated)

	if cartId != "" {
		session, err := grpc.CallSessionV2(jwtResult, cartId)
		if err == nil {
			print.PrintSessionResponseAndAddToClipboard(session)
		}
	}
}

func getJwt(partnerAccId *string, opaqueToken *string) model.Result {
	if partnerAccId != nil && *partnerAccId != model.BLANK {
		baerer := call.AppJwtCall()
		opaque := call.OpaqueTokenCall(baerer, *partnerAccId)
		return call.PrincipalTokenCall(opaque.OpaqueToken)
	}
	return call.PrincipalTokenCall(*opaqueToken)
}
