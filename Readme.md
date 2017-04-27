Y-IndexOf : Easy and modular indexOf in go
==========================================

[![Build Status](https://travis-ci.org/yazgazan/y-indexof.svg?branch=master)](https://travis-ci.org/yazgazan/y-indexof)
[![Go Report Card](https://goreportcard.com/badge/github.com/yazgazan/y-indexof)](https://goreportcard.com/report/github.com/yazgazan/y-indexof)
[![Go version](https://img.shields.io/badge/go-1.6%2B-brightgreen.svg)](https://github.com/yazgazan/jaydiff)

## Description

Y-IndexOf is an easy to use and modular indexOf server.

## Installing

Installing Y-IndexOf is as simple as typing `go get -u github.com/yazgazan/y-indexof`.
It is possible that go yells at you saying you are missing `magic.h`.
In such case, install the libmagic-dev package from your linux distribution repositories.

## Using

Start by initializing an index-of directory :

```bash
mkdir indexof && cd indexof
y-indexof init
```

Edit the config file `indexof.toml` and start the server :

```bash
y-indexof start
```

## Requirements

- [Go](http://golang.org/)
- libmagic-dev (from your linux distrib repositories)

## Compatibility

It has not been tested on Windows and BSD.

## License

Beerware (from [http://people.freebsd.org/~phk/](http://people.freebsd.org/~phk/))

