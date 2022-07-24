package finance

import "time"

type FinanceRecordRepo interface {
	QueryFinanceRecord(userAddress string, offsetIdx int) []FinanceRecord

	// 保存资产变动记录
	SaveFinanceRecord(r FinanceRecord)

	// 删除资产记录
	DeleteRecords(userAddress string, deleteRecords []int)

	// 保存合并后的资产记录
	SaveMergedFinanceRecords(rs []FinanceRecord)

	// 查询已合并的资产记录
	QueryMergedFinanceRecords(userAddress string) []FinanceRecord
}

type FinanceRecord struct {
	Id                 int
	FromAddress        string
	ToAddress          string
	Event              string
	ChangedNu          float64
	ChangedVipDuration int64
	CreateTime         time.Time
}
