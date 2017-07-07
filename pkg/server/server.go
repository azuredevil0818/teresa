package server

import (
	"crypto/tls"
	"fmt"
	"net"
	"strings"

	"golang.org/x/net/context"

	log "github.com/Sirupsen/logrus"
	"github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"github.com/jinzhu/gorm"
	"github.com/luizalabs/teresa-api/models/storage"
	"github.com/luizalabs/teresa-api/pkg/server/app"
	"github.com/luizalabs/teresa-api/pkg/server/auth"
	"github.com/luizalabs/teresa-api/pkg/server/k8s"
	st "github.com/luizalabs/teresa-api/pkg/server/storage"
	"github.com/luizalabs/teresa-api/pkg/server/team"
	"github.com/luizalabs/teresa-api/pkg/server/user"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

type Options struct {
	Port    string
	TLSCert *tls.Certificate
	Auth    auth.Auth
	DB      *gorm.DB
	Storage st.Storage
	K8s     k8s.Client
}

type Server struct {
	listener   net.Listener
	grpcServer *grpc.Server
}

type serverStreamWrapper struct {
	grpc.ServerStream
	ctx context.Context
}

func (w *serverStreamWrapper) Context() context.Context {
	return w.ctx
}

func streamInterceptor(a auth.Auth, uOps user.Operations) grpc.StreamServerInterceptor {
	return func(srv interface{}, stream grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
		if strings.HasSuffix(info.FullMethod, "Login") {
			return handler(srv, stream)
		}

		ctx := stream.Context()
		user, err := authorize(ctx, a, uOps)
		if err != nil {
			return err
		}

		ctx = context.WithValue(ctx, "user", user)
		wrap := &serverStreamWrapper{stream, ctx}
		return handler(srv, wrap)
	}
}

func unaryInterceptor(a auth.Auth, uOps user.Operations) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		if strings.HasSuffix(info.FullMethod, "Login") {
			return handler(ctx, req)
		}

		user, err := authorize(ctx, a, uOps)
		if err != nil {
			return nil, err
		}

		ctx = context.WithValue(ctx, "user", user)
		return handler(ctx, req)
	}
}

func authorize(ctx context.Context, a auth.Auth, uOps user.Operations) (*storage.User, error) {
	md, ok := metadata.FromContext(ctx)
	if !ok {
		return nil, auth.ErrPermissionDenied
	}
	if len(md["token"]) < 1 || md["token"][0] == "" {
		return nil, auth.ErrPermissionDenied
	}
	email, err := a.ValidateToken(md["token"][0])
	if err != nil {
		return nil, err
	}
	return uOps.GetUser(email)
}

func (s *Server) Run() error {
	return s.grpcServer.Serve(s.listener)
}

func recFunc(p interface{}) (err error) {
	log.WithField("panic", p).Error("teresa-server recovered")
	return status.Errorf(codes.Unknown, "Internal Server Error")
}

func New(opt Options) (*Server, error) {
	l, err := net.Listen("tcp", fmt.Sprintf(":%s", opt.Port))
	if err != nil {
		return nil, err
	}

	uOps := user.NewDatabaseOperations(opt.DB, opt.Auth)
	recOpts := []grpc_recovery.Option{
		grpc_recovery.WithRecoveryHandler(recFunc),
	}
	sOpts := []grpc.ServerOption{
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			unaryInterceptor(opt.Auth, uOps),
			grpc_recovery.UnaryServerInterceptor(recOpts...),
		)),
		grpc.StreamInterceptor(grpc_middleware.ChainStreamServer(
			streamInterceptor(opt.Auth, uOps),
			grpc_recovery.StreamServerInterceptor(recOpts...),
		)),
	}
	if opt.TLSCert != nil {
		creds := credentials.NewServerTLSFromCert(opt.TLSCert)
		sOpts = append(sOpts, grpc.Creds(creds))
	}

	s := grpc.NewServer(sOpts...)

	us := user.NewService(uOps)
	us.RegisterService(s)

	tOps := team.NewDatabaseOperations(opt.DB, uOps)
	t := team.NewService(tOps)
	t.RegisterService(s)

	appOps := app.NewOperations(tOps, opt.K8s, opt.Storage)
	a := app.NewService(appOps)
	a.RegisterService(s)

	return &Server{listener: l, grpcServer: s}, nil
}
