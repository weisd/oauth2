package qq

import (
	"github.com/weisd/oauth2"
	"github.com/weisd/oauth2/internal"
)

var Endpoint = oauth2.Endpoint{
	AuthURL:  "https://graph.qq.com/oauth2.0/authorize",
	TokenURL: "https://graph.qq.com/oauth2.0/token",
}

var TokenParser = internal.JsonTokenParser
