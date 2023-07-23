# Music Store
Developing a RESTful API with Go and Gin  
This project implements a basic CRUD for working with a PostgreSQL database using Docker

You can run it using the following commands:
- To run server `go run .`
- To run tests `go test -v`
- To run DB `docker run -it --name some-postgres -e POSTGRES_PASSWORD=pass -e POSTGRES_USER=user -e POSTGRES_DB=db postgres`