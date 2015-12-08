package qq

import (
	"encoding/json"
	"errors"
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
	res, err := client.PostForm("https://graph.qq.com/oauth2.0/me", params)

	defer res.Body.Close()

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	var v map[string]string

	fmt.Println(string(data))

	err = json.Unmarshal(data, &v)
	if err != nil {
		return "", err
	}

	uid, ok := v["openid"]
	if !ok {
		return "", errors.New("uid not found")
	}

	return "", nil
}
