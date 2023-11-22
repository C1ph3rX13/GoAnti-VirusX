package Utils

import (
	"fmt"
	"math/rand"
	"strings"
)

func RandomByte(length int, random *rand.Rand) ([]byte, error) {
	// 创建一个长度为length的字节切片，用来存储生成的随机字节
	randByte := make([]byte, length)

	// 调用Rand类型的Read方法，将随机字节写入切片中
	_, err := random.Read(randByte)
	if err != nil {
		return nil, fmt.Errorf("failed to generate random bytes: %w", err)
	}

	return randByte, nil
}

func RandomStrings(length int, random *rand.Rand) string {
	// 将letters定义为常量，去掉数字字符
	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	// 使用strings.Builder类型
	var rStrings strings.Builder
	// 预分配内存
	rStrings.Grow(length)
	for i := 0; i < length; i++ {
		// 使用rand.Int31n函数
		rStrings.WriteByte(letters[random.Int31n(int32(len(letters)))])
	}

	return rStrings.String()
}
