# go-currency-converter

# Requirements

Docker environment must be setup and running

# Capabilities
The application can be used to convert amounts from EUR to USD.
Authentication is required for conversion functionality.
An API key can be retrieved a new project.

# Running the application

In order to run the app, execute

`docker-compose up`

Currency conversion API requests require fixer_access_key to be specified in config/config.yaml file.
This property needs to be set to a valid fixer API key.

This application sets up an ongoing poller to update the conversion rates hourly.

# Accessing Swagger documentation
Swagger docs are located in docs/swagger.json and can be accessed through http://localhost:17249/swagger/index.html
when running the application locally

# Outstanding changes
- Unit and integration tests. The application is largely test-friendly with DI practices used throughout. Database calls can be mocked and tested using sqlmock
- Keeping an API key unencrypted in a yaml file is not secure. A secrets vault could be used instead.
- EUR currency string specified in a lot of places and needs to be centralized for ease of maintenance.
- Setup from_currency + to_currency composite key for conversion_rates table
- Some classes and methods are missing conventional comments

# Debugging the application

Debugging the application can be done by running delve server through
`docker-compose -f docker-compose.debug.yml up --build`
And then attaching your IDE using port 40000.