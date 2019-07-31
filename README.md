[![Go Report Card](https://goreportcard.com/badge/github.com/ernst01/go-start-http-api)](https://goreportcard.com/report/github.com/ernst01/go-start-http-api)

# go-start-http-api

Boilerplates for getting started with an HTTP Service/API in Go.

## Project Layout

This project follows the Standard Go Project layout from [golang-standards/project-layout](https://github.com/golang-standards/project-layout)

## HTTP Service Design

When I first started looking into creating cleaner HTTP services I stumbled upon Mat Ryer's [How I write Go HTTP services after seven years](https://medium.com/statuscode/how-i-write-go-http-services-after-seven-years-37c208122831). I found his blog post inspiring and decided the emulate most of his recommendations.

## Dockerfile

The Dockerfile has 3 steps: The first one builds your application, the second one is meant for testing, and the final step produces a small image from `scratch`. As best practice the final image runs as non-root.  

## Makefile

A Makefile contains a set of directives that helps developers use and understand your application. 

```
> make help

Choose a command to run in go-start-http-api:

  run       Runs your application
  install   Installs your dependencies
  build     Builds your application
  test      Runs your tests if any
  cover     Checks your code coverage
```

## HTTP Responses

I wrote a simple package to help with sending error and success responses. [Ckeck it out!](https://github.com/ernst01/common/tree/master/pkg/response)

## Feedback

I've been using a similar code structure for the past 2 years and wanted to share it with the community. Feedback is more than welcome :) Feel free to create Github Issues and Pull requests as you see fit.

