# go-currency-converter

# Requirements

Docker environment must be setup and running

# Running the application

In order to run the app, execute

`docker-compose up`

# Debugging the application

Debugging the application can be done by running delve server through
`docker-compose -f docker-compose.debug.yml up --build`
And then attaching your IDE using port 40000.