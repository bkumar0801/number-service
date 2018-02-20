package appcontext

/*
Result ...
*/
type Result struct {
	Numbers []int
	Error   error
}

/*
ResponseBuilder ...
*/
type ResponseBuilder interface {
	Query(urls []string) []int
	Get(requesturl string) Result
}
