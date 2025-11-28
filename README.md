# Golang-to-do
API для управления CRUD TO-DO листа:
## Стек:
- Go
- Gin
- PostgreSQL
- Gorm
- jwt-go
- bcrypt
- godotenv

## Структура проекта
```go
golang-to-do/
│── cmd/                  # Точка входа приложения
│   └── app/
│       └── main.go        # запуск
│
│── internal/              # внутренняя логика
│   ├── models/            # модели
│   │   ├── task.go
│   │   └── user.go
│   │
│   ├── handlers/           # Слушатели (контроллеры)
│   │   ├── auth_handler.go      # Слушатель для управления авторизацией
│   │   └── task_handler.go      # Слушатель для управления задачами
│   │
│   ├── repositories/       # Репозитории
│   │   ├── auth_repository.go   # Запросы с БД, связанные с пользователями
│   │   ├── db.go                # Подключение к бд
│   │   └── task_repository.go   # Запросы в бд для задач
│   │
│   ├── requests/           # Запросы валидации для входящих данных
│   │   ├── login_request.go     
│   │   ├── register_request.go 
│   │   └── task_request.go      
│   │
│   ├── routes              # Роутинг (сборка всех эндпойнтов)
│   │   └── api.go
│   │
│   └── services            # Сервисы логики
│       ├── auth_service.go    # Логика авторизации
│       └── task_service.go    # Логика управления задачами
│
│── pkg/                    # Вспомогательные пакеты
│   └── middleware/           # Мидлвари для GIN
│       └── auth_middleware.go
│ 
│ 
│── go.mod
│── go.sum
```

## Gin роутинг
Для запуска gin локального сервера, нужно проинициализировать и настроить роутер. А также создать эндойнты и назначить на них обработчики (контроллеры). Все это делаем в файле `routers/api.go`

### Функция конструктор для создания роутера
```go
func NewRouter(taskHandler *handlers.TaskHandler, authHandler *handlers.AuthHandler) *Router {
	return &Router{
		taskHandler: taskHandler,
		authHandler: authHandler,
	}
}
```
### Метод для установки эндпойнтов (кратко)
```go
func (r *Router) SetupRoutes() *gin.Engine {
    // Инициализация роутера
    e := gin.Default()
    // Создание групп 
    api := e.Group("/api")
    tasks := api.Group("/tasks")
    // Вешание мидлвара на группу роутов
    tasks.Use(middleware.AuthMiddleware(r.authHandler.As.Ar))
    // Формирование эндпойнтов и их обработчиков
    {
	    tasks.GET("/", r.taskHandler.GetAll)
	    tasks.GET("/:id", r.taskHandler.GetByID)
	    tasks.PUT("/:id", r.taskHandler.Update)
	    tasks.POST("/", r.taskHandler.Create)
	    tasks.DELETE("/:id", r.taskHandler.Delete)
    }

    return e
}
```
p.s фигурные скобки чисто для группировки, их можно не использовать.
```go 
// Пример использования мидлвара на отдельном роуте
auth.GET("/me", middleware.AuthMiddleware(r.authHandler.As.Ar), r.authHandler.Me)
```
### Запуск
Для запуска используем метод Run с переданным портом в формате `":port"`
```go 
	r := routes.NewRouter(th, ah)
	e := r.SetupRoutes()
	e.Run(":8080")
```
### Handler 
`Handler` - слушатель, обработчик входящих эндпойнтов, аналог контроллеров в Laravel. Каждый метод хэндлера должен иметь сигнатуру `gin.Context`- это объект, который представляет один http запрос (Аналог Request), который содержит всю инфу о нем. Его параметры, тело, заголовки и тд
```go
func (th *TaskHandler) GetAll(c *gin.Context) {
	tasks, err := th.TaskService.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, tasks)
}
```
Для того, чтобы отдать ответ используем `c.JSON(200,data)`. Это аналог `response()->json($data,200)`, но только без `return`.
<br> Примеры функций gin.Context:
 - `c.Param("id")` - получение параметра из эндпойнта `/task/1`
 - `c.ShouldBindJSON(&input)` - распаковывает тело запроса в структуру input
 - `c.PostForm("name") - содержит данные формы`
 - `c.JSON(код, data)` - отправляет JSON ответ
 - `c.GetHeader("Authorization")` - получение заголовка запроса
 - `c.Set()` ставит данные глобально в gin.Context
 - `c.Get()`- получает данные из глобального хранилища в gin.Context

 c.Set() , c.Get() используются для хранения переменных в рамках одного запроса. Т.е можно в мидлваре закидывать туда данные и потом юзать в хендлере.

### JWT токен и авторизация по нему
Для реализации авторизации по JWT токену используется пакет `github.com/golang-jwt/jwt/v5`. Пример реализации метода Login AuthService. Все проверки по email и сверка паролей происходит вручную.
```go
func (as *AuthService) Login(lr requests.LoginRequest) (string, error) {
	user, err := as.Ar.GetByEmail(lr.Email)
	if err != nil {
		return "", err
	}

	err = as.checkPasswordHash(lr.Password, user.Password)

	if err != nil {
		return "", err
	}

	token, err := as.GenerateToken(user.ID, user.Email)

	return token, err
}
``` 
Метод `GenerateToken`. Для того, чтобы проверить jwt токен на подлинность нужен `JWT_SECRET `- специальный ключ, который используется для шифрования и генерации токена. Он храниться в `env`. 
<br> `claims` - данные, зашитые в токен. Для того чтобы их упаковать в правильный массив используется `jwt.MapClaims{}`. ДАлее для генерации токена используем:
```go
token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
```
Где `jwt.SigningMethodHS256` - метод шифрования. А `claims` - данные.
Но чтобы токен функционировал и был коректным, его нужно подписать. (Так сказать закрыть на ключ). Далее при использовании, этот же токен будет проверяться на подлинность (открываться) с помощью jwt_secret.
```go
return token.SignedString([]byte(secret))
```

Полный метод `GenerateToken`
```go
func (as *AuthService) GenerateToken(id int64, email string) (string, error) {
	secret := os.Getenv("JWT_SECRET")
	claims := jwt.MapClaims{
		"id":    id,
		"email": email,
		"exp":   time.Now().Add(time.Hour * 24).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(secret))
}
```

Проверка токена выполняется с помощью методов `jwt.Parse()` - он преобразовывает текстовый токен в структурный вид и проверяет его подпись. Итоговый токен имеет различные методы для проверок. Например:
- `token.Method.(*jwt.SigningMethodHMAC)` - проверяет зашифрован ли токен указанным способом. 
- `jwt.ErrSignatureInvalid`- возвращает ошибку токена.
- `token.Valid` - проверяет валидность токена и возвращает true/false.
-  `token.Claims.(jwt.MapClaims)` - получаем вшитые данные. `.(jwt.MapClaims)` - приведение данных к указанному типу.

<br>Полный код проверки:
```go
//Декодируем и проверяем токен
token, err := jwt.Parse(tokenStr,
	func(token *jwt.Token) (interface{}, error) {
	//Получаем инфу о методе шифрования токена
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, jwt.ErrSignatureInvalid
			}
		//Возвращаем секрет для проверки подписи
		return jwtSecret, nil
	})

		//Проверяем результаты разкодирования токена и его валидность
		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
			c.Abort()
			return
		}
```
P.S `jwtSecret` секрет лучше всего вытягивать из енв в начале мидлвар-функции и преобразовывать его в байтовый вид, перед тем как передавать его в функцию:
```go
jwtSecret := []byte(os.Getenv("JWT_SECRET"))
```
### Middleware
`Middleware` в gin работают по такому же принципу как в `Laravel`. Пример мидлвара.
```go
func AuthMiddleware(ar *repositories.AuthRepository) gin.HandlerFunc {
	return func(c *gin.Context) {}
}
```
Все мидлвари должны реализовывать колбэк функцию с gin.Context. И дальше в нем пишем то же, что и в обычных слушателях. Но только все проверки нужно прописывать вручную. Примеры некоторых проверок:
```go
//Получаем Authorization заголовок из запроса
authHeader := c.GetHeader("Authorization")
if authHeader == "" {
		c.JSON(401, gin.H{"error": "Authorization header is missing"})
		c.Abort()
		return
	}
// Разделяем содержимое заголовка и проверяем формат
parts := strings.SplitN(authHeader, " ", 2)
if len(parts) != 2 || parts[0] != "Bearer" {
		c.JSON(401, gin.H{"error": "Authorization header format must be Bearer {token}"})
		c.Abort()
		return
	}
```
Команды для управления запросом
```go
c.Abort() // Прерывает http запрос
c.Next()  // Продолжает выполнение http запроса
```
### Request валидация
Валидация реализуется через структуры + теги, а потом используются специальные пакеты, интегрированные в gin.
Пример структуры запроса:
```go
type RegisterRequest struct {
	Name            string `json:"name" binding:"required,min=3,max=50"`
	Lastname        string `json:"lastname" binding:"required,min=3,max=50"`
	Patronymic      string `json:"patronymic" binding:"omitempty,max=50"`
	Email           string `json:"email" binding:"required,email"`
	Age             int    `json:"age" binding:"omitempty,min=0,max=150"`
	Password        string `json:"password" binding:"required,min=6,max=100"`
	PasswordConfirm string `json:"password_confirm" binding:"required,eqfield=Password"`
}
```
И менно ее заполняет данными и затем запускает валидацию метод `c.ShouldBindJSON`. Сами правила валидации описываются в тег `binding:""`. Правила почти такие же, как и в `Laravel`.
P.S Валидация происходит в момент вызова `c.ShouldBindJSON`
