# Flight Aggregator

The goal of the api is to get every flight to a destination and sort it by : 
    - price,
    - departure date
    - travel time
  
# exercices : 
- create a server
- set 2 routes:
  - GET /health : to verify the healthiness of the server 
    - set the status response to 200 : w.WriteHeader(http.StatusCreated)
  - GET /flight.
- try to get the data of both apis from the server (client requests).
  - transform the data into structs
  - and organize the code to process the data in 2 repositories and extract the data using the same interface.
- return the flights orders by price
- now you want to sort by price, time_travel or departure_date :
  - pass this information by query/params or body,
  - create the algorithms,
  - verify the output
- Create tests for :
  - your sorting algorithms,
  - your flight service :
    - mock the repositories to make the tests.

# to help you

## pre-setup

  - install Docker Compose and start the project with: docker compose up
  - air is setup to auto reload the project on every modification !
  - a make file is here to run the tests with gotestsum :
    - install it with : `go install gotest.tools/gotestsum@latest`

## Run the base project: 
- `docker compose build`
- `docker compose up`

## test 
- `make test`

## access the apis : 
- j-server1 :
  - docker : http://j-server1:4001
  - localhost : http://localhost:4001
- j-server2 : 
  - docker : http://j-server2:4001
  - localhost : http://localhost:4001


startup : 
- use the Viper library [Link Text](https://github.com/spf13/viper),
- get every env variables with : viper.AutomaticEnv() 
- then select with : viper.Get("MY_VAR")

tests : 
- use Testify : [Link Text](https://github.com/stretchr/testify)