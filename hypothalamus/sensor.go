package hypothalamus

import (
	"time"
)

type Sensor interface {
	Halt() error
	SenseContinuous(interval time.Duration) (<-chan *Environmental, error)
}
