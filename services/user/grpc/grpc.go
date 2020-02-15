package grpc

import (
	"context"
	"fmt"
	"net"

	"github.com/go-pg/pg/v9"
	pb "github.com/weslenng/petssenger/protos"
	"github.com/weslenng/petssenger/services/user/config"
	"github.com/weslenng/petssenger/services/user/models"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type userServer struct{}

func (*userServer) AuthUser(
	ctx context.Context,
	req *pb.AuthUserRequest,
) (*pb.AuthUserResponse, error) {
	user := req.GetUser()
	_, err := models.AuthUserByID(user)
	if err != nil {
		if err == pg.ErrNoRows {
			return nil, status.Errorf(
				codes.PermissionDenied,
				fmt.Sprintf("The user %v was not found", user),
			)
		}

		panic(err)
	}

	return &pb.AuthUserResponse{
		Authed: true,
	}, nil
}

// UserServerListen is a helper function to listen an user gRPC server
func UserServerListen() (net.Listener, error) {
	lis, err := net.Listen("tcp", config.Default.Addr)
	if err != nil {
		return nil, err
	}

	server := grpc.NewServer()
	pb.RegisterUserServer(server, &userServer{})
	if err := server.Serve(lis); err != nil {
		lis.Close()
		return nil, err
	}

	return lis, nil
}