package service

import (
	"context"
	"fmt"
	"regexp"

	sgldb "github.com/gitJaesik/stream_go_srvs/streamgolib/db"
	pbsas "github.com/gitJaesik/stream_go_srvs/streamgolib/gen/proto/go/stream_auth_server/v1"
	"github.com/gitJaesik/stream_go_srvs/streamgolib/logger"
)

type streamAuthServer struct {
	mongoDb sgldb.StreamGoLibDB
	pbsas.UnimplementedStreamAuthServiceServer
}

var regUser *regexp.Regexp
var regPasswordAlphaLow *regexp.Regexp
var regPasswordAlphaUp *regexp.Regexp
var regPasswordNumberic *regexp.Regexp
var regPasswordSpecial *regexp.Regexp

func init() {
	regUser = regexp.MustCompile(`^([a-zA-Z0-9])+([a-zA-Z0-9\\-])*([a-zA-Z0-9])+$`)

	// https://stackoverflow.com/questions/19605150/regex-for-password-must-contain-at-least-eight-characters-at-least-one-number-a
	regPasswordAlphaLow = regexp.MustCompile(`([a-z])`)
	regPasswordAlphaUp = regexp.MustCompile(`([A-Z])`)
	regPasswordNumberic = regexp.MustCompile(`([0-9])`)
	regPasswordSpecial = regexp.MustCompile(`([^A-Za-z0-9])`)
}

func NewStreamAuthServerServer(mongoDb sgldb.StreamGoLibDB) *streamAuthServer {
	return &streamAuthServer{
		mongoDb: mongoDb,
	}
}

func (sas *streamAuthServer) Echo(c context.Context, eq *pbsas.EchoRequest) (*pbsas.EchoResponse, error) {
	str := fmt.Sprintf("Get value: (%q)\n", eq.Message)
	logger.Logger.Infow("Echo", "EchoRequest Message", str)

	return &pbsas.EchoResponse{Message: eq.Message + " response"}, nil
}

func (sas *streamAuthServer) SignIn(c context.Context, sir *pbsas.SignInRequest) (*pbsas.SignInResponse, error) {
	// str := fmt.Sprintf("Get value: (%q)\n", eq.Message)
	// logger.Logger.Infow("Echo", "EchoRequest Message", str)

	// return &pbsas.EchoResponse{Message: eq.Message + " response"}, nil

	return &pbsas.SignInResponse{JwtToken: "jwttoken"}, nil
}
func (sas *streamAuthServer) SignOut(c context.Context, sir *pbsas.SignOutRequest) (*pbsas.SignOutResponse, error) {
	// str := fmt.Sprintf("Get value: (%q)\n", eq.Message)
	// logger.Logger.Infow("Echo", "EchoRequest Message", str)

	// return &pbsas.EchoResponse{Message: eq.Message + " response"}, nil

	return &pbsas.SignOutResponse{Success: false, Result: ""}, nil
}

func (sas *streamAuthServer) SignUp(c context.Context, sir *pbsas.SignUpRequest) (*pbsas.SignUpResponse, error) {
	// str := fmt.Sprintf("Get value: (%q)\n", eq.Message)
	// logger.Logger.Infow("Echo", "EchoRequest Message", str)

	// return &pbsas.EchoResponse{Message: eq.Message + " response"}, nil

	return &pbsas.SignUpResponse{Success: false, Result: ""}, nil
}

func (sas *streamAuthServer) GetPlayerInfo(c context.Context, gpir *pbsas.GetPlayerInfoRequest) (*pbsas.GetPlayerInfoResponse, error) {
	// str := fmt.Sprintf("Get value: (%q)\n", gpir.UserId)
	logger.Logger.Infow("GetPlayerInfo", "received user_id", gpir.PlayerId)

	/*
		p, ok := peer.FromContext(c)
		if ok {
			logger.Logger.Debugw("GetPlayerInfo", "peer", p)
		}
		userID := c.Value(stream_authentication.UserIDKey{})
	*/

	validate := func() error {
		// if gpir.PlayerId == 0 {
		// 	// TODO: do not use magic var
		// 	err := errors.New("INVALID_PARAMETERS_EMPTY_PACKET")
		// 	return err
		// }
		return nil
	}

	err := ValidationHandler(c, validate)
	if err != nil {
		return &pbsas.GetPlayerInfoResponse{}, nil

	}
	// if err != nil {
	// 	responseErr := stream_error.CreateGRPCError(codes.Internal, err.Error(), pbsas.StreamErrorCodes_ASR_Fail)
	//  return nil, responseErr
	// }

	mongoResponse, mongoErr := sas.mongoDb.GetPlayerInfo(c, gpir)
	if mongoErr != nil {
		return nil, mongoErr
	}
	logger.Logger.Infow("GetPlayerInfo", "mongoResponse", mongoResponse)

	/*
		for _, docs := range mongoResponse{
		}
	*/

	// return &pbsas.GetPlayerInfoResponse{}, nil
	return mongoResponse.(*pbsas.GetPlayerInfoResponse), nil
}
