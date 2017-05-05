package access

import(
  "github.com/jinzhu/gorm"
  "github.com/wscherfel/fitlogic-backend/models"
)


// CounterMeasureDAO is a data access object to a database containing models.CounterMeasures
type CounterMeasureDAO struct {
	db *gorm.DB
}

// NewCounterMeasureDAO creates a new Data Access Object for the
// models.CounterMeasure model.
func NewCounterMeasureDAO(db *gorm.DB) *CounterMeasureDAO {
	return &CounterMeasureDAO{
		db: db,
	}
}

// Create will create single models.CounterMeasure in database.
func (dao *CounterMeasureDAO) Create(m *models.CounterMeasure) (error) {
	if err := dao.db.Create(m).Error; err != nil {
		return err
	}
	return nil
}

// Update will update a record of models.CounterMeasure in DB
func (dao *CounterMeasureDAO) Update(m *models.CounterMeasure, id uint) (*models.CounterMeasure, error) {
	oldVal, err := dao.ReadByID(id)
	if err != nil {
		return nil, err
	}

	if err := dao.db.Model(&oldVal).Updates(m).Error; err != nil {
		return nil, err
	}
	return oldVal, nil
}

// Delete will soft-delete a single models.CounterMeasure
func (dao *CounterMeasureDAO) Delete(m *models.CounterMeasure) (error) {
	if err := dao.db.Delete(m).Error; err != nil {
		return err
	}
	return nil
}

// GetAll will return all records of models.CounterMeasure in database
func (dao *CounterMeasureDAO) GetAll() ([]models.CounterMeasure, error) {
	m := []models.CounterMeasure{}
	if err := dao.db.Find(&m).Error; err != nil {
		return nil, err
	}

	return m, nil
}

// AddRisksAssociation will add
// an association to model given by parameter
func (dao *CounterMeasureDAO) AddRisksAssociation(m *models.CounterMeasure, asocVal *models.Risk) (*models.CounterMeasure, error) {
	if err := dao.db.Model(&m).Association("Risks").Append(asocVal).Error; err != nil {
		return nil, err
	}

	return m, nil
}

// RemoveRisksAssociation will remove
// an association from model given by parameter
func (dao *CounterMeasureDAO) RemoveRisksAssociation(m *models.CounterMeasure, asocVal *models.Risk) (*models.CounterMeasure, error) {
	if err := dao.db.Model(&m).Association("Risks").Delete(asocVal).Error; err != nil {
		return nil, err
	}

	return m, nil
}

// GetAllAssociatedRisks will get all
// an association from model given by parameter
func (dao *CounterMeasureDAO) GetAllAssociatedRisks(m *models.CounterMeasure) ([]models.Risk, error) {
	retVal := []models.Risk{}

	dao.db.Model(&m).Association("Risks").Find(&retVal)
	return retVal, nil
}

// ReadByID will find models.CounterMeasure by ID given by parameter
func (dao *CounterMeasureDAO) ReadByID(id uint) (*models.CounterMeasure, error) {
	m := &models.CounterMeasure{}
	if err := dao.db.First(&m, id).Error; err != nil {
		return nil, err
	}

	return m, nil
}

// UserDAO is a data access object to a database containing models.Users
type UserDAO struct {
	db *gorm.DB
}

// NewUserDAO creates a new Data Access Object for the
// models.User model.
func NewUserDAO(db *gorm.DB) *UserDAO {
	return &UserDAO{
		db: db,
	}
}

// Create will create single models.User in database.
func (dao *UserDAO) Create(m *models.User) (error) {
	if err := dao.db.Create(m).Error; err != nil {
		return err
	}
	return nil
}

// Read will find all DB records matching
// values in a model given by parameter
func (dao *UserDAO) Read(m *models.User) ([]models.User, error) {
	retVal := []models.User{}
	if err := dao.db.Where(m).Find(&retVal).Error; err != nil {
		return nil, err
	}
	return retVal, nil
}

// Update will update a record of models.User in DB
func (dao *UserDAO) Update(m *models.User, id uint) (*models.User, error) {
	oldVal, err := dao.ReadByID(id)
	if err != nil {
		return nil, err
	}

	if err := dao.db.Model(&oldVal).Updates(m).Error; err != nil {
		return nil, err
	}
	return oldVal, nil
}

// Delete will soft-delete a single models.User
func (dao *UserDAO) Delete(m *models.User) (error) {
	if err := dao.db.Delete(m).Error; err != nil {
		return err
	}
	return nil
}

// GetAll will return all records of models.User in database
func (dao *UserDAO) GetAll() ([]models.User, error) {
	m := []models.User{}
	if err := dao.db.Find(&m).Error; err != nil {
		return nil, err
	}

	return m, nil
}

// AddProjectsAssociation will add
// an association to model given by parameter
func (dao *UserDAO) AddProjectsAssociation(m *models.User, asocVal *models.Project) (*models.User, error) {
	if err := dao.db.Model(&m).Association("Projects").Append(asocVal).Error; err != nil {
		return nil, err
	}

	return m, nil
}

// RemoveProjectsAssociation will remove
// an association from model given by parameter
func (dao *UserDAO) RemoveProjectsAssociation(m *models.User, asocVal *models.Project) (*models.User, error) {
	if err := dao.db.Model(&m).Association("Projects").Delete(asocVal).Error; err != nil {
		return nil, err
	}

	return m, nil
}

// GetAllAssociatedProjects will get all
// an association from model given by parameter
func (dao *UserDAO) GetAllAssociatedProjects(m *models.User) ([]models.Project, error) {
	retVal := []models.Project{}

	dao.db.Model(&m).Association("Projects").Find(&retVal)
	return retVal, nil
}

// AddRisksAssociation will add
// an association to model given by parameter
func (dao *UserDAO) AddRisksAssociation(m *models.User, asocVal *models.Risk) (*models.User, error) {
	if err := dao.db.Model(&m).Association("Risks").Append(asocVal).Error; err != nil {
		return nil, err
	}

	return m, nil
}

// RemoveRisksAssociation will remove
// an association from model given by parameter
func (dao *UserDAO) RemoveRisksAssociation(m *models.User, asocVal *models.Risk) (*models.User, error) {
	if err := dao.db.Model(&m).Association("Risks").Delete(asocVal).Error; err != nil {
		return nil, err
	}

	return m, nil
}

// GetAllAssociatedRisks will get all
// an association from model given by parameter
func (dao *UserDAO) GetAllAssociatedRisks(m *models.User) ([]models.Risk, error) {
	retVal := []models.Risk{}

	dao.db.Model(&m).Association("Risks").Find(&retVal)
	return retVal, nil
}

// ReadByEmail will find all records
// matching the value given by parameter
func (dao *UserDAO) ReadByEmail(m string) ([]models.User, error) {
	retVal := []models.User{}
	if err := dao.db.Where(&models.User{Email: m}).Find(&retVal).Error; err != nil {
		return nil, err
	}

	return retVal, nil
}

// ReadByID will find models.User by ID given by parameter
func (dao *UserDAO) ReadByID(id uint) (*models.User, error) {
	m := &models.User{}
	if err := dao.db.First(&m, id).Error; err != nil {
		return nil, err
	}

	return m, nil
}

// ProjectDAO is a data access object to a database containing models.Projects
type ProjectDAO struct {
	db *gorm.DB
}

// NewProjectDAO creates a new Data Access Object for the
// models.Project model.
func NewProjectDAO(db *gorm.DB) *ProjectDAO {
	return &ProjectDAO{
		db: db,
	}
}

// Create will create single models.Project in database.
func (dao *ProjectDAO) Create(m *models.Project) (error) {
	if err := dao.db.Create(m).Error; err != nil {
		return err
	}
	return nil
}

// Update will update a record of models.Project in DB
func (dao *ProjectDAO) Update(m *models.Project, id uint) (*models.Project, error) {
	oldVal, err := dao.ReadByID(id)
	if err != nil {
		return nil, err
	}

	if err := dao.db.Model(&oldVal).Updates(m).Error; err != nil {
		return nil, err
	}
	return oldVal, nil
}

// Delete will soft-delete a single models.Project
func (dao *ProjectDAO) Delete(m *models.Project) (error) {
	if err := dao.db.Delete(m).Error; err != nil {
		return err
	}
	return nil
}

// GetAll will return all records of models.Project in database
func (dao *ProjectDAO) GetAll() ([]models.Project, error) {
	m := []models.Project{}
	if err := dao.db.Find(&m).Error; err != nil {
		return nil, err
	}

	return m, nil
}

// AddUsersAssociation will add
// an association to model given by parameter
func (dao *ProjectDAO) AddUsersAssociation(m *models.Project, asocVal *models.User) (*models.Project, error) {
	if err := dao.db.Model(&m).Association("Users").Append(asocVal).Error; err != nil {
		return nil, err
	}

	return m, nil
}

// RemoveUsersAssociation will remove
// an association from model given by parameter
func (dao *ProjectDAO) RemoveUsersAssociation(m *models.Project, asocVal *models.User) (*models.Project, error) {
	if err := dao.db.Model(&m).Association("Users").Delete(asocVal).Error; err != nil {
		return nil, err
	}

	return m, nil
}

// GetAllAssociatedUsers will get all
// an association from model given by parameter
func (dao *ProjectDAO) GetAllAssociatedUsers(m *models.Project) ([]models.User, error) {
	retVal := []models.User{}

	dao.db.Model(&m).Association("Users").Find(&retVal)
	return retVal, nil
}

// AddRisksAssociation will add
// an association to model given by parameter
func (dao *ProjectDAO) AddRisksAssociation(m *models.Project, asocVal *models.Risk) (*models.Project, error) {
	if err := dao.db.Model(&m).Association("Risks").Append(asocVal).Error; err != nil {
		return nil, err
	}

	return m, nil
}

// RemoveRisksAssociation will remove
// an association from model given by parameter
func (dao *ProjectDAO) RemoveRisksAssociation(m *models.Project, asocVal *models.Risk) (*models.Project, error) {
	if err := dao.db.Model(&m).Association("Risks").Delete(asocVal).Error; err != nil {
		return nil, err
	}

	return m, nil
}

// GetAllAssociatedRisks will get all
// an association from model given by parameter
func (dao *ProjectDAO) GetAllAssociatedRisks(m *models.Project) ([]models.Risk, error) {
	retVal := []models.Risk{}

	dao.db.Model(&m).Association("Risks").Find(&retVal)
	return retVal, nil
}

// ReadByID will find models.Project by ID given by parameter
func (dao *ProjectDAO) ReadByID(id uint) (*models.Project, error) {
	m := &models.Project{}
	if err := dao.db.First(&m, id).Error; err != nil {
		return nil, err
	}

	return m, nil
}

// RiskDAO is a data access object to a database containing models.Risks
type RiskDAO struct {
	db *gorm.DB
}

// NewRiskDAO creates a new Data Access Object for the
// models.Risk model.
func NewRiskDAO(db *gorm.DB) *RiskDAO {
	return &RiskDAO{
		db: db,
	}
}

// Create will create single models.Risk in database.
func (dao *RiskDAO) Create(m *models.Risk) (error) {
	if err := dao.db.Create(m).Error; err != nil {
		return err
	}
	return nil
}

// Update will update a record of models.Risk in DB
func (dao *RiskDAO) Update(m *models.Risk, id uint) (*models.Risk, error) {
	oldVal, err := dao.ReadByID(id)
	if err != nil {
		return nil, err
	}

	if err := dao.db.Model(&oldVal).Updates(m).Error; err != nil {
		return nil, err
	}
	return oldVal, nil
}

// Delete will soft-delete a single models.Risk
func (dao *RiskDAO) Delete(m *models.Risk) (error) {
	if err := dao.db.Delete(m).Error; err != nil {
		return err
	}
	return nil
}

// GetAll will return all records of models.Risk in database
func (dao *RiskDAO) GetAll() ([]models.Risk, error) {
	m := []models.Risk{}
	if err := dao.db.Find(&m).Error; err != nil {
		return nil, err
	}

	return m, nil
}

// AddProjectsAssociation will add
// an association to model given by parameter
func (dao *RiskDAO) AddProjectsAssociation(m *models.Risk, asocVal *models.Project) (*models.Risk, error) {
	if err := dao.db.Model(&m).Association("Projects").Append(asocVal).Error; err != nil {
		return nil, err
	}

	return m, nil
}

// RemoveProjectsAssociation will remove
// an association from model given by parameter
func (dao *RiskDAO) RemoveProjectsAssociation(m *models.Risk, asocVal *models.Project) (*models.Risk, error) {
	if err := dao.db.Model(&m).Association("Projects").Delete(asocVal).Error; err != nil {
		return nil, err
	}

	return m, nil
}

// GetAllAssociatedProjects will get all
// an association from model given by parameter
func (dao *RiskDAO) GetAllAssociatedProjects(m *models.Risk) ([]models.Project, error) {
	retVal := []models.Project{}

	dao.db.Model(&m).Association("Projects").Find(&retVal)
	return retVal, nil
}

// ReadByID will find models.Risk by ID given by parameter
func (dao *RiskDAO) ReadByID(id uint) (*models.Risk, error) {
	m := &models.Risk{}
	if err := dao.db.First(&m, id).Error; err != nil {
		return nil, err
	}

	return m, nil
}