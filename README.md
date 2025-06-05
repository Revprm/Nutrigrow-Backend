# Nutrigrow-Backend

## Overview
A Backend Repository for [Nutrigrow](https://github.com/Logiqode/Nutrigrow) Project. This backend is built with Go, Gin, and GORM, following a Clean Architecture approach. It provides a robust API for managing user data, stunting records, news, food recipes, and their ingredients.

## Features
- **User Management**: Register, Login, Refresh Token, User Profile (Me), Update/Delete User, Email Verification.
- **Stunting Records**: Create, Retrieve (by ID, by User ID, latest by User ID), Update, Delete stunting records.
- **Stunting Prediction**: API endpoint to predict stunting status using an external ML model.
- **News Management**: Create, Retrieve (by ID, by Category, All with Pagination), Update, Delete news articles.
- **News Category Management**: Create, Retrieve (by ID, by Name, All), Update, Delete news categories.
- **Food Recipes Management**: Create, Retrieve (by ID, by Name, All with Pagination, by Ingredient), Update, Delete food recipes.
- **Ingredient Management**: Create, Retrieve (by ID, by Name, All with Pagination), Update, Delete ingredients.
- **Database Migrations & Seeding**: Tools to manage database schema and populate initial data.
- **Authentication**: JWT-based authentication for secure API access.
- **Logging**: Built-in system for monitoring and tracking system queries with a web interface.

## Prerequisite 🏆
- Go Version `>= go 1.20`
- PostgreSQL Version `>= version 15.0`
- Docker & Docker Compose (for Dockerized setup)

## How To Use
1. Clone the repository
  ```bash
  git clone https://github.com/Revprm/Nutrigrow-Backend.git
  ```
2. Navigate to the project directory:
  ```bash
  cd go-gin-clean-starter
  ```
3. Copy the example environment file and configure it:
  ```bash 
  cp .env.example .env
  ```
There are 2 ways to do running
### With Docker
1. Build Docker
  ```bash
  make up
  ```
2. Run Initial UUID V4 for Auto Generate UUID
  ```bash
  make init-uuid
  ```
3. Run Migration and Seeder
  ```bash
  make migrate-seed
  ```

### Without Docker
1. Configure `.env` with your PostgreSQL credentials:
  ```bash
  DB_HOST=localhost
  DB_USER=postgres
  DB_PASS=
  DB_NAME=
  DB_PORT=5432
  ```
2. Open the terminal and follow these steps:
  - If you haven't downloaded PostgreSQL, download it first.
  - Run:
    ```bash
    psql -U postgres
    ```
  - Create the database according to what you put in `.env` => if using uuid-ossp or auto generate (check file **/entity/user.go**):
    ```bash
    CREATE DATABASE your_database;
    \c your_database
    CREATE EXTENSION IF NOT EXISTS "uuid-ossp"; // remove default:uuid_generate_v4() if you not use you can uncomment code in user_entity.go
    \q
    ``` 
3. Run the application:
  ```bash
  go run main.go
  ```

## Run Migrations, Seeder, and Script
To run migrations, seed the database, and execute a script while keeping the application running, use the following command:

```bash
go run main.go --migrate --seed --run --script:example_script
```

- ``--migrate`` will apply all pending migrations.
- ``--seed`` will seed the database with initial data.
- ``--script:example_script`` will run the specified script (replace ``example_script`` with your script name).
- ``--run`` will ensure the application continues running after executing the commands above.

#### Migrate Database 
To migrate the database schema 
```bash
go run main.go --migrate
```
This command will apply all pending migrations to your PostgreSQL database specified in `.env`

#### Seeder Database 
To seed the database with initial data:
```bash
go run main.go --seed
```
This command will populate the database with initial data using the seeders defined in your application.

### API Documentation
You can explore the available API endpoints and their usage through the Postman Documentation:
- [Postman Documentation](https://documenter.getpostman.com/view/39901805/2sB2qUnPvX)
