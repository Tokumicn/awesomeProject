package randsleep

import (
	"golang.org/x/exp/rand"
	"time"
)

func init() {
	rand.Seed(uint64(time.Now().UnixNano()))
}

// RandSleep 随机休眠几秒钟
func RandSleep(n int) {
	if n == 0 {
		n = 1
	}
	sleepNum := rand.Intn(n)
	time.Sleep(time.Duration(sleepNum) * time.Second)
}
