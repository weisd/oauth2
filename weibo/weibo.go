package weibo

import (
	"golang.org/x/oauth2"
)

// docs http://open.weibo.com/wiki/%E6%8E%88%E6%9D%83%E6%9C%BA%E5%88%B6%E8%AF%B4%E6%98%8E
var Endpoint = oauth2.Endpoint{
	AuthURL:  "https://api.weibo.com/oauth2/authorize",
	TokenURL: "https://api.weibo.com/oauth2/access_token",
}
