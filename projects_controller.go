package v1

import (
    "gohub/app/models/project"
    "gohub/app/policies"
    "gohub/app/requests"
    "gohub/pkg/response"

    "github.com/gin-gonic/gin"
)

type ProjectsController struct {
    BaseAPIController
}

func (ctrl *ProjectsController) Index(c *gin.Context) {
    projects := project.All()
    response.Data(c, projects)
}

func (ctrl *ProjectsController) Show(c *gin.Context) {
    projectModel := project.Get(c.Param("id"))
    if projectModel.ID == 0 {
        response.Abort404(c)
        return
    }
    response.Data(c, projectModel)
}

func (ctrl *ProjectsController) Store(c *gin.Context) {

    request := requests.ProjectRequest{}
    if ok := requests.Validate(c, &request, requests.ProjectSave); !ok {
        return
    }

    projectModel := project.Project{
        FieldName:      request.FieldName,
    }
    projectModel.Create()
    if projectModel.ID > 0 {
        response.Created(c, projectModel)
    } else {
        response.Abort500(c, "创建失败，请稍后尝试~")
    }
}

func (ctrl *ProjectsController) Update(c *gin.Context) {

    projectModel := project.Get(c.Param("id"))
    if projectModel.ID == 0 {
        response.Abort404(c)
        return
    }

    if ok := policies.CanModifyProject(c, projectModel); !ok {
        response.Abort403(c)
        return
    }

    request := requests.ProjectRequest{}
    bindOk, errs := requests.Validate(c, &request, requests.ProjectSave)
    if !bindOk {
        return
    }
    if len(errs) > 0 {
        response.ValidationError(c, errs)
        return
    }

    projectModel.FieldName = request.FieldName
    rowsAffected := projectModel.Save()
    if rowsAffected > 0 {
        response.Data(c, projectModel)
    } else {
        response.Abort500(c, "更新失败，请稍后尝试~")
    }
}

func (ctrl *ProjectsController) Delete(c *gin.Context) {

    projectModel := project.Get(c.Param("id"))
    if projectModel.ID == 0 {
        response.Abort404(c)
        return
    }

    if ok := policies.CanModifyProject(c, projectModel); !ok {
        response.Abort403(c)
        return
    }

    rowsAffected := projectModel.Delete()
    if rowsAffected > 0 {
        response.Success(c)
        return
    }

    response.Abort500(c, "删除失败，请稍后尝试~")
}