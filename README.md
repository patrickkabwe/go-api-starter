## Introduction

This project is a Go API starter project. It provides a basic structure for building RESTful APIs in Go using the net/http package and gorm for database operations. The project is organized in a way that separates concerns and makes it easy to add new features.  

### Author

[Patrick Kabwe](patrickckabwe@gmail.com)

### License

This project is licensed under the MIT License.  

### Folder Structure

* `cmd/api`: This is the main entry point of the application. It contains the `api.go` file which sets up and starts the HTTP server.

* `internal`: This directory contains the business logic of the application. It is further divided into auth and user subdirectories.  

  * `auth`: Contains the authentication logic of the application. It includes the `routes.go` file which sets up the routes for user registration and login, and the `routes_test.go` file which contains tests for the authentication routes.  
  * `user`: Contains the user management logic of the application. It includes the `routes.go` file which sets up the route for fetching user details.  
  
* `utils`: This directory contains utility functions that are used across the application.  
* `types`: This directory contains the data types and structs used in the application.  

### Contribution

Contributions are welcome. Please make sure to read the contributing guide before making a pull request.  

### Getting Started

To get started with this project, clone the repository and run `go run cmd/api/api.go`.  

### Testing

To run the tests, use the `go test` command followed by the directory containing the tests. For example, `go test internal/auth`.  

Please note that this is a basic structure and you may need to add or modify it according to your project requirements.