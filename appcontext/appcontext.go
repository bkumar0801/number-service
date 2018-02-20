package appcontext

/*
Result ...
*/
type Result struct {
	Numbers []int // array of numbers received from host server
	Error   error // if not error the value will be nil
}

/*
ResponseBuilder ...
*/
type ResponseBuilder interface {
	Query(urls []string) []int
	Get(requesturl string) Result
}
