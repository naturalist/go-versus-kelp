The Go programming language has been very aggressively advertised lately. A quick look at the titles on Hacker News is enough to convince anyone that they must switch to Go immediately or risk being left behind.

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
1. Perl Kelp has one code file [app.psgi](https://github.com/naturalist/go-versus-kelp/blob/master/app.psgi) and one config file.

