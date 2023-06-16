# go-jwt-authentication

This repository provides a simple JWT authentication system for user login and sign up. It includes user validation by checking for existing users and is built using Gin-Gonic.

## Architecture

The application follows the following architecture:

```
routes
|__auth
   |__user

Controllers - Contains all the control logic of the application
Database - Contains the connection to the local MongoDB
Helpers - Utility functions and helper modules
Models - Contains model structures for the application
Routes - Contains routes for routing requests to the appropriate controllers
main.go - Entry point of the application
```

## Step 1 - Build User Model

To begin, create the user model:

```
controllers
database
helpers
routes
main.go
models
  |_ userModel.go
```

## Step 2 - Initialize Router in main.go

Next, initialize the router in the `main.go` file.

## Step 3 - Define Routers under the Routes folder

Define the routers for authentication and user management in separate files under the `routes` folder:

```
controllers
database
helpers
routes
  |_ authRouter.go
  |_ userRouter.go
main.go
models
```

This structure organizes the application's components in a clear and maintainable manner, separating concerns and promoting modularity.
