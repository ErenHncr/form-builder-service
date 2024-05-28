## Go API Structure
This is a simple API structure for a Go application. The structure is designed to be modular and scalable. The structure is as follows:
  - **api** contains the code for the backend of the application. 
  - **storage** contains the code for database management.
  - **types** contains the code for structuring the data.
  - **utils** contains the code for utility functions.

Entry point of the application is in the ```main.go``` file. The ```main.go``` file initializes the database and starts the server. The server is started on port ```3000``` by default. ```listenaddr``` parameter in the ```main.go``` file can be changed to start the server on a different port.


TODOs

[x] - Add MustHaveKey method to Question struct
[x] - Add GetQuestion, DeleteQuestion and UpdateQuestion methods
[x] - Separate create and update validation functions
[ ] - Add minimum and maximum methods to QuestionLabel struct for validation
[ ] - Add environment variables and read with https://github.com/joho/godotenv
[ ] - Add postgres and mongodb database
  https://vercel.com/docs/storage/vercel-postgres
  [ ] - Migration
    - https://www.freecodecamp.org/news/database-migration-golang-migrate/
    - https://github.com/pressly/goose and https://sqlc.dev
    - https://github.com/jackc/tern and https://github.com/jackc/pgx
  [ ] - Add database connection and query methods
  [ ] - Add database migration methods
  [ ] - Add database seeding methods
[ ] - Create Question and Form structs in the types package.
[ ] - Send a custom validation code for different cases
[ ] - Separate question struct into different files