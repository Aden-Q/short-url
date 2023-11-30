# Short URL
A fully functional URL shortener using MySQL as the storage server and Redis as the cache server. Also, there is rate limiting in place.

We use the command runner [just](https://github.com/casey/just) as a makefile alternative.

We use `direnv` and .envrc so that we can create new migrations with less pain.

The current components include:
+ A bare bone Web app 
+ Version bump and autotag for release
+ A simple CI workflow running with GitHub Actions

## API Endpoints

1. URL shortening

`POST /api/v1/data/shorten`

+ Request parameter: {longURL: string}
+ Rreturn: shortURL

2. URL redirecting

`GET /api/v1/:shortURL`

+ Return: longURL for HTTP redirection

## Build and Run Locally

We use docker-compose.

First use docker-compose up to start a web server, a mysql server, and a redis server by running:

```bash
$ docker-compose up -d
```

To connect to the mysql server, run:

```bash
mysql -h 127.0.0.1 -P 3306 -u root -p db
```

To connect to the redis server, run:

```
redis-cli
```


## License

[MIT License](./LICENSE)
