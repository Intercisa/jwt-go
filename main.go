package main

import (
	"flag"

	"bsipiczki.com/jwt-go/call"
	"bsipiczki.com/jwt-go/model"
	"bsipiczki.com/jwt-go/print"
	"bsipiczki.com/jwt-go/util"
)

func main() {
	defaultToken := util.GetEnv(model.DEFAULT_TOKEN_ENV_KEY)

	opaqueToken := flag.String(model.OpaqueFlag, defaultToken, model.OpaqueInfo)
	templated := flag.Bool(model.TemplateFlag, false, model.TemplateInfo)
	partnerAccId := flag.String(model.PartnerAccIdFlag, model.BLANK, model.PartnerAccIdInfo)
	flag.Parse()

	jwtResult := getJwt(partnerAccId, opaqueToken)
	print.PrintJWTAndAddToClipboard(jwtResult, *templated)
}

func getJwt(partnerAccId *string, opaqueToken *string) model.Result {
	if partnerAccId != nil && *partnerAccId != model.BLANK {
		baerer := call.AppJwtCall()
		opaque := call.OpaqueTokenCall(baerer, *partnerAccId)
		return call.PrincipalTokenCall(opaque.OpaqueToken)
	}
	return call.PrincipalTokenCall(*opaqueToken)
}
