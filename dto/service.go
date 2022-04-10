package dto

import (
	"github.com/e421083458/yqb_monitor/public"
	"github.com/gin-gonic/gin"
)

type ServiceListInput struct {
	Info     string `json:"info" form:"info" comment:"关键词" example:"" validate:""`                      //关键词
	PageNo   int    `json:"page_no" form:"page_no" comment:"页数" example:"1" validate:"required"`        //页数
	PageSize int    `json:"page_size" form:"page_size" comment:"每页条数" example:"20" validate:"required"` //每页条数
}

func (param *ServiceListInput) BindValidParam(c *gin.Context) error {
	return public.DefaultGetValidParams(c, param)
}

type ServiceItemOutput struct {
	ID          int64       `json:"id" form:"id"`
	LoadType    int       `json:"load_type" form:"load_type"`
	ServiceName string    `json:"service_name" form:"service_name"`
	ServiceDesc string    `json:"service_desc" form:"service_desc"`
}

type ServiceListOutput struct {
	Total int64                   `json:"total" form:"total" comment:"总数" validate:""`
	List  []ServiceItemOutput     `json:"list" form:"list" comment:"列表" validate:""`
}

