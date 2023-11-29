package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	rootDir, err := os.Getwd() // получаем текущую рабочую директорию
	if err != nil {
		fmt.Println("Ошибка при получении текущей директории:", err)
		return
	}

	var goFiles []string
	err = filepath.Walk(rootDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Println("Ошибка при поиске файлов:", err)
			return err
		}

		// Ищем только файлы с расширением .go
		if !info.IsDir() && filepath.Ext(path) == ".go" {
			goFiles = append(goFiles, path)
		}

		return nil
	})
	if err != nil {
		fmt.Println("Ошибка при поиске файлов:", err)
		return
	}

	f, err := os.Create("result.txt")
	if err != nil {
		fmt.Println("Ошибка при открытии файла с результатом:", err)
		return
	}
	defer f.Close()

	for _, file := range goFiles {
		// Определяем относительный путь к файлу относительно корня проекта
		relPath, err := filepath.Rel(rootDir, file)
		if err != nil {
			fmt.Println("Ошибка при определении относительного пути:", err)
			return
		}

		// Читаем содержимое файла
		content, err := os.ReadFile(file)
		if err != nil {
			fmt.Println("Ошибка при чтении файла:", err)
			return
		}

		// Выводим относительный путь к файлу и его содержимое
		fmt.Fprintf(f, "Относительный путь: %s\nСодержимое:\n%s\n\n\n\n\n\n", relPath, strings.TrimSpace(string(content)))
	}
}
