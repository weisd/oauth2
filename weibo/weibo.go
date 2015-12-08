package weibo

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/weisd/oauth2"
	"github.com/weisd/oauth2/internal"
	"io/ioutil"
	"net/url"
)

// docs http://open.weibo.com/wiki/%E6%8E%88%E6%9D%83%E6%9C%BA%E5%88%B6%E8%AF%B4%E6%98%8E
var Endpoint = oauth2.Endpoint{
	AuthURL:  "https://api.weibo.com/oauth2/authorize",
	TokenURL: "https://api.weibo.com/oauth2/access_token",
}

var TokenParser = internal.JsonTokenParser

func GetTokenUid(conf *oauth2.Config, tok *oauth2.Token) (uid string, err error) {

	client := conf.Client(oauth2.NoContext, tok)
	params := url.Values{}
	params.Set("access_token", tok.AccessToken)

	res, err := client.PostForm("https://api.weibo.com/oauth2/get_token_info", params)
	if err != nil {
		return "", err
	}

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

	uid, ok := v["uid"]
	if !ok {
		return "", errors.New("uid not found")
	}

	return uid, nil
}
