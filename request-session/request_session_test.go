package request_session

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestSetGet(t *testing.T) {

	tests := []interface{}{
		"123456",
		"abcdefghij",
		"ABCDEFG",
		0.1234,
		0,
		-1,
		1111111111,
	}

	var wg sync.WaitGroup
	for _, v := range tests {
		wg.Add(1)
		go func(v interface{}) {
			Set(v)
			defer Delete()

			func(v interface{}) {
				got, ok := Get()
				if !ok || v != got {
					t.Errorf("get requestid failed. got:%v, want:%v", got, v)
				}
				t.Logf("set requestid:%v", v)
			}(v)
			wg.Done()
		}(v)
	}
	wg.Wait()
}

func TestConcurrency(t *testing.T) {
	var wg sync.WaitGroup
	count := 100000
	for i := 0; i < count; i++ {
		wg.Add(1)
		go func(v interface{}) {
			Set(v)
			defer Delete()

			func(v interface{}) {
				got, ok := Get()
				if !ok || v != got {
					t.Errorf("get requestid failed. got:%v, want:%v", got, v)
				}
			}(v)
			wg.Done()
		}(i)
	}
	t.Logf("test concurrency count:%v", count)

	wg.Wait()
}

func BenchmarkSet(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Set(i)
	}
}

func BenchmarkGet(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Get()
	}
}

func BenchmarkDelete(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Delete()
	}
}

func BenchmarkGetGoID(b *testing.B) {
	for i := 0; i < b.N; i++ {
		getGoID()
	}
}

func TestSync(t *testing.T) {
	m := make(map[int]int) //concurrent map writes
	for i := 0; i < 1000; i++ {
		i := i
		go func() {
			m[i] = i
		}()
	}
	time.Sleep(time.Second * 10)
	for i := 0; i < 1000000; i++ {
		i := 0
		go func() {
			if m[i] != i {
				fmt.Printf("key:%v  value:%v\n", i, m[i])
			}
		}()
	}
}

func TestSync2(t *testing.T) {
	m := sync.Map{}
	for i := 0; i < 1000000; i++ {
		i := i
		go func() {
			m.Store(i, i)
		}()
	}

	time.Sleep(time.Second * 10)
	for i := 0; i < 1000000; i++ {
		i := i
		go func() {
			if value, ok := m.Load(i); !ok || value != i {
				fmt.Printf("key:%v  value:%v\n", i, value)
			}
		}()
	}
}
