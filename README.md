# go-jwt-authentication

A simple JWT authentication for users logging in and signing up

Contains validation as well by checking for existing users

Works using Gin-Gonic

There will be 2 main routes:
routes
|______auth
    |______user

# Architecture - Routes | Model | Controller
- Create Architecture of the application

|_controllers - Contains all the control logic of the application
|_database - Contains connection towards local mongo db
|_helpers
|_models - Contains model structures for the required application
|_routes - Contains routes for routing requests towards required controllers
|_main.go


# Step 1 - Build user model

|_controllers
|_database 
|_helpers
|_models
|   |____userModel.go
|_routes
|_main.go

# Step 2 - Initialise router in main.go

# Step 3 - Define routers under routes folder

|_controllers
|_database 
|_helpers
|_models
|   |____userModel.go
|_routes
    |___authRouter.go
    |___userRouter.go
|_main.go