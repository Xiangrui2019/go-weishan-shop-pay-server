package utils

import (
	"fmt"
	"net/http"
	"net/url"
	"os"
)

func SendWeixinNotify(title string, content string, mapurl string) {
	userid := os.Getenv("WEIXIN_NOTIFY_USERKEY")

	_, err := http.Get(fmt.Sprintf("http://wxmsg.dingliqc.com/send?title=%s&msg=%s&userIds=%s&url=%s", url.QueryEscape(title), url.QueryEscape(content), url.QueryEscape(userid), mapurl))

	if err != nil {
		panic(err)
	}
}
