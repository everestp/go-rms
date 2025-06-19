# Restaurant Management System (Go-RMS)

This is a Restaurant Management System built with Go (Golang), Gin web framework, and MongoDB. The application provides RESTful APIs for managing users, foods, menus, tables, orders, order items, and invoices.

## Features

- **User Management:** Signup, login, and user listing.
- **Food Management:** CRUD operations for food items.
- **Menu Management:** Manage restaurant menus.
- **Table Management:** Manage restaurant tables.
- **Order Management:** Create and update orders.
- **Order Items:** Manage items within orders.
- **Invoice Management:** Generate and update invoices.
- **Authentication:** Middleware-protected routes for secure access.

## Project Structure

```
.
├── main.go
├── go.mod
├── go.sum
├── controllers/
├── database/
├── helpers/
├── middleware/
├── models/
└── routes/
```

- **controllers/**: Business logic for each resource.
- **database/**: MongoDB connection and collection helpers.
- **helpers/**: Utility functions (e.g., token handling).
- **middleware/**: Authentication middleware.
- **models/**: Data models for MongoDB collections.
- **routes/**: API route definitions.

## Getting Started

### Prerequisites

- Go 1.18+
- MongoDB running locally on `mongodb://localhost:27017`

### Installation

1. Clone the repository:
    ```sh
    git clone https://github.com/yourusername/go-rms.git
    cd go-rms
    ```

2. Install dependencies:
    ```sh
    go mod tidy
    ```

3. Start MongoDB if not already running.

4. Run the application:
    ```sh
    go run main.go
    ```

5. The server will start on `http://localhost:8000` by default.

## API Endpoints

### User

- `POST /user/signup` - Register a new user
- `POST /user/login` - User login
- `GET /users` - List users
- `GET /users/:user_id` - Get user by ID

### Food

- `GET /foods` - List foods
- `GET /foods/:food_id` - Get food by ID
- `POST /foods` - Add new food
- `PATCH /foods/:food_id` - Update food

### Menu

- `GET /menus`
- `GET /menus/:menu_id`
- `POST /menus`
- `PATCH /menus/:menu_id`

### Table

- `GET /tables`
- `GET /tables/:table_id`
- `POST /tables`
- `PATCH /tables/:table_id`

### Order

- `GET /orders`
- `GET /orders/:order_id`
- `POST /orders`
- `PATCH /orders/:order_id`

### Order Items

- `GET /orderItems`
- `GET /orderItems/:orderItem_id`
- `GET /orderItems-order/:order_id`
- `POST /orderItems`
- `PATCH /orderItems/:orderItem_id`

### Invoice

- `GET /invoices`
- `GET /invoices/:invoice_id`
- `POST /invoices`
- `PATCH /invoices/:invoice_id`

## License

This project is licensed under the MIT License.

---

**Note:** This project is under active development. Contributions