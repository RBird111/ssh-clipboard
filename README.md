# SSH Clipboard

A WIP `ssh` wrapper which starts up a server that writes incoming
data into the local clipboard. The server gracefully shuts down
once the `ssh` command exits.

## Requirements

[Docker](https://www.docker.com/)

[Go](https://go.dev/)

## Nice to have

[xsel](https://github.com/kfish/xsel)
If you don't have `xsel` you just have to modify the clipboard command in `main.go`

## How to run

```bash
make
```

You'll SSH directly into the newly created container and be prompted for a password:

```txt
root@localhost's password: root
```

A file called `rand.txt` gets created when the image is built so once inside the
container you can copy to your clipboard like so:

```txt
root@559d5c49fe3c:~# cat rand.txt > /dev/tcp/localhost/8888
```

Once you exit the container will be shut down and removed and the program will shutdown.
