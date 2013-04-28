The Go programming language has been very aggressively advertised lately. A quick look at the titles on Hacker News is enough to convince anyone that they must switch to Go immediately or risk being left behind.

![brogrammers](https://raw.github.com/naturalist/go-versus-kelp/master/brogramming_go.jpg)

I tried Go and I thought it was nice, but the few things I didn't like about it made me abandon it. I won't be discussing any details here *cough* no triadic operator *cough*, but instead I am going to try to figure out, with the help of benchmarks and common sense, if Go is suitable for web development.

First, let's make sure we all understand each other here. It's 2013 and any server side language is only good for server side stuff. No contemporary web application will render HTML using the sever side. This is the front-end's job. The back end needs to parse and return JSON. That's all.

No analysis matters, unless it uses a comparison. This is why, I've chosen Perl and its Kelp web framework to see how Go compares.

The Test
---------

1. Create an http server that takes two routes:
1.1 `/get?p=<json>` - Takes a JSON structure, parses it and returns "OK" if everything looks well. For example: `/get?p={"x":"bar","y":"foo"}` will parse the JSON and return "OK".
1.1 `/put/:x/:y` - Takes a URL with two variables, creates a JSON structure and returns it back. For example: `/put/bar/foo` should return `{"x":"bar","y":"foo"}`
1. Put it behind nginx.
1. Siege the hell out of it.
1. Compare the results for Go versus Perl Kelp.
1. Make conclusions.

The Setup
---------

1. Install the latest stable Debian in a virtual machine
1. Install nginx
1. Install Go
1. Install [Kelp](https://metacpan.org/module/Kelp) and [Kelp::Module::JSON::XS](https://metacpan.org/module/Kelp::Module::JSON::XS)

The Code
--------

1. The Go code is contained in one file [server.go](https://github.com/naturalist/go-versus-kelp/blob/master/serve.go)
1. Perl Kelp has one code file [app.psgi](https://github.com/naturalist/go-versus-kelp/blob/master/app.psgi) and one [config file](https://github.com/naturalist/go-versus-kelp/blob/master/conf/config.pl).
1. Set up nginx to look for an upstream at `127.0.0.1:8080`
1. Prepare a [file with URLs](https://github.com/naturalist/go-versus-kelp/blob/master/urls.txt) for `siege`.

The Benchmarks
--------------

### Go

1. Run the http server: `./server`
1. Siege the server with a single user:

        > siege -b -c 1 -t 20s -f urls.txt

        Transactions:               4207 hits
        Availability:             100.00 %
        Elapsed time:              19.56 secs
        Data transferred:           0.05 MB
        Response time:              0.00 secs
        Transaction rate:         215.08 trans/sec
        Throughput:             0.00 MB/sec
        Concurrency:                0.94
        Successful transactions:        4207
        Failed transactions:               0
        Longest transaction:            1.83
        Shortest transaction:           0.00

1. Siege the server with four concurrent users:

        > siege -b -c 4 -t 20s -f urls.txt

        Transactions:              20960 hits
        Availability:             100.00 %
        Elapsed time:              19.42 secs
        Data transferred:           0.27 MB
        Response time:              0.00 secs
        Transaction rate:        1079.30 trans/sec
        Throughput:             0.01 MB/sec
        Concurrency:                3.98
        Successful transactions:       20960
        Failed transactions:               0
        Longest transaction:            0.80
        Shortest transaction:           0.00

### Perl Kelp

1. Run the http server: `plackup -E deployment -s Starman -p 8080`
1. Siege the server with a single user:

        > siege -b -c 1 -t 20s -f urls.txt

        Transactions:               4651 hits
        Availability:             100.00 %
        Elapsed time:              19.33 secs
        Data transferred:           0.06 MB
        Response time:              0.00 secs
        Transaction rate:         240.61 trans/sec
        Throughput:             0.00 MB/sec
        Concurrency:                1.00
        Successful transactions:        4651
        Failed transactions:               0
        Longest transaction:            1.91
        Shortest transaction:           0.00

1. Siege the server with four concurrent users:

        > siege -b -c 4 -t 20s -f urls.txt

        Transactions:              11842 hits
        Availability:             100.00 %
        Elapsed time:              19.48 secs
        Data transferred:           0.15 MB
        Response time:              0.01 secs
        Transaction rate:         607.91 trans/sec
        Throughput:             0.01 MB/sec
        Concurrency:                3.98
        Successful transactions:       11842
        Failed transactions:               0
        Longest transaction:            0.02
        Shortest transaction:           0.00

Speed
-----

Go and Kelp show approximately equal results when sieged by a single user. Go, however, is about 45% faster when four concurrent users are sieging the server. This was well expected, of course, after all Go is a lower lever compiled language, and Perl is a scripting language.

**Winner: Go**

JSON encoding/decoding
----------------------

Go doesn't handle arbitrary JSON structures. You can't just send any JSON and expect Go's oddly named function `Unmarshal` to figure it out. You have to have the exact structure, so you can first define it.

```go
type Message struct {
	X string `json:"x"`
	Y string `json:"y"`
}
```

The above definition allows for this structure `{"X":"bar","Y":"foo"}`. If you (or a client using your API) sends another key, `{"X":"bar","Y":"foo","Z":"baz"}` for example, the "Z" key will be lost. Perl, on the other hand will keep it. Anyone who has worked at a company with more than three server-side developers will tell you, that it is quite often that you don't know the exact structure of your JSON input and output. Yes, you can define a long structure covering every single option, but that sounds like a big pain. Perl makes it easier to deal with JSON.

**Winner: Perl**

Length and readability
----------------------

The Go code is 54 lines and pretty verbose. Structural definitions, handlers, oh my ... Kelp's code is 21 lines and pretty obvious:

```perl
get '/put/:x/:y' => sub {
    my ( $self, $x, $y ) = @_;
    { x => $x, y => $y };
};
```
It's very clear what's happening here, you define an HTTP GET route that catches '/put/:x/:y', then right away you have x and y as variables.

Go, while also clear to any developer with enough experience, requires more writing and more reading. This is where some of you may say "But Perl has all those weird symbols, and semicolons and stuff". My answer to you is that if these kind of things bother you, then you're not fit for a programmer and you should consider a career in fashion design and beauty.

**Winner: Perl Kelp**

Final conclusions
-----------------

When it comes to web development, writing Go for faster performance is not worth the trouble. It gets too verbose and its JSON handling is too strict. For the sake of rapid development and your own sanity, consider using a scripted language and an elegant web framework such as [Kelp](https://metacpan.org/module/Kelp).
