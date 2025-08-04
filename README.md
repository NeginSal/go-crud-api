## 🚀 Go CRUD API

A simple and clean RESTful API built with **Golang (Gin)**, **MongoDB**, and **JWT Authentication**.  
Now fully Dockerized for better development and deployment experience.  
This API allows you to **register**, **log in**, and **manage users securely**.

---

## 📦 Features

- ✅ User registration and login with hashed passwords (bcrypt)
- ✅ JWT-based authentication middleware
- ✅ Full CRUD operations on users
- ✅ Input validation with `go-playground/validator`
- ✅ Swagger documentation (`swaggo/gin-swagger`)
- ✅ MongoDB integration using `mongo-driver`
- ✅ Docker + Docker Compose support

---

## 🐳 Run with Docker (Recommended)

### 1. Clone the repository

```
git clone https://github.com/NeginSal/go-crud-api.git
cd go-crud-api
```
### 2. Set environment variables in .env file
```
PORT=8080
MONGO_URI=mongodb://mongo:27017
JWT_SECRET=your_super_secret_key
```
### 3. Run with Docker Compose
```
docker-compose up --build
```
Access the API at: `http://localhost:8080`
Swagger UI at: `http://localhost:8080/swagger/index.html`



### 1. Clone the repository

```bash
git clone https://github.com/NeginSal/go-crud-api.git
cd go-crud-api
```
### 2. Install dependencies
```
 go mod tidy
```
### 3. Set up environment variables
```
MONGODB_URI=
JWT_SECRET=your_jwt_secret
PORT=8080
```
### 4. Run the application
```
go run main.go
```
## 🧪 API Documentation
Once the server is running, you can access Swagger UI at: ``` http://localhost:8080/swagger/index.html```

The API is documented using swaggo/swag. To regenerate docs after changes : ``` swag init ```

## 🔐 Authentication
- Register with /signup
- Log in with /login
- Use the returned token (JWT) in the Authorization header for protected routes: Authorization: Bearer <your-token>

## 📁 Project Structure
go-crud-api/

│

├── main.go

├── config/          ← Database settings and ...

│   └── db.go

├── controller/      ← Controllers (logic related to APIs)

│   └── user.go

├── model/           ← Data structures and models

│   └── user.go

├── middleware/      ← Middlewares (like auth)

│   └── authMiddleware.go

├── routes/          ← Routing

│   └── userRoutes.go

├── utils/           ← Helper functions (like ValidateToken)

│   └── jwt.go

├── go.mod / go.sum  ← Package information

└── README.md        ← Project description 


## 📖 Learn More
You can learn more about this project in this link [go-crud-api](https://dev.to/negin/a-crud-api-with-go-using-the-gin-framework-and-mongodb-379e).

## 🤝 Contributing
Pull requests are welcome! For major changes, please open an issue first to discuss what you would like to   change.