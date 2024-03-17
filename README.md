# Blackhole

This is a very simple web & SMTP server which will respond to all requests sent to it and log the output to stdout. This is designed for testing some Postal interactions with live SMTP and HTTP servers.

## Usage

Start the server with the following:

```
docker run --pull always \
           -p 80:8080 \
           -p 25:2525 \
           ghcr.io/postalserver/blackhole:latest
```

### SMTP

You can send email to various addresses to illicit different responses:

* `accept@` - message will be accepted with a 250 
* `softfail@` - message will be rejected with a 450 ("Mailbox unavailable at the moment")
* `later@` - message will be rejected with a 450 ("Try again in 250 seconds")
* `hardfail@` - message will be rejected with a 550 ("Invalid recipient address") after receiving data
* `anything-else@` - message will be rejected with a 550 ("Invalid recipient address") afer `RCPT TO`

### HTTP

The following URLs are available. 

* `/200` or `/ok` - returns a 200 OK
* `/403` or `/forbidden` - returns a 403 Forbidden
* `/500` or `/internal-server-error` - returns a 500 Internal Server Error
* All other URLs will return a 404 Not Found
