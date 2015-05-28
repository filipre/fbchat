# fbchat
<!-- _This Go package is part of the [Facebook Onlinetime Logger]() project. Please follow the link if you are interested into it._ -->

[Go](http://golang.org/) package to access the Facebook messenger and to gain data from it like currently online users.

## Usage and Example

Create a `fbchat.Client` and provide your `c_user`, `datr` and `xs` values from your *facebook.com* cookie file.

### Currently Online

```go
c, err := fbchat.NewClient(cUser, cookie, &http.Client{})
online, err := c.ReqOnline()
```

### Friend's Messenger Information

```go
c, err := fbchat.NewClient(cUser, cookie, &http.Client{})
friends, err := c.ReqFriends(cUser1, cUser2, ...)
```

## TODO

- godoc & more comments in source code
- tests

## Licence (MIT)

Copyright (c) 2015 [René Filip](http://github.com/filipre)

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
