# Task Manager API (на GO+PostgreSQL)

Этот проект — REST API для управления задачами  
Позволяет **создавать, просматривать, обновлять и удалять задачи**.  

## Функционал: 

1.Создание задач
2.Получение списка задач
3.Обновление задачи
4.Удаление задач


### 1️ **Склонировать репозиторий**
```sh
git clone https://github.com/Mutter0815/todo-app-rest-api.git
cd todo-app-rest-api
``` 
### 2 **Добавить данные для подключения БД**
Создай файл `.env` и заполни его:
```
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=пароль
DB_NAME=taskdb
```
### 3 **Запустить сервер**
```sh
go run main.go
```

---

## API Эндпоинты

### Получить все задачи
```http
GET/tasks
```
### Получить задачу по ID
```http
GET /task/{id}
```

### Добавить задачу
```http
POST /tasks
```
**Пример тела запроса**
```json
{
    "title":"Закончить проект",
    "description":"Дедлайн завтра",
    "is_completed":"false",
    "due_date":"2025-02-10T23:59:59Z"
}
```
### Обновить задачу
```http
PUT /tasks/{id}
```

### Удалить задачу
```http
DELETE /tasks/{id}
```