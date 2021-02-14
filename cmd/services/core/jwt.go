package main

import (
	"time"

	"github.com/lovung/GoCleanArchitecture/pkg/jwtutil"
)

const (
	halfDayHour = 12
	oneDayHour  = 24
)

func (s *application) initJWTSession(secret string) {
	jwtutil.InitJWTSession(
		secret,
		time.Hour*halfDayHour,
		time.Hour*oneDayHour,
	)
}
