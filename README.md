# The Task
Create a production ready web service which combines two existing web services.
Fetch a random name from https://names.mcquay.me/api/v0
Fetch a random Chuck Norris joke from http://api.icndb.com/jokes/random?firstName=John&lastName=Doe&limitTo=nerdy
Combine the results and return them to the user.

# Getting Started

## Clone the repository:
```shell
$ git clone https://github.com/wildsurfer/task
```
## Build Docker image:
```shell
$ cd task
$ docker build -t task .
```
## Run:
```shell
$ docker run -it --rm -p 5000:5000 kek .
$ curl "http://localhost:5000"
```
## Try:
```shell
$ docker run -it --rm -p 5000:5000 kek .
```
