# My Transactions API

This is a sample API for managing monetary transactions made by users. It provides endpoints for creating transactions and listing transactions with pagination and filtering.

## Table of Contents

- [Requirements](#requirements)
- [Project Structure](#project-structure)
- [Usage](#usage)
  - [Using `go run`](#using-go-run)
  - [Using `start_project.sh`](#using-start_projectsh)
- [API Endpoints](#api-endpoints)
- [Swagger Documentation](#swagger-documentation)
- [Contributing](#contributing)
- [License](#license)

## Requirements

Before you begin, ensure you have met the following requirements:

- [Go](https://golang.org/) installed on your system.
- [Docker](https://www.docker.com/) installed (optional, for running the application in a Docker container).

## Project Structure

- `cmd/`: Contains the application's main entry point (`main.go`).
- `internal/`: Contains the application logic, models, repositories, and handlers.
- `scripts/`: Contains the `setup_db.sh` script for setting up the MongoDB database.
- `docs/`: Contains Swagger documentation files generated using `swag`.
- `Dockerfile`: Defines the Docker image for the application.
- `start_project.sh`: Shell script to set up the database, build, and run the application.

## Usage

### Using `go run`

To run the application using `go run`, follow these steps:

1. Navigate to the root directory of the project.

2. Run the following command to start the application:

   ```bash
   go run cmd/main.go


### Using start_project.sh
To start the project using the provided start_project.sh script, follow these steps:

1. Navigate to the root directory of the project.

2. Make sure the setup_db.sh script is in the scripts/ directory and is executable:

```Bash
chmod +x scripts/setup_db.sh
```
3. Run the following command to start the entire project, including database setup, Docker image build, and application run:

```Bash
./start_project.sh
```

4. The application will be accessible at http://localhost:8080.

## API Endpoints
The API provides the following endpoints:

- `POST /transactions/CreateTransaction`: Create a new transaction.
- `GET /transactions/ListTransactions`: List transactions with optional pagination and filtering parameters.
Detailed API documentation with request and response examples can be found in the Swagger documentation.


### Request Parameters 
- `POST  /transactions/CreateTransaction `

#### Request
To create a new transaction, send a POST request to the /transactions endpoint with the following JSON request body:
```json
{
    "origin": "mobile-ios",
    "user_id": 123,
    "amount": 50.0,
    "operation": "credit",
    "created_at": "2023-01-15T14:30:00Z"
}

```

- `origin (string, required)`: The origin where the transaction was initiated (e.g., "mobile-ios," "desktop-web").

- `user_id (integer, required)`: The ID of the user who made the transaction.

- `amount (float, required)`: The monetary amount of the transaction.

- `operation (string, required)`: The type of operation, which can be either "credit" or "debit."

- `created_at (string, required)`: The date and time when the transaction was created in ISO 8601 format (e.g., "2023-01-15T14:30:00Z").

#### Sample Response

```json
{
    "ID": 1,
    "Origin": "mobile-ios",
    "UserID": 123,
    "Amount": 50.0,
    "Operation": "credit",
    "CreatedAt": "2023-01-15T14:30:00Z"
}

```

- `GET /transactions/ListTransactions `

#### Filtering Parameters

`origin` (optional):

Description: Filter transactions by their origin.
Example: /transactions?origin=mobile-ios

`userId` (optional):

Description: Filter transactions by the user's ID.
Example: /transactions?userId=123

`operation` (optional):

Description: Filter transactions by the type of operation (credit or debit).
Example: /transactions?operation=credit

`startDate` and `endDate` (optional):

Description: Filter transactions by a date range. Provide both startDate and endDate parameters.
Example: /transactions?startDate=2023-01-01&endDate=2023-12-31

#### Pagination Parameters

`page` (optional):

Description: Specify the page number for paginated results. Default is 1.
Example: /transactions?page=2

`pageSize` (optional):

Description: Specify the number of transactions per page. Default is 10.
Example: /transactions?pageSize=20

#### Response

The ListTransactions endpoint returns a JSON response containing the list of transactions that match the specified filters and pagination criteria.

```json
{
    "data": [
        {
            "ID": 1,
            "Origin": "mobile-ios",
            "UserID": 123,
            "Amount": 50.0,
            "Operation": "credit",
            "CreatedAt": "2023-01-15T14:30:00Z"
        },
        {
            "ID": 2,
            "Origin": "desktop-web",
            "UserID": 456,
            "Amount": 25.0,
            "Operation": "debit",
            "CreatedAt": "2023-02-10T09:45:00Z"
        },
        ...
    ],
    "totalRecords": 100,
    "page": 1,
    "pageSize": 10
}

```

#### Usage Examples
Here are some example requests to demonstrate how to use the ListTransactions API endpoint:

List all transactions:
- Request: `/transactions`

List transactions for a specific origin:
- Request: `/transactions?origin=mobile-ios`

List transactions for a specific user:
- Request: `/transactions?userId=123`

List credit transactions within a date range:
- Request: `/transactions?operation=credit&startDate=2023-01-01&endDate=2023-12-31`

Paginate through the results:
- Request: `/transactions?page=2&pageSize=20`

## Swagger Documentation
Swagger documentation is available for exploring and testing the API. You can access it at:

http://localhost:8080/swagger/index.html (when the application is running)