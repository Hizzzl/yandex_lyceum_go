```markdown
# Calc Service

Calc Service — это простой и надежный веб-сервис для вычисления арифметических выражений. Пользователи могут отправлять арифметические выражения через HTTP-запросы и получать результаты вычислений в формате JSON.

## Начало работы

### Требования

- **Go:** Версия 1.16 или выше

### Установка

1. **Клонируйте репозиторий:**

   ```bash
   git clone https://github.com/yourusername/calc_service.git
   ```

2. **Перейдите в каталог проекта:**

   ```bash
   cd calc_service
   ```

3. **Запустите сервис:**

   ```bash
   go run ./cmd/calc_service/...
   ```

   После запуска, сервер будет доступен по адресу `http://localhost:8080`.

## Использование

### Endpoint

- **URL:** `/api/v1/calculate`
- **Метод:** `POST`
- **Content-Type:** `application/json`

### Тело запроса

Запрос должен содержать JSON-объект с полем `expression`, в котором указано арифметическое выражение для вычисления.

**Пример:**

```json
{
    "expression": "2+2*2"
}
```

### Примеры запросов

#### Успешный расчет

**Описание:** Вычисление выражения `2+2*2` должно вернуть результат `6`.

**Команда `curl`:**

```bash
curl --location 'http://localhost:8080/api/v1/calculate' \
--header 'Content-Type: application/json' \
--data '{
  "expression": "2+2*2"
}'
```

**Ответ:**

```json
{
    "result": "6"
}
```

**HTTP-статус:** `200 OK`

#### Ошибка 422: Неверные данные

**Описание:** Отправка выражения с недопустимыми символами, например, `2+2a*2`, должна вернуть ошибку.

**Команда `curl`:**

```bash
curl --location 'http://localhost:8080/api/v1/calculate' \
--header 'Content-Type: application/json' \
--data '{
  "expression": "2+2a*2"
}'
```

**Ответ:**

```json
{
    "error": "Expression is not valid"
}
```

**HTTP-статус:** `422 Unprocessable Entity`

#### Ошибка 500: Внутренняя ошибка сервера

**Описание:** В случае возникновения непредвиденной ошибки сервер вернет соответствующий ответ.

**Команда `curl`:**

*В реальной ситуации трудно симулировать внутреннюю ошибку сервера через простой `curl` запрос. Однако, вы можете изменить код сервера так, чтобы он вызывал ошибку, и протестировать соответствующий ответ.*

**Пример изменения кода для тестирования:**

```go
// Внутри функции calculateHandler, добавьте условие для вызова 500 ошибки
if reqBody.Expression == "trigger_error" {
    w.WriteHeader(http.StatusInternalServerError)
    json.NewEncoder(w).Encode(ErrorResponse{Error: "Internal server error"})
    return
}
```

**Команда `curl`:**

```bash
curl --location 'http://localhost:8080/api/v1/calculate' \
--header 'Content-Type: application/json' \
--data '{
  "expression": "trigger_error"
}'
```

**Ответ:**

```json
{
    "error": "Internal server error"
}
```

**HTTP-статус:** `500 Internal Server Error`