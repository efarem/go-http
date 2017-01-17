Simple HTTP Server written in Go

### Usage
`go build && ./go-http`

### Test
`curl --data-ascii '{"secret":"shotsfired"}' -X POST http://localhost:1337/hello/jonny --header "Content-Type:application/json" -i`
