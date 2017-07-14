## pg_query_cli

This is a cli wrapper for the normalization function of the excellent [pg_query_go](https://github.com/brianbroderick/pg_query_cli/tree/master/vendor/github.com/lfittl/pg_query_go) Go package. 

Assuming that your Go workspace is set up, run `go build`. If not, first visit [Golang](https://golang.org/) to set it up.
Building the file may take a couple of minutes because it's compiling parts of the PG lib. Once it's done, you can run it like this:

`./pg_query_cli "SELECT character FROM cartoons WHERE phrase = 'zoinks' OR phrase = 'ruh roh' LIMIT 10"`

It will return the normalized version. i.e. values will be converted to question marks like this:

`SELECT character FROM cartoons WHERE phrase = ? OR phrase = ? LIMIT ?`

This is helpful when aggregating query stats like long running queries, frequently used queries, etc.  It's important to note that it uses the actual PostgreSQL server source to parse SQL queries and return the internal PostgreSQL parse tree. 

FYI, there's an MRI specific Ruby gem at [pg_query](https://github.com/lfittl/pg_query).

## Authors

- [Brian Broderick](https://github.com/brianbroderick)


## License

Copyright (c) 2017, [Brian Broderick](https://github.com/brianbroderick)
pg_query_go package used by pg_query_cli is Copyright (c) 2015, Lukas Fittl <lukas@fittl.com>
pg_query_cli is licensed under the 3-clause BSD license, see LICENSE file for details.

This project includes code derived from the [PostgreSQL project](http://www.postgresql.org/),
see LICENSE.POSTGRESQL for details.
