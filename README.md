# Golang-to-do
API для управления CRUD TO-DO листа:
Стек:
- Go
- Gin
- PostgreSQL
- Gorm
- jwt-go
- bcrypt
- godotenv


## Подписание токена
Этот метод подписвает токен секретным ключом. SignedString берет данные header + payload, применяет к ним алгоритм подписи и подписывает по secret.
```go
secret := os.Getenv("JWT_SECRET")
return token.SignedString([]byte(secret))
```