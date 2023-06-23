# go-crud
Go CRUD App with GORM and PostgreSQL
This repository contains a CRUD (Create, Read, Update, Delete) application built using the Go programming language, GORM as the ORM (Object-Relational Mapping) library, and PostgreSQL as the database. The application manages users and posts, allowing you to perform various operations on these entities.

Features
User management: Create, retrieve, update, and delete users.
Post management: Create, retrieve, update, and delete posts.
Database integration: Utilizes PostgreSQL as the database backend.
ORM support: Uses GORM library for database operations, providing an easy-to-use interface for interacting with the database.
Prerequisites
Before running the application, ensure the following requirements are met:

Go 1.16 or above is installed.
PostgreSQL database is set up and running.
GORM and the required dependencies are installed. You can install them by running the following command:
go
Copy code
go get -u github.com/go-gorm/gorm
Installation
Clone the repository:

bash
Copy code
git clone https://github.com/your-username/your-repository.git
Change into the project directory:

bash
Copy code
cd your-repository
Install the dependencies:

go
Copy code
go mod download
Set up the database:

Create a new PostgreSQL database.
Update the database configuration in the config.go file with your database connection details.
Build the application:

go
Copy code
go build
Run the application:

bash
Copy code
./your-repository
The application will start running on http://localhost:8080.

Usage
Once the application is up and running, you can use the following endpoints to interact with the API:

GET /users - Retrieve all users.

GET /users/{id} - Retrieve a specific user by ID.

POST /users - Create a new user. Provide user details in the request body.

PUT /users/{id} - Update an existing user by ID. Provide user details in the request body.

DELETE /users/{id} - Delete a user by ID.

GET /posts - Retrieve all posts.

GET /posts/{id} - Retrieve a specific post by ID.

POST /posts - Create a new post. Provide post details in the request body.

PUT /posts/{id} - Update an existing post by ID. Provide post details in the request body.

DELETE /posts/{id} - Delete a post by ID.

Note: Replace {id} with the actual ID of the user or post you want to retrieve, update, or delete.
