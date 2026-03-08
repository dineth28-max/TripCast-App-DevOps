# Comprehensive Implementation Guide: TripCast Auth Service

Welcome to the `auth-service` documentation! This guide explains *how* the service is built, *why* it's structured this way, and *how* to troubleshoot common issues. It's written specifically to help you learn the implementation details.

---

## 1. System Architecture (Layered Pattern)

We used a **4-Layer Architecture** pattern. This separates concerns so the code is easy to read, test, and maintain. 

Here is the flow of an API Request:
**`Routes` ➔ `Controllers` ➔ `Services` ➔ `Repository` (➔ MongoDB)**

### Layer Breakdown:
1.  **`routes/` (The Map):**
    *   Defines the API endpoints (e.g., `POST /auth/login`).
    *   It only maps a URL to a specific function in the `conteollers` folder.
2.  **`conteollers/` (The Traffic Cop):**
    *   Catch the HTTP request from the user.
    *   Read the JSON Body (`ctx.ShouldBindJSON`).
    *   Call the `services/` layer to do the actual work.
    *   Return the JSON response (`200 OK`, `400 Bad Request`, etc.).
3.  **`services/` (The Brain - Business Logic):**
    *   *This is where the real work happens.* 
    *   It hashes passwords (`bcrypt`), generates JWT tokens, and enforces business rules (like checking if an email already exists).
4.  **`repository/` (The Data Vault):**
    *   This is the ONLY layer that talks to the database (MongoDB). 
    *   If we ever switch from MongoDB to PostgreSQL, we *only* have to rewrite this folder. The rest of the app doesn't care!

---

## 2. Analyzing & Fixing the Errors You Encountered

You encountered three specific HTTP error codes while testing. Here is why they happened and how the code handles them:

### A. The `400 Bad Request` (EOF Error)
*   **What you did:** You sent a `POST` request to `/auth/register` without putting any JSON text in the Postman "Body" tab.
*   **What happened:** The server tried to read the JSON text to figure out the `email` and `password`, but reached the "End of File" (EOF) because it was blank.
*   **How we fixed the code quality:** In `conteollers/auth.go`, if `ctx.ShouldBindJSON` fails, we now return a highly descriptive error message:
    `"Invalid request format or missing required fields. Ensure you are sending JSON with email, password (min 6 chars), and name."`

### B. The `404 Not Found` Error
*   **What you did:** You sent a `GET` request to the base URL `http://localhost:8888/`
*   **What happened:** We hadn't defined a route for the root path `/`. 
*   **How we fixed it:** We added a `router.GET("/")` in `routes/auth_routes.go` that returns a friendly "Welcome to the API" message.

### C. The `401 Unauthorized` Error
*   **What you did:** You called `GET /auth/validate` without passing the token.
*   **What happened:** The API requires a JWT token to prove who you are.
*   **How it works in code:** In `conteollers/auth.go`, the `Validate` function explicitly looks for the HTTP Header: `Authorization: Bearer <token_string>`. If it's missing or invalid, it rejects the request with `401`.

---

## 3. The API Gateway Strategy

Currently, you are hitting the endpoints directly on port `8888`:
`http://localhost:8888/auth/login`

**Future Strategy (API Gateway):**
In a microservices architecture, the React Frontend will **never** talk directly to Port `8888`. Instead:
1.  React talks to the **API Gateway** (e.g., running on Port `5000`).
2.  React asks the Gateway for `/api/auth/login`.
3.  The Gateway looks at its map and says: *"Ah, anything starting with `/api/auth` goes to the auth-service container on port `8888`!"*
4.  The Gateway strips the `/api` and forwards the request to `auth-service` as just `/auth/login`.

*This is why we removed the `/api` prefix from our `routes_auth.go` file!*

---

## 4. How to Run Locally

### Option 1: Docker (Recommended)
This spins up both your Go App and (eventually) your MongoDB.
```bash
# From the project root (TripCast-App-DevOps/)
docker-compose up --build -d auth-service
```

### Option 2: Pure Go (For Development)
Ensure your local MongoDB Server (Compass) is running on port 27017.
```bash
cd services/auth-service
go run main.go
```

The server will start on `http://localhost:8888`. Every time you change Go code, stop the server (`Ctrl+C`) and run `go run main.go` again to see the changes.
