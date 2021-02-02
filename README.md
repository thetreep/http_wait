
<h1 align="center">
  <br>
  <a href="https://thetreep.com"><img src="https://www.thetreep.com/uploads/1/1/0/1/110164989/logo-the-treep-carr-vert-texte-blanc-web_1.png" alt="the Treep" width="200"></a>
  <br>
  HTTP Wait
  <br>
</h1>

<h4 align="center">A minimal CLI tool that waits for a specific <a href="https://developer.mozilla.org/en-US/docs/Web/HTTP/Status" target="_blank">HTTP status code</a> to be returned.</h4>

<p align="center">
  <a href="#key-features">Key Features</a> •
  <a href="#how-to-use">How To Use</a> •
  <a href="#download">Download</a> •
</p>

## Features

* Configurable port, timeout, polling interval, and response code.
* That's it !

## How To Use

You can either install it via the go cli, or download it from the [releases](https://github.com/thetreep/http_wait/releases)

```bash
# Install directly with go get
$ go get github.com/thetreep/http_wait

# Check out the built-in doc!
$ http_wait -h
Usage of http_wait:
      --i int               alias to interval (default 1000)
      --interval int        Polling interval, in milliseconds. Defaults to 1000ms (default 1000)
      --p int               alias to port (default 80)
      --port int            Port to use. Defaults to 80 (default 80)
      --rc int              alias to response-code (default 200)
      --response-code int   Response code to wait for. Defaults to 200 (default 200)
      --t int               alias to timeout (default 10000)
      --timeout int         Timeout before returning a non 0 exit code, in milliseconds. Defaults to 10000ms (default 10000)
      --u string            alias to uri (default "http://localhost")
      --uri string          URI, with protocol and no port. Defaults to http://localhost (default "http://localhost")
```

## Download

You can [download](https://github.com/thetreep/http_wait/releases) the latest installable version of http_wait for Windows, macOS and Linux (ARM builds included).

---

> [thetreep.com](https://thetreep.com) &nbsp;&middot;&nbsp;
> GitHub [@thetreep](https://github.com/thetreep)

