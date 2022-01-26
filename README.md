# Go Todo REST API Example
A RESTful API example for simple application with Go

It is a just simple tutorial or example for making simple RESTful API with Go using **gorilla/mux** (A nice mux library) and **gorm** (An ORM for Go)

## Installation & Run

Before running API server, you should set the database config with yours or set the your database config with my values on [config.go](https://github.com/charlitoro/go-db-connection-api/blob/master/config/config.go)
```go
func GetConfig() *Config {
	return &Config{
		DB: &DBConfig{
			Dialect:  "mysql",
			Username: "admin",
			Password: "PWD00@",
			Name:     "todoapp",
			Charset:  "utf8",
		},
	}
}
```

```bash
# Build and Run
go run .

# Or
go build .

# API Endpoint : http://localhost:3000
```


## API

#### /projects
* `GET` : Get all projects
* `POST` : Create a new project

#### /projects/:title
* `GET` : Get a project
* `PUT` : Update a project
* `DELETE` : Delete a project

#### /projects/:title/archive
* `PUT` : Archive a project
* `DELETE` : Restore a project 

#### /projects/:title/tasks
* `GET` : Get all tasks of a project
* `POST` : Create a new task in a project

#### /projects/:title/tasks/:id
* `GET` : Get a task of a project
* `PUT` : Update a task of a project
* `DELETE` : Delete a task of a project

#### /projects/:title/tasks/:id/complete
* `PUT` : Complete a task of a project
* `DELETE` : Undo a task of a project

