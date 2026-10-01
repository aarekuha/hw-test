package hw03frequencyanalysis

import (
	"sort"
	"strings"
)

const resultLimit = 10

type Item struct {
	Key   string
	Value int
}

func Top10(src string) []string {
	// Посчитать количество слов
	counter := make(map[string]int)
	for _, value := range strings.Fields(src) {
		counter[value]++
	}
	// Отсортировать по количеству найденных
	items := make([]Item, 0, len(counter))
	for key, value := range counter {
		items = append(items, Item{
			Key:   key,
			Value: value,
		})
	}
	sort.Slice(items, func(i int, j int) bool {
		if items[i].Value == items[j].Value {
			return items[i].Key < items[j].Key
		}
		return items[i].Value > items[j].Value
	})
	// Формирование результата
	limit := min(resultLimit, len(items))
	res := make([]string, 0, resultLimit)
	for _, item := range items[:limit] {
		res = append(res, item.Key)
	}
	return res
}
