# Todoapp

This is a full-stack TodoApp built using Go, the Gin framework, Vue.js with TypeScript, Docker, Swagger, and Cypress for testing.

## Table of Contents

- [Introduction](#Introduction)
- [Features](#features)
- [Installation](#installation)
- [API Documentation](#api-documentation)
- [Testing](#testing)

## Introduction

The Fullstack TodoApp is a web application that allows users to create, manage, and track their tasks. The backend is built using Go with the Gin framework, which provides a performant and lightweight web server. The frontend is developed using Vue.js with TypeScript, providing a seamless and interactive user experience. Docker is used to containerize the application for easy deployment, and Swagger is integrated to provide API documentation. Cypress is used for end-to-end testing to ensure the app's functionality and stability.

## Features

- Create, update, and delete tasks
- Mark tasks as completed
- List all tasks
- API documentation using Swagger
- End-to-end testing using Cypress

## Installation

Make sure you have the following tools installed before proceeding:

- [Go](https://golang.org/)
- [Node.js](https://nodejs.org/)
- [Docker](https://www.docker.com/)
- [Docker Compose](https://docs.docker.com/compose/install/)
- [Cypress](https://www.cypress.io/)

To quickly set up and run the Todoapp, follow these steps:

1. Clone this repository to your local machine:

```bash
git clone https://github.com/codescalersinternships/todoapp-omar
cd todoapp-omar
```

2. Start the Todoapp using Docker Compose:

```bash
docker-compose up -d
```

## API Documentation

The Fullstack TodoApp provides a RESTful API for performing CRUD operations (Create, Read, Update, and Delete) on tasks. Below are the endpoints available in the API:

- Create a Task
Request
```bash
POST /task
Content-Type: application/json

{
  "title": "Task Title",
}
```
Response
```bash
HTTP/1.1 201 Created
Content-Type: application/json

{
  "id": 1,
  "title": "Task Title",
  "is_completed": false
}
```

- Get All Tasks
Request
```bash
GET /task
```
Response
```bash
HTTP/1.1 200 OK
Content-Type: application/json

[
  {
    "id": 1,
    "title": "Task 1",
    "is_completed": false
  },
  {
    "id": 2,
    "title": "Task 2",
    "is_completed": true
  },
  // More tasks...
]
```

- Update a Task
Request
```bash
PUT /task/:id
Content-Type: application/json

{
  "title": "Updated Task Title",
  "is_completed": true
}
```
Response
```bash
HTTP/1.1 200 OK
Content-Type: application/json

{
  "id": 1,
  "title": "Updated Task Title",
  "is_completed": true
}
```

- Delete a Task
Request
```bash
DELETE /api/tasks/:id
```
Response
```bash
HTTP/1.1 200
```

For detailed information about the API endpoints and request/response parameters, you can access the Swagger documentation at [todo swagger docs](http://localhost:8080/docs/index.html) after starting the application. The Swagger UI provides an interactive interface to explore the API and test the endpoints directly from the browser.

## Testing

First install all dependencies and you can test either frontend or backend:

- backend
```bash
cd backend
go mod download
go test ./...
```

- frontend
```bash
cd frontend
npm install
npm run test
```