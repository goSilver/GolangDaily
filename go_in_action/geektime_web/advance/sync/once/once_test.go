package once

import "testing"

func TestOnceClose_Close(t *testing.T) {
	o := CloseOnce{}
	for i := 0; i < 10; i++ {
		o.Close()
	}
}
