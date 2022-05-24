package limiter

import (
	"clean-gin-template/config"
	"github.com/didip/tollbooth"
	"github.com/didip/tollbooth/limiter"
)

//type limiterSpec struct {
//
//}

func New(cfg *config.Config) *limiter.Limiter {
	newLimiter := tollbooth.NewLimiter(2, nil)
	newLimiter.SetBurst(5)

	return newLimiter
}
