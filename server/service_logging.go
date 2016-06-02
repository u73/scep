package scepserver

import (
	"time"

	"github.com/go-kit/kit/log"
	"golang.org/x/net/context"
)

type loggingService struct {
	logger log.Logger
	Service
}

// NewLoggingService creates adds logging to the SCEP service
func NewLoggingService(logger log.Logger, s Service) Service {
	return &loggingService{logger, s}
}

func (mw loggingService) GetCACaps(ctx context.Context) (caps []byte, err error) {
	defer func(begin time.Time) {
		_ = mw.logger.Log(
			"method", "GetCACaps",
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())
	caps, err = mw.Service.GetCACaps(ctx)
	return
}

func (mw loggingService) GetCACert(ctx context.Context) (cert []byte, err error) {
	defer func(begin time.Time) {
		_ = mw.logger.Log(
			"method", "GetCACert",
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())
	cert, err = mw.Service.GetCACert(ctx)
	return
}

func (mw loggingService) PKIOperation(ctx context.Context, data []byte) (certRep []byte, err error) {
	defer func(begin time.Time) {
		_ = mw.logger.Log(
			"method", "PKIOperation",
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())
	certRep, err = mw.Service.PKIOperation(ctx, data)
	return
}
