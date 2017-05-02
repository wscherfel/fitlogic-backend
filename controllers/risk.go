package controllers

import (
	"github.com/wscherfel/fitlogic-backend/access"
	"github.com/labstack/echo"
	"github.com/wscherfel/fitlogic-backend/models"
	"github.com/wscherfel/fitlogic-backend"
	"time"
	"github.com/wscherfel/fitlogic-backend/common"
	"net/http"
	"github.com/dgrijalva/jwt-go"
	"strconv"
)

type RiskControllerConfig struct {
	RiskDao *access.RiskDAO
	UserDao *access.UserDAO
	ProjectDao *access.ProjectDAO
	CmDao *access.CounterMeasureDAO
}

type RiskController struct {
	RiskControllerConfig
}

func NewRiskController(conf RiskControllerConfig) *RiskController {
	return &RiskController{
		RiskControllerConfig: conf,
	}
}

type RiskAPI struct {
	ID uint
	Value float64
	Cost int
	Probability float64
	Risk float64

	Name string
	Description string
	Category string
	Threat string
	Status string
	Trigger string
	Impact float64

	Start string
	End string

	UserID uint

	Projects []uint `json:"omitempty"`

	CounterMeasures []uint `json:"omitempty"`
}

func MapAPIToRisk(req RiskAPI) (models.Risk, error){
	start, err := time.Parse(fitlogic.TimeFormat, req.Start)
	if err != nil {
		return models.Risk{}, err
	}
	end, err := time.Parse(fitlogic.TimeFormat, req.Start)
	if err != nil {
		return models.Risk{}, err
	}

	return models.Risk{
		Value: req.Value,
		Cost: req.Cost,
		Probability: req.Probability,
		Risk: req.Risk,
		Name: req.Name,
		Description: req.Description,
		Category: req.Category,
		Threat: req.Threat,
		Status: req.Status,
		Trigger: req.Trigger,
		Impact: req.Impact,

		Start: common.JSONTime{start},
		End: common.JSONTime{end},

		UserID: req.UserID,
	}, nil
}

func MapRiskToAPI(r models.Risk) (RiskAPI) {
	proj := []uint{}
	for _, p := range r.Projects {
		proj = append(proj, p.ID)
	}

	cms := []uint{}
	for _, c := range r.CounterMeasures {
		cms = append(cms, c.ID)
	}

	return RiskAPI{
		ID: r.ID,
		Value: r.Value,
		Cost: r.Cost,
		Probability: r.Probability,
		Risk: r.Risk,
		Name: r.Name,
		Description: r.Description,
		Category: r.Category,
		Threat: r.Threat,
		Status: r.Status,
		Trigger: r.Trigger,
		Impact: r.Impact,
		Start: r.Start.Time.Format(fitlogic.TimeFormat),
		End: r.End.Time.Format(fitlogic.TimeFormat),
		UserID: r.UserID,
		Projects: proj,
		CounterMeasures: cms,
	}
}

func (c *RiskController) Create(ctx echo.Context) error {
	userID, role, err := common.GetUserIdAndRoleFromToken(ctx.Get("user").(*jwt.Token))
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, common.CreateError(err))
	}
	req := RiskAPI{}
	err = common.BindAndValid(ctx, &req)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, common.CreateError(err))
	}

	risk, err := MapAPIToRisk(req)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, common.CreateError(err))
	}


	if role > models.RoleAdmin && risk.UserID != userID {
		return ctx.JSON(http.StatusUnauthorized, common.CreateError(common.ErrUnsufficientPrivileges))
	}

	err = c.RiskDao.Create(&risk)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, common.CreateError(err))
	}

	return ctx.JSON(http.StatusOK, risk)
}

func (c *RiskController) GetAll(ctx echo.Context) error {
	_, _, err := common.GetUserIdAndRoleFromToken(ctx.Get("user").(*jwt.Token))
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, common.CreateError(err))
	}

	all, err := c.RiskDao.GetAll()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, common.CreateError(err))
	}

	return ctx.JSON(http.StatusOK, all)
}

func (c *RiskController) ReadByID(ctx echo.Context) error {
	pathIDuint64, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, common.CreateError(common.ErrIdInPathWrongFormat))
	}
	pathID := uint(pathIDuint64)

	_, _, err = common.GetUserIdAndRoleFromToken(ctx.Get("user").(*jwt.Token))
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, common.CreateError(err))
	}

	risk, err := c.RiskDao.ReadByID(pathID)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, common.CreateError(err))
	}

	risk.Projects, err = c.RiskDao.GetAllAssociatedProjects(risk)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, common.CreateError(err))
	}

	risk.CounterMeasures, err = c.RiskDao.GetAllAssociatedCounterMeasures(risk)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, common.CreateError(err))
	}

	return ctx.JSON(http.StatusOK, risk)
}

func (c *RiskController) UpdateByID(ctx echo.Context) error {
	pathIDuint64, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, common.CreateError(common.ErrIdInPathWrongFormat))
	}
	pathID := uint(pathIDuint64)
	jwtID, JWTRole, err := common.GetUserIdAndRoleFromToken(ctx.Get("user").(*jwt.Token))
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, common.CreateError(err))
	}
	riskCheck, err := c.RiskDao.ReadByID(pathID)
	if JWTRole >= models.RoleManager && jwtID != riskCheck.UserID {
		return ctx.JSON(http.StatusUnauthorized, common.CreateError(common.ErrUnsufficientPrivileges))
	}

	req := RiskAPI{}
	err = common.BindAndValid(ctx, &req)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, common.CreateError(err))
	}

	risk, err := MapAPIToRisk(req)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, common.CreateError(err))
	}

	risk.ID = pathID

	newVals, err := c.RiskDao.Update(&risk, pathID)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, common.CreateError(err))
	}

	return ctx.JSON(http.StatusOK, newVals)
}

func (c *RiskController) DeleteByID(ctx echo.Context) error {
	pathIDuint64, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, common.CreateError(common.ErrIdInPathWrongFormat))
	}
	pathID := uint(pathIDuint64)
	jwtID, JWTRole, err := common.GetUserIdAndRoleFromToken(ctx.Get("user").(*jwt.Token))
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, common.CreateError(err))
	}
	riskCheck, err := c.RiskDao.ReadByID(pathID)
	if JWTRole >= models.RoleManager && jwtID != riskCheck.UserID {
		return ctx.JSON(http.StatusUnauthorized, common.CreateError(common.ErrUnsufficientPrivileges))
	}

	risk, err := c.RiskDao.ReadByID(pathID)
	if err != nil {
		return ctx.JSON(http.StatusNotFound, common.CreateError(err))
	}

	err = c.RiskDao.Delete(risk)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, common.CreateError(err))
	}

	return ctx.NoContent(http.StatusOK)
}

func (c *RiskController) AssignCms(ctx echo.Context) error {
	pathIDuint64, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, common.CreateError(common.ErrIdInPathWrongFormat))
	}
	pathID := uint(pathIDuint64)
	jwtID, JWTRole, err := common.GetUserIdAndRoleFromToken(ctx.Get("user").(*jwt.Token))
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, common.CreateError(err))
	}
	risk, err := c.RiskDao.ReadByID(pathID)
	if JWTRole >= models.RoleManager && jwtID != risk.UserID {
		return ctx.JSON(http.StatusUnauthorized, common.CreateError(common.ErrUnsufficientPrivileges))
	}

	ids := common.IDsRequest{}
	err = common.BindAndValid(ctx, &ids)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, common.CreateError(common.ErrUnsufficientPrivileges))
	}

	for _, id := range ids.IDs {
		cm, err := c.CmDao.ReadByID(id)
		if err != nil {
			continue
		}
		risk, _ = c.RiskDao.AddCounterMeasuresAssociation(risk, cm)
	}

	return ctx.NoContent(http.StatusOK)
}

func (c *RiskController) UnAssignCms(ctx echo.Context) error {
	pathIDuint64, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, common.CreateError(common.ErrIdInPathWrongFormat))
	}
	pathID := uint(pathIDuint64)
	jwtID, JWTRole, err := common.GetUserIdAndRoleFromToken(ctx.Get("user").(*jwt.Token))
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, common.CreateError(err))
	}
	risk, err := c.RiskDao.ReadByID(pathID)
	if JWTRole >= models.RoleManager && jwtID != risk.UserID {
		return ctx.JSON(http.StatusUnauthorized, common.CreateError(common.ErrUnsufficientPrivileges))
	}

	ids := common.IDsRequest{}
	err = common.BindAndValid(ctx, &ids)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, common.CreateError(common.ErrUnsufficientPrivileges))
	}

	for _, id := range ids.IDs {
		cm, err := c.CmDao.ReadByID(id)
		if err != nil {
			continue
		}
		risk, _ = c.RiskDao.RemoveCounterMeasuresAssociation(risk, cm)
	}

	return ctx.NoContent(http.StatusOK)
}
