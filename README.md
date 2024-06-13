## Go API Structure
This is a simple API structure for a Go application. The structure is designed to be modular and scalable. The structure is as follows:
  - **api** contains the code for the backend of the application. 
  - **storage** contains the code for database management.
  - **types** contains the code for structuring the data.
  - **utils** contains the code for utility functions.

Entry point of the application is in the ```main.go``` file. The ```main.go``` file initializes the database and starts the server. The server is started on port ```3000``` by default. ```listenaddr``` parameter in the ```main.go``` file can be changed to start the server on a different port.


### Benchmarking the API
[**wrk**](https://github.com/wg/wrk) is a modern HTTP benchmarking tool capable of generating significant load when run on a single multi-core CPU. It combines a multithreaded design with scalable event notification systems such as epoll and kqueue. An optional LuaJIT script can perform HTTP request generation, response processing, and custom reporting.

#### Installation
```bash
brew install wrk
```

#### Usage
```GET /questions```

12 threads, 500 connections, 30 seconds
```bash
wrk -t12 -c500 -d30s --latency http://127.0.0.1:3000/questions
```

### To Do

- [x] Create Question struct in the types package.
- [x] Add MustHaveKey method to Question struct
- [x] Add GetQuestion, DeleteQuestion and UpdateQuestion methods
- [x] Separate create and update validation functions
- [x] Add minimum and maximum methods to QuestionLabel struct for validation
- [x] Add minimum and maximum key length validation
- [x] Add GET method for questions
- [x] Add environment variables and read with https://github.com/joho/godotenv
- [ ] Add postgres and mongodb database
  - [ ] Add seeding database
  - [x] Add mongodb storage
    - [x] Add database connection
    - [x] Add query methods
      - [x] Add Create method
      - [x] Add Get method
      - [x] Add Delete method
      - [x] Add List method
        - [x] Add pagination
        - [x] Add sorting
          - [x] add default sorting for createdAt field
          - [x] add custom sorting for different fields (updatedAt, createdAt)
      - [x] Add Update method
  - [ ] Add postgres storage
    https://vercel.com/docs/storage/vercel-postgres
    - [ ] Migration
      https://www.freecodecamp.org/news/database-migration-golang-migrate/
      https://github.com/pressly/goose and https://sqlc.dev
      https://github.com/jackc/tern and https://github.com/jackc/pgx
    - [ ] Add database connection and query methods
    - [ ] Add database migration methods
- [ ] Create Form struct in the types package.
- [x] add swagger documentation
  - [x] add swagger documentation for questions
- [ ] add CI/CD pipeline
  - [ ] add github actions
  - [ ] add dockerfile
  - [ ] add gcp cloud run deployment
- [ ] Add validation for other question keys