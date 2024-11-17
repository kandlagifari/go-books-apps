# Go Books API

This is a simple API for managing books and categories built with the Gin framework in Go. The API allows users to:

- Authenticate using JWT tokens.
- Manage categories API.
- Manage books API.

## Table of Contents
- [Installation](#installation)
- [Usage](#usage)
- [Negative Test](#negative-test)

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

3. Set up your `config/.env` file with necessary environment variables:
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

### Endpoint 1: Authentication API

This API uses JWT (JSON Web Tokens) for user authentication. You need to include the token in the `Authorization` header in each request to access protected endpoints.

#### 1. User Register
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
    ![Alt text](images/01_post-user-register.png)

#### 2. User Login
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
    ![Alt text](images/02_post-user-login.png)

---

### Endpoint 2: Categories API

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
      },
      ...
    ]
    ```
    ![Alt text](images/04_get-all-categories.png)

#### 2. Create a Category
- **POST** `/api/categories`
  - **Description**: Creates a new category.
  - **Request Body**:
    ```json
    {
      "name": "Science"
    }
    ```
  - **Response**:
    ```json
    {
      "message": "Category created successfully"
    }
    ```
    ![Alt text](images/03_post-category.png)

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
    ![Alt text](images/05_get-category-by-id.png)

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
    ![Alt text](images/06_put-category-by-id.png)
    ---
    **Verify**
    ![Alt text](images/07_get-category-after-update.png)

#### 5. Delete Category by ID
- **DELETE** `/api/categories/:id`
  - **Description**: Deletes a category by its ID.
  - **Response**:
    ```json
    {
      "message": "Category deleted successfully"
    }
    ```
    ![Alt text](images/08_delete-category-by-id.png)
    ---
    **Verify**
    ![Alt text](images/09_get-category-after-delete.png)
    
    ![Alt text](images/10_get-all-categories-after-delete.png)

#### 6. Get Books by Category ID
- **GET** `/api/categories/:id/books`
  - **Description**: Retrieves all books in a specific category.
  - **Response**:
    ```json
    [
     {
        "id": 3,
        "title": "Sword Art Online",
        "description": "Sword Art Online (SAO), one of the most recent games on the console, offers a gateway into the wondrous world of Aincrad, a vivid, medieval landscape where users can do anything within the limits of imagination.",
        "image_url": "https://cdn.myanimelist.net/images/anime/11/39717.jpg",
        "release_year": 2012,
        "price": 16,
        "total_page": 130,
        "thickness": "tebal",
        "category_id": 2,
        "created_at": "2024-11-17T20:27:38.3768Z",
        "created_by": "user1",
        "modified_at": "2024-11-17T13:27:38.377213Z",
        "modified_by": ""
      },
      {
        "id": 4,
        "title": "Sword Art Online II",
        "description": "Approached by officials to assist in investigating the murders, Kazuto assumes his persona of Kirito once again and logs into Gun Gale Online, intent on stopping the killer.",
        "image_url": "https://cdn.myanimelist.net/images/anime/1223/121999.jpg",
        "release_year": 2014,
        "price": 16,
        "total_page": 95,
        "thickness": "tipis",
        "category_id": 2,
        "created_at": "2024-11-17T20:27:43.781472Z",
        "created_by": "user1",
        "modified_at": "2024-11-17T13:27:43.782148Z",
        "modified_by": ""
      },
      ...
    ]
    ```
    ![Alt text](images/19_get-books-by-category-id.png)

---

### Endpoint 3: Books API

#### 1. Get All Books
- **GET** `/api/books`
  - **Description**: Retrieves a list of all books.
  - **Response**:
    ```json
    [
      {
        "id": 1,
        "title": "Dr. Stone",
        "description": "Blinding green light strikes the Earth and petrifies mankind around the world—turning every single human into stone.",
        "image_url": "https://cdn.myanimelist.net/images/anime/1613/102576.jpg",
        "release_year": 2019,
        "price": 16,
        "total_page": 120,
        "thickness": "tebal",
        "category_id": 1,
        "created_at": "2024-11-17T18:50:07.163542Z",
        "created_by": "user1",
        "modified_at": "2024-11-17T11:50:07.164832Z",
        "modified_by": ""
      },
      ...
    ]
    ```
    ![Alt text](images/12_get-all-books.png)

#### 2. Create a Book
- **POST** `/api/books`
  - **Description**: Creates a new book.
  - **Request Body**:
    ```json
    {
      "title": "Dr. Stone",
      "description": "Blinding green light strikes the Earth and petrifies mankind around the world—turning every single human into stone.",
      "image_url": "https://cdn.myanimelist.net/images/anime/1613/102576.jpg",
      "release_year": 2019,
      "price": 16,
      "total_page": 120,
      "category_id": 1
    }
    ```
  - **Response**:
    ```json
    {
      "message": "Book created successfully"
    }
    ```
    ![Alt text](images/11_post-book.png)

#### 3. Get Book by ID
- **GET** `/api/books/:id`
  - **Description**: Retrieves a single book by its ID.
  - **Response**:
    ```json
    {
      "category_id": 1,
      "created_at": "2024-11-17T18:50:07.163542Z",
      "created_by": "user1",
      "description": "Blinding green light strikes the Earth and petrifies mankind around the world—turning every single human into stone.",
      "id": 1,
      "image_url": "https://cdn.myanimelist.net/images/anime/1613/102576.jpg",
      "modified_at": "2024-11-17T11:50:07.164832Z",
      "modified_by": "",
      "price": 16,
      "release_year": 2019,
      "thickness": "tebal",
      "title": "Dr. Stone",
      "total_page": 120
    }
    ```
    ![Alt text](images/13_get-book-by-id.png)

#### 4. Update Book by ID
- **PUT** `/api/books/:id`
  - **Description**: Updates an existing book.
  - **Request Body**:
    ```json
    {
      "title": "Dr. Stone: Stone Wars",
      "description": "Senkuu has made it his goal to bring back two million years of human achievement and revive the entirety of those turned to statues.",
      "image_url": "https://cdn.myanimelist.net/images/anime/1711/110614.jpg",
      "release_year": 2021,
      "price": 16,
      "total_page": 90,
      "category_id": 1
    }
    ```
  - **Response**:
    ```json
    {
      "message": "Book updated successfully"
    }
    ```
    ![Alt text](images/14_put-book-by-id.png)
    ---
    **Verify**
    ![Alt text](images/15_get-book-after-update.png)

#### 5. Delete Book by ID
- **DELETE** `/api/books/:id`
  - **Description**: Deletes a book by its ID.
  - **Response**:
    ```json
    {
      "message": "Book deleted successfully"
    }
    ```
    ![Alt text](images/16_delete-book-by-id.png)
    ---
    **Verify**
    ---
    ![Alt text](images/17_get-book-after-delete.png)
    
    ![Alt text](images/18_get-all-books-after-delete.png)

## Negative Test

### Endpoint 1: Authentication API

#### 1. User Login Error
  ![Alt text](images/00_user-login-error.png)
  
#### 2. No Token Authorization Error
  ![Alt text](images/00_no-token-authorization-error.png)

---

### Endpoint 2: Categories API

#### 1. Category Name Unique Error
  ![Alt text](images/00_category-name-unique-error.png)

#### 2. Delete Category Not Found Error
  ![Alt text](images/00_delete-category-not-found-error.png)

#### 3. Update Category Not Found Error
  ![Alt text](images/00_update-category-not-found-error.png)

---

### Endpoint 3: Books API

#### 1. Book Title Unique Error
  ![Alt text](images/00_book-title-unique-error.png)

#### 2. Create Book Min Year Error
  ![Alt text](images/00_create-book-min-year-error.png)

#### 3. Create Book Max Year Error
  ![Alt text](images/00_create-book-max-year-error.png)

#### 4. Delete Book Not Found Error
  ![Alt text](images/00_delete-book-not-found-error.png)

#### 5. Update Book Not Found Error
  ![Alt text](images/00_update-book-not-found-error.png)
