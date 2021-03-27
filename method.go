package nni

type Method interface {
	withMethod() string
}

type GET struct{}
type POST struct{}
type PUT struct{}
type PATCH struct{}
type HEAD struct{}
type TRACE struct {}
type OPTION struct {}

func (m GET) withMethod() string{return "get"}
func (m POST) withMethod() string{return "post"}
func (m PUT) withMethod() string{return "put"}
func (m PATCH) withMethod() string{return "patch"}
func (m HEAD) withMethod() string{return "head"}
func (m TRACE) withMethod() string{return "trace"}
func (m OPTION) withMethod() string{return "option"}


