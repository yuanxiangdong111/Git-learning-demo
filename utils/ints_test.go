package utils

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"testing"
	"time"
)

func IntToByte1(num int64) []byte {
	var buffer bytes.Buffer
	err := binary.Write(&buffer, binary.BigEndian, num)
	if err != nil {
		return nil
	}
	return buffer.Bytes()
}

func IntToByte2(num int64) []byte {
	var buf = make([]byte, 8)
	binary.BigEndian.PutUint64(buf, uint64(num))
	return buf
}

func TestIntToByteTime(t *testing.T) {
	var a []int = []int{100, 1000, 10000, 100000, 1000000}
	for _, n := range a {
		testSlice(n)
	}
	fmt.Println("************")
}

func TestIntToByte(t *testing.T) {
	fmt.Println(bytes.Equal(IntToByte1(10), IntToByte2(10)))

}

func testSlice(n int) {
	fmt.Println("************")
	fmt.Println("N : ", n)
	t1 := time.Now()

	for i := 0; i < n; i++ {
		IntToByte1(10)
	}

	elapsed1 := time.Since(t1)
	fmt.Println("no cap :", elapsed1)

	t2 := time.Now()
	for i := 0; i < n; i++ {
		IntToByte2(10)
	}
	elapsed2 := time.Since(t2)
	fmt.Println("has cap :", elapsed2)
}
