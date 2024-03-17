# Blackhole

This is a very simple web & SMTP server which will response to all requests sent to it and log the output to stdout. This is designed for testing some Postal interactions with live SMTP and HTTP servers.

## Usage

Start the server with the following:

```
docker run --pull always \
           -p 80:8080 \
           -p 25:2525 \
           ghcr.io/postalserver/blackhole:latest
```

### SMTP

You can send email to `accept@anydomain.com` which will be accepted. Any other username will be blocked at the `RCPT TO` stage.

### HTTP

The following URLs are available. 

* `/200` - returns a 200 OK
* `/403` - returns a 403 Forbidden
* `/500` - returns a 500 Internal Server Error
* All other URLs will return a 404 Not Found
