Simple HTTP Server written in Go

### Usage
`go run server.go`

### Test
`curl --data-ascii '{"secret":"shotsfired"}' -X POST http://localhost:1337/hello/jonny --header "Content-Type:application/json" -i`
