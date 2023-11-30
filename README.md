# Short URL
A fully functional URL shortener using MySQL as the storage server and Redis as the cache server.

The current components include:

+ A web server for URL shortening and redirection
+ [TODO] logging
+ [TODO] connection pooling
+ [TODO] Redis cache server
+ [TODO] a rate limiter middleware
+ [TODO] logging

## Usages

We provide two API endpoints: one for URL shortening and another for URL redirection.

### URL shortening

`POST /api/v1/data/shorten`

+ Request parameter: {longURL: string}
+ Rreturn: shortURL

Example usage:

```bash
$ curl -iX POST 'http://localhost:8080/api/v1/data/shorten?longURL=www.google.com'
```

### URL redirection 

`GET /api/v1/:shortURL`

+ Return: longURL for HTTP redirection

Example usage:

`shortURL`: The URL you get from the previous step, or any short URL in the database

```bash
$ curl -i 'http://localhost:8080/api/v1/shortURL'
```

## Command Runner

We use the command runner [just](https://github.com/casey/just) as an alternative to makefile. Execute `just help` to access a complete list of available commands.

## Build and Run Locally

We use docker-compose to set up local development pretty easily. It exposes a web server forwarded to host port 8080, a MySQL server forwarded to host port 3306, and a Redis server forwarded to host port 6379. To launch, simply run:

```bash
$ docker-compose up -d
```
Then you are ready to go.


To connect to the mysql server, run:

```bash
$ just mysql
```

To connect to the redis server, run:

```bash
$ just redis
```

## License

[MIT License](./LICENSE)
