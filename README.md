# gin-jwt

## Overview

Skeleton code for API with JWT authentication in GO.

Direct Dependencies:
- Gin
- Gorm
- Godotenv
- Bcrypy
- Golang-jwt

## CI/CD architecture

Platform: Github Actions

The pipeline is triggered by a push commit in `develop` branch. This includes the PRs merged into the branch. `develop` branch is a protected branch also known as the integration branch. All feature branches should be checked out and merged back from `develop` branch after being tested and approved by the relevent code owners. Currently CI/CD setup performs the following steps:

### CI
- Checkout repository
- Perform unit testing and linting (no unit tests in project as of this moment)
- Login to Docker Hub
- Builds Docker image using approved code
- Pushes Docker image to docker hub

### CD
- Deploy to EC2 (Explained all the steps, Does not actually deploy)
- Deploy to Kubernetes (Explained all the steps, Does not actually deploy)

### Health check
- Make an HTTP request to health check url using curl or any other utility
- If response is in 2xx range then returns 0 else return non zero exit code
- As non zero exit code will fail this job hence pipeline fails


## Routes

### 1. /ping

- Description: This route serves as a health check endpoint to check if the application is running and healthy.
- Method: GET
- Response: JSON-encoded response with a success message.

### 2. /signup

- Description: Allows users to register by providing their email, password, username, and display name in JSON format.
- Method: POST
- Request Body: JSON with email, password, username, and displayname.
- Response: JSON response indicating the success or failure of the registration process.

### 3. /login

- Description: Handles user authentication by expecting email and password in JSON format. If authentication is successful, it sets a JWT token in the cookie.
- Method: POST
- Request Body: JSON with email and password.
- Response: JSON response with the JWT token upon successful login.

### 4. /dummy

- Description: This route demonstrates a protected page that requires authentication. Users must be logged in using the "/login" route to access this endpoint.
- Method: GET
- Authentication: JWT token must be provided in the request headers.

## Getting Started

1. Clone the repository:

```shell
git clone https://github.com/yourusername/gin-jwt.git
cd gin-jwt
```
2. Install dependencies:

```shell
go mod tidy
```
3. Set up environment variables by creating a .env file with the following content:

```shell
# .env
SECRET=your_secret_key
DB=your_data_source_name
```
4. Start the application:

```shell
go run main.go
```

The application will be accessible at http://localhost:8080.
