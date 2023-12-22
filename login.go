package main

import (
	"errors"
	"github.com/gocolly/colly"
	"github.com/spf13/viper"
	"strings"
)

func login() (err error) {
	userId := viper.GetString("usr-id")
	userPwd := viper.GetString("usr-pwd")

	c := colly.NewCollector()

	// attach callbacks after login
	c.OnResponse(func(r *colly.Response) {
		// alert('用户名或密码错误');
		body := string(r.Body)
		if !strings.Contains(body, userId) {
			invalidAccount := "用户名或密码错误"
			if strings.Contains(body, invalidAccount) {
				err = errors.New(invalidAccount)
				return
			} else {
				panic("登录错误！详情：" + body)
			}
		}
	})

	const url = "http://eds.newtouch.cn/eds3/DefaultLogin.aspx?lan=zh-cn"
	err = c.Post(url, map[string]string{"UserId": userId, "UserPsd": userPwd})

	return
}
