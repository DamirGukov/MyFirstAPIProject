First API Project
====================

Running the Project
-------------------

To run the project, you can use the following command:
```
go run cmd/main.go
```
The application will start on port 8000.

Usage
-----

The API provides the following resources:

* `GET /items` - get a list of all items
* `POST /items` - create a new item
* `GET /items/{id}` - get an item by its ID
* `PUT /items/{id}` - update an item by its ID
* `DELETE /items/{id}` - delete an item by its ID

Documentation
-------------

The API documentation is created using Swagger and is located in the `/docs` folder. You can open it in your browser at `http://localhost:8000/swagger/index.html` after running the application.

Makefile
--------

The project includes a Makefile that contains several targets to automate some common tasks related to building, running, and deploying the application using Docker and Swagger. Here's a brief explanation of each target:

* `docker-run`: This target runs the `docker run` command with the `-p` flag to map the container's port 8000 to the host's port 8000. It depends on the `docker-build` target, which means that the `docker-build` target will be executed before this target.
* `docker-build`: This target builds a Docker image using the `docker build` command and tags it with the name `go-app`.
* `swag-init`: This target initializes Swagger documentation using the `swag init` command. It generates a `docs` folder with the Swagger documentation for the API defined in the `cmd/main.go` file.
* `run`: This target runs the Go application using the `go run` command.

