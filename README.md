## Run containers

```bash
$ docker-compose up -d
```

## Run project locally

```bash
$ make start
```

or

```bash
$ ./scripts/start.sh
```

## Destroy project locally

```bash
$ make destroy
```

```bash
$ ./scripts/destroy.sh
```

## Install dependencies from toolkit (private repo)

```bash
$ GOPRIVATE=github.com/quessapp go get -u -f github.com/quessapp/toolkit
```
