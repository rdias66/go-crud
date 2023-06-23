# Go CRUD App with GORM and PostgreSQL

This repository contains a CRUD (Create, Read, Update, Delete) application built using the Go programming language, GORM as the ORM (Object-Relational Mapping) library, and PostgreSQL as the database. The application manages users and posts, allowing you to perform various operations on these entities.

## Features

- User management: Create, retrieve, update, and delete users.
- Post management: Create, retrieve, update, and delete posts.
- Database integration: Utilizes PostgreSQL as the database backend.
- ORM support: Uses GORM library for database operations, providing an easy-to-use interface for interacting with the database.

## Prerequisites

Before running the application, ensure the following requirements are met:

- Go 1.16 or above is installed.
- PostgreSQL database is set up and running.
- GORM and the required dependencies are installed. You can install them by running the following command:

go get -u github.com/go-gorm/gorm


## Installation

1. Clone the repository:

git clone https://github.com/rdias66/go-crud


2. Change into the project directory:

cd go-crud


3. Install the dependencies:

go mod download


4. Set up the database:
   - Create a new PostgreSQL database.
   - Update the database configuration in the `config.go` file with your database connection details.

5. Build the application:

go build


6. Run the application:

./go-crud


   The application will start running on `http://localhost:8080`.

## Usage

Once the application is up and running, you can use the following endpoints to interact with the API:

- **GET** `/users` - Retrieve all users.
- **GET** `/users/{id}` - Retrieve a specific user by ID.
- **POST** `/users` - Create a new user. Provide user details in the request body.
- **PUT** `/users/{id}` - Update an existing user by ID. Provide user details in the request body.
- **DELETE** `/users/{id}` - Delete a user by ID.

- **GET** `/posts` - Retrieve all posts.
- **GET** `/posts/{id}` - Retrieve a specific post by ID.
- **POST** `/posts` - Create a new post. Provide post details in the request body.
- **PUT** `/posts/{id}` - Update an existing post by ID. Provide post details in the request body.
- **DELETE** `/posts/{id}` - Delete a post by ID.

Note: Replace `{id}` with the actual ID of the user or post you want to retrieve, update, or delete.

## Contributing

Contributions are welcome! If you find any issues or have suggestions for improvement, please feel free to open an issue or submit a pull request.

When contributing to this repository, please ensure that your changes are well-documented and follow the established coding conventions.

## License

This project is licensed under the [MIT License](LICENSE).

## Acknowledgments

- The creators of GORM, PostgreSQL, and the Go programming language for their excellent tools and libraries.
- The open-source community for their contributions and support.

