package access

import(
  "github.com/jinzhu/gorm"
  "time"
  "github.com/wscherfel/fitlogic-backend/models"
  )


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

// Read will find all DB records matching
// values in a model given by parameter
func (dao *ProjectDAO) Read(m *models.Project) ([]models.Project, error) {
	retVal := []models.Project{}
	if err := dao.db.Where(m).Find(&retVal).Error; err != nil {
		return nil, err
	}
	return retVal, nil
}

// ReadT will return a transaction that
// can be used to find DB records matching with models
func (dao *ProjectDAO) ReadT(m *models.Project) (*gorm.DB, error) {
	retVal := dao.db.Where(m)
	return retVal, retVal.Error
}

// ReadByIDT will return a transaction that
// an be used to find DB record with ID given by parameter
func (dao *ProjectDAO) ReadByIDT(id uint) (*gorm.DB, error) {
	//m := &models.Project{}
	retVal := dao.db.Where("ID = ?", id)

	return retVal, retVal.Error
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

// UpdateAllFields will update ALL fields of models.Project in db
// with values given in the models.Project by parameter
func (dao *ProjectDAO) UpdateAllFields(m *models.Project) (*models.Project, error) {
	if err := dao.db.Save(&m).Error; err != nil {
		return nil, err
	}
	return m, nil
}

// Delete will soft-delete a single models.Project
func (dao *ProjectDAO) Delete(m *models.Project) (error) {
	if err := dao.db.Delete(m).Error; err != nil {
		return err
	}
	return nil
}

// GetUpdatedAfter will return all models.Projects that were
// updated after given timestamp
func (dao *ProjectDAO) GetUpdatedAfter(timestamp time.Time) ([]models.Project, error) {
	m := []models.Project{}
	if err := dao.db.Where("updated_at > ?", timestamp).Find(&m).Error; err != nil {
		return nil, err
	}
	return m, nil
}

// GetAll will return all records of models.Project in database
func (dao *ProjectDAO) GetAll() ([]models.Project, error) {
	m := []models.Project{}
	if err := dao.db.Find(&m).Error; err != nil {
		return nil, err
	}

	return m, nil
}

// ExecuteCustomQueryT will execute a query string
// given by parameter on DB and return the transaction
func (dao *ProjectDAO) ExecuteCustomQueryT(query string) (*gorm.DB, error) {
	retVal := dao.db.Where(query)

	return retVal, retVal.Error
}

// ReadByDescription will find all records
// matching the value given by parameter
func (dao *ProjectDAO) ReadByDescription(m string) ([]models.Project, error) {
	retVal := []models.Project{}
	if err := dao.db.Where(&models.Project{Description: m}).Find(&retVal).Error; err != nil {
		return nil, err
	}

	return retVal, nil
}

// ReadByDescriptionT will return a transaction that
// can be used to find all models matching the value given by parameter
func (dao *ProjectDAO) ReadByDescriptionT(m string) (*gorm.DB, error) {
	retVal := dao.db.Where(&models.Project{Description: m})

	return retVal, retVal.Error
}

// DeleteByDescription deletes all records in database with
// Description the same as parameter given
func (dao *ProjectDAO) DeleteByDescription(m string) (error) {
	if err := dao.db.Where(&models.Project{Description: m}).Delete(&models.Project{}).Error; err != nil {
		return err
	}
	return nil
}

// EditByDescription will edit all records in database
// with the same Description as parameter given
// using model given by parameter
func (dao *ProjectDAO) EditByDescription(m string, newVals *models.Project) (error) {
	if err := dao.db.Table("projects").Where(&models.Project{Description: m}).Updates(newVals).Error; err != nil {
		return err
	}
	return nil
}

// SetDescription will set Description
// to a value given by parameter
func (dao *ProjectDAO) SetDescription(m *models.Project, newVal string) (*models.Project, error) {
	m.Description = newVal
	record, err := dao.ReadByID((m.ID))
	if err != nil {
		return nil, err
	}

	if err := dao.db.Model(&record).Updates(m).Error; err != nil {
		return nil, err
	}

	return record, nil
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

	if err := dao.db.Model(&m).Related(&retVal).Error; err != nil {
		return nil, err
	}
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

	if err := dao.db.Model(&m).Related(&retVal).Error; err != nil {
		return nil, err
	}
	return retVal, nil
}

// SetStart will set a Start property of a model
// to value given by parameter
func (dao *ProjectDAO) SetStart(m *models.Project, str time.Time) (*models.Project, error) {
	m.Start = str
	var err error
	m, err = dao.Update(m, m.ID)
	if err != nil {
		return nil, err
	}

	return m, nil
}

// SetEnd will set a End property of a model
// to value given by parameter
func (dao *ProjectDAO) SetEnd(m *models.Project, str time.Time) (*models.Project, error) {
	m.End = str
	var err error
	m, err = dao.Update(m, m.ID)
	if err != nil {
		return nil, err
	}

	return m, nil
}

// ReadByIsFinished will find all records
// matching the value given by parameter
func (dao *ProjectDAO) ReadByIsFinished(m bool) ([]models.Project, error) {
	retVal := []models.Project{}
	if err := dao.db.Where(&models.Project{IsFinished: m}).Find(&retVal).Error; err != nil {
		return nil, err
	}

	return retVal, nil
}

// ReadByIsFinishedT will return a transaction that
// can be used to find all models matching the value given by parameter
func (dao *ProjectDAO) ReadByIsFinishedT(m bool) (*gorm.DB, error) {
	retVal := dao.db.Where(&models.Project{IsFinished: m})

	return retVal, retVal.Error
}

// DeleteByIsFinished deletes all records in database with
// IsFinished the same as parameter given
func (dao *ProjectDAO) DeleteByIsFinished(m bool) (error) {
	if err := dao.db.Where(&models.Project{IsFinished: m}).Delete(&models.Project{}).Error; err != nil {
		return err
	}
	return nil
}

// EditByIsFinished will edit all records in database
// with the same IsFinished as parameter given
// using model given by parameter
func (dao *ProjectDAO) EditByIsFinished(m bool, newVals *models.Project) (error) {
	if err := dao.db.Table("projects").Where(&models.Project{IsFinished: m}).Updates(newVals).Error; err != nil {
		return err
	}
	return nil
}

// SetIsFinished will set IsFinished
// to a value given by parameter
func (dao *ProjectDAO) SetIsFinished(m *models.Project, newVal bool) (*models.Project, error) {
	m.IsFinished = newVal
	record, err := dao.ReadByID((m.ID))
	if err != nil {
		return nil, err
	}

	if err := dao.db.Model(&record).Updates(m).Error; err != nil {
		return nil, err
	}

	return record, nil
}

// ReadByName will find all records
// matching the value given by parameter
func (dao *ProjectDAO) ReadByName(m string) ([]models.Project, error) {
	retVal := []models.Project{}
	if err := dao.db.Where(&models.Project{Name: m}).Find(&retVal).Error; err != nil {
		return nil, err
	}

	return retVal, nil
}

// ReadByNameT will return a transaction that
// can be used to find all models matching the value given by parameter
func (dao *ProjectDAO) ReadByNameT(m string) (*gorm.DB, error) {
	retVal := dao.db.Where(&models.Project{Name: m})

	return retVal, retVal.Error
}

// DeleteByName deletes all records in database with
// Name the same as parameter given
func (dao *ProjectDAO) DeleteByName(m string) (error) {
	if err := dao.db.Where(&models.Project{Name: m}).Delete(&models.Project{}).Error; err != nil {
		return err
	}
	return nil
}

// EditByName will edit all records in database
// with the same Name as parameter given
// using model given by parameter
func (dao *ProjectDAO) EditByName(m string, newVals *models.Project) (error) {
	if err := dao.db.Table("projects").Where(&models.Project{Name: m}).Updates(newVals).Error; err != nil {
		return err
	}
	return nil
}

// SetName will set Name
// to a value given by parameter
func (dao *ProjectDAO) SetName(m *models.Project, newVal string) (*models.Project, error) {
	m.Name = newVal
	record, err := dao.ReadByID((m.ID))
	if err != nil {
		return nil, err
	}

	if err := dao.db.Model(&record).Updates(m).Error; err != nil {
		return nil, err
	}

	return record, nil
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

// Read will find all DB records matching
// values in a model given by parameter
func (dao *RiskDAO) Read(m *models.Risk) ([]models.Risk, error) {
	retVal := []models.Risk{}
	if err := dao.db.Where(m).Find(&retVal).Error; err != nil {
		return nil, err
	}
	return retVal, nil
}

// ReadT will return a transaction that
// can be used to find DB records matching with models
func (dao *RiskDAO) ReadT(m *models.Risk) (*gorm.DB, error) {
	retVal := dao.db.Where(m)
	return retVal, retVal.Error
}

// ReadByIDT will return a transaction that
// an be used to find DB record with ID given by parameter
func (dao *RiskDAO) ReadByIDT(id uint) (*gorm.DB, error) {
	//m := &models.Risk{}
	retVal := dao.db.Where("ID = ?", id)

	return retVal, retVal.Error
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

// UpdateAllFields will update ALL fields of models.Risk in db
// with values given in the models.Risk by parameter
func (dao *RiskDAO) UpdateAllFields(m *models.Risk) (*models.Risk, error) {
	if err := dao.db.Save(&m).Error; err != nil {
		return nil, err
	}
	return m, nil
}

// Delete will soft-delete a single models.Risk
func (dao *RiskDAO) Delete(m *models.Risk) (error) {
	if err := dao.db.Delete(m).Error; err != nil {
		return err
	}
	return nil
}

// GetUpdatedAfter will return all models.Risks that were
// updated after given timestamp
func (dao *RiskDAO) GetUpdatedAfter(timestamp time.Time) ([]models.Risk, error) {
	m := []models.Risk{}
	if err := dao.db.Where("updated_at > ?", timestamp).Find(&m).Error; err != nil {
		return nil, err
	}
	return m, nil
}

// GetAll will return all records of models.Risk in database
func (dao *RiskDAO) GetAll() ([]models.Risk, error) {
	m := []models.Risk{}
	if err := dao.db.Find(&m).Error; err != nil {
		return nil, err
	}

	return m, nil
}

// ExecuteCustomQueryT will execute a query string
// given by parameter on DB and return the transaction
func (dao *RiskDAO) ExecuteCustomQueryT(query string) (*gorm.DB, error) {
	retVal := dao.db.Where(query)

	return retVal, retVal.Error
}

// ReadByOwner will find all records
// matching the value given by parameter
func (dao *RiskDAO) ReadByOwner(m int) ([]models.Risk, error) {
	retVal := []models.Risk{}
	if err := dao.db.Where(&models.Risk{Owner: m}).Find(&retVal).Error; err != nil {
		return nil, err
	}

	return retVal, nil
}

// ReadByOwnerT will return a transaction that
// can be used to find all models matching the value given by parameter
func (dao *RiskDAO) ReadByOwnerT(m int) (*gorm.DB, error) {
	retVal := dao.db.Where(&models.Risk{Owner: m})

	return retVal, retVal.Error
}

// DeleteByOwner deletes all records in database with
// Owner the same as parameter given
func (dao *RiskDAO) DeleteByOwner(m int) (error) {
	if err := dao.db.Where(&models.Risk{Owner: m}).Delete(&models.Risk{}).Error; err != nil {
		return err
	}
	return nil
}

// EditByOwner will edit all records in database
// with the same Owner as parameter given
// using model given by parameter
func (dao *RiskDAO) EditByOwner(m int, newVals *models.Risk) (error) {
	if err := dao.db.Table("risks").Where(&models.Risk{Owner: m}).Updates(newVals).Error; err != nil {
		return err
	}
	return nil
}

// SetOwner will set Owner
// to a value given by parameter
func (dao *RiskDAO) SetOwner(m *models.Risk, newVal int) (*models.Risk, error) {
	m.Owner = newVal
	record, err := dao.ReadByID((m.ID))
	if err != nil {
		return nil, err
	}

	if err := dao.db.Model(&record).Updates(m).Error; err != nil {
		return nil, err
	}

	return record, nil
}
// ReadByName will find all records
// matching the value given by parameter
func (dao *RiskDAO) ReadByName(m string) ([]models.Risk, error) {
	retVal := []models.Risk{}
	if err := dao.db.Where(&models.Risk{Name: m}).Find(&retVal).Error; err != nil {
		return nil, err
	}

	return retVal, nil
}

// ReadByNameT will return a transaction that
// can be used to find all models matching the value given by parameter
func (dao *RiskDAO) ReadByNameT(m string) (*gorm.DB, error) {
	retVal := dao.db.Where(&models.Risk{Name: m})

	return retVal, retVal.Error
}

// DeleteByName deletes all records in database with
// Name the same as parameter given
func (dao *RiskDAO) DeleteByName(m string) (error) {
	if err := dao.db.Where(&models.Risk{Name: m}).Delete(&models.Risk{}).Error; err != nil {
		return err
	}
	return nil
}

// EditByName will edit all records in database
// with the same Name as parameter given
// using model given by parameter
func (dao *RiskDAO) EditByName(m string, newVals *models.Risk) (error) {
	if err := dao.db.Table("risks").Where(&models.Risk{Name: m}).Updates(newVals).Error; err != nil {
		return err
	}
	return nil
}

// SetName will set Name
// to a value given by parameter
func (dao *RiskDAO) SetName(m *models.Risk, newVal string) (*models.Risk, error) {
	m.Name = newVal
	record, err := dao.ReadByID((m.ID))
	if err != nil {
		return nil, err
	}

	if err := dao.db.Model(&record).Updates(m).Error; err != nil {
		return nil, err
	}

	return record, nil
}

// ReadByThreat will find all records
// matching the value given by parameter
func (dao *RiskDAO) ReadByThreat(m string) ([]models.Risk, error) {
	retVal := []models.Risk{}
	if err := dao.db.Where(&models.Risk{Threat: m}).Find(&retVal).Error; err != nil {
		return nil, err
	}

	return retVal, nil
}

// ReadByThreatT will return a transaction that
// can be used to find all models matching the value given by parameter
func (dao *RiskDAO) ReadByThreatT(m string) (*gorm.DB, error) {
	retVal := dao.db.Where(&models.Risk{Threat: m})

	return retVal, retVal.Error
}

// DeleteByThreat deletes all records in database with
// Threat the same as parameter given
func (dao *RiskDAO) DeleteByThreat(m string) (error) {
	if err := dao.db.Where(&models.Risk{Threat: m}).Delete(&models.Risk{}).Error; err != nil {
		return err
	}
	return nil
}

// EditByThreat will edit all records in database
// with the same Threat as parameter given
// using model given by parameter
func (dao *RiskDAO) EditByThreat(m string, newVals *models.Risk) (error) {
	if err := dao.db.Table("risks").Where(&models.Risk{Threat: m}).Updates(newVals).Error; err != nil {
		return err
	}
	return nil
}

// SetThreat will set Threat
// to a value given by parameter
func (dao *RiskDAO) SetThreat(m *models.Risk, newVal string) (*models.Risk, error) {
	m.Threat = newVal
	record, err := dao.ReadByID((m.ID))
	if err != nil {
		return nil, err
	}

	if err := dao.db.Model(&record).Updates(m).Error; err != nil {
		return nil, err
	}

	return record, nil
}

// ReadByTrigger will find all records
// matching the value given by parameter
func (dao *RiskDAO) ReadByTrigger(m string) ([]models.Risk, error) {
	retVal := []models.Risk{}
	if err := dao.db.Where(&models.Risk{Trigger: m}).Find(&retVal).Error; err != nil {
		return nil, err
	}

	return retVal, nil
}

// ReadByTriggerT will return a transaction that
// can be used to find all models matching the value given by parameter
func (dao *RiskDAO) ReadByTriggerT(m string) (*gorm.DB, error) {
	retVal := dao.db.Where(&models.Risk{Trigger: m})

	return retVal, retVal.Error
}

// DeleteByTrigger deletes all records in database with
// Trigger the same as parameter given
func (dao *RiskDAO) DeleteByTrigger(m string) (error) {
	if err := dao.db.Where(&models.Risk{Trigger: m}).Delete(&models.Risk{}).Error; err != nil {
		return err
	}
	return nil
}

// EditByTrigger will edit all records in database
// with the same Trigger as parameter given
// using model given by parameter
func (dao *RiskDAO) EditByTrigger(m string, newVals *models.Risk) (error) {
	if err := dao.db.Table("risks").Where(&models.Risk{Trigger: m}).Updates(newVals).Error; err != nil {
		return err
	}
	return nil
}

// SetTrigger will set Trigger
// to a value given by parameter
func (dao *RiskDAO) SetTrigger(m *models.Risk, newVal string) (*models.Risk, error) {
	m.Trigger = newVal
	record, err := dao.ReadByID((m.ID))
	if err != nil {
		return nil, err
	}

	if err := dao.db.Model(&record).Updates(m).Error; err != nil {
		return nil, err
	}

	return record, nil
}

// SetStart will set a Start property of a model
// to value given by parameter
func (dao *RiskDAO) SetStart(m *models.Risk, str time.Time) (*models.Risk, error) {
	m.Start = str
	var err error
	m, err = dao.Update(m, m.ID)
	if err != nil {
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

	if err := dao.db.Model(&m).Related(&retVal).Error; err != nil {
		return nil, err
	}
	return retVal, nil
}

// ReadByValue will find all records
// matching the value given by parameter
func (dao *RiskDAO) ReadByValue(m int) ([]models.Risk, error) {
	retVal := []models.Risk{}
	if err := dao.db.Where(&models.Risk{Value: m}).Find(&retVal).Error; err != nil {
		return nil, err
	}

	return retVal, nil
}

// ReadByValueT will return a transaction that
// can be used to find all models matching the value given by parameter
func (dao *RiskDAO) ReadByValueT(m int) (*gorm.DB, error) {
	retVal := dao.db.Where(&models.Risk{Value: m})

	return retVal, retVal.Error
}

// DeleteByValue deletes all records in database with
// Value the same as parameter given
func (dao *RiskDAO) DeleteByValue(m int) (error) {
	if err := dao.db.Where(&models.Risk{Value: m}).Delete(&models.Risk{}).Error; err != nil {
		return err
	}
	return nil
}

// EditByValue will edit all records in database
// with the same Value as parameter given
// using model given by parameter
func (dao *RiskDAO) EditByValue(m int, newVals *models.Risk) (error) {
	if err := dao.db.Table("risks").Where(&models.Risk{Value: m}).Updates(newVals).Error; err != nil {
		return err
	}
	return nil
}

// SetValue will set Value
// to a value given by parameter
func (dao *RiskDAO) SetValue(m *models.Risk, newVal int) (*models.Risk, error) {
	m.Value = newVal
	record, err := dao.ReadByID((m.ID))
	if err != nil {
		return nil, err
	}

	if err := dao.db.Model(&record).Updates(m).Error; err != nil {
		return nil, err
	}

	return record, nil
}

// ReadByDescription will find all records
// matching the value given by parameter
func (dao *RiskDAO) ReadByDescription(m string) ([]models.Risk, error) {
	retVal := []models.Risk{}
	if err := dao.db.Where(&models.Risk{Description: m}).Find(&retVal).Error; err != nil {
		return nil, err
	}

	return retVal, nil
}

// ReadByDescriptionT will return a transaction that
// can be used to find all models matching the value given by parameter
func (dao *RiskDAO) ReadByDescriptionT(m string) (*gorm.DB, error) {
	retVal := dao.db.Where(&models.Risk{Description: m})

	return retVal, retVal.Error
}

// DeleteByDescription deletes all records in database with
// Description the same as parameter given
func (dao *RiskDAO) DeleteByDescription(m string) (error) {
	if err := dao.db.Where(&models.Risk{Description: m}).Delete(&models.Risk{}).Error; err != nil {
		return err
	}
	return nil
}

// EditByDescription will edit all records in database
// with the same Description as parameter given
// using model given by parameter
func (dao *RiskDAO) EditByDescription(m string, newVals *models.Risk) (error) {
	if err := dao.db.Table("risks").Where(&models.Risk{Description: m}).Updates(newVals).Error; err != nil {
		return err
	}
	return nil
}

// SetDescription will set Description
// to a value given by parameter
func (dao *RiskDAO) SetDescription(m *models.Risk, newVal string) (*models.Risk, error) {
	m.Description = newVal
	record, err := dao.ReadByID((m.ID))
	if err != nil {
		return nil, err
	}

	if err := dao.db.Model(&record).Updates(m).Error; err != nil {
		return nil, err
	}

	return record, nil
}

// ReadByStatus will find all records
// matching the value given by parameter
func (dao *RiskDAO) ReadByStatus(m string) ([]models.Risk, error) {
	retVal := []models.Risk{}
	if err := dao.db.Where(&models.Risk{Status: m}).Find(&retVal).Error; err != nil {
		return nil, err
	}

	return retVal, nil
}

// ReadByStatusT will return a transaction that
// can be used to find all models matching the value given by parameter
func (dao *RiskDAO) ReadByStatusT(m string) (*gorm.DB, error) {
	retVal := dao.db.Where(&models.Risk{Status: m})

	return retVal, retVal.Error
}

// DeleteByStatus deletes all records in database with
// Status the same as parameter given
func (dao *RiskDAO) DeleteByStatus(m string) (error) {
	if err := dao.db.Where(&models.Risk{Status: m}).Delete(&models.Risk{}).Error; err != nil {
		return err
	}
	return nil
}

// EditByStatus will edit all records in database
// with the same Status as parameter given
// using model given by parameter
func (dao *RiskDAO) EditByStatus(m string, newVals *models.Risk) (error) {
	if err := dao.db.Table("risks").Where(&models.Risk{Status: m}).Updates(newVals).Error; err != nil {
		return err
	}
	return nil
}

// SetStatus will set Status
// to a value given by parameter
func (dao *RiskDAO) SetStatus(m *models.Risk, newVal string) (*models.Risk, error) {
	m.Status = newVal
	record, err := dao.ReadByID((m.ID))
	if err != nil {
		return nil, err
	}

	if err := dao.db.Model(&record).Updates(m).Error; err != nil {
		return nil, err
	}

	return record, nil
}

// SetEnd will set a End property of a model
// to value given by parameter
func (dao *RiskDAO) SetEnd(m *models.Risk, str time.Time) (*models.Risk, error) {
	m.End = str
	var err error
	m, err = dao.Update(m, m.ID)
	if err != nil {
		return nil, err
	}

	return m, nil
}

// ReadByCost will find all records
// matching the value given by parameter
func (dao *RiskDAO) ReadByCost(m int) ([]models.Risk, error) {
	retVal := []models.Risk{}
	if err := dao.db.Where(&models.Risk{Cost: m}).Find(&retVal).Error; err != nil {
		return nil, err
	}

	return retVal, nil
}

// ReadByCostT will return a transaction that
// can be used to find all models matching the value given by parameter
func (dao *RiskDAO) ReadByCostT(m int) (*gorm.DB, error) {
	retVal := dao.db.Where(&models.Risk{Cost: m})

	return retVal, retVal.Error
}

// DeleteByCost deletes all records in database with
// Cost the same as parameter given
func (dao *RiskDAO) DeleteByCost(m int) (error) {
	if err := dao.db.Where(&models.Risk{Cost: m}).Delete(&models.Risk{}).Error; err != nil {
		return err
	}
	return nil
}

// EditByCost will edit all records in database
// with the same Cost as parameter given
// using model given by parameter
func (dao *RiskDAO) EditByCost(m int, newVals *models.Risk) (error) {
	if err := dao.db.Table("risks").Where(&models.Risk{Cost: m}).Updates(newVals).Error; err != nil {
		return err
	}
	return nil
}

// SetCost will set Cost
// to a value given by parameter
func (dao *RiskDAO) SetCost(m *models.Risk, newVal int) (*models.Risk, error) {
	m.Cost = newVal
	record, err := dao.ReadByID((m.ID))
	if err != nil {
		return nil, err
	}

	if err := dao.db.Model(&record).Updates(m).Error; err != nil {
		return nil, err
	}

	return record, nil
}


// ReadByProbability will find all records
// matching the value given by parameter
func (dao *RiskDAO) ReadByProbability(m float64) ([]models.Risk, error) {
	retVal := []models.Risk{}
	if err := dao.db.Where(&models.Risk{Probability: m}).Find(&retVal).Error; err != nil {
		return nil, err
	}

	return retVal, nil
}

// ReadByProbabilityT will return a transaction that
// can be used to find all models matching the value given by parameter
func (dao *RiskDAO) ReadByProbabilityT(m float64) (*gorm.DB, error) {
	retVal := dao.db.Where(&models.Risk{Probability: m})

	return retVal, retVal.Error
}

// DeleteByProbability deletes all records in database with
// Probability the same as parameter given
func (dao *RiskDAO) DeleteByProbability(m float64) (error) {
	if err := dao.db.Where(&models.Risk{Probability: m}).Delete(&models.Risk{}).Error; err != nil {
		return err
	}
	return nil
}

// EditByProbability will edit all records in database
// with the same Probability as parameter given
// using model given by parameter
func (dao *RiskDAO) EditByProbability(m float64, newVals *models.Risk) (error) {
	if err := dao.db.Table("risks").Where(&models.Risk{Probability: m}).Updates(newVals).Error; err != nil {
		return err
	}
	return nil
}

// SetProbability will set Probability
// to a value given by parameter
func (dao *RiskDAO) SetProbability(m *models.Risk, newVal float64) (*models.Risk, error) {
	m.Probability = newVal
	record, err := dao.ReadByID((m.ID))
	if err != nil {
		return nil, err
	}

	if err := dao.db.Model(&record).Updates(m).Error; err != nil {
		return nil, err
	}

	return record, nil
}


// ReadByCategory will find all records
// matching the value given by parameter
func (dao *RiskDAO) ReadByCategory(m string) ([]models.Risk, error) {
	retVal := []models.Risk{}
	if err := dao.db.Where(&models.Risk{Category: m}).Find(&retVal).Error; err != nil {
		return nil, err
	}

	return retVal, nil
}

// ReadByCategoryT will return a transaction that
// can be used to find all models matching the value given by parameter
func (dao *RiskDAO) ReadByCategoryT(m string) (*gorm.DB, error) {
	retVal := dao.db.Where(&models.Risk{Category: m})

	return retVal, retVal.Error
}

// DeleteByCategory deletes all records in database with
// Category the same as parameter given
func (dao *RiskDAO) DeleteByCategory(m string) (error) {
	if err := dao.db.Where(&models.Risk{Category: m}).Delete(&models.Risk{}).Error; err != nil {
		return err
	}
	return nil
}

// EditByCategory will edit all records in database
// with the same Category as parameter given
// using model given by parameter
func (dao *RiskDAO) EditByCategory(m string, newVals *models.Risk) (error) {
	if err := dao.db.Table("risks").Where(&models.Risk{Category: m}).Updates(newVals).Error; err != nil {
		return err
	}
	return nil
}

// SetCategory will set Category
// to a value given by parameter
func (dao *RiskDAO) SetCategory(m *models.Risk, newVal string) (*models.Risk, error) {
	m.Category = newVal
	record, err := dao.ReadByID((m.ID))
	if err != nil {
		return nil, err
	}

	if err := dao.db.Model(&record).Updates(m).Error; err != nil {
		return nil, err
	}

	return record, nil
}

// AddCounterMeasuresAssociation will add
// an association to model given by parameter
func (dao *RiskDAO) AddCounterMeasuresAssociation(m *models.Risk, asocVal *models.CounterMeasure) (*models.Risk, error) {
	if err := dao.db.Model(&m).Association("CounterMeasures").Append(asocVal).Error; err != nil {
		return nil, err
	}

	return m, nil
}

// RemoveCounterMeasuresAssociation will remove
// an association from model given by parameter
func (dao *RiskDAO) RemoveCounterMeasuresAssociation(m *models.Risk, asocVal *models.CounterMeasure) (*models.Risk, error) {
	if err := dao.db.Model(&m).Association("CounterMeasures").Delete(asocVal).Error; err != nil {
		return nil, err
	}

	return m, nil
}

// GetAllAssociatedCounterMeasures will get all
// an association from model given by parameter
func (dao *RiskDAO) GetAllAssociatedCounterMeasures(m *models.Risk) ([]models.CounterMeasure, error) {
	retVal := []models.CounterMeasure{}

	if err := dao.db.Model(&m).Related(&retVal).Error; err != nil {
		return nil, err
	}
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

// Read will find all DB records matching
// values in a model given by parameter
func (dao *CounterMeasureDAO) Read(m *models.CounterMeasure) ([]models.CounterMeasure, error) {
	retVal := []models.CounterMeasure{}
	if err := dao.db.Where(m).Find(&retVal).Error; err != nil {
		return nil, err
	}
	return retVal, nil
}

// ReadT will return a transaction that
// can be used to find DB records matching with models
func (dao *CounterMeasureDAO) ReadT(m *models.CounterMeasure) (*gorm.DB, error) {
	retVal := dao.db.Where(m)
	return retVal, retVal.Error
}

// ReadByIDT will return a transaction that
// an be used to find DB record with ID given by parameter
func (dao *CounterMeasureDAO) ReadByIDT(id uint) (*gorm.DB, error) {
	//m := &models.CounterMeasure{}
	retVal := dao.db.Where("ID = ?", id)

	return retVal, retVal.Error
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

// UpdateAllFields will update ALL fields of models.CounterMeasure in db
// with values given in the models.CounterMeasure by parameter
func (dao *CounterMeasureDAO) UpdateAllFields(m *models.CounterMeasure) (*models.CounterMeasure, error) {
	if err := dao.db.Save(&m).Error; err != nil {
		return nil, err
	}
	return m, nil
}

// Delete will soft-delete a single models.CounterMeasure
func (dao *CounterMeasureDAO) Delete(m *models.CounterMeasure) (error) {
	if err := dao.db.Delete(m).Error; err != nil {
		return err
	}
	return nil
}

// GetUpdatedAfter will return all models.CounterMeasures that were
// updated after given timestamp
func (dao *CounterMeasureDAO) GetUpdatedAfter(timestamp time.Time) ([]models.CounterMeasure, error) {
	m := []models.CounterMeasure{}
	if err := dao.db.Where("updated_at > ?", timestamp).Find(&m).Error; err != nil {
		return nil, err
	}
	return m, nil
}

// GetAll will return all records of models.CounterMeasure in database
func (dao *CounterMeasureDAO) GetAll() ([]models.CounterMeasure, error) {
	m := []models.CounterMeasure{}
	if err := dao.db.Find(&m).Error; err != nil {
		return nil, err
	}

	return m, nil
}

// ExecuteCustomQueryT will execute a query string
// given by parameter on DB and return the transaction
func (dao *CounterMeasureDAO) ExecuteCustomQueryT(query string) (*gorm.DB, error) {
	retVal := dao.db.Where(query)

	return retVal, retVal.Error
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

	if err := dao.db.Model(&m).Related(&retVal).Error; err != nil {
		return nil, err
	}
	return retVal, nil
}

// ReadByName will find all records
// matching the value given by parameter
func (dao *CounterMeasureDAO) ReadByName(m string) ([]models.CounterMeasure, error) {
	retVal := []models.CounterMeasure{}
	if err := dao.db.Where(&models.CounterMeasure{Name: m}).Find(&retVal).Error; err != nil {
		return nil, err
	}

	return retVal, nil
}

// ReadByNameT will return a transaction that
// can be used to find all models matching the value given by parameter
func (dao *CounterMeasureDAO) ReadByNameT(m string) (*gorm.DB, error) {
	retVal := dao.db.Where(&models.CounterMeasure{Name: m})

	return retVal, retVal.Error
}

// DeleteByName deletes all records in database with
// Name the same as parameter given
func (dao *CounterMeasureDAO) DeleteByName(m string) (error) {
	if err := dao.db.Where(&models.CounterMeasure{Name: m}).Delete(&models.CounterMeasure{}).Error; err != nil {
		return err
	}
	return nil
}

// EditByName will edit all records in database
// with the same Name as parameter given
// using model given by parameter
func (dao *CounterMeasureDAO) EditByName(m string, newVals *models.CounterMeasure) (error) {
	if err := dao.db.Table("counter_measures").Where(&models.CounterMeasure{Name: m}).Updates(newVals).Error; err != nil {
		return err
	}
	return nil
}

// SetName will set Name
// to a value given by parameter
func (dao *CounterMeasureDAO) SetName(m *models.CounterMeasure, newVal string) (*models.CounterMeasure, error) {
	m.Name = newVal
	record, err := dao.ReadByID((m.ID))
	if err != nil {
		return nil, err
	}

	if err := dao.db.Model(&record).Updates(m).Error; err != nil {
		return nil, err
	}

	return record, nil
}

// ReadByDescription will find all records
// matching the value given by parameter
func (dao *CounterMeasureDAO) ReadByDescription(m string) ([]models.CounterMeasure, error) {
	retVal := []models.CounterMeasure{}
	if err := dao.db.Where(&models.CounterMeasure{Description: m}).Find(&retVal).Error; err != nil {
		return nil, err
	}

	return retVal, nil
}

// ReadByDescriptionT will return a transaction that
// can be used to find all models matching the value given by parameter
func (dao *CounterMeasureDAO) ReadByDescriptionT(m string) (*gorm.DB, error) {
	retVal := dao.db.Where(&models.CounterMeasure{Description: m})

	return retVal, retVal.Error
}

// DeleteByDescription deletes all records in database with
// Description the same as parameter given
func (dao *CounterMeasureDAO) DeleteByDescription(m string) (error) {
	if err := dao.db.Where(&models.CounterMeasure{Description: m}).Delete(&models.CounterMeasure{}).Error; err != nil {
		return err
	}
	return nil
}

// EditByDescription will edit all records in database
// with the same Description as parameter given
// using model given by parameter
func (dao *CounterMeasureDAO) EditByDescription(m string, newVals *models.CounterMeasure) (error) {
	if err := dao.db.Table("counter_measures").Where(&models.CounterMeasure{Description: m}).Updates(newVals).Error; err != nil {
		return err
	}
	return nil
}

// SetDescription will set Description
// to a value given by parameter
func (dao *CounterMeasureDAO) SetDescription(m *models.CounterMeasure, newVal string) (*models.CounterMeasure, error) {
	m.Description = newVal
	record, err := dao.ReadByID((m.ID))
	if err != nil {
		return nil, err
	}

	if err := dao.db.Model(&record).Updates(m).Error; err != nil {
		return nil, err
	}

	return record, nil
}

// ReadByCost will find all records
// matching the value given by parameter
func (dao *CounterMeasureDAO) ReadByCost(m int) ([]models.CounterMeasure, error) {
	retVal := []models.CounterMeasure{}
	if err := dao.db.Where(&models.CounterMeasure{Cost: m}).Find(&retVal).Error; err != nil {
		return nil, err
	}

	return retVal, nil
}

// ReadByCostT will return a transaction that
// can be used to find all models matching the value given by parameter
func (dao *CounterMeasureDAO) ReadByCostT(m int) (*gorm.DB, error) {
	retVal := dao.db.Where(&models.CounterMeasure{Cost: m})

	return retVal, retVal.Error
}

// DeleteByCost deletes all records in database with
// Cost the same as parameter given
func (dao *CounterMeasureDAO) DeleteByCost(m int) (error) {
	if err := dao.db.Where(&models.CounterMeasure{Cost: m}).Delete(&models.CounterMeasure{}).Error; err != nil {
		return err
	}
	return nil
}

// EditByCost will edit all records in database
// with the same Cost as parameter given
// using model given by parameter
func (dao *CounterMeasureDAO) EditByCost(m int, newVals *models.CounterMeasure) (error) {
	if err := dao.db.Table("counter_measures").Where(&models.CounterMeasure{Cost: m}).Updates(newVals).Error; err != nil {
		return err
	}
	return nil
}

// SetCost will set Cost
// to a value given by parameter
func (dao *CounterMeasureDAO) SetCost(m *models.CounterMeasure, newVal int) (*models.CounterMeasure, error) {
	m.Cost = newVal
	record, err := dao.ReadByID((m.ID))
	if err != nil {
		return nil, err
	}

	if err := dao.db.Model(&record).Updates(m).Error; err != nil {
		return nil, err
	}

	return record, nil
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

// ReadT will return a transaction that
// can be used to find DB records matching with models
func (dao *UserDAO) ReadT(m *models.User) (*gorm.DB, error) {
	retVal := dao.db.Where(m)
	return retVal, retVal.Error
}

// ReadByIDT will return a transaction that
// an be used to find DB record with ID given by parameter
func (dao *UserDAO) ReadByIDT(id uint) (*gorm.DB, error) {
	//m := &models.User{}
	retVal := dao.db.Where("ID = ?", id)

	return retVal, retVal.Error
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

// UpdateAllFields will update ALL fields of models.User in db
// with values given in the models.User by parameter
func (dao *UserDAO) UpdateAllFields(m *models.User) (*models.User, error) {
	if err := dao.db.Save(&m).Error; err != nil {
		return nil, err
	}
	return m, nil
}

// Delete will soft-delete a single models.User
func (dao *UserDAO) Delete(m *models.User) (error) {
	if err := dao.db.Delete(m).Error; err != nil {
		return err
	}
	return nil
}

// GetUpdatedAfter will return all models.Users that were
// updated after given timestamp
func (dao *UserDAO) GetUpdatedAfter(timestamp time.Time) ([]models.User, error) {
	m := []models.User{}
	if err := dao.db.Where("updated_at > ?", timestamp).Find(&m).Error; err != nil {
		return nil, err
	}
	return m, nil
}

// GetAll will return all records of models.User in database
func (dao *UserDAO) GetAll() ([]models.User, error) {
	m := []models.User{}
	if err := dao.db.Find(&m).Error; err != nil {
		return nil, err
	}

	return m, nil
}

// ExecuteCustomQueryT will execute a query string
// given by parameter on DB and return the transaction
func (dao *UserDAO) ExecuteCustomQueryT(query string) (*gorm.DB, error) {
	retVal := dao.db.Where(query)

	return retVal, retVal.Error
}


// ReadByName will find all records
// matching the value given by parameter
func (dao *UserDAO) ReadByName(m string) ([]models.User, error) {
	retVal := []models.User{}
	if err := dao.db.Where(&models.User{Name: m}).Find(&retVal).Error; err != nil {
		return nil, err
	}

	return retVal, nil
}

// ReadByNameT will return a transaction that
// can be used to find all models matching the value given by parameter
func (dao *UserDAO) ReadByNameT(m string) (*gorm.DB, error) {
	retVal := dao.db.Where(&models.User{Name: m})

	return retVal, retVal.Error
}

// DeleteByName deletes all records in database with
// Name the same as parameter given
func (dao *UserDAO) DeleteByName(m string) (error) {
	if err := dao.db.Where(&models.User{Name: m}).Delete(&models.User{}).Error; err != nil {
		return err
	}
	return nil
}

// EditByName will edit all records in database
// with the same Name as parameter given
// using model given by parameter
func (dao *UserDAO) EditByName(m string, newVals *models.User) (error) {
	if err := dao.db.Table("users").Where(&models.User{Name: m}).Updates(newVals).Error; err != nil {
		return err
	}
	return nil
}

// SetName will set Name
// to a value given by parameter
func (dao *UserDAO) SetName(m *models.User, newVal string) (*models.User, error) {
	m.Name = newVal
	record, err := dao.ReadByID((m.ID))
	if err != nil {
		return nil, err
	}

	if err := dao.db.Model(&record).Updates(m).Error; err != nil {
		return nil, err
	}

	return record, nil
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

// ReadByEmailT will return a transaction that
// can be used to find all models matching the value given by parameter
func (dao *UserDAO) ReadByEmailT(m string) (*gorm.DB, error) {
	retVal := dao.db.Where(&models.User{Email: m})

	return retVal, retVal.Error
}

// DeleteByEmail deletes all records in database with
// Email the same as parameter given
func (dao *UserDAO) DeleteByEmail(m string) (error) {
	if err := dao.db.Where(&models.User{Email: m}).Delete(&models.User{}).Error; err != nil {
		return err
	}
	return nil
}

// EditByEmail will edit all records in database
// with the same Email as parameter given
// using model given by parameter
func (dao *UserDAO) EditByEmail(m string, newVals *models.User) (error) {
	if err := dao.db.Table("users").Where(&models.User{Email: m}).Updates(newVals).Error; err != nil {
		return err
	}
	return nil
}

// SetEmail will set Email
// to a value given by parameter
func (dao *UserDAO) SetEmail(m *models.User, newVal string) (*models.User, error) {
	m.Email = newVal
	record, err := dao.ReadByID((m.ID))
	if err != nil {
		return nil, err
	}

	if err := dao.db.Model(&record).Updates(m).Error; err != nil {
		return nil, err
	}

	return record, nil
}

// ReadByPassword will find all records
// matching the value given by parameter
func (dao *UserDAO) ReadByPassword(m string) ([]models.User, error) {
	retVal := []models.User{}
	if err := dao.db.Where(&models.User{Password: m}).Find(&retVal).Error; err != nil {
		return nil, err
	}

	return retVal, nil
}

// ReadByPasswordT will return a transaction that
// can be used to find all models matching the value given by parameter
func (dao *UserDAO) ReadByPasswordT(m string) (*gorm.DB, error) {
	retVal := dao.db.Where(&models.User{Password: m})

	return retVal, retVal.Error
}

// DeleteByPassword deletes all records in database with
// Password the same as parameter given
func (dao *UserDAO) DeleteByPassword(m string) (error) {
	if err := dao.db.Where(&models.User{Password: m}).Delete(&models.User{}).Error; err != nil {
		return err
	}
	return nil
}

// EditByPassword will edit all records in database
// with the same Password as parameter given
// using model given by parameter
func (dao *UserDAO) EditByPassword(m string, newVals *models.User) (error) {
	if err := dao.db.Table("users").Where(&models.User{Password: m}).Updates(newVals).Error; err != nil {
		return err
	}
	return nil
}

// SetPassword will set Password
// to a value given by parameter
func (dao *UserDAO) SetPassword(m *models.User, newVal string) (*models.User, error) {
	m.Password = newVal
	record, err := dao.ReadByID((m.ID))
	if err != nil {
		return nil, err
	}

	if err := dao.db.Model(&record).Updates(m).Error; err != nil {
		return nil, err
	}

	return record, nil
}


// ReadByRole will find all records
// matching the value given by parameter
func (dao *UserDAO) ReadByRole(m int) ([]models.User, error) {
	retVal := []models.User{}
	if err := dao.db.Where(&models.User{Role: m}).Find(&retVal).Error; err != nil {
		return nil, err
	}

	return retVal, nil
}

// ReadByRoleT will return a transaction that
// can be used to find all models matching the value given by parameter
func (dao *UserDAO) ReadByRoleT(m int) (*gorm.DB, error) {
	retVal := dao.db.Where(&models.User{Role: m})

	return retVal, retVal.Error
}

// DeleteByRole deletes all records in database with
// Role the same as parameter given
func (dao *UserDAO) DeleteByRole(m int) (error) {
	if err := dao.db.Where(&models.User{Role: m}).Delete(&models.User{}).Error; err != nil {
		return err
	}
	return nil
}

// EditByRole will edit all records in database
// with the same Role as parameter given
// using model given by parameter
func (dao *UserDAO) EditByRole(m int, newVals *models.User) (error) {
	if err := dao.db.Table("users").Where(&models.User{Role: m}).Updates(newVals).Error; err != nil {
		return err
	}
	return nil
}

// SetRole will set Role
// to a value given by parameter
func (dao *UserDAO) SetRole(m *models.User, newVal int) (*models.User, error) {
	m.Role = newVal
	record, err := dao.ReadByID((m.ID))
	if err != nil {
		return nil, err
	}

	if err := dao.db.Model(&record).Updates(m).Error; err != nil {
		return nil, err
	}

	return record, nil
}

// ReadBySkills will find all records
// matching the value given by parameter
func (dao *UserDAO) ReadBySkills(m string) ([]models.User, error) {
	retVal := []models.User{}
	if err := dao.db.Where(&models.User{Skills: m}).Find(&retVal).Error; err != nil {
		return nil, err
	}

	return retVal, nil
}

// ReadBySkillsT will return a transaction that
// can be used to find all models matching the value given by parameter
func (dao *UserDAO) ReadBySkillsT(m string) (*gorm.DB, error) {
	retVal := dao.db.Where(&models.User{Skills: m})

	return retVal, retVal.Error
}

// DeleteBySkills deletes all records in database with
// Skills the same as parameter given
func (dao *UserDAO) DeleteBySkills(m string) (error) {
	if err := dao.db.Where(&models.User{Skills: m}).Delete(&models.User{}).Error; err != nil {
		return err
	}
	return nil
}

// EditBySkills will edit all records in database
// with the same Skills as parameter given
// using model given by parameter
func (dao *UserDAO) EditBySkills(m string, newVals *models.User) (error) {
	if err := dao.db.Table("users").Where(&models.User{Skills: m}).Updates(newVals).Error; err != nil {
		return err
	}
	return nil
}

// SetSkills will set Skills
// to a value given by parameter
func (dao *UserDAO) SetSkills(m *models.User, newVal string) (*models.User, error) {
	m.Skills = newVal
	record, err := dao.ReadByID((m.ID))
	if err != nil {
		return nil, err
	}

	if err := dao.db.Model(&record).Updates(m).Error; err != nil {
		return nil, err
	}

	return record, nil
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

	if err := dao.db.Model(&m).Related(&retVal).Error; err != nil {
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