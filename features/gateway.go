package features

import "github.com/cs-sea/golang-sdk/common"

type Gateway interface {
	Post(path string, request interface{}) (*common.Response, error)
}
