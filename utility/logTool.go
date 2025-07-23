package utility

import (
	"context"

	"github.com/gogf/gf/v2/os/glog"
)

func Clog(v ...any) {
	glog.Skip(1).Line(true).Info(context.TODO(), v...)
}
