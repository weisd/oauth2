package qq

import (
	"fmt"
	"github.com/weisd/oauth2"
	"github.com/weisd/oauth2/internal"
	"io/ioutil"
	"net/url"
)

var Endpoint = oauth2.Endpoint{
	AuthURL:  "https://graph.qq.com/oauth2.0/authorize",
	TokenURL: "https://graph.qq.com/oauth2.0/token",
}

var TokenParser = internal.UrlTokenParser

func GetTokenUid(conf *oauth2.Config, tok *oauth2.Token) (uid string, err error) {

	client := conf.Client(oauth2.NoContext, tok)
	params := url.Values{}
	params.Set("access_token", tok.AccessToken)
	res, err := client.PostForm("https://graph.z.qq.com/moc2/me", params)

	defer res.Body.Close()

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	fmt.Println(string(data))

	vals, err := url.ParseQuery(string(data))
	if err != nil {
		return "", err
	}

	return vals.Get("openid"), nil
}
