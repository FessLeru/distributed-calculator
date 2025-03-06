# distributed-calculator
 Яндекс Лицей проект Голанг (Шорников Данила)

## Описание проекта

Этот проект представляет собой распределенный калькулятор, состоящий из оркестратора и воркеров. Оркестратор управляет задачами, а воркеры выполняют вычисления. Воркеры периодически запрашивают задачи у оркестратора, выполняют их и отправляют результаты обратно.

Проект реализован на языке Go и демонстрирует базовые принципы распределенных вычислений.

## Установка и запуск

### 1. Клонирование репозитория

Для начала работы с проектом, клонируйте репозиторий:

```bash
git clone https://github.com/FessLeru/distributed-calculator.git
cd distributed-calculator
```
### 2. Установка Postman
Для тестирования API рекомендуется использовать Postman. Скачайте и установите Postman с официального сайта (https://www.postman.com).

### 3. Запуск оркестратора и агента
Для запуска оркестратора выполните следующую команду в терминале:
```bash
cd ./cmd/orchestrator/
go run main.go
```
Откройте второй терминал с помощью + сверху и введите команду
```bash
cd ./cmd/agent/
go run main.go
```

### 4. Настройка и тестирование в Postman
Откройте Postman.

Создайте новый запрос, выбрав New -> HTTP.

Введите URL для отправки запроса: http://localhost:8080/api/v1/calculate.

Выберите метод POST.

Перейдите в раздел Body, выберите raw и введите один из следующих JSON-запросов:

Запрос с корректным выражением (код 200 + результат):
```json
{ "expression": "2+2*2" }
```

Функции
worker(): Функция, которая запускается в отдельной горутине и периодически запрашивает задачи у оркестратора, выполняет их и отправляет результаты обратно.

fetchTask(): Функция для получения задачи от оркестратора.

executeTask(): Функция для выполнения арифметической операции, указанной в задаче.

submitTaskResult(): Функция для отправки результата выполнения задачи обратно оркестратору.


