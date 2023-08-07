package pool

import "github.com/gernhan/xml/concurrent"

type Pools struct {
	DbPool         *concurrent.ThreadPool
	ProcessingPool *concurrent.ThreadPool
	UploadingPool  *concurrent.ThreadPool
	CommonPool     *concurrent.ThreadPool
	BillRunPool    *concurrent.ThreadPool
}

var pools Pools

func GetPools() Pools {
	return pools
}

func InitPools() {
	pools = Pools{
		DbPool:         concurrent.NewThreadPool(20),
		ProcessingPool: concurrent.NewThreadPool(20),
		UploadingPool:  concurrent.NewThreadPool(2000),
		CommonPool:     concurrent.NewThreadPool(20),
		BillRunPool:    concurrent.NewThreadPool(1),
	}
}
