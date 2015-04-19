Example TCP Proxy
=================
An example TCP proxy in Go. It listens on an address and proxies them to a
single backend address.

Usage
-----
You need to specify a -listen and -backend option. For example:

    go-proxy-example -listen=0.0.0.0:80 -backend=10.1.1.2:80

This would proxy all local connections to port 80 to port 80 on the machine at 10.1.1.2.

License
-------
Copyright (c) 2015 Ryan Bourgeois. Licensed under BSD-Modified. See the [LICENSE][1] file for a copy of the license.

[1]: https://raw.githubusercontent.com/BlueDragonX/go-proxy-example/master/LICENSE "LICENSE"
