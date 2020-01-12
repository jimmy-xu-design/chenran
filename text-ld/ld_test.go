package text_ld

import (
	"testing"
	"fmt"
)

func TestLD(t *testing.T) {

	A := "今天我看到了天上有两个太阳，吓死人了, oh my god"
	B := "今天我看到天上有三个大太阳，吓人啊, wa my good"

	H, _, _, I, N := LD(A, B)
	ACC := float64(H-I)/float64(N)
	Corr := float64(H)/float64(N)

	fmt.Printf("Acc:%f, Corr:%f\n", ACC, Corr)
}
