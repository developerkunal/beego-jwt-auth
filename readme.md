# 🚀 **Beego JWT Authentication Example**

This repository demonstrates **JWT authentication** in a **Beego application** using middleware to validate tokens. It includes:

✅ **Middleware for JWT authentication**  
✅ **Extracting Go’s standard context from Beego**  
✅ **Public & protected API routes**

---

## 📌 **Project Structure**

```
/beego-jwt-auth
│── main.go             # Entry point
│── auth/
│   ├── auth.go         # JWT token generation
│── middleware/
│   ├── middleware.go   # JWT validation middleware
│── go.mod              # Go module file
│── go.sum              # Dependencies
│── README.md           # Documentation
```

---

## 🚀 **Setup & Installation**

### **1️⃣ Clone the repository**
```sh
git clone https://github.com/developerkunal/beego-jwt-auth.git
cd beego-jwt-auth
```

### **2️⃣ Install dependencies**
```sh
go mod tidy
```

### **3️⃣ Run the server**
```sh
go run main.go
```
Server will start on **http://localhost:8080** 🚀

---

## 🔥 **API Endpoints**

### ✅ **Public Route (No Authentication Required)**
```sh
curl -X GET http://localhost:8080/
```
**Response:**
```json
🏠 Welcome to the Beego API
```

### 🔑 **Generate JWT Token**
```sh
curl -X GET http://localhost:8080/login
```
**Response:**
```json
{
  "token": "your.jwt.token"
}
```

### 🔐 **Access Protected Route (With Token)**
```sh
curl -X GET http://localhost:8080/protected -H "Authorization: Bearer your.jwt.token"
```
If token is **valid**, response:
```json
🔓 Token Validated!
```
If token is **invalid or missing**, response:
```json
⛔ Invalid or expired token: ERROR
```

---

## 🛠 **How It Works?**

### **1️⃣ Extract Standard Go Context from Beego**
```go
func GetStandardContext(ctx *beegoCtx.Context) context.Context {
	if ctx != nil && ctx.Request != nil {
		return ctx.Request.Context()
	}
	return context.Background()
}
```

### **2️⃣ Middleware to Validate JWT**
```
func JWTMiddleware(ctx *beegoCtx.Context) {
stdCtx := GetStandardContext(ctx) // Convert Beego context to standard Go context

	token := ctx.Input.Header("Authorization") // Get token from request header
	if token == "" {
		ctx.Output.SetStatus(http.StatusUnauthorized)
		ctx.Output.Body([]byte("⛔ Missing token"))
		return
	}

	// Remove "Bearer " prefix if present
	token = strings.TrimPrefix(token, "Bearer ")

	fmt.Println("🔍 Received Token:", token) // Debugging line

	// Ensure validator is properly initialized before use
	if jwtValidator == nil {
		ctx.Output.SetStatus(http.StatusInternalServerError)
		ctx.Output.Body([]byte("❌ JWT Validator not initialized"))
		return
	}

	// Validate the token
	_, err := jwtValidator.ValidateToken(stdCtx, token)
	if err != nil {
		ctx.Output.SetStatus(http.StatusUnauthorized)
		ctx.Output.Body([]byte("⛔ Invalid or expired token: " + err.Error()))
		return
	}

	// If validation succeeds, allow request
	ctx.Output.SetStatus(http.StatusOK)
	ctx.Output.Body([]byte("🔓 Token Validated!"))
}
```
---

## 📚 **Resources**

- [Beego - Official Github](https://github.com/beego/beego)
- [JWT.io - Official Website](https://jwt.io/)
- [Go JWT Middleware](https://github.com/auth0/go-jwt-middleware)
