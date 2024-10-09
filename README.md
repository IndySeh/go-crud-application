## Simple Go Application

This is a basic Go application structured to demonstrate clean architecture principles, with the aim of separating concerns into different layers such as handlers, database interactions, middleware, logging, and utilities. The project also uses MySQL for database operations and logs application activities to different files for easier debugging.

### Directory Structure

```bash
.
├── cmd
│   └── api
│       └── main.go                   # Entry point of the application
├── go.mod                             # Go module file
├── go.sum                             # Go dependencies' checksum
├── internals
│   ├── db
│   │   └── db.go                      # Database connection setup
│   ├── handlers
│   │   └── user_handlers.go           # HTTP handlers for user-related actions
│   ├── middleware
│   │   └── logging.go                 # Middleware for logging API requests
│   ├── repository
│   │   └── user_repository.go         # Repository layer for interacting with the database
│   └── utils
│       └── error.go                   # Utility functions (e.g., error handling)
├── logs
│   ├── api-requests.log               # Log for API requests
│   ├── db.log                         # Log for database-related events
│   ├── error.log                      # Log for errors
│   └── info.log                       # General information log
├── pkg
│   ├── logging
│   │   └── logging.go                 # Logging package for handling application-wide logging
│   └── types
│       └── types.go                   # Shared data structures (e.g., types)
└── README.md                          # Project documentation (this file)
```

### Features

- **Handlers**: Manage incoming HTTP requests for users.
- **Middleware**: Adds request logging for every API call.
- **Repository Layer**: Isolates the database interaction logic.
- **Utility Functions**: Standard error handling and response structure.
- **Logging**: Application logs are saved in `logs/` for easier debugging.
  
### Setup Instructions

1. **Clone the repository**:
   ```bash

   git clone https://github.com/IndySeh/go-crud-application.git
   cd go-crud-application
   ```

2. **Install dependencies**:
   Ensure Go is installed on your machine. Install required dependencies by running:
   ```bash
   go mod tidy
   ```

3. **Set .env**:
   Create a .env file:  `cmd/api/.env`:
   
   ```bash
   
  DB_DRIVER=mysql
  DB_USER=root
  DB_PASSWORD=root
  DB_NAME=crud_app
    ```

4. **Run the application**:
```bash

   Start the server using:
   ```bash
   go run cmd/api/main.go
   ```

   The API should be running at `http://localhost:8080`.

5. **API Endpoints**:
   - `GET /api/users`: Fetch all users.
   - `GET /api/users{id}`: Fetch users by id.
   - `POST /api/users`: Add a new user.
   - `DEL /api/users{id}`: Delete a user.
   - `PUT /api/users`: Update a user

6. **Logs**:
   Application logs (API requests, database queries, errors) are saved under the `logs/` directory.


### Project Structure Explanation

- **`cmd/api/main.go`**: This is the entry point of the Go application where the server is started and routes are defined.
- **`internals/`**: Contains the core functionality of the application (handlers, repository, database setup).
  - **`handlers/`**: Defines HTTP handlers that process incoming requests.
  - **`repository/`**: Encapsulates database interactions, such as fetching or inserting data.
  - **`db/`**: Handles database connections.
  - **`middleware/`**: Defines middleware functions like logging to intercept and process requests.
  - **`utils/`**: Contains utility functions such as error handling.
- **`pkg/`**: Contains reusable packages (e.g., logging).
- **`logs/`**: Directory where log files are stored.

### Contributing

Feel free to contribute to this project by submitting a pull request, reporting issues, or improving documentation.
