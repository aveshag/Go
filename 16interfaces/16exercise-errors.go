// https://go.dev/tour/methods/20
package main

import (
	"fmt"
	"math"
)

const DELTA = 0.0000001
const INITIAL_Z = 100.0

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
}

// here second argument is error so Sprintf will be in infinite loop if do not cast e to float64
// https://stackoverflow.com/questions/27474907/why-would-a-call-to-fmt-sprinte-inside-the-error-method-result-in-an-infinit
func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0.0, ErrNegativeSqrt(x)
	}
	z := INITIAL_Z

	step := func() float64 {
		return z - (z*z-x)/(2*z)
	}

	for zz := step(); math.Abs(zz-z) > DELTA; {
		z = zz
		zz = step()
	}
	return z, nil
}

func main() {
	fmt.Println(Sqrt(2))
	_, err := Sqrt(-2)
	if err != nil {
		fmt.Println(err)
	}
}
