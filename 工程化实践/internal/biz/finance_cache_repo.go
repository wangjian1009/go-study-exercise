package biz

import "time"

type FinanceCacheRepo interface {

	// 查询用户资产缓存
	QueryFinanceCache(userAddress string) (*UserFinanceCache, error)

	// 创建缓存
	CreateFinanceCache(userAddress string, offsetIdx int, nuAmount float64, vipExpireTime time.Time)

	// 更新数据库缓存的值
	UpdateFinanceCache(userAddress string, offsetIdx int, nuAmount float64, vipExpireTime time.Time)
}
