package jsonrpc2

import (
	"encoding/json"
	"reflect"
	"strings"
)

type Service struct {
	impl any

	conn *Conn

	methods map[string]*reflect.Value
}

func NewService(impl any, conn *Conn) *Service {
	ret := &Service{
		impl:    impl,
		conn:    conn,
		methods: map[string]*reflect.Value{},
	}

	ret.initMethods()

	return ret
}

func (s *Service) initMethods() {
	rv := reflect.ValueOf(s.impl)
	rt := reflect.TypeOf(s.impl)
	for i := 0; i < rv.NumField(); i++ {
		fieldValue := rv.Field(i)
		field := rt.Field(i)
		if !fieldValue.IsNil() {
			if name, ok := parseStructTag(field); ok {
				s.methods[name] = &fieldValue
			}
		}
	}

	if len(s.methods) == 0 {
		panic("no JSON-RPC 2.0 method found")
	}

	log.Debug("registered methods", s.methods)
}

func parseStructTag(rsf reflect.StructField) (string, bool) {
	if rsf.Type.Kind() != reflect.Func {
		return "", false
	}

	value, ok := rsf.Tag.Lookup("jsonrpc2")
	if !ok {
		return "", false
	}

	name := rsf.Name

	components := strings.Split(value, ",")
	for _, component := range components {
		if strings.HasPrefix(component, "name=") {
			name = component[len("name="):]
		}
	}

	return name, true
}

func (s *Service) Run() {
	for {
		select {
		case <-s.conn.ctx.Done():
			return
		case req := <-s.conn.Requests():
			go s.handleRequest(req)
		}
	}
}

func (s *Service) handleRequest(req *Request) {
	method, ok := s.methods[req.Method]
	if !ok {
		s.respond(req, nil, newMethodNotFound(req.Method))
		return
	}

	numIn := reflect.TypeOf(method).NumIn()
	// func(ctx context.Context, param *struct{}) (result*struct{}, error)

	method.Call()
}

func (s *Service) respond(req *Request, result json.RawMessage, err *Error) {
}
