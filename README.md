# Go infotecs

Тестовое задание на Go разработчика в ИнфоТеКС


# Запуск api
1) **Измените путь в БД в environment/.env**
2) Скомпилируйте api -```go build main.go```
3) Запустите api

P.S - миграции и сидинг бд автоматический


# Тестирование
```bash 
    go test -v ./...
```

# Документация к API

### 1. Поиск кошелька по номеру

```http request
GET /api/{number}/balance
```
### path параметры
| Параметр | Тип      | Описание                        |
|:---------|:---------|:--------------------------------|
| `number` | `string` | **Обязательно**. Номер кошелька |


### Ответ
```json
{
  "number": string,
  "balance": float
}
```
### Статус коды
| Код | Описание             |
|:----|:---------------------|
| 200 | `OK`                 |
| 400 | `BAD REQUEST` |
| 404 | `NOT FOUND`          |
| 405 | `METHOD NOT ALLOWED` |

### 2. перевод денег с одного кошелька на другой

```http request
GET /api/send
```

### Тело запроса
```json
{
  "from": string,
  "to": string,
  "amount": float
}
```

### Ответ
```json
{
  "detail": "success"
}
```
### Статус коды
| Код | Описание             |
|:----|:---------------------|
| 200 | `OK`                 |
| 400 | `BAD REQUEST`        |
| 404 | `NOT FOUND`          |
| 405 | `METHOD NOT ALLOWED` |

### 3. Список последних транзакций

```http request
GET /api/transactions?count=
```

### Query параметры
| Параметр | Тип   | Описание                                         |
|:---------|:------|:-------------------------------------------------|
| `count`  | `int` | **Обязательно**. количество выводимых транзакций |


### Ответ
```json
[
  {
    "sender": string,
    "recipient": string,
    "amount": float,
    "time": timestamp
  }
]
```
### Статус коды
| Код | Описание             |
|:----|:---------------------|
| 200 | `OK`                 |
| 400 | `BAD REQUEST`        |
| 404 | `NOT FOUND`          |
| 405 | `METHOD NOT ALLOWED` |
