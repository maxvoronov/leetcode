package reverse_integer

import "math"

func ReverseV2(x int) int {
	if x >= -9 && x <= 9 {
		return x
	}

	if x > math.MaxInt32 || x < math.MinInt32 {
		return 0 // Int32 overflow
	}

	// Assume the environment does not allow you to store 64-bit integers by conditions
	var result, num, mul int32
	numLen := int(math.Log10(math.Abs(float64(x))))

	for i := numLen; i >= 0; i-- {
		num = int32(x % 10)
		mul = int32(math.Pow10(i))

		if math.MaxInt32/mul < result/mul+num || math.MinInt32/mul > result/mul+num {
			return 0 // Int32 overflow
		}

		result += num * mul
		x /= 10
	}

	return int(result)
}

func Reverse(x int) int {
	if x >= -9 && x <= 9 {
		return x
	}

	if x > math.MaxInt32 {
		return 0 // Int32 overflow
	}

	val := int32(x)
	if val < 0 {
		val *= -1 // Make absolute value
	}

	// Assume the environment does not allow you to store 64-bit integers by conditions
	var result int32
	numLen := int(math.Log10(float64(val))) + 1
	for i := 1; i <= numLen; i++ {
		n := val % 10
		d := int32(math.Pow(10, float64(numLen-i)))
		if math.MaxInt32/d < result/d+n {
			return 0 // Int32 overflow
		}

		result += n * d
		val /= 10
	}

	if x < 0 {
		result *= -1 // Apply sign form source value
	}

	return int(result)
}
