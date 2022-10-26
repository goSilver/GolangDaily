package unsafe

import (
	"fmt"
	"testing"
	"unsafe"
)

func TestPrintFieldOffset(t *testing.T) {
	fmt.Println(unsafe.Sizeof(User{}))
	PrintFieldOffset(User{})

	fmt.Println(unsafe.Sizeof(UserV1{}))
	PrintFieldOffset(UserV1{})

	fmt.Println(unsafe.Sizeof(UserV2{}))
	PrintFieldOffset(UserV2{})
}
