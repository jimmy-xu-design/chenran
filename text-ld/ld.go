package text_ld

import "fmt"

func Min(a, b, c int) int {
	tmp := a

	if tmp > b {
		tmp = b
	}

	if tmp > c {
		tmp = c
	}

	return tmp
}

func build_ld_matrix(result []rune, tag []rune) [][]int {
	len1 := len(result)
	len2 := len(tag)

	var ld [][]int

	ld = make([][]int, len1 + 1)
	for i := 0; i < len1 + 1; i++ {
		ld[i] = make([]int, len2 + 1)
	}

	// LD(0, 0) = 0
	ld[0][0] = 0

    // LD(i, 0) = i
	for i := 0; i < len1; i++ {
		ld[i + 1][0] = i + 1
	}

	// LD(0, j) = j
	for j := 0; j < len2; j++ {
		ld[0][j + 1] = j
	}

	// ai=bj  => LD(i,j)=LD(i-1,j-1)
	// ai!=bj => LD(i,j)=Min(LD(i-1,j-1),LD(i-1,j),LD(i,j-1)) + 1
	for i := 0; i < len1; i++ {
		for j := 0; j < len2; j++ {
			if result[i] == tag[j] {
				ld[i + 1][j + 1] = ld[i][j]
			} else {
				ld[i + 1][j + 1] = Min(ld[i][j], ld[i][j + 1], ld[i + 1][j]) + 1
			}
		}
	}

	return ld
}

func LD(A string, B string) (H int, D int, S int, I int, N int) {
	str1 := []rune(A)
	str2 := []rune(B)

	N = len(str1)
	ld := build_ld_matrix(str1, str2)

	// 从终点回溯
	index1 := len(str1)
	index2 := len(str2)

	var resultA []rune
	var resultB []rune

	for index1 > 0 && index2 > 0 {

		// ai=bj 回溯至左上角
		if str1[index1 - 1] == str2[index2 - 1] {
			resultA = append(resultA, str1[index1 - 1])
			resultB = append(resultB, str2[index2 - 1])

			index1--
			index2--

			H++
		} else {
			// 从坐上，上，左中找最小值
			if index1 == 0{
				resultA = append(resultA, rune('_'))
				resultB = append(resultB, str2[index2 - 1])

				index2--
				D++
			} else if index2 == 0 {
				resultA = append(resultA, str1[index1 - 1])
				resultB = append(resultB, rune('_'))

				index1--
				I++
			} else {
				min := Min(ld[index1 - 1][index2 - 1], ld[index1 - 1][index2], ld[index1][index2 - 1])
				if min == ld[index1 - 1][index2 - 1] {
					resultA = append(resultA, str1[index1 - 1])
					resultB = append(resultB, rune('*'))

					index1--
					index2--
					S++
				} else if min == ld[index1][index2 - 1] {
					resultA = append(resultA, rune('_'))
					resultB = append(resultB, str2[index2 - 1])

					index2--
					D++
				} else {
					resultA = append(resultA, str1[index1 - 1])
					resultB = append(resultB, rune('_'))

					index1--
					I++
				}
			}
		}
	}

	fmt.Print("A: ")
	for i := len(resultA) - 1; i >= 0; i-- {
		fmt.Print(string(resultA[i]))
	}
	fmt.Print("\nB: ")
	for i := len(resultB) - 1; i >= 0; i-- {
		fmt.Print(string(resultB[i]))
	}
	fmt.Println()

	return H, D, S, I, N
}