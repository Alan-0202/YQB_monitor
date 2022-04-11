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
	IsDelete    int8      `json:"is_delete" form:"is_deleted"`
}

type ServiceListOutput struct {
	Total int64                   `json:"total" form:"total" comment:"总数" validate:""`
	List  []ServiceItemOutput     `json:"list" form:"list" comment:"列表" validate:""`
}


//delete
type ServiceDeleteInput struct {
	ID int64 `json:"id" form:"id" comment:"服务ID" example:"56" validate:"required"` //服务ID
}

func(params *ServiceDeleteInput) BindValidParam(c *gin.Context) error{
	return public.DefaultGetValidParams(c, params)
}


// Add http service
type ServiceAddHttpInput struct {
	ServiceName string `json:"service_name" form:"service_name" comment:"服务名" example:"" validate:"required"` //服务名
	ServiceDesc string `json:"service_desc" form:"service_desc" comment:"服务描述" example:"" validate:"required,max=255,min=1"`     //服务描述
}

func(p *ServiceAddHttpInput) BindValidParam(c *gin.Context) error {
	return public.DefaultGetValidParams(c, p)
}
