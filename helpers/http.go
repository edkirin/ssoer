package helpers

import (
	"fmt"

	beego "github.com/beego/beego/v2/server/web"
)

func InternalError(c beego.Controller, errMessage string) {
	c.CustomAbort(500, fmt.Sprintf("Internal server error: %s", errMessage))
}
