package controllers

import (
	"github.com/wscherfel/fitlogic-backend/access"
	"github.com/wscherfel/fitlogic-backend/models"
	"github.com/labstack/echo"
	"github.com/wscherfel/fitlogic-backend/common"
	"net/http"
	"github.com/dgrijalva/jwt-go"
	"strconv"
)

type CmControllerConfig struct {
	CmDao *access.CounterMeasureDAO
}

// CmController is a controller for CounterMeasures
// it is currently not used, but has simple CRUD implemented
type CmController struct {
	CmControllerConfig
}

type CmAPI struct {
	ID uint
	Name string
	Description string
	Cost int
	Risks []uint `json:"omitempty"`
}

func MapCounterMeasureToAPI(cm models.CounterMeasure) CmAPI {
	riskIDs := []uint{}
	for _, r := range cm.Risks {
		riskIDs = append(riskIDs, r.ID)
	}

	return CmAPI{
		ID: cm.ID,
		Name: cm.Name,
		Description: cm.Description,
		Cost: cm.Cost,
		Risks: riskIDs,
	}
}

func MapAPIToCounterMeasure(api CmAPI) (models.CounterMeasure) {
	return models.CounterMeasure{
		//Model: gorm.Model{ID:api.ID},
		Name: api.Name,
		Description: api.Description,
		Cost: api.Cost,
	}
}

func NewCounterMeasureController(c CmControllerConfig) *CmController {
	return &CmController{
		CmControllerConfig: c,
	}
}

// Create will create a new countermeasure
func (c *CmController) Create(ctx echo.Context) error {
	_, _, err := common.GetUserIdAndRoleFromToken(ctx.Get("user").(*jwt.Token))
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, common.CreateError(err))
	}

	req := CmAPI{}
	err = common.BindAndValid(ctx, &req)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, common.CreateError(err))
	}

	cm := MapAPIToCounterMeasure(req)
	err = c.CmDao.Create(&cm)

	return ctx.JSON(http.StatusOK, cm)
}

// GetAll will return all countermeasures
func (c *CmController) GetAll(ctx echo.Context) error {
	_, _, err := common.GetUserIdAndRoleFromToken(ctx.Get("user").(*jwt.Token))
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, common.CreateError(err))
	}

	cms, err := c.CmDao.GetAll()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, common.CreateError(err))
	}

	return ctx.JSON(http.StatusOK, cms)
}

// ReadByID will return detail of countermeasure
func (c *CmController) ReadByID(ctx echo.Context)error {
	pathIDuint64, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, common.CreateError(common.ErrIdInPathWrongFormat))
	}
	pathID := uint(pathIDuint64)
	_, role, err := common.GetUserIdAndRoleFromToken(ctx.Get("user").(*jwt.Token))
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, common.CreateError(err))
	}
	if role >= models.RoleManager { // the >= condition is for possibility of adding new user roles
		return ctx.JSON(http.StatusUnauthorized, common.CreateError(common.ErrUnsufficientPrivileges))
	}

	cm, err := c.CmDao.ReadByID(pathID)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, common.CreateError(err))
	}

	cm.Risks, err = c.CmDao.GetAllAssociatedRisks(cm)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, common.CreateError(err))
	}

	return ctx.JSON(http.StatusOK, cm)
}

// UpdateByID will update countermeasure with ID in path
// to new values sent in request body
func (c *CmController) UpdateByID(ctx echo.Context) error {
	pathIDuint64, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, common.CreateError(common.ErrIdInPathWrongFormat))
	}
	pathID := uint(pathIDuint64)
	_, JWTRole, err := common.GetUserIdAndRoleFromToken(ctx.Get("user").(*jwt.Token))
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, common.CreateError(err))
	}
	if JWTRole >= models.RoleUser{
		return ctx.JSON(http.StatusUnauthorized, common.CreateError(common.ErrUnsufficientPrivileges))
	}

	req := CmAPI{}
	err = common.BindAndValid(ctx, &req)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, common.CreateError(err))
	}

	cm := MapAPIToCounterMeasure(req)
	newVals, err := c.CmDao.Update(&cm, pathID)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, common.CreateError(err))
	}

	return ctx.JSON(http.StatusOK, newVals)
}

// DeleteByID will delete countermeasure with id in path if logged user has
// sufficient privileges
func (c *CmController) DeleteByID(ctx echo.Context) error {
	pathIDuint64, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, common.CreateError(common.ErrIdInPathWrongFormat))
	}
	pathID := uint(pathIDuint64)
	_, JWTRole, err := common.GetUserIdAndRoleFromToken(ctx.Get("user").(*jwt.Token))
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, common.CreateError(err))
	}
	if JWTRole >= models.RoleUser{
		return ctx.JSON(http.StatusUnauthorized, common.CreateError(common.ErrUnsufficientPrivileges))
	}

	cm, err := c.CmDao.ReadByID(pathID)
	if err != nil {
		return ctx.JSON(http.StatusNotFound, common.CreateError(err))
	}

	err = c.CmDao.Delete(cm)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, common.CreateError(err))
	}

	return ctx.NoContent(http.StatusOK)
}
