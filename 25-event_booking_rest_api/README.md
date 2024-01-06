# Event Booking API

## Overview

The Event Booking API is a GoLang-based web server designed for creating, managing, and querying events. It provides a RESTful interface for CRUD operations on event data. The API utilizes SQLite for database management and JWT (JSON Web Tokens) for secure authentication. This project serves as a practical implementation to demonstrate GoLang's capabilities in building scalable and maintainable web services.

## API Endpoints

Below is the list of available API endpoints in the application:

### User Authentication and Registration

- **Login**
  - **Endpoint:** `/login`
  - **Method:** `POST`
  - **Description:** Authenticate a user and provide a token.

- **Register**
  - **Endpoint:** `/register`
  - **Method:** `POST`
  - **Description:** Register a new user.

### Event Management

- **List Events**
  - **Endpoint:** `/events`
  - **Method:** `GET`
  - **Description:** Retrieve a list of all upcoming events.

- **Create Event**
  - **Endpoint:** `/events`
  - **Method:** `POST`
  - **Auth Required**
  - **Description:** Create a new event. Requires user authentication.

- **Get Event**
  - **Endpoint:** `/events/<id>`
  - **Method:** `GET`
  - **Description:** Retrieve details of a specific event by its ID.

- **Update Event**
  - **Endpoint:** `/events/<id>`
  - **Method:** `PUT`
  - **Auth Required**
  - **Description:** Update an existing event. Requires user authentication.

- **Delete Event**
  - **Endpoint:** `/events/<id>`
  - **Method:** `DELETE`
  - **Auth Required**
  - **Description:** Delete an existing event. Requires user authentication.

### Event Booking

- **Register for Event**
  - **Endpoint:** `/events/<id>/register`
  - **Method:** `POST`
  - **Auth Required**
  - **Description:** Register the authenticated user for an event.

- **Cancel Event Registration**
  - **Endpoint:** `/events/<id>/register`
  - **Method:** `DELETE`
  - **Auth Required**
  - **Description:** Cancel the authenticated user's registration for an event.

## Project Structure Overview

- /cmd
    - /server
        - main.go : Entry point of the application.
- /internal
    - /app
        - /handlers : Contains HTTP handlers for web requests.
        - /middlewares : Contains middleware functions.
        - /models : Defines data structures and models.
        - /services : Contains business logic.
    - /config
        - config.go : Manages configuration.
- /pkg
    - /database
        - db.go : Manages database connections and interactions.
    - /auth
        - jwt.go : Handles authentication logic.
- /scripts
    - init_db.sql : SQL scripts for database initialization.
- go.mod and go.sum : Manage dependencies.
- .env : Stores environment variables.
- README.md : Provides project documentation.


## Explanation of Structure:


    

- `/cmd/server/main.go` : 
    - **Purpose:** Serves as the entry point of the application.
    - **Shouldn't do:** Should not contain business logic or direct database handling.
    **Responsibilities:**
        - Entry point of the application.
        - Initializes main components like database connections, configurations, and HTTP server.
        - Sets up routing and middleware.

- `/internal`
    - `/app`
    - This is where the core application code resides. It's split into different subdirectories for organization.

        - `/handlers`
        **Purpose:** Contains the HTTP handlers that respond to web requests.
        **Shouldn't do:** Shouldn't implement business logic or interact with the database directly.
        **Responsibilities:**
            - Parsing requests, calling business logic functions (services), and returning responses.
            - Core application code.
    
        - `/middlewares`
            **Purpose:** Contains middleware functions.
            **Shouldn't do:** Shouldn't handle business logic or database operations.
            **Responsibilities:**
            - Executing pre/post-processing logic on HTTP requests (like JWT token validation).
    
        - `/models`
            **Purpose:** Defines data structures and models.
            **Shouldn't do:** Shouldn't contain logic for database or business processes.
            **Responsibilities:**
            - Representing entities (like events, users) and their attributes.
    
        - `/services`
            **Purpose:** Contains business logic.
            **Shouldn't do:** Shouldn't directly handle HTTP requests/responses or database connection logic.
            **Responsibilities:**
            - Implementing core functionalities (like CRUD operations for events).
    - `/config`
    **Purpose:** Manages configuration.
    **Shouldn't do:** Shouldn't be involved in business logic or request handling.
    **Responsibilities:**
        - Loading and providing access to configurations like database paths, API keys, etc.


- `/pkg`
This directory is for libraries/packages that can be potentially reused across different projects.

    - `/database`
**Purpose:** Manages database connections and interactions.
**Shouldn't do:** Shouldn't implement business logic or handle HTTP requests.
**Responsibilities:**
        - Establishing database connection, executing SQL queries, etc.

    - `/auth`
**Purpose:** Handles authentication-related logic.
**Shouldn't do:** Shouldn't deal with database operations or business logic unrelated to authentication.
**Responsibilities:**
    - Managing JWT tokens, including creation and validation.

- `/scripts/init_db.sql`
    **Purpose:** Contains SQL scripts for database initialization.
    **Shouldn't do:** Shouldn't contain application logic or configuration details.
    **Responsibilities:**
        - Creating tables, initializing data, etc.

- `.env`

    **Purpose:** Stores environment variables.
    **Shouldn't do:** Shouldn't be committed to version control if it contains sensitive data.
    **Responsibilities:**
    - Holds sensitive/non-static configuration data like database credentials, secret keys, etc.




## Install Dependencies:

```shell
go get -u
```


## Run the app 
```shell
air
```


## Build the app 
```shell
go build ./cmd/server/main.go
```

<!-- central error handling -->
<!-- validation -->
<!-- create the database if it doens't exist -->
<!-- create tables if doesn't exist -->
<!-- using services -->
<!-- using handlers -->
<!-- using models -->