package controller

import (
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
		outItem := dto.ServiceItemOutput{
			ID: listItem.ID,
			LoadType: listItem.LoadType,
			ServiceName: listItem.ServiceName,
			ServiceDesc: listItem.ServiceDesc,
		}
		outList = append(outList, outItem)
	}
	out := &dto.ServiceListOutput{
		Total: total,
		List: outList,
	}
	middleware.ResponseSuccess(c, out)
}
