# Learning go

## First steps

`First steps` folder includes 6 small scripts I managed using Golang

- even
- geometry
- max
- multithread
- strings
- sum

To run a single script go to `first_steps` folder and run `go run even.go` in example

### Even

Write a function which takes an integer and halves it and returns true if it was even or false if it was odd.
For example half(1) should return (0, false) and half(2) should return (1, true).

### Geometry

Just a plain code to measure perim and area of `rect` and `circle` by using `interface` and `structs`

### Max

Write a function with one variadic parameter that finds the greatest number in a list of numbers.

### Multithread

Just a plain code to print `ping pong` string by using `channels`

### Strings

There are most common `string` fuctions you might use

### Sum

Sum is a function which takes a slice of numbers and adds them together.
What would its function signature look like in Go?

# Rest API

 We are building a simple REST-API with Golang a router package called Mux. The main resource we're going to work with is `post`
 
 Go to `rest_api` folder, run `go install`, then run `go run main.go` and use the Postman for testing purposes.
 
*Available endpoints*

## Create a post

*POST* `http:localhost:8000/posts`

```
{
  "title": "yours title",
  "body": "yours body text"
}
```

## Get all posts

*GET* `http:localhost:8000/posts`

## Get a single post

*GET* `http:localhost:8000/posts/1`

## Update a single post

*PUT* `http:localhost:8000/posts/1`

```
{
  "title": "new awesome title",
  "body": "new awesome body"
}
```

## Delete a single post

*DELETE* `http:localhost:8000/posts/1`

