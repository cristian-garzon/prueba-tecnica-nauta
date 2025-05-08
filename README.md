# Marine Logistics API

technical test nauta


## Features

- write in database, but read in cahce, with this we have an inprovement when we read the data, because we don't need to go to the database. The Historical data that we save in cache, we can adjust if is necesary. In the first time it feature doesn't seem necesary, but, if we need to save data such as invoices, orders, etc..., we need a data retention too long, so, if we have for example, 1 year from historical data, the queries could have a delay, so, in those cases, is util have a reader in cache. Just if we don't get the data, we can go to the database.

- We have a circuitBreaker, in down database case, it help us if the database is saturated, the historical will be unavailable by 1 minute (by default), the max retries and delay time are ajust.

## Prerequisites

- Go 1.21 or higher
- PostgreSQL 12 or higher
- Docker

## run project

To run the database and migrations:

```bash
# Start the database and run migrations
docker-compose up -d postgres migrations

# To stop the services
docker-compose down
```

The database will be available at:
- Host: localhost
- Port: 5432
- User: homestead
- Password: secret
- Database: marine-logistics

## Project Structure

```
.
├── app/
│   ├── domain/
│   │   ├── model/
│   │   │   ├── Dto/           # Data  Transfer Objects for API responses
│   │   │   ├── entities/      # Database entities and their mappings
│   │   │   ├── queries/       # SQL queries for database operations
│   │   │   └── types/         # Custom types functions that we use in all service
│   ├── infrastructure/
│   │   ├── clients/           # Database client and connection management
│   │   ├── config/            # Application configuration
│   │   ├── repositories/      # Data access layer (SQL and Cache implementations)
│   │   ├── circuitbreaker/    # Circuit breaker pattern implementation
│   │   └── server/            # HTTP server and routing
├── migrations/                 # Database migration files
└── README.md
```

### Folder Structure Explanation

#### Domain Layer (`app/domain/`)
- **Dto/**: Contains Data Transfer Objects that define the structure of API responses
  - `BookingDto`: Booking response structure
  - `ContainerDto`: Container response structure
  - `OrderDto`: Order response structure
  - `InvoiceDto`: Invoice response structure

- **entities/**: Database entity models and their mappings
  - `Booking`: Booking database entity
  - `Container`: Container database entity
  - `Order`: Order database entity
  - `Invoice`: Invoice database entity

- **queries/**: SQL queries for database operations
  - Contains all SQL queries used by the repositories
  - Organized by entity type (bookings, containers, orders, invoices)

- **types/**: Custom types functions that we use in all service

#### Infrastructure Layer (`app/infrastructure/`)
- **clients/**: Database connection management
  - `Pool`: PostgreSQL connection pool
  - Connection configuration and management

- **config/**: Application configuration
  - Database configuration
  - Server configuration
  - Environment variables

- **repositories/**: Data access implementations
  - `NautaSqlRepository`: PostgreSQL implementation
  - `NautaCacheRepository`: In-memory cache implementation
  - Repository interfaces implementation

- **circuitbreaker/**: Circuit breaker pattern
  - Handles database failures
  - Prevents cascading failures
  - Configurable retry and timeout settings

- **server/**: HTTP server implementation
  - Route definitions
  - Middleware
  - Request handlers

#### Migrations
- Contains database migration files
- `V1__create_bookings.sql`: Initial schema creation
  - Creates tables for clients, bookings, containers, orders, and invoices
  - Sets up foreign key relationships
  - Defines indexes and constraints
  - we use flyway, so, you need to have the next format if you want to create migrations "V#__COMMENT.sql"

## Running Tests

To run the tests:

```bash
go test ./...
```

The tests use a real PostgreSQL database, so make sure your database is properly configured before running tests.

## API Endpoints

### Bookings

- `GET /bookings` - Get all bookings
- `POST /bookings` - Create a new booking
  ```json
  {
    "booking_id": "BK7",
    "client_id": 1,
    "containers": [
      {
        "container_id": "CON23",
        "container_type": "CH4KPO",
        "description": "motorbikes",
        "weight": 135
      }
    ],
    "orders": [
      {
        "purchase_id": "PUR19",
        "total_amount": 540,
        "description": "first payment",
        "invoices": [
          {
            "invoice_id": "INv25",
            "amount": 520,
            "payment_date": "2025-05-05T19:40:00Z"
          },
          {
            "invoice_id": "INV26",
            "amount": 20,
            "payment_date": "2025-05-05T19:40:00Z"
          }
        ]
      }
    ]
  }
  ```
- `GET /bookings/{id}` - Get booking by ID

### Containers

- `GET /containers` - Get all containers
- `GET /containers/order/{orderId}` - Get containers by order ID
- `GET /containers/booking/{bookingId}` - Get containers by booking ID
- `GET /containers/email/{email}` - Get containers by client email

### Orders

- `GET /orders` - Get all orders
- `GET /orders/container/{containerId}` - Get orders by container ID
- `GET /orders/booking/{bookingId}` - Get orders by booking ID
- `GET /orders/email/{email}` - Get orders by client email

### Invoices

- `GET /invoices` - Get all invoices
- `GET /invoices/order/{orderId}` - Get invoices by order ID

## Error Handling

The application uses custom error types and the circuit breaker pattern to handle failures gracefully:

- `ErrNoBookingsFound` - When no bookings are found
- `ErrNoContainersFound` - When no containers are found
- `ErrNoOrdersFound` - When no orders are found
- `ErrNotClientFound` - When no client is found
- `ErrQueryError` - When there's a database query error
