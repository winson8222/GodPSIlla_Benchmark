namespace go viewersapi

struct Request {}

struct Response {
	1: required list<string> viewerslist;
}

service ViewerService {
	Response getuniqueviewernames8020(1: Request req)
	Response getuniqueviewernames2080(1: Request req)
	Response getuniqueviewernames8099920(1: Request req)
	Response getuniqueviewernames2099980(1: Request req)
	Response getuniqueviewernames80k20k(1: Request req)
	Response getuniqueviewernames20k80k(1: Request req)
}

