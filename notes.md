## `curl localhost:6000 | jq`

This will give you output in nice JSON format

<hr>

## `curl localhost:6000 -XDELETE -v`

This is give a DELETE request to server

<hr>

## `curl localhost:6000 -d '{}'`

This is for POST request to the server.

<hr>

## `$ curl -v localhost:6000/1 -XPUT`

This is for PUT method with id 1

<hr>

## `curl -v localhost:6000/1 -XPUT -d '{"name": "samosa", "description" : "spicy"}'`

This is used to PUT and update the data for the id 1