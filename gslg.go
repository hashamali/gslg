package gslg

import (
	"context"
	"path"
	"time"

	"github.com/google/uuid"
	"github.com/hashamali/gsl"
	"github.com/rs/zerolog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/peer"
	"google.golang.org/grpc/status"
)

// Interceptor will return a new gRPC interceptor for logging.
func Interceptor(logger gsl.Log) func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		start := time.Now()
		l := new(ctx, info.FullMethod)
		defer l.send(ctx, logger, start)
		resp, err := handler(ctx, req)
		l.Error = err
		return resp, err
	}
}

type log struct {
	ID         string
	RemoteIP   string
	Service    string
	Method     string
	StatusCode int
	Latency    float64
	Error      error
}

func (l *log) MarshalZerologObject(zle *zerolog.Event) {
	zle.Str("source", "grpc")
	zle.Str("id", l.ID)
	zle.Str("remote_ip", l.RemoteIP)
	zle.Str("service", l.Service)
	zle.Str("method", l.Method)
	zle.Int("status_code", l.StatusCode)
	zle.Float64("latency", l.Latency)

	if l.Error != nil {
		zle.Err(l.Error)
	}
}

func (l *log) send(ctx context.Context, logger gsl.Log, start time.Time) {
	l.StatusCode = int(status.Code(l.Error))
	l.Latency = float64(time.Since(start).Nanoseconds()) / 1000000.0

	if l.Error != nil {
		if logger != nil {
			logger.Errorw(l, "")
		}
	} else {
		if logger != nil {
			logger.Infow(l, "")
		}
	}
}

func new(ctx context.Context, fullMethod string) *log {
	l := &log{
		ID:      uuid.New().String(),
		Service: path.Dir(fullMethod)[1:],
		Method:  path.Base(fullMethod),
	}

	p, ok := peer.FromContext(ctx)
	if ok {
		l.RemoteIP = p.Addr.String()
	}

	return l
}
