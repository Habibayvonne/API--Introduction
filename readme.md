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