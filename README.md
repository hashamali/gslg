# gslg: gRPC Interceptor for gsl
[![godoc](https://godoc.org/github.com/hashamali/gslg?status.svg)](http://godoc.org/github.com/hashamali/gslg)
[![sec](https://img.shields.io/github/workflow/status/hashamali/gslg/security?label=security&style=flat-square)](https://github.com/hashamali/gslg/actions?query=workflow%3Asecurity)
[![go-report](https://goreportcard.com/badge/github.com/hashamali/gslg)](https://goreportcard.com/report/github.com/hashamali/gslg)
[![license](https://badgen.net/github/license/hashamali/gslg)](https://opensource.org/licenses/MIT)

A [gsl](https://github.com/hashamali/gsl) gRPC interceptor.

## API

* `Interceptor(logger gsl.Log)`: Creates a new gRPC interceptor with the given logger.