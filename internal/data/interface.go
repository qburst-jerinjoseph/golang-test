package data

import "context"

//Repo methods to communicate with models
type Repo interface {
	GetSample(c context.Context) ([]string, error)
}
