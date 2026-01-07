package api

import (
	"time"

	"github.com/go-admin-team/go-admin-core/logger"
	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func NewGrpcServer(opts ...grpc.ServerOption) *grpc.Server {
	// panic recovery
	recoveryHandler := func(p interface{}) (err error) {
		logger.Error(
			"grpc panic recovered",
			zap.Any("panic", p),
			zap.Stack("stack"),
		)
		return status.Error(codes.Internal, "internal server error")
	}

	// 1️⃣ 新建 gRPC 专用 zap logger
	grpcLogger, _ := zap.NewProduction() // 或者 NewDevelopment，根据需求

	// zap 日志选项
	zapOpts := []grpc_zap.Option{
		grpc_zap.WithDurationField(func(d time.Duration) zapcore.Field {
			return zap.Int64("grpc.time_ms", d.Milliseconds())
		}),
	}

	// 原生 gRPC 链式拦截器
	unaryInterceptor := grpc.ChainUnaryInterceptor(
		grpc_ctxtags.UnaryServerInterceptor(),
		grpc_zap.UnaryServerInterceptor(grpcLogger, zapOpts...),
		grpc_recovery.UnaryServerInterceptor(grpc_recovery.WithRecoveryHandler(recoveryHandler)),
	)

	streamInterceptor := grpc.ChainStreamInterceptor(
		grpc_ctxtags.StreamServerInterceptor(),
		grpc_zap.StreamServerInterceptor(grpcLogger, zapOpts...),
		grpc_recovery.StreamServerInterceptor(grpc_recovery.WithRecoveryHandler(recoveryHandler)),
	)

	baseOpts := []grpc.ServerOption{
		unaryInterceptor,
		streamInterceptor,
	}

	// 允许外部追加 option（如 TLS / creds / keepalive）
	baseOpts = append(baseOpts, opts...)

	return grpc.NewServer(baseOpts...)
}
