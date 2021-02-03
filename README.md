# iplogger

iplogger is an IP logger written in Go. It is intended for OSINT and also functions as URL shortener.

## Usage as a URL shortener

First we must create a URL.

`POST /urls -d '{"url": "http://www.rquinlivan.net"}'`

```
Response:
{"shortUrl": "e18993fe"}
```

Now we can visit `yoursite.com/link/e18993fe` which will 301 redirect to the URL we provided.

## Usage as an IP logger

When serving the 301, the server also logs the IP for each access. This can be useful for [OSINT](https://en.wikipedia.org/wiki/OSINT).
You can run your own instance of `iplogger` on a private server to collect information about your targets.