A Technical test  project 

As a free test,  i decide to create a project about birds, inspired on our recently build 
bird feeder 


Local Install:

Prepare / Postgres Database and Data for test

* docker run  --name postgres-db -p 5432:5432 -e POSTGRES_PASSWORD=onlydemopassword  -d postgres
* cat dump_bird_db.sql | docker exec -i postgres-db psql -U postgres

Testing 

Use any web client like Postman or the 2 methods get directly on the browser

###
List of Birds

GET http://localhost:8080/api/v1/list

###

Find One Bird By Id
GET http://localhost:8080/api/v1/byid?id=1


###
Add One Bird 

POST http://localhost:8080/api/v1/add
Content-Type: application/json

{
"specie": "Petirrojo Test 2",
"name": "Petirrojo",
"characteristics": "Beatiful Bird from north america"
}


###
PATCH http://localhost:8080/api/v1/update
Content-Type: application/json

{
"id":  100,
"specie": "PetiAvi- 2  ",
"name": "PetiAzul",
"characteristics": "Beaituful bird from north america"
}

###
List of Birds 

GET http://localhost:8080/api/v1/list

###
Delete One Bird by Id

DELETE http://localhost:8080/api/v1/delete/byid?id=23
###
