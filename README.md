# Shopping Cart API - Backend

This is the backend for a basic e-commerce shopping cart application. It provides a web service to manage users, items, shopping carts, and orders.

The project is built to follow a simple user flow:
1.  [cite_start]A new user signs up by creating a User account[cite: 9].
2.  [cite_start]The user logs in to get a token for their session[cite: 11].
3.  [cite_start]While shopping, the user adds Items to their Cart[cite: 13].
4.  [cite_start]The cart is converted into an order when the user is ready to checkout[cite: 16].
5.  [cite_start]The backend provides listing endpoints for Users, Items, Carts, and Orders[cite: 17].

## Technologies Used

* **Go:** The programming language for the backend.
* **Gin:** A web framework for building the API endpoints.
* **GORM:** An ORM for interacting with the database.
* **PostgreSQL:** The relational database to store application data.

## Getting Started

### Prerequisites

* Go (version 1.18 or higher)
* PostgreSQL database server

### Setup

1.  **Clone the repository:**
    ```sh
    git clone https://github.com/amit-shuklaa/ABCDE-Assignment
    cd shopping-cart-backend
    ```

2.  **Configure the database:**
    * Ensure your PostgreSQL server is running.
    * Create a new database for the project (e.g., `shopping_cart_db`).
    * Update the DSN (Data Source Name) in `database/database.go` with your database credentials.

3.  **Install dependencies and run migrations:**
    ```sh
    go mod tidy
    go run main.go
    ```
    This will install all necessary Go modules and automatically create the required tables in your database.

## API Endpoints

The following endpoints are available:

| Method   | Endpoint          | Description                                                                          |
|----------|-------------------|--------------------------------------------------------------------------------------|
| `POST`   | `/users`          | [cite_start]Creates a new User[cite: 20].                                                       |
| `GET`    | `/users`          | [cite_start]Lists all users[cite: 20].                                                          |
| `POST`   | `/users/login`    | [cite_start]Logs in an existing user based on username and password, returning a token[cite: 20]. |
| `POST`   | `/items`          | [cite_start]Creates a new Item[cite: 20].                                                       |
| `GET`    | `/items`          | [cite_start]Lists all items[cite: 20].                                                          |
| `POST`   | `/carts`          | [cite_start]Creates a cart and adds items to it[cite: 20].                                      |
| `GET`    | `/carts`          | [cite_start]Lists all carts[cite: 20].                                                          |
| `POST`   | `/orders`         | [cite_start]Converts a cart to an order[cite: 20].                                              |
| `GET`    | `/orders`         | [cite_start]Lists all orders[cite: 20].                                                         |

[cite_start]Note: The user's token must be present in the `Authorization` header for cart and order-related endpoints[cite: 21].

## Postman Collection

A Postman collection is provided in this repository. You can import the `shopping-cart-api.json` file into Postman to easily test all the API endpoints. It includes environment variables for `baseUrl` and `authToken` for a streamlined testing experience.
# Shopping Cart App - Frontend

This is the frontend for the shopping cart application, built as a single-page web app to interact with the Go backend API.

The application follows the user flow described in the project:
1.  [cite_start]A user can log in with their username and password[cite: 9]. [cite_start]If login fails, an alert is shown[cite: 9].
2.  [cite_start]On successful login, the user is taken to the List Items screen[cite: 9, 11].
3.  From the List Items screen, the user can:
    * [cite_start]Add items to their cart by clicking on them[cite: 13].
    * View all items in their cart via a "Cart" button.
    * View their placed orders via an "Order History" button.
    * Checkout and place an order via a "Checkout" button.

## Technologies Used

* **React:** The JavaScript library for building the user interface.
* **React Router:** For client-side routing between the login and item list pages.
* **Axios:** A promise-based HTTP client for making API calls to the backend.

## Getting Started

### Prerequisites

* Node.js and npm
* The `shopping-cart-backend` server must be running and accessible at `http://localhost:8080`.

### Setup

1.  **Clone the repository:**
    ```sh
    git clone https://github.com/amit-shuklaa/ABCDE-Assignment
    cd shopping-cart-frontend
    ```

2.  **Install dependencies:**
    ```sh
    npm install
    ```

3.  **Run the development server:**
    ```sh
    npm start
    ```
    The application will open in your browser at `http://localhost:3000`.

## User Interaction

* **Login:** Use the login form to authenticate with an existing user account.
* **Items:** The page will display all items fetched from the backend. Clicking an item adds it to your cart.
* **Cart Button:** Shows a window alert with a list of items currently in your cart (cart_id, item_id).
* **Order History Button:** Shows a window alert with a list of your placed order IDs.
* **Checkout Button:** Converts your cart into an order and shows a confirmation message.
