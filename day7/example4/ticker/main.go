package ticker

import "time"

func main() {
	for i := 0; i < 1024; i++ {
		go func() {
			select {
			case <-time.After(time.Microsecond):
			}
		}()
	}
	time.Sleep(time.Second * 100)
}
