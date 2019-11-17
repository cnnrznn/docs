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
	defer client.Close()
}

func TestClient(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(2)
	defer wg.Wait()

	for i := 0; i < 2; i++ {
		go fakeClient(&wg)
	}
}
