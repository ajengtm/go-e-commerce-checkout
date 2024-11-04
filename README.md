### Step-by-Step Instructions

### 1. Set Up PostgreSQL

- **Install and Start PostgreSQL**: Make sure PostgreSQL is installed and running on your system. You can download it from the [PostgreSQL website](https://www.postgresql.org/download/).
- **Create a Database**: Open the PostgreSQL command line interface (`psql`) or use a database management tool like pgAdmin, and execute the following commands to create a database and tables:

    ```sql
    CREATE DATABASE ecommerce;
    
    \\c ecommerce  -- Connect to the 'ecommerce' database
    
    CREATE TABLE items (
        id SERIAL PRIMARY KEY,
        name TEXT NOT NULL,
        price NUMERIC(10, 2) NOT NULL,
        stock NUMERIC(10, 2) NOT NULL
    );
    
    CREATE TABLE carts (
        user_id TEXT NOT NULL,
        item_id INT NOT NULL,
        quantity INT NOT NULL,
        PRIMARY KEY (user_id, item_id),
        FOREIGN KEY (item_id) REFERENCES items (id)
    );
    
    ```

- **Insert Sample Data** (optional): Populate the `items` table with some initial data for testing:

    ```sql
    INSERT INTO items (name, price, stock) VALUES
    ('Apple', 0.50, 10),
    ('Banana', 0.30, 10),
    ('Orange', 0.70, 10);
    
    ```


### 2. Set Up the Go Project

- **Install Go**: Ensure you have Go installed on your machine. You can download it from [the Go website](https://golang.org/dl/).
- **Create Project Structure**: Organize your Go project with the directory structure outlined earlier.
- **Initialize Go Modules**: In the root of your project, initialize a new Go module:

    ```bash
    go mod init ecommerce-app
    
    ```

- **Install Dependencies**: Ensure the required packages are available:

    ```bash
    go get github.com/gorilla/mux
    go get github.com/jackc/pgx/v5
    
    ```


### 3. Configure the Database Connection

- **Update Connection String**: In `cmd/server/main.go`, replace the database connection string with your actual database credentials:

    ```go
    databaseURL := "postgres://username:password@localhost:5432/ecommerce"
    
    ```

  Replace `username`, `password`, and `localhost` with your PostgreSQL credentials and host information.


### 4. Run the Application

- **Run the Server**: Start the Go application from the root directory of your project:

    ```bash
    go run cmd/server/main.go
    
    ```

  This command compiles and runs the application, starting the HTTP server on port 8080. You should see output indicating the server is running, such as "Server is running on port 8080".


### 5. Test the API

- **Use cURL or Postman**: Send HTTP requests to verify that the API is functioning as expected. Here are some example cURL commands:
    - **Add Item to Cart**:

        ```bash
        curl -X POST -H "Content-Type: application/json" -d '{"item_id": 1, "quantity": 2}' http://localhost:8080/cart/user123/add
        
        ```

    - **Checkout**:

        ```bash
        curl -X GET http://localhost:8080/cart/user123/checkout
        
        ```

    - **Process Payment**:

        ```bash
        curl -X POST -H "Content-Type: application/json" -d '{"amount": 1.00}' http://localhost:8080/cart/user123/pay
        ```


If everything is set up correctly, these commands will interact with your Go server, performing operations like adding items to the cart, checking out, and processing payments.