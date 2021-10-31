# DevKit

- this is a basic http server able to easy test frontend and backend code.
- default use bootstrap5 and vue3 lib, you can change to whatever you want, please change via `views/header.html` and `views/footer.html`
- support a basic api implementation

# How to

## Add new page with path "/newpage"

- `cp views/empty.html views/newpage.html`
- add content inside `views/newpage.html`
- add css inside `static/css/newpage.css`
- add js inside `static/js/newpage.js`

## Add new api

- add implementation to api.go, check example implemente `ping`, you can check api via `curl 'localhost:8080/api?action=ping'`

## Docker

- `Dockerfile` current use vendor, so run `go mod vendor` after clone the project
- `docker-compose up -d`

