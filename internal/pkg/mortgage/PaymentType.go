package mortgage

import (
	"time"
)

type PaymentType time.Duration

const dayDuration time.Duration = time.Hour * 24

const (
	PaymentTypeWeekly   = PaymentType(dayDuration * 7)
	PaymentTypeBiWeekly = PaymentType(dayDuration * 7 * 2)
	PaymentTypeMonthly  = PaymentType(float32(dayDuration) * 30.5)
)
