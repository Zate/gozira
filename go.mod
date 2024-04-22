module github.com/zate/gozira

go 1.18

replace (
	github.com/andygrunwald/go-jira/v2 => .
	github.com/zate/gozira => .
)

require (
	github.com/andygrunwald/go-jira/v2 v2.0.0-00010101000000-000000000000
	github.com/fatih/structs v1.1.0
	github.com/golang-jwt/jwt/v4 v4.5.0
	github.com/google/go-cmp v0.6.0
	github.com/google/go-querystring v1.1.0
	github.com/trivago/tgo v1.0.7
	golang.org/x/term v0.16.0
)

require golang.org/x/sys v0.16.0 // indirect
