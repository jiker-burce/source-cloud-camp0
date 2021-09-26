package cncamp

import "fmt"

func replaceByRepArr(oriArr []string, replaceArr []string) []string {
	for i := 0; i < len(oriArr); i++ {
		tmp := replaceArr[i]
		if tmp != "" {
			oriArr[i] = tmp
		}
	}

	return oriArr
}

func replaceByRepArr2(oriArr []string, replaceArr []string) []string {
	for i := 0; i < len(oriArr); i++ {
		if replaceArr[i] == "" {
			continue
		}
		oriArr[i] = replaceArr[i]
	}

	return oriArr
}

func replaceByIndex(oriArr []string, values map[int]string) []string {
	for key, value := range values {
		oriArr[key] = value
	}

	return oriArr
}

func swapValue() {
	a := 100
	a <<= 4 // 2的4次方 * 100 = 1600
	b := 100
	b >>= 4 // 100 / 2的4次方 = 16 （16为取整后的数据）
	//a = a ^ b
	//b = b ^ a
	//a = a ^ b

	//a, b = b, a

	fmt.Print(a, b)
}
