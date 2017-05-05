package controllers

import (
	"github.com/wscherfel/fitlogic-backend/access"
	"github.com/labstack/echo"
	"github.com/wscherfel/fitlogic-backend/models"
	"time"
	"github.com/wscherfel/fitlogic-backend/common"
	"net/http"
	"github.com/dgrijalva/jwt-go"
	"strconv"
	"github.com/spf13/viper"
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

	CounterMeasureUsed bool
	CounterMeasureCost int
	CounterMeasureDesc string
}

func MapAPIToRisk(req RiskAPI) (models.Risk, error){
	format := viper.GetString("TimeFormat")
	start, err := time.Parse(format, req.Start)
	if err != nil {
		return models.Risk{}, err
	}
	if !start.After(common.DateMin) {
		return models.Risk{}, common.ErrDateOutOfRange
	}
	end, err := time.Parse(format, req.End)
	if err != nil {
		return models.Risk{}, err
	}
	if !end.Before(common.DateMax) {
		return models.Risk{}, common.ErrDateOutOfRange
	}
	if end.Before(start) || end.Equal(start) {
		return models.Risk{}, common.ErrStartDateAfterEnd
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

		Start: req.Start,
		End: req.End,

		UserID: req.UserID,

		CounterMeasureUsed: req.CounterMeasureUsed,
		CounterMeasureCost: req.CounterMeasureCost,
		CounterMeasureDesc: req.CounterMeasureDesc,
	}, nil
}

func MapRiskToAPI(r models.Risk) (RiskAPI) {
	proj := []uint{}
	for _, p := range r.Projects {
		proj = append(proj, p.ID)
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
		Start: r.Start,
		End: r.End,
		UserID: r.UserID,
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


	if risk.UserID != userID {
		if role > models.RoleAdmin {
			return ctx.JSON(http.StatusUnauthorized, common.CreateError(common.ErrUnsufficientPrivileges))
		}
		_, err := c.UserDao.ReadByID(risk.UserID)
		if err != nil {
			return ctx.JSON(http.StatusInternalServerError, common.CreateError(err))
		}
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

/*func (c *RiskController) AssignCms(ctx echo.Context) error {
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
}*/

/*func (c *RiskController) UnAssignCms(ctx echo.Context) error {
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
*/
