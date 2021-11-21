#Go Concurrency Patterns

Implementation examples for the most common concurrency patterns on Go.

## Go Concurrency Patterns video from Rob Pike:
https://www.youtube.com/watch?v=f6kdp27TYZs

## Patterns

| Name                                                  | Description                                                           | Playground                                   |
|-------------------------------------------------------|-----------------------------------------------------------------------|----------------------------------------------|
| [boring](/boring/boring.go)                           | A simple hello world with goroutine                                   | [run](https://play.golang.org/p/8fyYDEqfgqf) | 
| [Fan in](/fan-in/fan-in.go)                           | Fan in pattern, join all the results from a datasource                | [run](https://play.golang.org/p/K2_nv3Kyahn) | 
| [Fan out](/fan-out/fan-out.go)                        | Fan out pattern, small jobs on gorutines to do the work concurrently  | [run](https://play.golang.org/p/8iwjE9azXF6) | 
| [Drop](/drop/drop.go)                                 | Drop pattern, useful pattern to drop request when a heavy loads happen | [run](https://play.golang.org/p/wwmN-9bwF9M) | 
