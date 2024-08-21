# Go Simple Bank App

This project is a practice project to learn Go Lang and other backend tools. It has the ability to connect to AWS EKS and ECR and then uses Kubernetes to deploy the app.

## Prerequisites

1. Go 1.16
2. Docker
3. migrate CLI (Install using `brew install golang-migrate/migrate`)

## Getting Started

To get the Go Simple Bank App running locally:

1. Clone this repository
2. Create a PostgreSQL instance using Docker

    ```bash
    make postgres
    ```

3. Create database in PostgreSQL

    ```bash
    make createdb
    ```

4. Run the Go Simple Bank App server

    ```bash
    make server
    ```

## Using Docker

```bash
docker compose up
```


## API Endpoints

### User Endpoints

- **Create User**: `POST /users`
- **Login User**: `POST /users/login`
- **Renew Access Token**: `POST /tokens/renew_access`

### Account Endpoints

- **Create Account**: `POST /accounts`
- **Get Account**: `GET /accounts/:id`
- **List Accounts**: `GET /accounts`
- **Update Account**: `PUT /accounts/:id`
- **Delete Account**: `DELETE /accounts/:id`

### Transfer Endpoints

- **Create Transfer**: `POST /transfers`

## Configuration

The application configuration is managed via the `util.Config` struct. Ensure you have the necessary environment variables set up for the configuration.

## Running Tests

To run the tests for this project, use the following command:

```bash
make test
