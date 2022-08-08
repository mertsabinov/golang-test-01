package array_01

func Sum(numbers []int) int {
	result := 0
	for _, number := range numbers {
		result += number
	}
	return result
}

func SumAll(numberList ...[]int) []int {
	// number list içerisinde dönen bir fonskyon yaz
	// içerisine girilen slice ların toplamını döndür
	// yeni bir listeye ekle
	var sums []int
	for _, n := range numberList {
		sums = append(sums, Sum(n))
	}
	return sums
}
