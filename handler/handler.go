package handler

import (
	"bufio"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/frannecki/laddergo/server"
)

type GoHandler interface {
	server.Handler
	IsGoHandler() // to differ from server.Handler
}

// satisfies the interface GoHandler
type goHandler struct {
	server.Handler
}

func (p *goHandler) IsGoHandler() {}

// a struct overriding default struct for the interface server.Handler
type overwrittenMethodsOnHandler struct {
	p server.Handler
}

func NewGoHandler() GoHandler {
	om := &overwrittenMethodsOnHandler{}
	p := server.NewDirectorHandler(om)
	om.p = p

	// fill with a specific handler
	return &goHandler{Handler: p}
}

func IsGoHandler(p goHandler) {}

type HttpRequest struct {
	Method   string            `json:"method"`
	Uri      string            `json:"uri"`
	Protocol string            `json:"protocol"`
	Headers  map[string]string `json:"headers"`
	Body     string            `json:"body"`
}

func parseHttpRequest(request_str string) HttpRequest {
	content_position := strings.Index(request_str, "\r\n\r\n")
	if content_position == -1 {
		content_position = len(request_str)
	} else {
		content_position += 4
	}

	content := request_str[content_position:]
	header := request_str[:content_position]

	var method string
	var uri string
	var protocol string
	var key string
	var value string
	var headers = make(map[string]string)

	scanner := bufio.NewScanner(strings.NewReader(header))
	scanner.Scan()
	fmt.Sscanf(scanner.Text(), "%s %s %s", &method, &uri, &protocol)

	for scanner.Scan() {
		fmt.Sscanf(scanner.Text(), "%s %s", &key, &value)
		key = key[:len(key)-1]
		headers[key] = value
	}

	return HttpRequest{Method: method, Uri: uri, Protocol: protocol, Headers: headers, Body: content}
}

func (p *overwrittenMethodsOnHandler) OnRequest(request string) string {
	req := parseHttpRequest(request)
	marshaled, err := json.Marshal(req)
	var response_str string
	if err != nil {
		err_str := err.Error()
		response_str = fmt.Sprintf("HTTP/1.1 500 Internal Server Error\r\nContent-Type: text/plain\r\nContent-Length: %d\r\n\r\n%s", len(err_str), err_str)
	} else {
		response_str = fmt.Sprintf("HTTP/1.1 200 OK\r\nContent-Type: application/json\r\nContent-Length: %d\r\n\r\n%s", len(marshaled), marshaled)
	}
	return response_str
}
