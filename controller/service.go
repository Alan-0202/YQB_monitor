package controller

import (
	"errors"
	"fmt"
	"github.com/e421083458/golang_common/lib"
	"github.com/e421083458/yqb_monitor/dao"
	"github.com/e421083458/yqb_monitor/dto"
	"github.com/e421083458/yqb_monitor/middleware"
	"github.com/gin-gonic/gin"
)

type ServiceController struct {

}

func ServiceRegister(group *gin.RouterGroup) {
	service := &ServiceController{}
	group.GET("service_list",service.ServiceList)
	group.GET("service_delete",service.ServiceDelete)
	group.POST("service_add_http", service.ServiceAddHttp)
}



// ServiceList godoc
// @Summary 服务列表
// @Description 服务列表
// @Tags 服务管理
// @ID /service/service_list
// @Accept  json
// @Produce  json
// @Param info query string false "关键词"
// @Param page_size query int true "每页个数"
// @Param page_no query int true "当前页数"
// @Success 200 {object} middleware.Response{data=dto.ServiceListOutput} "success"
// @Router /service/service_list [get]
func (service *ServiceController) ServiceList(c *gin.Context) {
	params := &dto.ServiceListInput{}
	if err := params.BindValidParam(c); err != nil {
		middleware.ResponseError(c, 2000, err)
		return
	}

	//db connection
	tx, err := lib.GetGormPool("default")
	if err != nil {
		middleware.ResponseError(c, 2001, err)
		return
	}

	serviceList := &dao.ServiceInfo{}
	list, total, err := serviceList.PageList(c, tx, params)
	if err != nil {
		middleware.ResponseError(c, 2002, err)
		return
	}

	outList := []dto.ServiceItemOutput{}
	for _, listItem := range list {
		if listItem.IsDelete == 1 {
			continue
		}
		outItem := dto.ServiceItemOutput{
			ID: listItem.ID,
			LoadType: listItem.LoadType,
			ServiceName: listItem.ServiceName,
			ServiceDesc: listItem.ServiceDesc,
			IsDelete: listItem.IsDelete,
		}
		outList = append(outList, outItem)
	}
	out := &dto.ServiceListOutput{
		Total: total,
		List: outList,
	}
	middleware.ResponseSuccess(c, out)
}


// ServiceDelete godoc
// @Summary 服务删除
// @Description 服务删除
// @Tags 服务管理
// @ID /service/service_delete
// @Accept  json
// @Produce  json
// @Param id query string true "服务ID"
// @Success 200 {object} middleware.Response{data=string} "ok"
// @Router /service/service_delete [get]
func (service *ServiceController) ServiceDelete (c *gin.Context) {
	params := &dto.ServiceDeleteInput{}
	if err := params.BindValidParam(c); err != nil {
		middleware.ResponseError(c, 2000, err)
		return
	}
	serviceInfo := &dao.ServiceInfo{
		ID: params.ID,
	}

	//DB
	tx, err := lib.GetGormPool("default")
	if err != nil {
		middleware.ResponseError(c, 2001, err)
		return
	}

	serviceInfo, err = serviceInfo.Find(c, tx, serviceInfo)
	if err != nil {
		middleware.ResponseError(c, 2002, err)
		return
	}

	serviceInfo.IsDelete = 1
	fmt.Println(serviceInfo)

	if err := serviceInfo.Save(c, tx); err != nil {
		middleware.ResponseError(c, 2003, err)
		return
	}
	middleware.ResponseSuccess(c,"ok")
}


// ServiceAddHTTP godoc
// @Summary 添加HTTP服务
// @Description 添加HTTP服务
// @Tags 服务管理
// @ID /service/service_add_http
// @Accept  json
// @Produce  json
// @Param body body dto.ServiceAddHttpInput true "body"
// @Success 200 {object} middleware.Response{data=string} "success"
// @Router /service/service_add_http [post]
func(t *ServiceController) ServiceAddHttp(c *gin.Context) {
	params := &dto.ServiceAddHttpInput{}
	if err := params.BindValidParam(c); err != nil {
		middleware.ResponseError(c, 2000, err)
		return
	}

	tx, err := lib.GetGormPool("default")
	if err != nil {
		middleware.ResponseError(c, 2001, err)
		return
	}

	tx = tx.Begin()

	fmt.Println(params.ServiceName)
	serviceInfo := &dao.ServiceInfo{ServiceName: params.ServiceName}
	fmt.Println(serviceInfo)
	if find, err := serviceInfo.Find(c, tx, serviceInfo); err == nil {

		if find.ServiceName == serviceInfo.ServiceName {
			tx.Rollback()
			middleware.ResponseError(c, 2002, errors.New("服务已存在"))
			return
		}
	}

	serviceNew := &dao.ServiceInfo{
		ServiceName: params.ServiceName,
		ServiceDesc: params.ServiceDesc,
	}

	if err := serviceNew.Save(c, tx); err != nil {
		tx.Rollback()
		middleware.ResponseError(c, 2005, err)
		return
	}

	tx.Commit()

	middleware.ResponseSuccess(c, "ok")
}