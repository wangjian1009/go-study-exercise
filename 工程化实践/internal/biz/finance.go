package finance

import "time"

type UserFinance struct {
	NuAmount      float64
	VipExpireTime time.Time
}

type UserFinanceCache struct {
	NuAmount      float64
	VipExpireTime time.Time
	CacheVersion  int
	OffsetIdx     int
}

func NewUserFinance(nuAmount float64, vipExpireTime time.Time) UserFinance {
	return UserFinance{
		NuAmount:      nuAmount,
		VipExpireTime: vipExpireTime,
	}
}
