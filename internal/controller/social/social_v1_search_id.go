package social

import (
	"context"

	v1 "zeal_be/api/social/v1"
	"zeal_be/internal/service"
)

func (c *ControllerV1) SearchId(ctx context.Context, req *v1.SearchIdReq) (res *v1.SearchIdRes, err error) {
	res = &v1.SearchIdRes{}
	res.User_list, res.Total = service.Social().SearchId(ctx, req.SearchReqInfo)
	return
}
