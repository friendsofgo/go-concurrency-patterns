# Go Concurrency Patterns

Implementation examples for the most common concurrency patterns on Go.

## Go Concurrency Patterns video from Rob Pike:

https://www.youtube.com/watch?v=f6kdp27TYZs

## Patterns

| Name                                                           | Description                                                                     | Playground                                   |
|----------------------------------------------------------------|---------------------------------------------------------------------------------|----------------------------------------------|
| [Boring](/boring/boring.go)                                    | A simple hello world with goroutine                                             | [run](https://play.golang.org/p/8fyYDEqfgqf) | 
| [Fan in](/fan-in/fan-in.go)                                    | Fan in pattern, join all the results from a datasource                          | [run](https://play.golang.org/p/K2_nv3Kyahn) | 
| [Fan out](/fan-out/fan-out.go)                                 | Fan out pattern, small jobs on goroutines to do the work concurrently           | [run](https://play.golang.org/p/8iwjE9azXF6) |
| [Fan out semaphore](/fan-out-semaphore/fan-out-semaphore.go)   | Fan out pattern as well, but with a semaphore to limit concurrency              | [run](https://go.dev/play/p/zTZ4VIDsN4k)     |
| [Drop](/drop/drop.go)                                          | Drop pattern, useful pattern to drop request when heavy loads happen            | [run](https://play.golang.org/p/wwmN-9bwF9M) | 
| [Retry Timeout](/retry-timeout/retry-timeout.go)               | Retry timeout pattern, used to manage failures when reaching resources          | [run](https://play.golang.org/p/NgbgDUFUvnY) |
| [Worker Pool](/worker-pool/worker-pool.go)                     | Worker pool pattern, used to control a limited set of workers                   | [run](https://go.dev/play/p/EJTzsNdVd_f) |