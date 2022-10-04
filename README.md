# Order service

- Order service implements functions of registering users, managing products and customer orders.

## Project Architecture
The project implements `Clean Architecture` which helps to separate concerns by organizing code into several layers with a very explicit rule which enables us to create a testable and maintainable project. [The Clean Architecture](https://blog.8thlight.com/uncle-bob/2012/08/13/the-clean-architecture.html).
### The project is divided into the following layers:
The project has the following layers:
#### 1. Presentation
This represents logic that consume the business logic from the `Usecase Layer`
and renders to the view. Here you can choose to render the view in e.g `rest`

#### 2. Usecases
The code in this layer contains application specific business rules.
This represents the pure business logic of the application.
The rules of the application also shouldn't rely on the UI or the persistence frameworks being used.

#### 3. Interfaces
Clean architecture dictates that dependandency should only point inwards therefore the inner layers(the usecase layer) should not have any idea of the implementations of the database, third party interactions. So this is just an interface.
This will ensures that the system is independent of a database and any third party agencies making it easier to switch them without affecting the business logic.

#### 4. Infrastructure
These are the `ports` that allow the system to talk to 'outside things' which
could be a `database` for persistence or a `web server` for the UI. None of
the inner use cases or domain entities should know about the implementation of
these layers and if we choose to change them they should not cause change to any of our business rules.

#### 5. Domain
Here we have `business objects` or `entities` and should represent and encapsulate the fundamental business rules.

## Technologies
- Golang 1.17
- Gorilla Mux
- Postgres

## How to use it

1. First clone the codefrom the repository

```bash
    user@user:~$ git clone
```
2. Set up your local environment by creating a `env.sh` and add your environment variables

```bash
    export DBPORT="5432"
    export DBPASSWORD="Shinnok1996"
    export DBUSER="postgres"
    export DBNAME="orderservice"
    export DBHOST="localhost"
    export AUTH_PASSWORD="basic-auth-password"
    export AUTH_USERNAME="basic-auth-username"
    export AIT_URL="https://api.sandbox.africastalking.com/version1/messaging"
    export AIT_API_KEY="217ebb97eebf8ba1c722a5c3f16e5f7ea83a87f370c2b5da5c985a523d3c406d"
    export AIT_USERNAME="sandbox"
    export OTP_ISSUER="OrderCustomer"
    export OTP_ACCOUNTNAME="victorineosewe@gmail.com"
```
3. Install dependencies.

```bash
    user@user:~$ go mod tidy
    user@user:~$ go generate ./...
```

4. Run the server

```bash
    user@user:~$ go run main.go
```