# Go CRUD API

A simple and clean RESTful API built with **Golang (Gin)**, **MongoDB**, and **JWT Authentication**.  
This API allows you to register, log in, and perform user management operations securely.

---

## 📦 Features

- User registration and login with hashed passwords
- JWT-based authentication
- Full CRUD operations for users
- Input validation using `go-playground/validator`
- Swagger documentation with `swaggo/gin-swagger`
- MongoDB integration using `mongo-driver`

---

## 🚀 Getting Started

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
.

├── controller/     ( Route handlers)

├── model/          ( Data models and DTOs)

├── routes/         ( Route groups)

├── utils/          ( JWT utilities)

├── docs/           ( Swagger docs (auto-generated) )

├── main.go         ( Entry point)

└── go.mod

## 🤝 Contributing
Pull requests are welcome! For major changes, please open an issue first to discuss what you would like to change.



