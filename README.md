# Completed endpoints
| No. | Endpoint                                      | Description                                                                    |
|-----|-----------------------------------------------|--------------------------------------------------------------------------------|
| 1   | `POST /v1/auth/login`                         | Initiates the login process.                                                   |
| 2   | `POST /v1/auth/login/verify`                  | Verifies the login credentials and completes the login process.                |
| 3   | `POST /v1/auth/register`                      | Initiates the user registration process.                                       |
| 4   | `POST /v1/auth/register/verify`               | Verifies the user registration details and completes the registration process. |
| 5   | `GET /v1/auth/user/:user_id`                  | Retrieves user metadata based on the specified user_id.                        |
| 6   | `PUT /v1/auth/user/:user_id`                  | Updates user metadata based on the specified user_id.                          |
| 7   | `GET /v1/auth/logout`                         | Performs user logout.                                                          |
| 8   | `GET /v1/dealership/:id`                      | Retrieves dealership information based on the specified id.                    |
| 9   | `GET /v1/dealership/user/:id`                 | Retrieves dealership information associated with the specified user_id.        |
| 10  | `PUT /v1/dealership/:id`                      | Updates dealership information based on the specified id.                      |
| 11  | `DELETE /v1/dealership/:id`                   | Soft deletes a dealership.                                                     |
| 12  | `POST /v1/auth/login`                         | Initiates the login process.                                                   |
| 13  | `POST /v1/auth/login/verify`                  | Verifies the login credentials and completes the login process.                |
| 14  | `POST /v1/auth/register`                      | Initiates the user registration process.                                       |
| 15  | `POST /v1/auth/register/verify`               | Verifies the user registration details and completes the registration process. |
| 16  | `GET /v1/auth/user/:user_id`                  | Retrieves user metadata based on the specified user_id.                        |
| 17  | `PUT /v1/auth/user/:user_id`                  | Updates user metadata based on the specified user_id.                          |
| 18  | `PUT /v1/auth/user/:user_id/dealership`       | Associates a user with a dealership.                                           |
| 19  | `POST /v1/auth/refresh`                       | Refreshes the authentication token.                                            |
| 20  | `GET /v1/auth/logout`                         | Performs user logout.                                                          |
| 21  | `GET /v1/dealership/:id`                      | Retrieves dealership information based on the specified id.                    |
| 22  | `GET /v1/dealership/user/:id`                 | Retrieves dealership information associated with the specified user_id.        |
| 23  | `PUT /v1/dealership/:id`                      | Updates dealership information based on the specified id.                      |
| 24  | `POST /v1/dealership/`                        | Creates a new dealership.                                                      |
| 25  | `DELETE /v1/dealership/:id`                   | Soft deletes a dealership.                                                     |
| 26  | `GET /v1/cars/`                               | Retrieves a list of cars.                                                      |
| 27  | `POST /v1/cars/`                              | Creates a new car.                                                             |
| 28  | `PUT /v1/cars/:car_id`                        | Updates car information based on the specified car_id.                         |
| 29  | `GET /v1/cars/:car_id`                        | Retrieves car information based on the specified car_id.                       |
| 30  | `DELETE /v1/cars/:car_id`                     | Deletes a car.                                                                 |
| 31  | `GET /v1/cars/search`                         | Searches for cars based on specific criteria.                                  |
| 32  | `GET /v1/cars/filter`                         | Filters cars based on specific criteria.                                       |
| 33  | `POST /v1/cars/brand`                         | Creates a new car brand.                                                       |
| 34  | `GET /v1/cars/brand`                          | Retrieves a list of car brands.                                                |
| 35  | `PUT /v1/cars/brand/:brand_id`                | Updates car brand information based on the specified brand_id.                 |
| 36  | `GET /v1/cars/brand/:brand_id`                | Retrieves car brand information based on the specified brand_id.               |
| 37  | `DELETE /v1/cars/brand/:brand_id`             | Deletes a car brand.                                                           |
| 38  | `POST /v1/cars/:car_id/extra`                 | Adds extra information to a car.                                               |
| 39  | `GET /v1/cars/:car_id/extra`                  | Retrieves extra information associated with the specified car_id.              |
| 40  | `PUT /v1/cars/extra/:extra_id`                | Updates extra information based on the specified extra_id.                     |
| 41  | `DELETE /v1/cars/extra/:extra_id`             | Deletes extra information based on the specified extra_id.                     |
| 42  | `POST /v1/cars/:car_id/image`                 | Uploads an image for a car.                                                    |
| 43  | `GET /v1/cars/:car_id/image`                  | Retrieves images associated with the specified car_id.                         |
| 44  | `PUT /v1/cars/image/:image_id`                | Updates image information based on the specified image_id.                     |
| 45  | `DELETE /v1/cars/image/:image_id`             | Deletes an image based on the specified image_id.                              |
| 46  | `GET /v1/cars/by/dealership/:dealership_id`   | Retrieves cars associated with the specified dealership_id.                    |
| 47  | `GET /v1/cars/by/brand/:brand_id`             | Retrieves cars associated with the specified brand_id.                         |
| 48  | `GET /v1/cars/by/dealer/:dealer_id`           | Retrieves cars associated with the specified dealer_id.                        |
| 49  | `GET /v1/cars/by/dealer/count`                | Retrieves the count of cars associated with each dealer.                       |
| 50  | `GET /v1/cars/by/dealership/count`            | Retrieves the count of cars associated with each dealership.                   |
| 51  | `GET /v1/cars/by/brand/count`                 | Retrieves the count of cars associated with each brand.                        |
| 52  | `GET /v1/spares/`                             | Retrieves a list of spares.                                                    |
| 53  | `POST /v1/spares/`                            | Creates a new spare.                                                           |
| 54  | `PUT /v1/spares/:spare_id`                    | Updates spare information based on the specified spare_id.                     |
| 55  | `GET /v1/spares/:spare_id`                    | Retrieves spare information based on the specified spare_id.                   |
| 56  | `DELETE /v1/spares/:spare_id`                 | Deletes a spare.                                                               |
| 57  | `GET /v1/spares/search`                       | Searches for spares based on specific criteria.                                |
| 58  | `GET /v1/spares/filter`                       | Filters spares based on specific criteria.                                     |
| 59  | `POST /v1/spares/image`                       | Uploads an image for a spare.                                                  |
| 60  | `GET /v1/spares/image`                        | Retrieves images associated with spares.                                       |
| 61  | `GET /v1/spares/:spare_id/image`              | Retrieves images associated with the specified spare_id.                       |
| 62  | `PUT /v1/spares/image/:image_id`              | Updates image information based on the specified image_id.                     |
| 63  | `DELETE /v1/spares/image/:image_id`           | Deletes an image based on the specified image_id.                              |
| 64  | `GET /v1/spares/by/brand/:brand_id`           | Retrieves spares associated with the specified brand_id.                       |
| 65  | `GET /v1/spares/by/category/:category_id`     | Retrieves spares associated with the specified category_id.                    |
| 66  | `GET /v1/spares/by/car/:car_id`               | Retrieves spares associated with the specified car_id.                         |
| 67  | `GET /v1/spares/by/dealer/:dealer_id`         | Retrieves spares associated with the specified dealer_id.                      |
| 68  | `GET /v1/spares/by/dealership/:dealership_id` | Retrieves spares associated with the specified dealership_id.                  |


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
