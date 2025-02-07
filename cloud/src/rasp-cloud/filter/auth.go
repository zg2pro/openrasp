//Copyright 2017-2020 Baidu Inc.
//
//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http: //www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.

package filter

import (
	"net/http"
	"rasp-cloud/models"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego/plugins/cors"
)

func init() {
	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Authorization", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type", "X-OpenRASP-Token"},
		ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		AllowCredentials: true,
	}))
	beego.InsertFilter("/v1/agent/*", beego.BeforeRouter, authAgent)
	beego.InsertFilter("/v1/api/*", beego.BeforeRouter, authApi)
	beego.InsertFilter("/v1/iast/auth", beego.BeforeRouter, authAgent)
	beego.InsertFilter("/v1/iast/version", beego.BeforeRouter, authAgent)
	beego.InsertFilter("/v1/user/islogin", beego.BeforeRouter, authApi)
	beego.InsertFilter("/v1/user/default", beego.BeforeRouter, authApi)
}

func authAgent(ctx *context.Context) {
	var (
		appId     = ctx.Input.Header("X-OpenRASP-AppID")
		appSecret = ctx.Input.Header("X-OpenRASP-AppSecret")
		app, err  = models.GetAppById(appId)
	)

	if appId == "" || app == nil || err != nil {
		ctx.Output.JSON(map[string]interface{}{
			"status": http.StatusBadRequest, "description": "Invalid X-OpenRASP-AppID value specified"},
			false, false)
		return
	}

	if appSecret != app.Secret {
		ctx.Output.JSON(map[string]interface{}{
			"status": http.StatusUnauthorized, "description": "Bad X-OpenRASP-AppSecret value, authentication failed"},
			false, false)
		return
	}
}

func authApi(ctx *context.Context) {
	// contentType := strings.ToLower(ctx.Input.Header("Content-Type"))
	// if !strings.Contains(contentType, "application/json") {
	// 	ctx.Output.JSON(map[string]interface{}{
	// 		"status": http.StatusBadRequest, "description": "Invalid Content-Type specified"},
	// 		false, false)
	// 	panic("")
	// }

	cookie := ctx.GetCookie(models.AuthCookieName)
	if has, err := models.HasCookie(cookie); !has || err != nil {
		token := ctx.Input.Header(models.AuthTokenName)
		if has, err = models.HasToken(token); !has || err != nil {
			ctx.Output.JSON(map[string]interface{}{
				"status": http.StatusUnauthorized, "description": http.StatusText(http.StatusUnauthorized)},
				false, false)
			panic("")
		}
	}
}
