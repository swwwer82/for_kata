package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

type Text struct {
	Content string
}

// Метод textModifier выполняет все необходимые преобразования текста
func (t *Text) textModifier() {
	// A. Удаление лишних пробелов
	t.Content = strings.Join(strings.Fields(t.Content), " ")

	// B. Перестановка символов вокруг знака минус
	runes := []rune(t.Content)
	for i := 1; i < len(runes)-1; i++ {
		if runes[i] == '-' {
			runes[i-1], runes[i+1] = runes[i+1], runes[i-1]
			runes = append(runes[:i], runes[i+1:]...) // Удаление знака минус
			i--                                       // Сдвиг индекса назад после удаления
		}
	}
	t.Content = string(runes)

	// C. Замена плюсов на восклицательные знаки
	t.Content = strings.ReplaceAll(t.Content, "+", "!")

	// D. Подсчет суммы цифр и их удаление
	sum := 0
	var result []rune
	for _, r := range t.Content {
		if unicode.IsDigit(r) {
			val, _ := strconv.Atoi(string(r))
			sum += val
		} else {
			result = append(result, r)
		}
	}

	t.Content = string(result)

	// Добавление суммы цифр, если она больше 0
	if sum > 0 {
		t.Content = t.Content + " " + strconv.Itoa(sum)
	}

	// Вывод результата
	fmt.Println(t.Content)
}

func main() {
	text := &Text{}
	// Создаем новый сканер для чтения из стандартного ввода
	scanner := bufio.NewScanner(os.Stdin)

	// Просим пользователя ввести строку
	fmt.Println("Введите строку:")

	for scanner.Scan() {
		text.Content = scanner.Text()
		text.textModifier()
	}
}
