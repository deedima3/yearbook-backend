package ping

import "context"

type PingService interface {
	Ping(ctx context.Context) string
}

type PingImpl struct {}

func(ps PingImpl) Ping(ctx context.Context) string {
	return "ini ping. saya wibu dan saya bangga"
}

func ProvidePingService() PingImpl{
	return PingImpl{}
}