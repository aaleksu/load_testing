# load_testing
stupid load tester in Go

It takes at least1 argument: url.

Can be run like this:

`go run ./load_testing.go http://example.com [int n: times] [int expected response body length]`

or

`go build`
+
`./load_testing http://example.com [int n: times] [int expected response body length]`
