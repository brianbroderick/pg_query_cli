This is a cli wrapper for portions of the [pg_query_go](https://github.com/lfittl/pg_query_go) Go Package. 

Assuming that your Go workspace is set up, run `go build`. If not, first visit [Golang](https://golang.org/) to set it up.
Building the file will take a couple of minutes because it's compiling parts of the PG lib.  
Once it's done, you can run it like this:

`./pg_query_cli "SELECT character FROM cartoons WHERE phrase = 'zoinks' OR phrase = 'ruh roh' LIMIT 10"`

It will return the normalized version. i.e. values will be converted to question marks like this:

`SELECT character FROM cartoons WHERE phrase = ? OR phrase = ? LIMIT ?`

This is to help when compiling query stats like long running queries, frequently used queries, etc.
