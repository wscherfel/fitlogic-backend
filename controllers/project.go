package controllers

import (
	"github.com/wscherfel/fitlogic-backend/access"
	"github.com/labstack/echo"
	"github.com/wscherfel/fitlogic-backend/common"
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"github.com/wscherfel/fitlogic-backend/models"
	"time"
	"github.com/wscherfel/fitlogic-backend"
	"strconv"
)

type ProjectControllerConfig struct {
	UserDao *access.UserDAO
	ProjectDao *access.ProjectDAO
	RiskDao *access.RiskDAO
}



type ProjectController struct {
	ProjectControllerConfig
}

func NewProjectController(config ProjectControllerConfig) *ProjectController {
	return &ProjectController{
		ProjectControllerConfig: config,
	}
}

type ProjectAPI struct {
	ID uint
	Name string `valid:"required"`
	Description string

	Start string `valid:"required"`
	End string `valid:"required"`

	ManagerID uint `valid:"required"`
}

type ProjectDetailAPI struct {
	ProjectAPI

	Users []models.User `json:"omitempty"`
	Risks []models.Risk `json:"omitempty"`
}

func MapProjectToAPI(project models.Project) (ProjectAPI) {
	return ProjectAPI{
		ID: project.ID,
		Name: project.Name,
		Description: project.Description,
		Start: project.Start.Time.Format(fitlogic.TimeFormat),
		End: project.End.Time.Format(fitlogic.TimeFormat),
		ManagerID: project.ManagerID,
	}
}

func MapAPIToProject(req ProjectAPI) (models.Project, error) {
	start, err := time.Parse(fitlogic.TimeFormat, req.Start)
	if err != nil {
		return models.Project{}, err
	}
	end, err := time.Parse(fitlogic.TimeFormat, req.End)
	if err != nil {
		return models.Project{}, err
	}

	return models.Project{
		IsFinished: false,
		ManagerID: req.ManagerID,
		Name: req.Name,
		Description: req.Description,

		Start: common.JSONTime{start},
		End: common.JSONTime{end},
	}, nil
}

func (c *ProjectController) Create(ctx echo.Context) error {
	userID, role, err := common.GetUserIdAndRoleFromToken(ctx.Get("user").(*jwt.Token))
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, common.CreateError(err))
	}
	if role > models.RoleManager {
		return ctx.JSON(http.StatusUnauthorized, common.CreateError(common.ErrUnsufficientPrivileges))
	}

	req := ProjectAPI{}
	err = common.BindAndValid(ctx, &req)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, common.CreateError(err))
	}

	manager, err := c.UserDao.ReadByID(req.ManagerID)

	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, common.CreateError(err))
	}

	if role == models.RoleManager { // role is manager
		if req.ManagerID != userID {
			return ctx.JSON(http.StatusUnauthorized, common.CreateError(common.ErrCannotCreateProjectForOthers))
		}
	} else { // role is admin, users were filtered above
		if req.ManagerID != userID {

			if manager.Role < models.RoleManager {
				return ctx.JSON(http.StatusInternalServerError, common.CreateError(common.ErrUnsufficientPrivileges))
			}
		}
	}

	project, err := MapAPIToProject(req)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, common.CreateError(err))
	}

	c.ProjectDao.Create(&project)
	c.ProjectDao.AddUsersAssociation(&project, manager)

	return ctx.JSON(http.StatusOK, project)
}

func (c *ProjectController) GetAll(ctx echo.Context) error {
	_, _, err := common.GetUserIdAndRoleFromToken(ctx.Get("user").(*jwt.Token))
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, common.CreateError(err))
	}

	projects, err := c.ProjectDao.GetAll()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, common.CreateError(err))
	}

	return ctx.JSON(http.StatusOK, projects)
}

func (c *ProjectController) UnAssignUsers(ctx echo.Context) error {
	pathIDuint64, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, common.CreateError(common.ErrIdInPathWrongFormat))
	}
	pathID := uint(pathIDuint64)
	userID, role, err := common.GetUserIdAndRoleFromToken(ctx.Get("user").(*jwt.Token))
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, common.CreateError(err))
	}
	if role > models.RoleManager {
		return ctx.JSON(http.StatusBadRequest, common.CreateError(common.ErrUnsufficientPrivileges))
	}

	project, err := c.ProjectDao.ReadByID(pathID)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, common.CreateError(err))
	}
	if role == models.RoleManager && project.ManagerID != userID {
		return ctx.JSON(http.StatusUnauthorized, common.CreateError(common.ErrUnsufficientPrivileges))
	}

	ids := common.IDsRequest{}
	err = common.BindAndValid(ctx, &ids)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, common.CreateError(common.ErrUnsufficientPrivileges))
	}

	for _, id := range ids.IDs {
		user, err := c.UserDao.ReadByID(id)
		if err != nil {
			continue
		}
		project, _ = c.ProjectDao.RemoveUsersAssociation(project, user)
	}

	return ctx.NoContent(http.StatusOK)
}

func (c *ProjectController) AssignUsers(ctx echo.Context) error {
	pathIDuint64, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, common.CreateError(common.ErrIdInPathWrongFormat))
	}
	pathID := uint(pathIDuint64)
	userID, role, err := common.GetUserIdAndRoleFromToken(ctx.Get("user").(*jwt.Token))
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, common.CreateError(err))
	}
	if role > models.RoleManager {
		return ctx.JSON(http.StatusBadRequest, common.CreateError(common.ErrUnsufficientPrivileges))
	}

	project, err := c.ProjectDao.ReadByID(pathID)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, common.CreateError(err))
	}
	if role == models.RoleManager && project.ManagerID != userID {
		return ctx.JSON(http.StatusUnauthorized, common.CreateError(common.ErrUnsufficientPrivileges))
	}

	ids := common.IDsRequest{}
	err = common.BindAndValid(ctx, &ids)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, common.CreateError(common.ErrUnsufficientPrivileges))
	}

	for _, id := range ids.IDs {
		user, err := c.UserDao.ReadByID(id)
		if err != nil {
			continue
		}
		project, _ = c.ProjectDao.AddUsersAssociation(project, user)
	}

	return ctx.NoContent(http.StatusOK)
}

func (c *ProjectController) AssignRisks(ctx echo.Context) error {
	pathIDuint64, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, common.CreateError(common.ErrIdInPathWrongFormat))
	}
	pathID := uint(pathIDuint64)
	userID, role, err := common.GetUserIdAndRoleFromToken(ctx.Get("user").(*jwt.Token))
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, common.CreateError(err))
	}
	if role > models.RoleManager {
		return ctx.JSON(http.StatusBadRequest, common.CreateError(common.ErrUnsufficientPrivileges))
	}

	project, err := c.ProjectDao.ReadByID(pathID)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, common.CreateError(err))
	}
	if role == models.RoleManager && project.ManagerID != userID {
		return ctx.JSON(http.StatusUnauthorized, common.CreateError(common.ErrUnsufficientPrivileges))
	}

	ids := common.IDsRequest{}
	err = common.BindAndValid(ctx, &ids)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, common.CreateError(common.ErrUnsufficientPrivileges))
	}

	for _, id := range ids.IDs {
		risk, err := c.RiskDao.ReadByID(id)
		if err != nil {
			continue
		}
		project, _ = c.ProjectDao.AddRisksAssociation(project, risk)
	}

	return ctx.NoContent(http.StatusOK)
}

func (c *ProjectController) UnAssignRisks(ctx echo.Context) error {
	pathIDuint64, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, common.CreateError(common.ErrIdInPathWrongFormat))
	}
	pathID := uint(pathIDuint64)
	userID, role, err := common.GetUserIdAndRoleFromToken(ctx.Get("user").(*jwt.Token))
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, common.CreateError(err))
	}
	if role > models.RoleManager {
		return ctx.JSON(http.StatusBadRequest, common.CreateError(common.ErrUnsufficientPrivileges))
	}

	project, err := c.ProjectDao.ReadByID(pathID)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, common.CreateError(err))
	}
	if role == models.RoleManager && project.ManagerID != userID {
		return ctx.JSON(http.StatusUnauthorized, common.CreateError(common.ErrUnsufficientPrivileges))
	}

	ids := common.IDsRequest{}
	err = common.BindAndValid(ctx, &ids)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, common.CreateError(common.ErrUnsufficientPrivileges))
	}

	for _, id := range ids.IDs {
		risk, err := c.RiskDao.ReadByID(id)
		if err != nil {
			continue
		}
		project, _ = c.ProjectDao.RemoveRisksAssociation(project, risk)
	}

	return ctx.NoContent(http.StatusOK)
}

func (c *ProjectController) UpdateByID(ctx echo.Context) error {
	pathIDuint64, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, common.CreateError(common.ErrIdInPathWrongFormat))
	}
	pathID := uint(pathIDuint64)
	jwtID, JWTRole, err := common.GetUserIdAndRoleFromToken(ctx.Get("user").(*jwt.Token))
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, common.CreateError(err))
	}
	projectCheck, err := c.ProjectDao.ReadByID(pathID)
	if JWTRole >= models.RoleManager && jwtID != projectCheck.ManagerID { // the >= condition is for possibility of adding new user roles
		return ctx.JSON(http.StatusUnauthorized, common.CreateError(common.ErrUnsufficientPrivileges))
	}

	req := ProjectAPI{}
	err = common.BindAndValid(ctx, &req)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, common.CreateError(err))
	}

	project, err := MapAPIToProject(req)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, common.CreateError(err))
	}

	project.ID = pathID
	newVals, err := c.ProjectDao.Update(&project, pathID)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, common.CreateError(err))
	}

	return ctx.JSON(http.StatusOK, *newVals)
}

func (c *ProjectController) ReadByID(ctx echo.Context) error {
	pathIDuint64, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, common.CreateError(common.ErrIdInPathWrongFormat))
	}
	pathID := uint(pathIDuint64)
	_, _, err = common.GetUserIdAndRoleFromToken(ctx.Get("user").(*jwt.Token))
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, common.CreateError(err))
	}

	project, err := c.ProjectDao.ReadByID(pathID)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, common.CreateError(err))
	}

	project.Users, err = c.ProjectDao.GetAllAssociatedUsers(project)
	/*if err != nil {
		return ctx.JSON(http.StatusInternalServerError, common.CreateError(err))
	}*/
	project.Risks, err = c.ProjectDao.GetAllAssociatedRisks(project)
	/*if err != nil {
		return ctx.JSON(http.StatusInternalServerError, common.CreateError(err))
	}*/

	return ctx.JSON(http.StatusOK, project)
}
