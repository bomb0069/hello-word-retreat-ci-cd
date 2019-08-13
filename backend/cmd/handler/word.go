package handler

import (
	"github.com/gin-gonic/gin"
)

func GetWordHandler(context *gin.Context) {
	context.String(200, `{text:"สวัสดีชาวโลก"}`)
}
