package client

import (
	"fmt"
	"sync"
	"time"

	pb "github.com/cnnrznn/docs/editor"

	"testing"
)

func fakeClient(wg *sync.WaitGroup) {
	defer wg.Done()

	client, err := New()
	if err != nil {
		fmt.Println(err)
		return
	}

	for i := 0; i < 100; i++ {
		client.Stream.Send(&pb.Op{Type: 0,
			Version: int64(i),
			Pos:     int64(i),
			Char:    int32('c')})
		fmt.Println(i)
	}

	time.Sleep(5 * time.Second)
}

func TestClient(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(2)
	defer wg.Wait()

	for i := 0; i < 2; i++ {
		go fakeClient(&wg)
	}
}
