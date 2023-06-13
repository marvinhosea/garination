# Completed endpoints

| No. | Endpoint                        | Description                                                                    |
|-----|---------------------------------|--------------------------------------------------------------------------------|
| 1   | `POST /v1/auth/login`           | Initiates the login process.                                                   |
| 2   | `POST /v1/auth/login/verify`    | Verifies the login credentials and completes the login process.                |
| 3   | `POST /v1/auth/register`        | Initiates the user registration process.                                       |
| 4   | `POST /v1/auth/register/verify` | Verifies the user registration details and completes the registration process. |
| 5   | `GET /v1/auth/user/:user_id`    | Retrieves user metadata based on the specified user_id.                        |
| 6   | `PUT /v1/auth/user/:user_id`    | Updates user metadata based on the specified user_id.                          |
| 7   | `GET /v1/auth/logout`           | Performs user logout.                                                          |
| 8   | `GET /v1/dealership/:id`        | Retrieves dealership information based on the specified id.                    |
| 9   | `GET /v1/dealership/user/:id`   | Retrieves dealership information associated with the specified user_id.        |
| 10  | `PUT /v1/dealership/:id`        | Updates dealership information based on the specified id.                      |
| 11  | `DELETE /v1/dealership/:id`     | Soft deletes a dealership                                                      |

# Other Tasks
- [x] Setup Liquibase for database migrations
- [x] Setup Prometheus infrastructure for metrics
- [x] Setup Grafana infrastructure for metrics visualization
- [ ] Setup Jaeger infrastructure for tracing
- [x] Send app-db metrics to Prometheus
- [x] Send app-api-gateway metrics to Prometheus
- [x] Set up Casdoor for authentication
- [x] Set up Redis for cache
- [x] Set up postgresql for persistence
- [ ] Write Authentication middleware
- [ ] Write Authorization/Permision middleware
- [ ] Write metrics middleware
- [ ] Write tracing code
- [ ] Write unit tests
- [ ] Write file upload code
