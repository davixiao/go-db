# go-db
In-memory key-value database in Golang that supports Append Only File (AOF) persistence, Redis serialization protocol specification (RESP), concurrency, and storage of strings and hashes

## getting-started

A guide for getting go-db up and running! 

### db

First, ensure dependencies are up to date

```go
go get .
```

To start the database server, enter the `/db` directory and run the module with

```go
go run .
```

The server should be up and running. Expect to see

```
Listening on port :6379
```

### example-client

`/example-client` is an example of how to interact with the database server. You are free to use any Redis client, but this is an example of using the new go-redis client.

First, ensure dependencies are up to date

```go
go get .
```

To start the database server, enter the `/example-client` directory and run the module with

```go
go run .
```

You should see the following information on the client side

```
key david
key2 does not exist
```

that tests set and get.

## encountered-errors

Sometimes you may encounter the following error on your database server

```
read tcp [::1]:6379->[::1]:51903: wsarecv: An existing connection was forcibly closed by the remote host.
```

It usually means that the server sent data to the client, but the connection has already been closed. This usually happens when the client closes the connection.
