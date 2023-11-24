# Assignment 3 - Weather Monitoring API

This repository contains the implementation of Assignment 3 for the "Scalable With Golang" course at Hacktiv8. The Weather Monitoring API is designed to track and update weather conditions at regular intervals.

## Features
- Continuously monitor and update weather conditions.
- Record water and wind data with their respective statuses.
- Utilizes a PostgreSQL database to store weather information.

## Technologies Used
- Go (Golang): The primary programming language for developing the API.
- Gin: The web framework used for API development.
- Gorm: The Go ORM used for interacting with the PostgreSQL database.
- PostgreSQL: The relational database used to persist weather information.

## How to Use
To get started with the Weather Monitoring API, follow these steps:

1. Clone this repository to your local machine.
2. Ensure you have Go installed on your system.
3. Set up a PostgreSQL database and update the database connection details in the `ConnectDatabase` function in the `models` package.
4. Run the command `go run main.go` to start the server.
5. The API will continuously update weather conditions at regular intervals.

## Endpoints
- **Get Weather Data**: `GET /weather`

## Database Setup
Make sure to set up a PostgreSQL database with the following configuration:

```sql
CREATE DATABASE db_assignment3;
```

Update the database connection details in the `ConnectDatabase` function in the `models` package.

## Notes
This project was created as part of completing the "Scalable With Golang" course at Hacktiv8. Feel free to modify this README.md to better suit your specific project details.