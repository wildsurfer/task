# The Task
Create a production ready web service which combines two existing web services.
Fetch a random name from https://names.mcquay.me/api/v0
Fetch a random Chuck Norris joke from http://api.icndb.com/jokes/random?firstName=John&lastName=Doe&limitTo=nerdy
Combine the results and return them to the user.

# The Solution

## Project structure
The project structure is taken from the [set of common historical and emerging project layout patterns in the Go ecosystem](https://github.com/golang-standards/project-layout). It may look ambiguous for such a small project at it really worth it for bigger projects.  

## Elapsed time
It took me approx. **_one_** hour to build the first working version and **_three_** more hours to refactor the code. 

## Known issues
There are some unhandled errors in deferred calls.

## Tests
To prevent over-engineering only one important function is tested. This is common for real-world programming. Also, note that tests run every time Docker image is being built.  

## Clone the repository:
```shell
$ git clone https://github.com/wildsurfer/task
```
## Build Docker image:
Docker image is built from the `scratch` image and is extremely small (~9mb) which is good for CI/CD pipelines.
```shell
$ cd task
$ docker build -f ./build/package/Dockerfile -t task .
```
## Run:
```shell
$ docker run -it --rm -p 5000:5000 task .
```
## Try:
```shell
$ docker run -it --rm -p 5000:5000 kek .
$ curl "http://localhost:5000"
```
## Try concurrent requests:
```shell
$ ab -n 1000 -c 100  "http://localhost:5001/"
```
