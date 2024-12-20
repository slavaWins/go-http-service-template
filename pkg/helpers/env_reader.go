package helpers

import (

    "log"
    "os"

    "github.com/joho/godotenv"
)


var envVariables map[string]string

func Init() {
}

// Get возвращает значение переменной окружения по ключу
func Env(key string) string {
    err := godotenv.Load()
    if err != nil {
    log.Fatalf("Ошибка загрузки .env файла: %v", err)
    }

    // Получаем значение переменной среды EXAMPLE_KEY
    exampleKey := os.Getenv(key)
    if exampleKey == "" {
       log.Fatal(key + " не найден в переменных окружения")
    }

    // Печатаем значение переменной в консоль
   // fmt.Println("EXAMPLE_KEY:", exampleKey)
    return exampleKey;
}
