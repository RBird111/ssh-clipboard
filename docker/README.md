# Instructions

## Build image

```sh
docker build -t ssh-test .
```

## Run container

```sh
docker run --rm --name clip-test -dp 8022:22 ssh-test
```

## SSH into container

```sh
ssh root@localhost -p 8022
```

**Password**: root
