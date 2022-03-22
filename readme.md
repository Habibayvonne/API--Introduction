# Web API

A web API is an application that runs on a server and offers resources e.g a database to the clients.

We will create a simple API and access it from the browser.

---

## Lesson 2 : MVC

There is a design paradigm called MVC:
M : Models
V : Views
C : Controller

In Go, we can create packages for each of them.

#### Models

Models are basically functions that deal with exchange of data. Let's say your API reads and writes data in the database, models will be functions that are responsible for this.  They can also be used to exchange data between your API and any other external resource which can be a database, cache, another API etc.

#### View

View are basically routes. These are endpoints that you give the client. So from the client's perspective, they can go to a certain route (view) to view data.

#### Controllers

These are functions that sit between views and models. When a request comes to a view (route), a controller is called. The controller then calls the relevant model(s), and then it processes the output from the model and then sends the response to the client.

If you are writing in a language that has pointers, the controllers can have direct access to resource connections e.g the database and then share the resources with the models as needed.



## Project Directories

Note that : this is a personal preference. You can explore other options online.

- controllers : will have controllers and views (routes) and also connections to outside resources e.g database

- models : functions that deal with data

- entities : structs to define the structure of data that is expected from and by the clients and also any struct you need e.g for models.

- utils : these will be utility functions - functions that can be used anywhere e.g a function that configure our logger, encrypt

---

## Lesson 3 : Middleware/ Gorrila Library

Gorrila is just a collection of libraries that can be used in web APIs. We'll start with mux which helps in creating routers (views)

With mux, you can specify the HTTP method (e.g GET, POST, PUT, PATCH, DELETE ) etc
e.g

```go
var router = mux.NewRouter()
router.HandleFunc("/users", controllers.UsersController).Methods("PUT", "POST")
```

This means, `/users` endpoint will only accept PUT and POST methods.

In the controller, if you want to know the method, you can use r.Method from the *http.Request



### Middleware

A middleware is a middleman between the router and the controller. For example, controller1 returns data as an array [1,2,3] while controller2 returns data as a number e.d 4. However, you want to maintain the data structure between you and the client so that controller1 returns 

```json
{
    "data": [1,2,3]
}
```

and controller2 returns

```json
{
    "data": 4
}
```

Instead of writing the logic for conversion in each controller, you can just create a middleware.

Another common use case for middlewares is authenticating requests. Therefore, before a request reaches a controller, it is authenticated first.

In the first use of middlware (controller -> middleware -> router), we just define the middleware as a function and call it from the controller.

In the second use of middleware (router -> middleware -> controller), we'll need to set the middleware on the router.

---

## Lesson 4 : Middleare 2, Connect to a Database