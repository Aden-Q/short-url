# Go Web App Template
This is a template to bootstrap building Web applications in Go. It is kept small and bare for better extensibility.

We use the command runner [just](https://github.com/casey/just) as a makefile alternative.

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

## License

[MIT License](./LICENSE)
