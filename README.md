# Completed endpoints

| No. | Endpoint                                         | Description                                                                    |
|-----|--------------------------------------------------|--------------------------------------------------------------------------------|
| 1   | `POST /v1/auth/login`                            | Initiates the login process.                                                   |
| 2   | `POST /v1/auth/login/verify`                     | Verifies the login credentials and completes the login process.                |
| 3   | `POST /v1/auth/register`                         | Initiates the user registration process.                                       |
| 4   | `POST /v1/auth/register/verify`                  | Verifies the user registration details and completes the registration process. |
| 5   | `GET /v1/auth/user/:user_id`                     | Retrieves user metadata based on the specified user_id.                        |
| 6   | `PUT /v1/auth/user/:user_id`                     | Updates user metadata based on the specified user_id.                          |
| 7   | `GET /v1/auth/logout`                            | Performs user logout.                                                          |
| 8   | `GET /v1/dealership/:id`                         | Retrieves dealership information based on the specified id.                    |
| 9   | `GET /v1/dealership/user/:id`                    | Retrieves dealership information associated with the specified user_id.        |
| 10  | `PUT /v1/dealership/:id`                         | Updates dealership information based on the specified id.                      |
| 11  | `DELETE /v1/dealership/:id`                      | Soft deletes a dealership`                                                     |
|     | `POST   /v1/auth/login`                          |                                                                                |
|     | `POST   /v1/auth/login/verify`                   |                                                                                |
|     | `POST   /v1/auth/register`                       |                                                                                |
|     | `POST   /v1/auth/register/verify`                |                                                                                |
|     | `GET    /v1/auth/user/:user_id`                  |                                                                                |
|     | `PUT    /v1/auth/user/:user_id`                  |                                                                                |
|     | `PUT    /v1/auth/user/:user_id/dealership`       |                                                                                |
|     | `POST   /v1/auth/refresh`                        |                                                                                |
|     | `GET    /v1/auth/logout`                         |                                                                                |
|     | `GET    /v1/dealership/:id`                      |                                                                                |
|     | `GET    /v1/dealership/user/:id`                 |                                                                                |
|     | `PUT    /v1/dealership/:id`                      |                                                                                |
|     | `POST   /v1/dealership/`                         |                                                                                |
|     | `DELETE /v1/dealership/:id`                      |                                                                                |
|     | `GET    /v1/cars/`                               |                                                                                |
|     | `POST   /v1/cars/`                               |                                                                                |
|     | `PUT    /v1/cars/:car_id`                        |                                                                                |
|     | `GET    /v1/cars/:car_id`                        |                                                                                |
|     | `DELETE /v1/cars/:car_id`                        |                                                                                |
|     | `GET    /v1/cars/search`                         |                                                                                |
|     | `GET    /v1/cars/filter`                         |                                                                                |
|     | `POST   /v1/cars/brand`                          |                                                                                |
|     | `GET    /v1/cars/brand`                          |                                                                                |
|     | `PUT    /v1/cars/brand/:brand_id`                |                                                                                |
|     | `GET    /v1/cars/brand/:brand_id`                |                                                                                |
|     | `DELETE /v1/cars/brand/:brand_id`                |                                                                                |
|     | `POST   /v1/cars/:car_id/extra`                  |                                                                                |
|     | `GET    /v1/cars/:car_id/extra`                  |                                                                                |
|     | `PUT    /v1/cars/extra/:extra_id`                |                                                                                |
|     | `DELETE /v1/cars/extra/:extra_id`                |                                                                                |
|     | `POST   /v1/cars/:car_id/image`                  |                                                                                |
|     | `GET    /v1/cars/:car_id/image`                  |                                                                                |
|     | `PUT    /v1/cars/image/:image_id`                |                                                                                |
|     | `DELETE /v1/cars/image/:image_id`                |                                                                                |
|     | `GET    /v1/cars/by/dealership/:dealership_id`   |                                                                                |
|     | `GET    /v1/cars/by/brand/:brand_id`             |                                                                                |
|     | `GET    /v1/cars/by/dealer/:dealer_id`           |                                                                                |
|     | `GET    /v1/cars/by/dealer/count`                |                                                                                |
|     | `GET    /v1/cars/by/dealership/count`            |                                                                                |
|     | `GET    /v1/cars/by/brand/count`                 |                                                                                |
|     | `GET    /v1/spares/`                             |                                                                                |
|     | `POST   /v1/spares/`                             |                                                                                |
|     | `PUT    /v1/spares/:spare_id`                    |                                                                                |
|     | `GET    /v1/spares/:spare_id`                    |                                                                                |
|     | `DELETE /v1/spares/:spare_id`                    |                                                                                |
|     | `GET    /v1/spares/search`                       |                                                                                |
|     | `GET    /v1/spares/filter`                       |                                                                                |
|     | `POST   /v1/spares/image`                        |                                                                                |
|     | `GET    /v1/spares/image`                        |                                                                                |
|     | `GET    /v1/spares/:spare_id/image`              |                                                                                |
|     | `PUT    /v1/spares/image/:image_id`              |                                                                                |
|     | `DELETE /v1/spares/image/:image_id`              |                                                                                |
|     | `GET    /v1/spares/by/brand/:brand_id`           |                                                                                |
|     | `GET    /v1/spares/by/category/:category_id`     |                                                                                |
|     | `GET    /v1/spares/by/car/:car_id`               |                                                                                |
|     | `GET    /v1/spares/by/dealer/:dealer_id`         |                                                                                |
|     | `GET    /v1/spares/by/dealership/:dealership_id` |                                                                                |


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
