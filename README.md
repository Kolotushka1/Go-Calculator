# Go Calculator Service
### Описание
Go Calculator Service — это простой HTTP-сервис для вычисления математических выражений. Сервис принимает POST-запросы с математическими выражениями и возвращает результат вычислений в формате JSON.

### Установка и запуск
1. Клонируйте репозиторий:
```bash
git clone https://github.com/Kolotushka1/Go-Calculator.git
cd Go-Calculator
```
2. Соберите и запустите сервер:
```bash
go run ./cmd/calc_service/...
```

### Использование API с помощью curl

#### Эндпоинт
URL: http://localhost:8080/api/v1/calculate<br>
Метод: POST<br>
Заголовок: Content-Type: application/json<br>
Тело запроса:<br>
```json
{
    "expression": "your_expression"
}
```
#### Примеры запросов (Windows CMD)
1. Успешный расчет
```bash
curl --location 'http://localhost:8080/api/v1/calculate' --header 'Content-Type: application/json' --data '{"expression":"2+2*2"}'
```
2. Некорректное выражение
```bash
curl -X POST http://localhost:8080/api/v1/calculate ^ -H "Content-Type: application/json" ^ -d "{\"expression\": \"2+a\"}"
```
3. Неправильный формат JSON
```bash
curl -X POST http://localhost:8080/api/v1/calculate ^ -H "Content-Type: application/json" ^ -d "invalid json"
```
### Коллекция запросов PostMan
В папке third_party вы можете найти txt файл, который можно импортировать в PostMan для тестирования запросов.
### Запуск тестов
В проекте уже реализованы автоматические тесты для калькулятора и HTTP-обработчиков.
1. Запустите тесты калькулятора:
```bash
go test ./internal/calculator
```
2. Запустите тесты обработчиков:
```bash
go test ./internal/handlers
```
3. Запустите все тесты одновременно:
```bash
go test ./...
```
### Лицензия
Этот проект лицензирован под MIT License.

### Контакты
Для вопросов и предложений, пожалуйста, создайте issue на GitHub.

## Дополнительная информация
Если у вас возникли проблемы с запуском сервера или использованием API, пожалуйста, убедитесь, что:
* Порт 8080 свободен и не используется другими приложениями.
* Вы используете корректный синтаксис JSON в запросах.
* Вы отправляете запросы методом POST к правильному эндпоинту.
* Если ничего не помогает, попробуйте воспользоваться PostMan