package method

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetMethod(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"method": "get",
	})
}

func PostMethod(ctx *gin.Context) {
	body, err := ioutil.ReadAll(ctx.Request.Body)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"err_msg": err,
		})
		return
	}
	var res map[string]interface{}
	err = json.Unmarshal(body, &res)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"err_msg": err,
		})
		return
	}
	code, ok := res["code"]
	if ok {
		httpCode, isNum := code.(float64)
		if isNum && httpCode >= 200 && httpCode < 600 {
			ctx.JSON(int(httpCode), res)
			return
		}
	}
	ctx.JSON(http.StatusOK, res)
}

func PutMethod(ctx *gin.Context) {
	PostMethod(ctx)
}

func PatchMethod(ctx *gin.Context) {
	PostMethod(ctx)
}

func DeleteMethod(ctx *gin.Context) {
	PostMethod(ctx)
}
