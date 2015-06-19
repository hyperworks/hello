# HELLO, WORLD!

This is a tiny GO program (and Docker image) that...

1. Waits for a data on port 8080
2. Looks for a `\r\n\r\n` sequence (indicating end of an HTTP packet) 
3. Spits out the "Hello, World!" (or program's arguments as text) HTTP response...
4. ... and immediately close the conneciton.

```http
HTTP/1.1 200 OK
Content-Type: text/plain
Connection: close
Content-Length: 13

Hello, World!
```

You'd be surprised the amount of time you can save with this when setting disparate parts
of your application services!

# USAGE

```sh
$ docker pull hyperworks/hello
$ docker run -p 8080:8080 hyperworks/hello
```

Or create a named container:

```sh
$ docker pull hyperworks/hello
$ docker create --name hello -p 8080:8080 hyperworks/hello
$ docker start hello
```

Or, even better, create TEN hello world services!

```sh
$ docker pull hyperworks/hello
$ for i in {0..9}
  do
    docker create --name hello$i -p 808$i:8080 hyperworks/hello Hello from node $i
    docker start hello$i
  done
```

# WHAT's THE USE?

* Testing NGINX upstream configuration.
* Testing load balancer configuration.
* Testing API gateways.
* Load-testing a middleware.
* Misc testing uses when setting up microservices.

# LICENSE

BSD

# TODO

* Persistent connection.
* Container w/o any Go stuff.

