# Go Books API

This is a simple API for managing books and categories built with the Gin framework in Go. The API allows users to:

- Authenticate using JWT tokens.
- Manage categories.
- Manage books.

## Table of Contents
- [Installation](#installation)
- [Usage](#usage)
- [API Endpoints](#api-endpoints)
- [Authentication](#authentication)
- [Example Request](#example-request)

## Installation

1. Clone the repository:
   ```shell
   git clone https://github.com/kandlagifari/go-books-apps.git
   cd go-books-apps
   ```

2.  Install dependencies:
    ```shell
    go mod tidy
    ```

3. Set up your `.env` file with necessary environment variables:
   ```txt
   # PostgreSQL
   DB_HOST=<your_host>
   DB_PORT=<your_port>
   DB_USER=<your_user>
   DB_PASSWORD=<your_password>
   DB_NAME=<your_database>

   # JWT Secret Key
   JWT_SECRET_KEY=<your_jwt_secret_key>
   ```

4. Run the migrations to set up the database and start web server:
   ```shell
   # With make
   make build && make run

   # Without make
   ./build.sh && ./bootstrap &
   ```

5. The server will be running on http://localhost:4321

## Usage

### Authentication Endpoints:

This API uses JWT (JSON Web Tokens) for user authentication. You need to include the token in the `Authorization` header in each request to access protected endpoints.

#### 1. User register:
- **POST** `/api/auth/register`: User register to access API endpoints.
  - **Request Body**:
    ```json
    {
      "username": "user1",
      "password": "password1"
    }
    ```
  - **Response**:
    ```json
    {
      "message": "User registered successfully",
      "user_id": 1
    }
    ```

#### 1. User login:
- **POST** `/api/auth/login`: Logs in a user and provides a JWT token.
  - **Request Body**:
    ```json
    {
      "username": "user1",
      "password": "password1"
    }
    ```
  - **Response**:
    ```json
    {
      "message": "Login successful",
      "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6InVzZXIxIiwiZXhwIjoxNzMxOTIzMDExfQ.qFbs28ECUpu0mOoXYo0J0GZ6hrF_7M5K08dLS0ApSi8"
    }
    ```

---

### API Categories Endpoints

#### 1. Get all Categories
- **GET** `/api/categories`
  - **Description**: Retrieves a list of all categories.
  - **Response**: 
    ```json
    [
      {
        "id": 1,
        "name": "Science",
        "created_at": "2024-11-17T17:20:22.666205Z",
        "created_by": "user1",
        "modified_at": "2024-11-17T10:20:22.667011Z",
        "modified_by": ""
      }
    ]
    ```

#### 2. Create a Category
- **POST** `/api/categories`
  - **Description**: Creates a new category.
  - **Request Body**:
    ```json
    {
      "name": "New Category"
    }
    ```
  - **Response**:
    ```json
    {
      "message": "Category created successfully"
    }
    ```

#### 3. Get Category by ID
- **GET** `/api/categories/:id`
  - **Description**: Retrieves a single category by its ID.
  - **Response**:
    ```json
    {
      "created_at": "2024-11-17T17:20:22.666205Z",
      "created_by": "user1",
      "id": 1,
      "modified_at": "2024-11-17T10:42:55.266629Z",
      "modified_by": "user1",
      "name": "Technology"
    }
    ```

#### 4. Update Category by ID
- **PUT** `/api/categories/:id`
  - **Description**: Updates an existing category.
  - **Request Body**:
    ```json
    {
      "name": "Technology"
    }
    ```
  - **Response**:
    ```json
    {
      "message": "Category updated successfully"
    }
    ```

#### 5. Delete Category by ID
- **DELETE** `/api/categories/:id`
  - **Description**: Deletes a category by its ID.
  - **Response**:
    ```json
    {
      "message": "Category deleted successfully"
    }
    ```

#### 6. Get Books by Category ID
- **GET** `/api/categories/:id/books`
  - **Description**: Retrieves all books in a specific category.
  - **Response**:
    ```json
    [
      {
        "id": 1,
        "title": "Book Title",
        "description": "Book Description",
        "release_year": 2020,
        "price": 20000,
        "total_page": 300,
        "thickness": "tebal",
        "category_id": 1,
        "created_at": "2024-01-01T12:00:00Z",
        "created_by": "admin"
      },
      ...
    ]
    ```

---

### API Books Endpoints

#### 1. Get All Books
- **GET** `/api/books`
  - **Description**: Retrieves a list of all books.
  - **Response**:
    ```json
    [
      {
        "id": 1,
        "title": "Book Title",
        "description": "Book Description",
        "release_year": 2020,
        "price": 20000,
        "total_page": 300,
        "thickness": "tebal",
        "category_id": 1,
        "created_at": "2024-01-01T12:00:00Z",
        "created_by": "admin"
      },
      ...
    ]
    ```

#### 2. Create a Book
- **POST** `/api/books`
  - **Description**: Creates a new book.
  - **Request Body**:
    ```json
    {
      "title": "Book Title",
      "description": "Book Description",
      "release_year": 2020,
      "price": 20000,
      "total_page": 300,
      "category_id": 1
    }
    ```
  - **Response**:
    ```json
    {
      "message": "Book created successfully"
    }
    ```

#### 3. Get Book by ID
- **GET** `/api/books/:id`
  - **Description**: Retrieves a single book by its ID.
  - **Response**:
    ```json
    {
      "id": 1,
      "title": "Book Title",
      "description": "Book Description",
      "release_year": 2020,
      "price": 20000,
      "total_page": 300,
      "thickness": "tebal",
      "category_id": 1,
      "created_at": "2024-01-01T12:00:00Z",
      "created_by": "admin"
    }
    ```

#### 4. Update Book by ID
- **PUT** `/api/books/:id`
  - **Description**: Updates an existing book.
  - **Request Body**:
    ```json
    {
      "title": "Updated Book Title",
      "description": "Updated Description",
      "release_year": 2021,
      "price": 25000,
      "total_page": 350,
      "category_id": 2
    }
    ```
  - **Response**:
    ```json
    {
      "message": "Book updated successfully"
    }
    ```

#### 5. Delete Book by ID
- **DELETE** `/api/books/:id`
  - **Description**: Deletes a book by its ID.
  - **Response**:
    ```json
    {
      "message": "Book deleted successfully"
    }
    ```

---