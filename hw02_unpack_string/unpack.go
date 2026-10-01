package hw02unpackstring

import (
	"errors"
	"strings"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(srcString string) (string, error) {
	var destString strings.Builder // Сюда постепенно собираем итоговую строку

	// Руну результата откладываем, потому что следующая руна может оказаться цифрой-повторителем
	var last rune    // Последняя руна результата, которую мы пока не записали
	hasLast := false // Есть ли сейчас отложенная руна в last
	escaped := false // Перед текущей руной был слэш '\'

	for _, current := range srcString {
		if escaped { // Обрабатываем руну после слэша
			if current != '\\' && !isDigit(current) { // После слэша можно экранировать только цифру или слэш
				return "", ErrInvalidString
			}
			if hasLast { // Перед новой руной результата записываем предыдущую отложенную руну один раз
				destString.WriteRune(last)
			}
			last = current // Экранированная цифра считается символом результата, а не повторителем
			hasLast = true
			escaped = false
			continue
		}

		switch {
		case current == '\\':
			escaped = true // Слэш сам не пишем, он только включает режим экранирования
		case isDigit(current):
			if !hasLast { // Цифра может повторять только предыдущую отложенную руну
				return "", ErrInvalidString
			}
			count := int(current - '0')
			destString.WriteString(strings.Repeat(string(last), count)) // Для count == 0 руна исчезнет
			hasLast = false
		default:
			if hasLast { // Перед новой руной результата записываем предыдущую отложенную руну один раз
				destString.WriteRune(last)
			}
			last = current
			hasLast = true
		}
	}

	if escaped { // Если строка закончилась сразу после слэша, это ошибка
		return "", ErrInvalidString
	}

	if hasLast { // После цикла последняя руна результата могла остаться отложенной
		destString.WriteRune(last)
	}

	return destString.String(), nil
}

func isDigit(r rune) bool {
	return r >= '0' && r <= '9'
}
