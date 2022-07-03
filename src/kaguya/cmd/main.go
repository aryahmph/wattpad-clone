package main

import (
	"fmt"
	"github.com/aryahmph/wattpad-clone/src/kaguya/common/env"
	"github.com/aryahmph/wattpad-clone/src/kaguya/common/grpc/pb"
	jwtCommon "github.com/aryahmph/wattpad-clone/src/kaguya/common/jwt"
	nsqProducer "github.com/aryahmph/wattpad-clone/src/kaguya/common/nsq/producer"
	mailProd "github.com/aryahmph/wattpad-clone/src/kaguya/common/nsq/producer/mail"
	dbCommon "github.com/aryahmph/wattpad-clone/src/kaguya/common/postgres"
	authDelivery "github.com/aryahmph/wattpad-clone/src/kaguya/internal/delivery/auth/grpc"
	userDelivery "github.com/aryahmph/wattpad-clone/src/kaguya/internal/delivery/user/grpc"
	userRepo "github.com/aryahmph/wattpad-clone/src/kaguya/internal/repository/user/postgres"
	authUc "github.com/aryahmph/wattpad-clone/src/kaguya/internal/usecase/auth"
	userUc "github.com/aryahmph/wattpad-clone/src/kaguya/internal/usecase/user"
	"github.com/go-playground/validator/v10"
	grpcMiddleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpcLogrus "github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus"
	grpcRecovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	grpcCtxTags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"log"
	"net"
)

func main() {
	cfg := env.LoadConfig()
	db := dbCommon.NewPostgres(cfg.PostgresMigrationPath, cfg.PostgresURL)
	logEntry, logOpts := initLogger()
	validate := validator.New()
	jwtManager := jwtCommon.NewJWTManager(cfg.AccessTokenKey)

	producer := nsqProducer.NewProducer(cfg.NSQAddr)
	mailProducer := mailProd.NewMailProducer(producer, cfg.NSQMailTopic)

	userRepository := userRepo.NewPostgresUserRepositoryImpl(db)
	userUsecase := userUc.NewUserUsecaseImpl(userRepository, validate, mailProducer)

	authUsecase := authUc.NewAuthUsecaseImpl(userRepository, jwtManager, validate, mailProducer)

	listen, err := net.Listen("tcp", fmt.Sprintf("%s:%d", cfg.Host, cfg.Port))
	if err != nil {
		log.Fatalln(err)
	}

	grpcServer := grpc.NewServer(
		grpcMiddleware.WithUnaryServerChain(
			grpcCtxTags.UnaryServerInterceptor(grpcCtxTags.WithFieldExtractor(grpcCtxTags.CodeGenRequestFieldExtractor)),
			grpcLogrus.UnaryServerInterceptor(logEntry, logOpts...),
			grpcRecovery.UnaryServerInterceptor(),
		),
	)

	pb.RegisterGRPCUserServer(grpcServer, userDelivery.NewGRPCUserServer(userUsecase))
	pb.RegisterGRPCAuthServer(grpcServer, authDelivery.NewGRPCAuthServer(authUsecase))

	err = grpcServer.Serve(listen)
	if err != nil {
		log.Fatalln(err)
	}
}

func initLogger() (*logrus.Entry, []grpcLogrus.Option) {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})
	entry := logrus.NewEntry(logger)
	logOpts := []grpcLogrus.Option{
		grpcLogrus.WithLevels(func(code codes.Code) logrus.Level {
			if code == codes.OK {
				return logrus.InfoLevel
			}
			return logrus.ErrorLevel
		}),
	}
	grpcLogrus.ReplaceGrpcLogger(entry)
	return entry, logOpts
}
