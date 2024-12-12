----------------------------------------------------------------------------------------------------------------

Maker-Checker Approval Process API
This is a simplified implementation of the Maker-Checker approval process for sending a message. 
The process requires validation by users before the message is sent to the recipient. 
The message is either approved or rejected by users in the system.

Features:
Maker initiates a message request.
Checker approves or rejects the message request.
If approved, the message is sent to the recipient.
If rejected, the message is discarded.

----------------------------------------------------------------------------------------------------------------

Technology Stack
Gin Framework (Web framework for Go)
PostgreSQL (Database)
Golang (Programming language)
Docker Compose (For PostgreSql )

----------------------------------------------------------------------------------------------------------------
Prerequisites
Docker and Docker Compose should be installed on your machine. You can download Docker from here.
Go version 1.22+.

1. Install dependencies
First, install the necessary Go dependencies.
   Run command:
   ``` 
     go mod tidy 
   ```

2. Docker Compose Setup
The docker-compose.yml file is provided to set up the PostgreSQL database. It will also create the necessary database and user.
  Run command:
   ``` 
    docker compose -f docker-compose-deps.yml up -d
   ```
  Database connection details are stored in the conf.json file that will be read at runtime by the Go application.
  After running Docker Compose,  migrations will automatically apply to create the necessary tables for your application.
  If we need to add new tables we can add in db.go

3. Run the application
     Run command:
    ```
      go run main.go
   ```

go run main.go
This will start the server on http://localhost:8002.
we can also change the port in conf.json file if we need to run on different port

----------------------------------------------------------------------------------------------------------------
Project Structure
.
1. ├── conf/                      # Configuration files
2. │   └── gbe_config.go          # DB and Email configurations
3. ├── rest/                      # Controller layer for API routes
4. │   └── user_controller.go     # main controller where all Endpoints are written
5. │   └── bootstrap.go           # we use bootstrap pattern
6. ├── repository/                # Database interaction layer
7. │   └── store.go               # db functions method declares
8. │   └── postgres               # main functions method defines and implements
9. │   └──└── db.go               # migrations and database setup
10. │   └──└── users.go            # main functions method defines and implements
11. ├── services/                  # Business logic layer
12. │   └── user_service.go        #main business logic functions
13. ├── main.go                    # Entry point for the application
14. ├── Dockerfile                 # Dockerfile to build the image
15. ├── docker-compose.yml         # Docker Compose file for DB
16. ├── go.mod                     # Go module dependencies
17. └── README.md                  # Project documentation

----------------------------------------------------------------------------------------------------------------
Explanation of Project Structure:
1. conf/: This directory contains configuration-related files such as gbe_config.go.go, which reads environment variables for DB and Email settings.
2. controllers/: Contains the HTTP request handlers (controllers) that define the API endpoints for managing message requests.
3. repository/: The repository layer interacts with the database to perform CRUD operations.
4. services/: Contains the business logic for handling message requests, including approval and rejection logic.

----------------------------------------------------------------------------------------------------------------

API Endpoints
The API exposes the following endpoints:
We can also check the postman collection in the directory with the name of Maker-Checker.postman_collection.json

POST /api/v1/message-request

Description: Initiates a new message request.
Request Body:

    {
      "user_id": "user1",
      "recipient": "user2",
      "message": "Hello World"
    }

Response:

    {
      "message_id": 7,
      "sender": "Frank",
      "recipient": "mbilalsorathia@gmail.com",
      "message": "hey its frank here",
      "status": "Pending",
      "created_at": "2024-12-10T23:33:00.918952+04:00",
      "updated_at": "2024-12-10T23:33:00.919367+04:00"
    }

PATCH /api/v1/message-request

Description: Approves a message request (can only be done by Checkers).
Request Body:

    {
      "request_id":7,
      "user_id":"user2",
      "status":"Approve"
    }

Response:

    {
    "message_id": 7,
    "sender": "Frank",
    "recipient": "mbilalsorathia@gmail.com",
    "message": "hey its frank here",
    "status": "Approve", // or Reject
    "created_at": "2024-12-10T23:33:00.918952+04:00",
    "updated_at": "2024-12-10T23:33:00.919367+04:00"
    }

GET /api/v1/users/message-request

Description: Retrieves the current status of the message request.
Response:

    {
      "request_id": 1,
      "status": "Pending",
      "sender": "user1",
      "recipient": "user2",
      "message": "Hello World",
      "created_at": "2024-12-10T12:34:56Z"
    }

----------------------------------------------------------------------------------------------------------------

