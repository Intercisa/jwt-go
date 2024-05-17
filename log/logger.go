package log

import (
	"encoding/json"
	"fmt"
	"strings"

	"bsipiczki.com/jwt-go/model"
	"bsipiczki.com/jwt-go/terminal"
	"github.com/atotto/clipboard"
)

func PrintJWTAndAddToClipboard(jwtResult model.Result, templated bool) {
	err := clipboard.WriteAll(jwtResult.PrintEGJwt())
	if err != nil {
		fmt.Println("Failed to copy to clipboard:", err)
		return
	}

	fmt.Println(jwtResult.PrintEGJwt())

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

	fmt.Println(strings.Repeat(".", terminal.GetTermWidth()))

	prettyJSON, err := json.MarshalIndent(header, "", "    ")
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", prettyJSON)
}
