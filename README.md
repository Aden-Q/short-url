# Short URL
A fully functional URL shortener using MySQL as the storage server and Redis as the cache server.

The current components include:

+ A web server for URL shortening and redirection
+ [TODO] Redis cache server
+ [TODO] generate doc with swagger
+ [TODO] logging
+ [TODO] connection pooling
+ [TODO] integration test running on CI
+ [TODO] a rate limiter middleware

## Usages

We provide two API endpoints: one for URL shortening and another for URL redirection.

### URL shortening

`POST /api/v1/data/shorten`

+ Request parameter: {longURL: string}
+ Rreturn: shortURL

Example usage:

```bash
$ curl -iX POST 'http://localhost:8080/api/v1/data/shorten?longURL=www.google.com'
HTTP/1.1 200 OK
Content-Type: application/json; charset=utf-8
Date: Fri, 01 Dec 2023 05:10:33 GMT
Content-Length: 18

{"shortURL":"aA4"}
```

### URL redirection

`shortURL`: The request URI path

`GET /api/v1/:shortURL`

+ Request parameter: no, using URI path binding instead
+ Return: longURL for HTTP redirection

Example usage:

```bash
$ curl -i 'http://localhost:8080/api/v1/aA4'
HTTP/1.1 301 Moved Permanently
Content-Type: text/html; charset=utf-8
Location: /api/v1/www.google.com
Date: Fri, 01 Dec 2023 05:11:54 GMT
Content-Length: 57

<a href="/api/v1/www.google.com">Moved Permanently</a>.
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
