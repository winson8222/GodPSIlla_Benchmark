namespace go postersapi

struct Request {}

struct Response {
    1: required list<string> posterslist;
}

service PosterService {
    Response getuniqueusernames8020(1: Request req)
    Response getuniqueusernames2080(1: Request req)
    Response getuniqueusernames8099920(1: Request req)
    Response getuniqueusernames2099980(1: Request req)
    Response getuniqueusernames80k20k(1: Request req)
    Response getuniqueusernames20k80k(1: Request req)
}

