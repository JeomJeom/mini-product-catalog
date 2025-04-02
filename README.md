# Product Catalog App
This is a full-stack product catalog application built with a Go (Gin, GORM) backend and an Angular frontend. It uses PostgreSQL as the database and Docker Compose to orchestrate the services.

## Prerequisites
Docker Desktop (or Docker Engine)

Docker Compose V2 (integrated with Docker CLI) or Docker Compose V1 installed separately

## Project Structure
```
├── thalesapi/                   # Go backend application
│   ├── cmd/main.go             # Application entry point
│   ├── configs/             # Configuration loading logic (if any)
│   ├── data/                # (Optional) Data folder, e.g. for file uploads
│   ├── datasets/            # SQL scripts for dummy data (e.g., products.sql)
│   ├── db/                  # Database connection and migration logic
│   ├── internal/            # Internal packages (business logic, domain)
│   ├── middleware/          # Custom Gin middleware
│   ├── router/              # Gin router setup
│   ├── wire/                # Google Wire DI setup
│   ├── .env                 # Environment variables
│   ├── .gitignore           # Git ignore file
│   ├── Dockerfile           # Dockerfile for building the Go app image
│   ├── go.mod               # Go module definition
│   ├── go.sum               # Go module dependencies
├── thalesui/                   # Angular frontend application
│   ├── src/
│   └── Dockerfile              # Dockerfile for the frontend app
└── docker-compose.yml          # Docker Compose file to run all services
```
## Configuration
The application uses environment variables to configure the database connection. In the docker-compose.yml file, the relevant environment variables for the backend service are:

### Running the Application Locally
Clone the repository:
```
git clone https://github.com/JeomJeom/mini-product-catalog.git
cd mini-product-catalog
```

Build and run all services using Docker Compose:

### If you're using Docker Compose V2, run:

```
docker compose up --build
```
### If you're using Docker Compose V1, run:

```
docker-compose up --build
```
This command will build the Docker images for the backend and frontend, start the PostgreSQL database, and wire everything together.

Access the Application:

### Backend API:
The backend runs on http://localhost:8990

### Frontend App:
The Angular frontend runs on http://localhost:4200 (if configured in your docker-compose file)

Note: Adjust the ports in docker-compose.yml if you want to use different ports.

## Troubleshooting
### Database Connection Issues:
If you see errors related to connecting to the database, verify that the environment variables (especially DB_PORT and DB_HOST) are correctly set and match the values in docker-compose.yml.

### Container Logs:
Use the following command to view logs for any service:

```
docker compose logs <service_name>
```
Replace <service_name> with db, app, or your frontend service name.

### Stopping the Application:

To stop and remove the containers, press Ctrl+C in the terminal running Docker Compose, and then execute:

```
docker compose down -v
```
### Additional Notes
Docker Compose Version Warning:
If you see a warning about the version attribute being obsolete in docker-compose.yml, simply remove the version key to avoid confusion.

### Modifications:
Feel free to modify the Dockerfiles, environment variables, or Docker Compose configuration as needed for your development workflow.

