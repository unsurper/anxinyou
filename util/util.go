package util

import (
	"anxinyou/model"
	"math/rand"
	"sort"
	"time"
)

func RandomString(n int) string {
	//随机字符串
	var letters = []byte("asdfghjklzxcvbnmqwertyuiopASDFGHJKLZXCVBNMQWERTYUIOP")
	result := make([]byte, n)

	rand.Seed(time.Now().Unix())
	for i := range result {
		result[i] = letters[rand.Intn(len(letters))]
	}
	return string(result)
}

func SortByID(u []model.Chat) {
	sort.Slice(u, func(i, j int) bool { // asc
		return u[i].ID < u[j].ID
	})
}
