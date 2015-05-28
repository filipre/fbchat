# timelogger

<!-- _This Go package is part of the [Facebook Onlinetime Logger]() project. Please follow the link if you are interested into it._ -->

[Go](http://golang.org/) package to `Start` and `Stop` a logging job by a given Facebook user. It is build on top of the [fbchat](http://github.com/filipre/fbchat) package.

## Usage and Examples

Create a `fbchat.Client` and provide your `c_user`, `datr` and `xs` values from your *facebook.com* cookie file. The package makes sure, that only one instance per `c_user` is running. If a job is already running for a specific user, then it will restart it.

### Start a Job

```go
t := timelogger.New()
go t.Start(cUser, cookie, time.Duration(interval), saveCh, errCh)
```
### Stop a Job

```go
t.Stop(cUser)
```
