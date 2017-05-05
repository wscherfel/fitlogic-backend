package controllers

import (
	"github.com/wscherfel/fitlogic-backend/access"
	"github.com/labstack/echo"
	"github.com/wscherfel/fitlogic-backend/common"
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"github.com/wscherfel/fitlogic-backend/models"
	"time"
	"strconv"
	"github.com/spf13/viper"
)

type ProjectControllerConfig struct {
	UserDao *access.UserDAO
	ProjectDao *access.ProjectDAO
	RiskDao *access.RiskDAO
}

// ProjectController is a controller that handles endpoints that are bound
// to projects
type ProjectController struct {
	ProjectControllerConfig
}

func NewProjectController(config ProjectControllerConfig) *ProjectController {
	return &ProjectController{
		ProjectControllerConfig: config,
	}
}

// ProjectAPI is a structure of requests for project API endpoints
type ProjectAPI struct {
	ID uint
	Name string `valid:"required"`
	Description string

	Start string `valid:"required"`
	End string `valid:"required"`

	IsFinished bool

	ManagerID uint `valid:"required"`
}

// ProjectDetailAPI is a structure that is returned when /projects/:id
// is called
type ProjectDetailAPI struct {
	ProjectAPI

	Users []models.User `json:"omitempty"`
	Risks []models.Risk `json:"omitempty"`
}

// MapProjectToAPI will map project to API structure, currently not used
func MapProjectToAPI(project models.Project) (ProjectAPI) {
	return ProjectAPI{
		ID: project.ID,
		Name: project.Name,
		Description: project.Description,
		Start: project.Start,
		End: project.End,
		ManagerID: project.ManagerID,
	}
}

// MapAPIToProject will map request values to DB model
func MapAPIToProject(req ProjectAPI) (models.Project, error) {
	format := viper.GetString("TimeFormat")
	start, err := time.Parse(format, req.Start)
	if err != nil {
		return models.Project{}, err
	}
	if !start.After(common.DateMin) {
		return models.Project{}, common.ErrDateOutOfRange
	}
	end, err := time.Parse(format, req.End)
	if err != nil {
		return models.Project{}, err
	}
	if !end.Before(common.DateMax) {
		return models.Project{}, common.ErrDateOutOfRange
	}
	if end.Before(start) || end.Equal(start) {
		return models.Project{}, common.ErrStartDateAfterEnd
	}

	return models.Project{
		IsFinished: req.IsFinished,
		ManagerID: req.ManagerID,
		Name: req.Name,
		Description: req.Description,

		Start: req.Start,
		End: req.End,
	}, nil
}

// Create will create a new project
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
	project.IsFinished = false

	c.ProjectDao.Create(&project)
	c.ProjectDao.AddUsersAssociation(&project, manager)
	for i := range project.Users {
		project.Users[i].Password = ""
	}

	return ctx.JSON(http.StatusOK, project)
}

// GetAll will return all projects
func (c *ProjectController) GetAll(ctx echo.Context) error {
	_, _, err := common.GetUserIdAndRoleFromToken(ctx.Get("user").(*jwt.Token))
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, common.CreateError(err))
	}

	projects, err := c.ProjectDao.GetAll()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, common.CreateError(err))
	}
	for i := range projects {
		for j := range projects[i].Users {
			projects[i].Users[j].Password = ""
		}
	}

	return ctx.JSON(http.StatusOK, projects)
}

// UnAssignUsers will remove association between users with sent IDs
// and project with ID in path
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
		if project.ManagerID == id {
			continue
		}
		project, _ = c.ProjectDao.RemoveUsersAssociation(project, user)
	}

	return ctx.NoContent(http.StatusOK)
}

// AssignUsers will add association between users with sent IDs
// and project with ID in path
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
		return ctx.JSON(http.StatusBadRequest, common.CreateError(err))
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

// AssignRisks will add association between risks with sent IDs
// and project with ID in path
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
		return ctx.JSON(http.StatusBadRequest, err)
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

// UnAssignRisks will remove association between risks with sent IDs
// and project with ID in path
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

// UpdateByID will update project with ID in path
// to new values sent in request body
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

// DeleteByID will delete project with id in path if logged user has
// sufficient privileges
func (c *ProjectController) DeleteByID(ctx echo.Context) error {
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

	project, err := c.ProjectDao.ReadByID(pathID)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, common.CreateError(err))
	}
	c.ProjectDao.Delete(project)

	return ctx.NoContent(http.StatusOK)
}

// GetRisksOfProjects will return risks of projects with sent IDs
func (c *ProjectController) GetRisksOfProjects(ctx echo.Context) error {
	_, _, err := common.GetUserIdAndRoleFromToken(ctx.Get("user").(*jwt.Token))
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, common.CreateError(err))
	}

	ids := common.IDsRequest{}
	err = common.BindAndValid(ctx, &ids)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}

	ret := []models.Risk{}
	usedRisksMap := make(map[uint]bool)

	for _, id := range ids.IDs {
		project, _ := c.ProjectDao.ReadByID(id)
		risks, _ := c.ProjectDao.GetAllAssociatedRisks(project)

		for _, risk := range risks {
			if _, ok := usedRisksMap[risk.ID]; ok {
				continue
			}
			ret = append(ret, risk)
			usedRisksMap[risk.ID] = true
		}
	}

	return ctx.JSON(http.StatusOK, ret)
}

// ReadByID will return detail of project with ID in path
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
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, common.CreateError(err))
	}

	for i := range project.Users {
		project.Users[i].Password = ""
	}

	project.Risks, err = c.ProjectDao.GetAllAssociatedRisks(project)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, common.CreateError(err))
	}

	return ctx.JSON(http.StatusOK, project)
}
