package utils

import (
	"sync"
	"sync/atomic"
	"time"
)

var (
	timestampTimer sync.Once
	Timestamp      uint32
)

func StartTimeStampUpdater() {
	timestampTimer.Do(func() {

		atomic.StoreUint32(&Timestamp, uint32(time.Now().Unix()))
		go func(sleep time.Duration) {
			ticker := time.NewTicker(sleep)
			defer ticker.Stop()

			for t := range ticker.C {
				atomic.StoreUint32(&Timestamp, uint32(t.Unix()))
			}
		}(1 * time.Second)
	})
}
