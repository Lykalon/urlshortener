package lib

import (
	"crypto/rand"
	"math/big"
	"strings"
)

func Generate() int64 {
	a, _ := rand.Int(rand.Reader, big.NewInt(int64(^uint64(0) >> 1)))

	//исправление последовательности битов в сгенерированном числе
	//необходимо чтобы ниже в методах не выходить за диапозон
	//значений массива по словарю. Элементов 63 всего, а макс число по маске 63
	//без обработки словил панику
	value := a.Int64()

	//оставляем только значащие биты (61-64 создают разные числа, которые алгоритм
	//сожрет как одну последовательность). Н а 3, так как и так без знака возврат
	value = value >> 3

	for i := range 10 {
		if value >> (6 * i) & 0b111111 == 0b111111 {
			value = value & ^(0b111111 << (6 * i))
		}

		//проверка на крайние "_"
		if (i == 0 || i == 9) && value >> (6 * i) & 0b111111 == 0b000000 {
			value = value | (0b000001 << (6 * i))
		}
	}

	//проверка на "__"
	for i := range 9 {
		if value >> (6 * i) & 0b111111 == 0b000000 && value >> (6 * (i + 1)) & 0b111111 == 0b000000 {
			value = value | (0b000001 << (6 * i))
		}
	}
	return value
}

func Encode(input string) (output int64) {
	var dict = "_abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	for i := range 10 {
		output = output | (int64(strings.Index(dict, string(input[i]))) << (6 * i))
	}

	return output
}

func Decode(input int64) (output string) {
	var dict = "_abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	str := make([]byte, 10)

	for i := range 10 {
		str[i] = dict[(input >> (6 * i)) & 0b111111]
	}

	return(string(str))
}