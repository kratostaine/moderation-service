# Moderation Service

[![Build Status](https://travis-ci.org/kratostaine/moderation-service.svg?branch=master)](https://travis-ci.org/kratostaine/moderation-service)

Moderation service is a REST API which provides the functionality of moderating any text to find out
if it contains any objectionable content or not.

It is an API written in **Go**.

## To run
In the root of the repo folder, please run the below commands:

1.  To get the dependencies:

        go get -t -v ./...
    
2. To build:

        go build
        
3. To run the app:

        ./moderation-service
        
4. To run tests:

        go test
        
The service would run on port `9000` currently, and the endpoints can be accessed locally on http://localhost:9000/

## Currently available endpoints
**GET** */health* :- Provides the health of the service, which can be used to check for its heart-beat.

**GET** */validation/{text}* :- Validates if the input text is objectionable or not.

## To Do
- Add a database to store data for objectionable content.
- Add endpoints to get, create, update and delete objectionable data set so that the same can be updated dynamically.
- Refactor swagger integration

