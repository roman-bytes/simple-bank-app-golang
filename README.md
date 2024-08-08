# Go Simple Bank App

## Prerequisites

1. Go 1.16
2. Docker

## Getting Started

To get the Go Simple Bank App running locally:

1. Clone this repository
2. Create a PostgreSQL instance using Docker

```base
    make postgres
```
3. Create database in PostgreSQL

```base
    make createdb
```

4. Generate SQL querys

```base
    make sqlc
```
5. Run the Go Simple Bank App server

```base
    make server
```

## API ENDPOINTS

Testing deploy