package memqueue_test

import (
	"testing"

	"github.com/go-msgqueue/msgqueue"
	"github.com/go-msgqueue/msgqueue/memqueue"
)

func BenchmarkCallAsync(b *testing.B) {
	q := memqueue.NewQueue(&msgqueue.Options{
		Handler:    func() {},
		BufferSize: 1000000,
	})
	defer q.Close()

	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			q.Call()
		}
	})
}

func BenchmarkNamedMessage(b *testing.B) {
	q := memqueue.NewQueue(&msgqueue.Options{
		Redis:      redisRing(),
		Handler:    func() {},
		BufferSize: 1000000,
	})
	defer q.Close()

	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			msg := msgqueue.NewMessage()
			msg.Name = "myname"
			q.Add(msg)
		}
	})
}
