package grpc

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"os/exec"

	"context"
	"strings"
	"time"

	"bsipiczki.com/jwt-go/model"

	"github.com/jhump/protoreflect/grpcreflect"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"

	// Register gzip compressor so compressed responses will work
	_ "google.golang.org/grpc/encoding/gzip"
	// Register xds so xds and xds-experimental resolver schemes work
	_ "google.golang.org/grpc/xds"

	"github.com/fullstorydev/grpcurl"
)

const (
	App  = "grpcurl"
	Arg0 = "-H"
	Arg1 = "-d"
)

var (
	format = "json"
)

func CallSession(jwtResult model.Result, cartId string) model.Session {
	sessionReqData := model.GetSessionReData(cartId)

	jsonData, err := json.Marshal(sessionReqData)
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
		os.Exit(1)
	}

	sessionReqStr := string(jsonData)

	cmd := exec.Command(App,
		Arg0,
		jwtResult.PrintEGJwtWithAuth(),
		Arg1,
		sessionReqStr,
		model.TravelerAPIEndPoint,
		model.CreateCheckoutSession,
	)

	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr

	err = cmd.Run()
	if err != nil {
		fmt.Printf("Error: %s\n", stderr.String())
		os.Exit(1)
	}
	var session model.Session

	err = json.Unmarshal([]byte(out.Bytes()), &session)
	if err != nil {
		fmt.Println("Error unmarshaling JSON:", err)
		os.Exit(1)
	}

	return session
}

func CallSessionV2(jwtResult model.Result, cartId string) (model.Session, error) {
	headers := jwtResult.PrintEGJwtWithAuth()

	data := model.GetSessionReData(cartId)
	return callGrpc[model.GetSessionReq, model.Session](headers, model.TravelerAPIEndPoint, model.CreateCheckoutSession, data)
}

func callGrpc[D interface{}, R interface{}](headers string, endpoint string, srevice string, data D) (R, error) {
	var zero R
	jsonData, err := json.Marshal(data)
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
		return zero, fail(err, "Error during marshaling", srevice)
	}

	verbosityLevel := 0
	ctx := context.Background()

	dial := func() (*grpc.ClientConn, error) {
		var zero *grpc.ClientConn

		dialTime := 10 * time.Second

		ctx, cancel := context.WithTimeout(ctx, dialTime)
		defer cancel()
		var opts []grpc.DialOption

		network := "tcp"
		var creds credentials.TransportCredentials

		insecure := false
		tlsConf, err := grpcurl.ClientTLSConfig(insecure, "", "", "")
		if err != nil {
			return zero, fail(err, "Failed to create TLS config")
		}

		creds = credentials.NewTLS(tlsConf)

		cc, err := grpcurl.BlockingDial(ctx, network, endpoint, creds, opts...)
		if err != nil {
			return zero, fail(err, "Failed to dial target host %q", endpoint)
		}
		return cc, nil
	}

	var cc *grpc.ClientConn
	var descSource grpcurl.DescriptorSource
	var refClient *grpcreflect.Client

	md := grpcurl.MetadataFromHeaders([]string{headers})
	refCtx := metadata.NewOutgoingContext(ctx, md)
	cc, err = dial()
	if err != nil {
		return zero, fail(err, "Failed to dial target host %q", endpoint)
	}
	refClient = grpcreflect.NewClientAuto(refCtx, cc)
	refClient.AllowMissingFileDescriptors()
	reflSource := grpcurl.DescriptorSourceFromServer(ctx, refClient)

	descSource = reflSource

	reset := func() {
		if refClient != nil {
			refClient.Reset()
			refClient = nil
		}
		if cc != nil {
			cc.Close()
			cc = nil
		}
	}
	defer reset()
	in := strings.NewReader(string(jsonData))

	options := grpcurl.FormatOptions{
		IncludeTextSeparator: true,
	}
	rf, formatter, err := grpcurl.RequestParserAndFormatter(grpcurl.Format(format), descSource, in, options)
	if err != nil {
		return zero, fail(err, "Failed to construct request parser and formatter for %q", format)
	}

	var out bytes.Buffer
	var session R
	h := &grpcurl.DefaultEventHandler{
		Out:            &out,
		Formatter:      formatter,
		VerbosityLevel: verbosityLevel,
	}

	err = grpcurl.InvokeRPC(ctx, descSource, cc, srevice, []string{headers}, h, rf.Next)
	if err != nil {
		return zero, fail(err, "Error invoking method %q", srevice)
	}

	if h.Status.Code() != codes.OK {
		grpcurl.PrintStatus(os.Stderr, h.Status, formatter)
		return zero, fail(err, "Error invoking method %q", srevice)
	}

	err = json.Unmarshal([]byte(out.Bytes()), &session)
	if err != nil {
		fmt.Println("Error unmarshaling JSON:", err)
		return zero, fail(err, "Error during marshaling", srevice)
	}

	return session, nil
}

func fail(err error, msg string, args ...interface{}) error {
	if err != nil {
		msg += ": %v"
		args = append(args, err)
	}
	fmt.Fprintf(os.Stderr, msg, args...)
	fmt.Fprintln(os.Stderr)
	if err != nil {
		return err
	} else {
		fmt.Fprintln(os.Stderr)
		return errors.New(msg)
	}
}
