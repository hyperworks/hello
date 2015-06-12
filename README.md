# HELLO, WORLD!

This is a tiny GO program that...

1. Waits for a data on a set port (the only argument)
2. Looks for a `\r\n\r\n` sequence (indicating end of an HTTP packet) 
3. Spits out the "Hello, World!" HTTP response...
4. ... and immediately close the conneciton.

```http
HTTP/1.1 200 OK
Content-Type: text/plain
Content-Length: 13

Hello, World!
```

You'd be surprised the amount of time you can save with this when setting disparate parts
of your application services!

# WHAT's THE USE?

* Testing NGINX upstream configuration.
* Testing load balancer configuration.
* Testing API gateways.
* Load-testing a middleware.
* Misc testing uses when setting up microservices.

# LICENSE

BSD

