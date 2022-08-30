package main

import (
	goredislib "github.com/go-redis/redis/v8"
	"github.com/go-redsync/redsync/v4"
	"github.com/go-redsync/redsync/v4/redis/goredis/v8"
	"go.uber.org/zap"
	"time"
)

func main() {
	logger, _ := zap.NewDevelopment()
	zap.ReplaceGlobals(logger)

	for i := 0; i < 20; i++ {
		go func() {
			client := goredislib.NewClient(&goredislib.Options{
				Addr: "localhost:6379",
			})
			pool := goredis.NewPool(client)
			rs := redsync.New(pool)
			mutex := rs.NewMutex("stock_sell")

			if err := mutex.Lock(); err != nil {
				zap.S().Error("redis 加锁失败")
			} else {
				zap.S().Info("redis 加锁成功")

			}
			time.Sleep(5 * time.Second)
			defer func(m *redsync.Mutex) {
				if ok, err := m.Unlock(); !ok || err != nil {
					zap.S().Error("redis 解锁失败")
				} else {
					zap.S().Info("redis 解锁成功")

				}
			}(mutex)
		}()
	}

	select {}

}
