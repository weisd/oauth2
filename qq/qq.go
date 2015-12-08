package qq

import (
	"encoding/json"
	"fmt"
	"github.com/weisd/oauth2"
	"github.com/weisd/oauth2/internal"
	"io/ioutil"
	"net/url"
	"strings"
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
	if err != nil {
		return "", err
	}

	defer res.Body.Close()

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	bodyStr := string(data)
	start := strings.Index(bodyStr, "{")
	end := strings.LastIndex(bodyStr, "}")
	jsonStr := string(bodyStr[start : end+1])

	fmt.Println(string(data))
	fmt.Println(jsonStr)

	var v map[string]string
	err = json.Unmarshal([]byte(jsonStr), &v)

	return v["openid"], nil
}
