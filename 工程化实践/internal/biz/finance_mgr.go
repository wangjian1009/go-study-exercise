package finance

type FinanceMgr struct {
	financeCacheRepo  FinanceCacheRepo
	financeRecordRepo FinanceRecordRepo
}

func NewFinanceMgr(
	financeCacheRepo FinanceCacheRepo,
	financeRecordRepo FinanceRecordRepo,
) FinanceMgr {
	mgr := FinanceMgr{
		financeCacheRepo:  financeCacheRepo,
		financeRecordRepo: financeRecordRepo,
	}

	return mgr
}
