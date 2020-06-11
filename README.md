# Weather API

Wheater API is a coding challenge.
It has the sole purpose of testing my current Go coding abilities, as well as learning by practice and grow as a developer.

## Versions

It consists of 5 stages or "versions". Each version starts from the previous one and has a higher complexity factor.

## Installation

The following command installs all the dependencies recursively.

```bash
go get ./...
```

## Usage

Versions 1 and 2:

```bash
cd <version name>
bee run
```

Versions >= 3:

```bash
cd <version name>
docker-compose up
```

## Requirements for each version

### General goal

- Create a Weather API using beego.

### Version 1

- City is a string. Example: Bogotá ✔️
- Country is a country code of two characters in lowercase. Example: co ✔️
- This endpoint should use an external API to get the proper info ✔️
- The data must be human readable ✔️
- Use environment variables for configuration ✔️
- Log errors to terminal ✔️
- The response must include the content type header (application/json) ✔️
- Functions must be tested using Convey ✔️

### Version 2

- Save each request in a database with a timestamp ✔️  
  Load the information from the persistence layer when it’s available. ✔️
- Store new values when the timestamp difference is higher than 300 seconds ✔️
- Dependency Injection (logger) ✔️
- Include fixtures ✔️
- Remove Convey library ✔️
- Add unit tests ✔️

### Version 3

- New endpoint that allows to add a scheduled job that will perform an hourly check of a city and persist it ✔️
  ```PUT /scheduler/weather
  Payload: {“city”: $City, “country”: $Country}
  Response: 202
- Add integration tests. ✔️
- Add a swagger description for the API. ✔️
- Add Dockerfile and docker-compose files. ✔️

### Version 4

- Add a second weather provider modifying the less code possible. At the beginning of the program we can choose which provider use based on a configuration variable. ✔️
- This provider is going to read the weather information from json files. Add some files to support at least two cities.
  Add integration tests. ✔️

### Version 5

- Add a pool of 5 workers (go routines) to perform the requests to the external service. If all the workers are busy the originating request should be blocked until one worker is available.

## License

[MIT](https://choosealicense.com/licenses/mit/)
