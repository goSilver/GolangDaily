package once

import (
	"fmt"
	"sync"
)

type CloseOnce struct {
	close sync.Once
}

// Close 当接收者是结构体时once不生效，指针类型接收者才生效
func (o *CloseOnce) Close() error {
	o.close.Do(func() {
		fmt.Println("close")
	})
	return nil
}
