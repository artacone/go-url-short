# go-url-short

URL shortener written in Go.

### Build
```shell
docker-compose up --build shortener db
```
By default, PostgreSQL database is used as a storage.
If `DB_USE` is set to `false` inside `docker-compose.yml`, then simple in-memory
storage is used.

### CI
All tests and linters are run on each push/pull-request

### Example
```shell
❯ grpcurl -plaintext -d '{"URL": "google.com"}' localhost:50000 api.URLShortener/Create
{
  "URL": "rD6blXM7cF"
}
❯ grpcurl -plaintext -d '{"URL": "rD6blXM7cF"}' localhost:50000 api.URLShortener/Get
{
  "URL": "google.com"
}
```
Table `urls` in database after several calls:

| url | short |
| :--- | :--- |
| yandex.ru | nCxaJ25gik |
| vk.ru/feed | aXFpipcn3L |
| t.me/artacone | MSffWO3ylk |
| ozon.ru | UhYQ2mMPjX |
| google.com | rD6blXM7cF |
