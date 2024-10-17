## SGBANK (Simple Go Bank)
A simple bank backend application built with Go, allowing users to create multiple accounts with different currency codes and manage transfer transactions between these accounts in a concurrent environment.

- RESTful APIs using JSON format, with request validation handled via [Gin](https://github.com/gin-gonic/gin)

- PostgreSQL, [Golang migrate](https://github.com/golang-migrate/migrate), and [sqlc](https://github.com/sqlc-dev/sqlc)

- Token-based authentication with JWT vs PASETO

- Write unit tests for database queries and API handlers

## How to Run
### Prerequisites
- Go 

- PostgreSQL

- ```Docker``` installed

- ```sqlc``` installed for code generation

- ```Golang migrate``` installed for database migration

- Linux/MacOS, or Windows with WSL2, should have ```make``` installed

### Launch Postgres container
```
make postgres
```

### Create Database
```
make createdb 
```

### Database Migration
```
make migrateup
```
```
make migratedown
```

### Generate Code with ```sqlc```
```
make sqlc
```

### Generate Mock Store for API Handler Testing
```
make mock
```

### Run All Unit Tests
```
make test
```

###  Run The Server

**./app.env**
```
DB_DRIVER=<"database driver">
DB_SOURCE=<"database source">
SERVER_ADDRESS=<"server address">
TOKEN_SYMMETRIC_KEY=<"secret key to sign token">
ACCESS_TOKEN_DURATION=<"access token duration">
```

```
make server
```

## Database Diagram
![sgbank](https://github.com/user-attachments/assets/a791039e-c755-4d30-8720-7850b3e65f32)<?xml version="1.0" standalone="no"?><svg version="1.1" xmlns="http://www.w3.org/2000/svg" width="967.5525071104554" height="436.8694316436253">

## API appendix

### User APIs

| Method | Endpoint       | Description                     | Request Body Example         | Response Body Example                                       | Authentication |
|--------|----------------|----------------------------------|----------------------|-------------------------------------------------------------|----------------|
| POST    | `/v1/users`        | Create a specific user          | `{"username": "nhhuy2002", "full_name": "Nguyen Nhut Huy", "email":"nguyennhuthuy02@gmail.com", "password": "9999999"}`         | `{"username": "nhhuy2002", "full_name": "Nguyen Nhut Huy", "email": "nguyennhuthuy02@gmail.com", "password_changed_at": "0001-01-01T00:00:00Z", "created_at": "2024-10-14T12:06:28.500453Z"}` | No             |
| POST    | `/v1/users/sign-in`   | Sign in  | `{"username": "nhhuy2002","password": "9999999"}` |  `{"access_token": "v2.local.FjXfgYue2N0OinFgH-OcSuDhwfRXJ_Y6qxXyGasAfD7ofQbmNbGIriNdX-qwKEeJ9z5dyTLToP_TVkLchQ8_gFzbul5kSAga6bW6iiIU9wusCAIa2tn09165-7an4mn1MEO4trvVyrUDjumQmIHUOslyGFWB0J-MUf0H-ekRNnXI4dWHAqhD3ExYqsQMdfbKz3VLom_8kAIIb9hbedBQ5XDocRmgwcodu-ydwepSyha_cd-rZNh2Q4H3a0Qr67ZDK43eerh8IERgkrMIZTI2ew.bnVsbA", "user": {"username": "nhhuy2002", "full_name": "Nguyen Nhut Huy", "email": "nguyennhuthuy02@gmail.com", "password_changed_at": "0001-01-01T00:00:00Z", "created_at": "2024-10-14T08:52:29.241677Z"}}`| No            |

### Account APIs

| Method | Endpoint       | Description                     | Request Body Example         | Response Body Example                                       | Authentication |
|--------|----------------|----------------------------------|----------------------|-------------------------------------------------------------|----------------|
| GET    | `/v1/accounts?page_id=1&page_size=5`        | Get all accounts owned by a specific user           | N/A                  | `{"accounts": [{"id": 1, "owner": "nhhuy2002", "balance": 0, "currency": "CAD", "created_at": "2024-10-14T12:07:56.383739Z"}]}`| Yes             |
| GET    | `/v1/accounts/:id`   | Get a specific account of the user  | N/A |  `{"account": {"id": 1, "owner": "nhhuy2002", "balance": 0, "currency": "CAD", "created_at": "2024-10-14T12:07:56.383739Z"}}` | Yes            |
| POST    | `/v1/accounts`   | Create a account with a currency code  | `{"Currency": "CAD"}` | `{"account": {"id": 1, "owner": "nhhuy2002", "balance": 0, "currency": "CAD", "created_at": "2024-10-14T12:07:56.383739Z"}}` | Yes            |

### Transfers APIs
| Method | Endpoint       | Description                     | Request Body Example         | Response Body Example                                       | Authentication |
|--------|----------------|----------------------------------|----------------------|-------------------------------------------------------------|----------------|
| POST    | `/v1/transfers`   | Transfer money between two accounts which have same currency code  | `{"from_account_id": 1, "to_account_id": 9, "amount": 300, "currency": "CAD"}` | `{"transfer": {"id": 30, "from_account_id": 1, "to_account_id": 9, "amount": 300, "created_at": "2024-10-14T12:16:45.771039Z"}, "from_account": {"id": 1, "owner": "nhhuy2002", "balance": 700, "currency": "CAD", "created_at": "2024-10-14T12:07:56.383739Z"}, "to_account": {"id": 9, "owner": "mppvlsv", "balance": 768, "currency": "CAD", "created_at": "2024-10-14T12:15:13.682382Z"}, "from_entry": {"id": 59, "account_id": 1, "amount": -300, "created_at": "2024-10-14T12:16:45.771039Z"}, "to_entry": {"id": 60, "account_id": 9, "amount": 300, "created_at": "2024-10-14T12:16:45.771039Z"}}` | Yes            |

### Notes:
- All responses are in JSON format as well.

- For endpoints marked with "Yes" in the Authentication column, a valid API key is required.

- The API key is sent using the `Authorization` header, formatted as follows: 
    ```
    authorization: bearer <"JWT token" or "PASETO token">
    ```

- Can use [REST Client](https://marketplace.visualstudio.com/items?itemName=humao.rest-client) extension to test these APIs

## References
[1]. Techschool. (n.d.). GitHub - techschool/simplebank: Backend master class: build a simple bank service in Go. GitHub. https://github.com/techschool/simplebank


