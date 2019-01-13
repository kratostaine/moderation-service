# Moderation Service

Moderation service is a REST API which provides the functionality of moderating any text to find out
if it contains any objectionable content or not.

It is an API written in **Go**.

## Currently available endpoints
**GET** */health* :- Provides the health of the service, which can be used to check for its heart-beat.

**GET** */validation/{text}* :- Validates if the input text is objectionable or not.

## To Do
- Add a database to store data for objectionable content.
- Add endpoints to get, create, update and delete objectionable data set so that the same can be updated dynamically.
- Add tests
- Add swagger integration
