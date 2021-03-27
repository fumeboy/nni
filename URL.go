package nni


type URL interface {
	isURL()
}
type Route struct { // the path is "/"

}
func (R Route) isURL() {}