package print

import (
	"encoding/json"
	"fmt"
	"strings"

	"bsipiczki.com/jwt-go/model"
	"bsipiczki.com/jwt-go/util"
)

func GetJWTJson(jwtResult model.Result, templated bool) []byte {
	var header interface{}

	headerJson := model.HeaderJson{
		ClientInfo:    model.ClientInfoValue,
		Authorization: jwtResult.PrintEGJwt(),
		DebugTrace:    model.DebugTraceValue,
	}

	if templated {
		header = model.TemplatedHeaderJson{
			HeaderJson:                    headerJson,
			ProgressiveDdeploymentVariant: model.ProgressiveDdeploymentVariantValue,
		}
	} else {
		header = headerJson
	}
	prettyJSON, err := json.MarshalIndent(header, "", "    ")
	if err != nil {
		panic(err)
	}

	return prettyJSON
}

func PrintJWTAndAddToClipboard(jwtResult model.Result, templated bool) {
	util.CopyToClippboard(jwtResult.PrintEGJwt())

	fmt.Println(jwtResult.PrintEGJwt())
	fmt.Println(strings.Repeat(".", util.GetTermWidth()))
	prettyJSON := GetJWTJson(jwtResult, templated)

	fmt.Printf("%s\n", prettyJSON)
}

func PrintSessionResponseAndAddToClipboard(session model.Session) {
	util.CopyToClippboard(session.CheckoutSession.CheckoutIdentifier.SessionID)

	fmt.Println(strings.Repeat(".", util.GetTermWidth()))

	prettyJSON, err := json.MarshalIndent(session.CheckoutSession.CheckoutIdentifier, "", "    ")
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", prettyJSON)
}
