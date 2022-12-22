package jsonrpc2

import (
	"context"
	"reflect"
	"strings"
	"sync"
)

type Service struct {
	impl any

	conns      []*Conn
	connsMutex sync.Mutex

	methods map[string]*reflect.Value
}

func NewService(impl any) *Service {
	ret := &Service{
		impl:       impl,
		conns:      []*Conn{},
		connsMutex: sync.Mutex{},
		methods:    map[string]*reflect.Value{},
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

func (s *Service) Run(ctx context.Context) {
	for {
		s.checkAllConn(ctx)
	}
}

func (s *Service) Listen(conn *Conn) {
	s.connsMutex.Lock()
	defer s.connsMutex.Unlock()
	s.conns = append(s.conns, conn)
}

func (s *Service) checkAllConn(ctx context.Context) {
	for i, conn := range s.conns {
		select {
		case <-ctx.Done():
			return
		case <-conn.Ctx.Done():
			s.connsMutex.Lock()
			defer s.connsMutex.Unlock()
			log.With(ctx).
				Debugf("Not listen on conn %v any more", conn)
			s.conns = append(s.conns[:i], s.conns[i+1:]...)
			return
		case req := <-conn.Requests():
			go s.handleRequest(
				context.WithValue(conn.Ctx, "conn", conn),
				conn, req)
		}
	}
}

func (s *Service) handleRequest(ctx context.Context, conn *Conn, req *Request) {
	method, ok := s.methods[req.Method]
	if !ok {
		s.returnMethodNotFound(ctx, conn, req.Id)
		return
	}

	numIn := reflect.TypeOf(method).NumIn()
	// func(ctx context.Context, param *struct{}) (result*struct{}, error)

	method.Call()
}
