package main

import (
	"log"
	"sync"
	"time"
)

const (
	PASS int = 1
	ERR  int = 2
)

type bucket struct {
	pass int
	err  int
}

// 以一秒为一个单位的滑动窗口计数器
type secWindowCounter struct {
	begin_time   int64    //统计周期开始时间
	begin_bucket int      //统计周期开始的桶的位置
	buckets      []bucket //滑动窗口由固定个桶组成，循环使用
	sync.RWMutex
}

// NewSecWindowCounter 创建一个滑动窗口
func NewSecWindowCounter(bucket_count int) *secWindowCounter {
	return &secWindowCounter{
		begin_time:   time.Now().Unix(),
		begin_bucket: 0,
		buckets:      make([]bucket, bucket_count),
	}
}

// AddEvent metrics 中非零代表产生相应事件
func (sw *secWindowCounter) AddEvent(state int) {
	sw.Lock()
	defer sw.Unlock()

	bucket_count := len(sw.buckets)
	time_now := time.Now().Unix()
	time_begin := time_now - int64(bucket_count)

	// 当前时间窗口需要更新
	if time_begin > sw.begin_time {
		// 计算需要移除的桶数量
		remove_bucket := int(time_begin - sw.begin_time)

		// 全部移除，那就从头开始
		if remove_bucket > bucket_count {
			for n := 0; n < bucket_count; n++ {
				sw.buckets[n].pass = 0
				sw.buckets[n].err = 0
			}
			sw.begin_time = time_now
			sw.begin_bucket = 0
		} else {
			for n := 0; n < remove_bucket; n++ {
				pos := (sw.begin_bucket + n) % bucket_count
				sw.buckets[pos].pass = 0
				sw.buckets[pos].err = 0
			}

			sw.begin_time = time_begin
			sw.begin_bucket = (sw.begin_bucket + remove_bucket) % bucket_count
		}
	}

	pos := (sw.begin_bucket + int(time_now-sw.begin_time)) % bucket_count

	switch state {
	case PASS:
		sw.buckets[pos].pass++
	case ERR:
		sw.buckets[pos].err++
	default:
		log.Fatal("err type")
	}
}

// GetData 获取最新统计信息
func (sw *secWindowCounter) GetData() (pass int, err int) {
	sw.RLock()
	defer sw.RUnlock()

	for i := 0; i < len(sw.buckets); i++ {
		pass += sw.buckets[i].pass
		err += sw.buckets[i].err
	}

	return pass, err
}
