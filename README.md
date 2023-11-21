# go-promises
Implementation of JavaScript Promises in Golang

# Install
````sh
go get -u github.com/agustin-del-pino/go-promises
```` 

## `gopromise.Promise[T]`
It's an interface that provides the `Await() (T, error)`. *This emulates the await operator*

## `gopromise.New(f func() (T, error)) Promise[T]`
It's a function that takes a function an returns in form of Promise. *This emulates the `new Promise(...)`*

## `gopromise.All(p []Promise[T]) Promise[[]T]`
It's a function that takes an **array of promises** and returns a new promise that waits for all promises get done. 
In case one of them got an error, the main promises will forward that error. *This emulates the `Promise.all`*

## `gopromise.AllSettled(p []Promise[T]) Promise[[]Result[T]]`
It's a function that takes an **array of promises** and returns a new promise that waits for all promises get done. 
The result of each promise will be wrapped in a `Result` instance. *This emulates the `Promise.allSettled`*
