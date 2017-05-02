package access

import(
  "github.com/jinzhu/gorm"
  "time"
  "github.com/wscherfel/fitlogic-backend/models"
  )


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

// Read is a mock implementation of Read method
func (mock *UserDAOMock) Read(m *models.User) ([]models.User, error) {
	ret := make([]models.User, 0, len(mock.db))
	for range mock.db {
		/*if val == *m {
			ret = append(ret, val)
		}*/
	}

	return ret, nil
}


// ReadByID is a mock implementation of ReadByID method
func (mock *UserDAOMock) ReadByID(id uint) (*models.User, error) {
	ret := mock.db[id]
	return &ret, nil
}

// ReadByIDT is a mock implementation of ReadByIDT method
func (mock *UserDAOMock) ReadByIDT(id uint) (*gorm.DB, error) {
	return nil, nil
}

// ReadT is a mock implementation of ReadT method
func (mock *UserDAOMock) ReadT(m *models.User) (*gorm.DB, error) {
	return nil, nil
}

// Update is a mock implementation of Update method
func (mock *UserDAOMock) Update(m *models.User, id uint) (*models.User, error) {
	m.UpdatedAt = time.Now()
	mock.db[id] = *m
	return m, nil
}

// UpdateAllFields is a mock implementation of UpdateAllFields method
func (mock *UserDAOMock) UpdateAllFields(m *models.User) (*models.User, error) {
	m.UpdatedAt = time.Now()
	mock.db[m.ID] = *m
	return m, nil
}

// Delete is a mock implementation of Delete method
func (mock *UserDAOMock) Delete(m *models.User) (error) {
	delete(mock.db, m.ID)
	return nil
}

// GetUpdatedAfter is a mock implementation of GetUpdatedAfter method
func (mock *UserDAOMock) GetUpdatedAfter(timestamp time.Time) ([]models.User, error) {
	ret := make([]models.User, 0, len(mock.db))
	for _, val :=  range mock.db {
		if val.UpdatedAt.After(timestamp) {
			ret = append(ret, val)
		}
	}

	return ret, nil
}

// GetAll is a mock implementation of GetAll method
func (mock *UserDAOMock) GetAll() ([]models.User, error) {
	ret := make([]models.User, 0, len(mock.db))
	for _, val := range mock.db {
		ret = append(ret, val)
	}

	return ret, nil
}

// ExecuteCustomQueryT is a mock implementation of ExecuteCustomQueryT method
func (mock *UserDAOMock) ExecuteCustomQueryT(query string) (*gorm.DB, error) {
	return nil, nil
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

// ReadByName is a mock implementation of ReadByName method
func (mock *UserDAOMock) ReadByName(m string) ([]models.User, error) {
	ret := make([]models.User, 0, len(mock.db))
	for _, val := range mock.db {
		if val.Name == m {
			ret = append(ret, val)
		}
	}

	return ret, nil
}

// ReadByNameT is a mock implementation of ReadByNameT method
func (mock *UserDAOMock) ReadByNameT(m string) (*gorm.DB, error) {
	return nil, nil
}

// DeleteByName is a mock implementation of DeleteByName method
func (mock *UserDAOMock) DeleteByName(m string) (error) {
	for _, val := range mock.db {
		if val.Name == m {
			delete(mock.db, val.ID)
		}
	}

	return nil
}

// EditByName is a mock implementation of EditByName method
func (mock *UserDAOMock) EditByName(m string, newVals *models.User) (error) {
	for _, val := range mock.db {
		if val.Name == m {
			id := val.ID
			val = *newVals
			val.ID = id
			val.UpdatedAt = time.Now()
		}
	}

	return nil
}

// SetName is a mock implementation of SetName method
func (mock *UserDAOMock) SetName(m *models.User, newVal string) (*models.User, error) {
	edit := mock.db[m.ID]
	edit.Name = newVal
	edit.UpdatedAt = time.Now()

	mock.db[m.ID] = edit
	return &edit, nil
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

// ReadByEmail is a mock implementation of ReadByEmail method
func (mock *UserDAOMock) ReadByEmail(m string) ([]models.User, error) {
	ret := make([]models.User, 0, len(mock.db))
	for _, val := range mock.db {
		if val.Email == m {
			ret = append(ret, val)
		}
	}

	return ret, nil
}

// ReadByEmailT is a mock implementation of ReadByEmailT method
func (mock *UserDAOMock) ReadByEmailT(m string) (*gorm.DB, error) {
	return nil, nil
}

// DeleteByEmail is a mock implementation of DeleteByEmail method
func (mock *UserDAOMock) DeleteByEmail(m string) (error) {
	for _, val := range mock.db {
		if val.Email == m {
			delete(mock.db, val.ID)
		}
	}

	return nil
}

// EditByEmail is a mock implementation of EditByEmail method
func (mock *UserDAOMock) EditByEmail(m string, newVals *models.User) (error) {
	for _, val := range mock.db {
		if val.Email == m {
			id := val.ID
			val = *newVals
			val.ID = id
			val.UpdatedAt = time.Now()
		}
	}

	return nil
}

// SetEmail is a mock implementation of SetEmail method
func (mock *UserDAOMock) SetEmail(m *models.User, newVal string) (*models.User, error) {
	edit := mock.db[m.ID]
	edit.Email = newVal
	edit.UpdatedAt = time.Now()

	mock.db[m.ID] = edit
	return &edit, nil
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

// ReadByPassword is a mock implementation of ReadByPassword method
func (mock *UserDAOMock) ReadByPassword(m string) ([]models.User, error) {
	ret := make([]models.User, 0, len(mock.db))
	for _, val := range mock.db {
		if val.Password == m {
			ret = append(ret, val)
		}
	}

	return ret, nil
}

// ReadByPasswordT is a mock implementation of ReadByPasswordT method
func (mock *UserDAOMock) ReadByPasswordT(m string) (*gorm.DB, error) {
	return nil, nil
}

// DeleteByPassword is a mock implementation of DeleteByPassword method
func (mock *UserDAOMock) DeleteByPassword(m string) (error) {
	for _, val := range mock.db {
		if val.Password == m {
			delete(mock.db, val.ID)
		}
	}

	return nil
}

// EditByPassword is a mock implementation of EditByPassword method
func (mock *UserDAOMock) EditByPassword(m string, newVals *models.User) (error) {
	for _, val := range mock.db {
		if val.Password == m {
			id := val.ID
			val = *newVals
			val.ID = id
			val.UpdatedAt = time.Now()
		}
	}

	return nil
}

// SetPassword is a mock implementation of SetPassword method
func (mock *UserDAOMock) SetPassword(m *models.User, newVal string) (*models.User, error) {
	edit := mock.db[m.ID]
	edit.Password = newVal
	edit.UpdatedAt = time.Now()

	mock.db[m.ID] = edit
	return &edit, nil
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

// ReadByRole is a mock implementation of ReadByRole method
func (mock *UserDAOMock) ReadByRole(m int) ([]models.User, error) {
	ret := make([]models.User, 0, len(mock.db))
	for _, val := range mock.db {
		if val.Role == m {
			ret = append(ret, val)
		}
	}

	return ret, nil
}

// ReadByRoleT is a mock implementation of ReadByRoleT method
func (mock *UserDAOMock) ReadByRoleT(m int) (*gorm.DB, error) {
	return nil, nil
}

// DeleteByRole is a mock implementation of DeleteByRole method
func (mock *UserDAOMock) DeleteByRole(m int) (error) {
	for _, val := range mock.db {
		if val.Role == m {
			delete(mock.db, val.ID)
		}
	}

	return nil
}

// EditByRole is a mock implementation of EditByRole method
func (mock *UserDAOMock) EditByRole(m int, newVals *models.User) (error) {
	for _, val := range mock.db {
		if val.Role == m {
			id := val.ID
			val = *newVals
			val.ID = id
			val.UpdatedAt = time.Now()
		}
	}

	return nil
}

// SetRole is a mock implementation of SetRole method
func (mock *UserDAOMock) SetRole(m *models.User, newVal int) (*models.User, error) {
	edit := mock.db[m.ID]
	edit.Role = newVal
	edit.UpdatedAt = time.Now()

	mock.db[m.ID] = edit
	return &edit, nil
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

// ReadBySkills is a mock implementation of ReadBySkills method
func (mock *UserDAOMock) ReadBySkills(m string) ([]models.User, error) {
	ret := make([]models.User, 0, len(mock.db))
	for _, val := range mock.db {
		if val.Skills == m {
			ret = append(ret, val)
		}
	}

	return ret, nil
}

// ReadBySkillsT is a mock implementation of ReadBySkillsT method
func (mock *UserDAOMock) ReadBySkillsT(m string) (*gorm.DB, error) {
	return nil, nil
}

// DeleteBySkills is a mock implementation of DeleteBySkills method
func (mock *UserDAOMock) DeleteBySkills(m string) (error) {
	for _, val := range mock.db {
		if val.Skills == m {
			delete(mock.db, val.ID)
		}
	}

	return nil
}

// EditBySkills is a mock implementation of EditBySkills method
func (mock *UserDAOMock) EditBySkills(m string, newVals *models.User) (error) {
	for _, val := range mock.db {
		if val.Skills == m {
			id := val.ID
			val = *newVals
			val.ID = id
			val.UpdatedAt = time.Now()
		}
	}

	return nil
}

// SetSkills is a mock implementation of SetSkills method
func (mock *UserDAOMock) SetSkills(m *models.User, newVal string) (*models.User, error) {
	edit := mock.db[m.ID]
	edit.Skills = newVal
	edit.UpdatedAt = time.Now()

	mock.db[m.ID] = edit
	return &edit, nil
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

// AddProjectsAssociation is a mock implementation of AddProjectsAssociation method
func (mock *UserDAOMock) AddProjectsAssociation(m *models.User, asocVal *models.Project) (*models.User, error) {
	edit := mock.db[m.ID]
	edit.Projects = append(edit.Projects, *asocVal)
	edit.UpdatedAt = time.Now()
	mock.db[m.ID] = edit

	return &edit, nil
}

// RemoveProjectsAssociation is a mock implementation of RemoveProjectsAssociation method
func (mock *UserDAOMock) RemoveProjectsAssociation(m *models.User, asocVal *models.Project) (*models.User, error) {
	a := m.Projects
	m.UpdatedAt = time.Now()
	deletedIndex := 0
	for j, val := range a {
		if val.ID == asocVal.ID {
			deletedIndex = j
		}
	}
	a[deletedIndex] = a[len(a)-1]
	a = a[:len(a)-1]

	return m, nil
}

// GetAllAssociatedProjects is a mock implementation of GetAllAssociatedProjects method
func (mock *UserDAOMock) GetAllAssociatedProjects(m *models.User) ([]models.Project, error) {
	return m.Projects, nil
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

// AddRisksAssociation is a mock implementation of AddRisksAssociation method
func (mock *UserDAOMock) AddRisksAssociation(m *models.User, asocVal *models.Risk) (*models.User, error) {
	edit := mock.db[m.ID]
	edit.Risks = append(edit.Risks, *asocVal)
	edit.UpdatedAt = time.Now()
	mock.db[m.ID] = edit

	return &edit, nil
}

// RemoveRisksAssociation is a mock implementation of RemoveRisksAssociation method
func (mock *UserDAOMock) RemoveRisksAssociation(m *models.User, asocVal *models.Risk) (*models.User, error) {
	a := m.Risks
	m.UpdatedAt = time.Now()
	deletedIndex := 0
	for j, val := range a {
		if val.ID == asocVal.ID {
			deletedIndex = j
		}
	}
	a[deletedIndex] = a[len(a)-1]
	a = a[:len(a)-1]

	return m, nil
}

// GetAllAssociatedRisks is a mock implementation of GetAllAssociatedRisks method
func (mock *UserDAOMock) GetAllAssociatedRisks(m *models.User) ([]models.Risk, error) {
	return m.Risks, nil
}


// ReadByID will find models.User by ID given by parameter
func (dao *UserDAO) ReadByID(id uint) (*models.User, error) {
	m := &models.User{}
	if err := dao.db.First(&m, id).Error; err != nil {
		return nil, err
	}

	return m, nil
}


// UserDAOMock is a mock DAO
type UserDAOMock struct{
	db map[uint]models.User
	lastID uint
}

// NewUserDAOMock is a factory function for NewUserDAOMock
func NewUserDAOMock() *UserDAOMock{
	return &UserDAOMock{
		db: make(map[uint]models.User),
	}
}

// Create will put a model into mock in-memory DB
func (mock *UserDAOMock) Create(m *models.User) (error) {
	created := false
	m.CreatedAt = time.Now()
	for !created{
		mock.lastID++
		if _, exists := mock.db[mock.lastID]; !exists {
			m.ID = mock.lastID
			mock.db[mock.lastID] = *m
			created = true
		}
	}
	return nil
}

type UserDAOInterface interface {
	 Create(m *models.User) (error)
	 Read(m *models.User) ([]models.User, error)
	 ReadT(m *models.User) (*gorm.DB, error)
	 ReadByIDT(id uint) (*gorm.DB, error)
	 Update(m *models.User, id uint) (*models.User, error)
	 UpdateAllFields(m *models.User) (*models.User, error)
	 Delete(m *models.User) (error)
	 GetUpdatedAfter(timestamp time.Time) ([]models.User, error)
	 GetAll() ([]models.User, error)
	 ExecuteCustomQueryT(query string) (*gorm.DB, error)
	 ReadByName(m string) ([]models.User, error)
	 ReadByNameT(m string) (*gorm.DB, error)
	 DeleteByName(m string) (error)
	 EditByName(m string, newVals *models.User) (error)
	 SetName(m *models.User, newVal string) (*models.User, error)
	 ReadByEmail(m string) ([]models.User, error)
	 ReadByEmailT(m string) (*gorm.DB, error)
	 DeleteByEmail(m string) (error)
	 EditByEmail(m string, newVals *models.User) (error)
	 SetEmail(m *models.User, newVal string) (*models.User, error)
	 ReadByPassword(m string) ([]models.User, error)
	 ReadByPasswordT(m string) (*gorm.DB, error)
	 DeleteByPassword(m string) (error)
	 EditByPassword(m string, newVals *models.User) (error)
	 SetPassword(m *models.User, newVal string) (*models.User, error)
	 ReadByRole(m int) ([]models.User, error)
	 ReadByRoleT(m int) (*gorm.DB, error)
	 DeleteByRole(m int) (error)
	 EditByRole(m int, newVals *models.User) (error)
	 SetRole(m *models.User, newVal int) (*models.User, error)
	 ReadBySkills(m string) ([]models.User, error)
	 ReadBySkillsT(m string) (*gorm.DB, error)
	 DeleteBySkills(m string) (error)
	 EditBySkills(m string, newVals *models.User) (error)
	 SetSkills(m *models.User, newVal string) (*models.User, error)
	 AddProjectsAssociation(m *models.User, asocVal *models.Project) (*models.User, error)
	 RemoveProjectsAssociation(m *models.User, asocVal *models.Project) (*models.User, error)
	 GetAllAssociatedProjects(m *models.User) ([]models.Project, error)
	 AddRisksAssociation(m *models.User, asocVal *models.Risk) (*models.User, error)
	 RemoveRisksAssociation(m *models.User, asocVal *models.Risk) (*models.User, error)
	 GetAllAssociatedRisks(m *models.User) ([]models.Risk, error)
	 ReadByID(id uint) (*models.User, error)
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

// Read is a mock implementation of Read method
func (mock *ProjectDAOMock) Read(m *models.Project) ([]models.Project, error) {
	ret := make([]models.Project, 0, len(mock.db))
	for range mock.db {
		/*if val == *m {
			ret = append(ret, val)
		}*/
	}

	return ret, nil
}


// ReadByID is a mock implementation of ReadByID method
func (mock *ProjectDAOMock) ReadByID(id uint) (*models.Project, error) {
	ret := mock.db[id]
	return &ret, nil
}

// ReadByIDT is a mock implementation of ReadByIDT method
func (mock *ProjectDAOMock) ReadByIDT(id uint) (*gorm.DB, error) {
	return nil, nil
}

// ReadT is a mock implementation of ReadT method
func (mock *ProjectDAOMock) ReadT(m *models.Project) (*gorm.DB, error) {
	return nil, nil
}

// Update is a mock implementation of Update method
func (mock *ProjectDAOMock) Update(m *models.Project, id uint) (*models.Project, error) {
	m.UpdatedAt = time.Now()
	mock.db[id] = *m
	return m, nil
}

// UpdateAllFields is a mock implementation of UpdateAllFields method
func (mock *ProjectDAOMock) UpdateAllFields(m *models.Project) (*models.Project, error) {
	m.UpdatedAt = time.Now()
	mock.db[m.ID] = *m
	return m, nil
}

// Delete is a mock implementation of Delete method
func (mock *ProjectDAOMock) Delete(m *models.Project) (error) {
	delete(mock.db, m.ID)
	return nil
}

// GetUpdatedAfter is a mock implementation of GetUpdatedAfter method
func (mock *ProjectDAOMock) GetUpdatedAfter(timestamp time.Time) ([]models.Project, error) {
	ret := make([]models.Project, 0, len(mock.db))
	for _, val :=  range mock.db {
		if val.UpdatedAt.After(timestamp) {
			ret = append(ret, val)
		}
	}

	return ret, nil
}

// GetAll is a mock implementation of GetAll method
func (mock *ProjectDAOMock) GetAll() ([]models.Project, error) {
	ret := make([]models.Project, 0, len(mock.db))
	for _, val := range mock.db {
		ret = append(ret, val)
	}

	return ret, nil
}

// ExecuteCustomQueryT is a mock implementation of ExecuteCustomQueryT method
func (mock *ProjectDAOMock) ExecuteCustomQueryT(query string) (*gorm.DB, error) {
	return nil, nil
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

	dao.db.Model(m).Association("Risks").Find(&retVal)

	return retVal, nil
}

// AddRisksAssociation is a mock implementation of AddRisksAssociation method
func (mock *ProjectDAOMock) AddRisksAssociation(m *models.Project, asocVal *models.Risk) (*models.Project, error) {
	edit := mock.db[m.ID]
	edit.Risks = append(edit.Risks, *asocVal)
	edit.UpdatedAt = time.Now()
	mock.db[m.ID] = edit

	return &edit, nil
}

// RemoveRisksAssociation is a mock implementation of RemoveRisksAssociation method
func (mock *ProjectDAOMock) RemoveRisksAssociation(m *models.Project, asocVal *models.Risk) (*models.Project, error) {
	a := m.Risks
	m.UpdatedAt = time.Now()
	deletedIndex := 0
	for j, val := range a {
		if val.ID == asocVal.ID {
			deletedIndex = j
		}
	}
	a[deletedIndex] = a[len(a)-1]
	a = a[:len(a)-1]

	return m, nil
}

// GetAllAssociatedRisks is a mock implementation of GetAllAssociatedRisks method
func (mock *ProjectDAOMock) GetAllAssociatedRisks(m *models.Project) ([]models.Risk, error) {
	return m.Risks, nil
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

// ReadByName is a mock implementation of ReadByName method
func (mock *ProjectDAOMock) ReadByName(m string) ([]models.Project, error) {
	ret := make([]models.Project, 0, len(mock.db))
	for _, val := range mock.db {
		if val.Name == m {
			ret = append(ret, val)
		}
	}

	return ret, nil
}

// ReadByNameT is a mock implementation of ReadByNameT method
func (mock *ProjectDAOMock) ReadByNameT(m string) (*gorm.DB, error) {
	return nil, nil
}

// DeleteByName is a mock implementation of DeleteByName method
func (mock *ProjectDAOMock) DeleteByName(m string) (error) {
	for _, val := range mock.db {
		if val.Name == m {
			delete(mock.db, val.ID)
		}
	}

	return nil
}

// EditByName is a mock implementation of EditByName method
func (mock *ProjectDAOMock) EditByName(m string, newVals *models.Project) (error) {
	for _, val := range mock.db {
		if val.Name == m {
			id := val.ID
			val = *newVals
			val.ID = id
			val.UpdatedAt = time.Now()
		}
	}

	return nil
}

// SetName is a mock implementation of SetName method
func (mock *ProjectDAOMock) SetName(m *models.Project, newVal string) (*models.Project, error) {
	edit := mock.db[m.ID]
	edit.Name = newVal
	edit.UpdatedAt = time.Now()

	mock.db[m.ID] = edit
	return &edit, nil
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

// ReadByDescription is a mock implementation of ReadByDescription method
func (mock *ProjectDAOMock) ReadByDescription(m string) ([]models.Project, error) {
	ret := make([]models.Project, 0, len(mock.db))
	for _, val := range mock.db {
		if val.Description == m {
			ret = append(ret, val)
		}
	}

	return ret, nil
}

// ReadByDescriptionT is a mock implementation of ReadByDescriptionT method
func (mock *ProjectDAOMock) ReadByDescriptionT(m string) (*gorm.DB, error) {
	return nil, nil
}

// DeleteByDescription is a mock implementation of DeleteByDescription method
func (mock *ProjectDAOMock) DeleteByDescription(m string) (error) {
	for _, val := range mock.db {
		if val.Description == m {
			delete(mock.db, val.ID)
		}
	}

	return nil
}

// EditByDescription is a mock implementation of EditByDescription method
func (mock *ProjectDAOMock) EditByDescription(m string, newVals *models.Project) (error) {
	for _, val := range mock.db {
		if val.Description == m {
			id := val.ID
			val = *newVals
			val.ID = id
			val.UpdatedAt = time.Now()
		}
	}

	return nil
}

// SetDescription is a mock implementation of SetDescription method
func (mock *ProjectDAOMock) SetDescription(m *models.Project, newVal string) (*models.Project, error) {
	edit := mock.db[m.ID]
	edit.Description = newVal
	edit.UpdatedAt = time.Now()

	mock.db[m.ID] = edit
	return &edit, nil
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

// AddUsersAssociation is a mock implementation of AddUsersAssociation method
func (mock *ProjectDAOMock) AddUsersAssociation(m *models.Project, asocVal *models.User) (*models.Project, error) {
	edit := mock.db[m.ID]
	edit.Users = append(edit.Users, *asocVal)
	edit.UpdatedAt = time.Now()
	mock.db[m.ID] = edit

	return &edit, nil
}

// RemoveUsersAssociation is a mock implementation of RemoveUsersAssociation method
func (mock *ProjectDAOMock) RemoveUsersAssociation(m *models.Project, asocVal *models.User) (*models.Project, error) {
	a := m.Users
	m.UpdatedAt = time.Now()
	deletedIndex := 0
	for j, val := range a {
		if val.ID == asocVal.ID {
			deletedIndex = j
		}
	}
	a[deletedIndex] = a[len(a)-1]
	a = a[:len(a)-1]

	return m, nil
}

// GetAllAssociatedUsers is a mock implementation of GetAllAssociatedUsers method
func (mock *ProjectDAOMock) GetAllAssociatedUsers(m *models.Project) ([]models.User, error) {
	return m.Users, nil
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

// ReadByIsFinished is a mock implementation of ReadByIsFinished method
func (mock *ProjectDAOMock) ReadByIsFinished(m bool) ([]models.Project, error) {
	ret := make([]models.Project, 0, len(mock.db))
	for _, val := range mock.db {
		if val.IsFinished == m {
			ret = append(ret, val)
		}
	}

	return ret, nil
}

// ReadByIsFinishedT is a mock implementation of ReadByIsFinishedT method
func (mock *ProjectDAOMock) ReadByIsFinishedT(m bool) (*gorm.DB, error) {
	return nil, nil
}

// DeleteByIsFinished is a mock implementation of DeleteByIsFinished method
func (mock *ProjectDAOMock) DeleteByIsFinished(m bool) (error) {
	for _, val := range mock.db {
		if val.IsFinished == m {
			delete(mock.db, val.ID)
		}
	}

	return nil
}

// EditByIsFinished is a mock implementation of EditByIsFinished method
func (mock *ProjectDAOMock) EditByIsFinished(m bool, newVals *models.Project) (error) {
	for _, val := range mock.db {
		if val.IsFinished == m {
			id := val.ID
			val = *newVals
			val.ID = id
			val.UpdatedAt = time.Now()
		}
	}

	return nil
}

// SetIsFinished is a mock implementation of SetIsFinished method
func (mock *ProjectDAOMock) SetIsFinished(m *models.Project, newVal bool) (*models.Project, error) {
	edit := mock.db[m.ID]
	edit.IsFinished = newVal
	edit.UpdatedAt = time.Now()

	mock.db[m.ID] = edit
	return &edit, nil
}


// ReadByManagerID will find all records
// matching the value given by parameter
func (dao *ProjectDAO) ReadByManagerID(m uint) ([]models.Project, error) {
	retVal := []models.Project{}
	if err := dao.db.Where(&models.Project{ManagerID: m}).Find(&retVal).Error; err != nil {
		return nil, err
	}

	return retVal, nil
}

// ReadByManagerIDT will return a transaction that
// can be used to find all models matching the value given by parameter
func (dao *ProjectDAO) ReadByManagerIDT(m uint) (*gorm.DB, error) {
	retVal := dao.db.Where(&models.Project{ManagerID: m})

	return retVal, retVal.Error
}

// DeleteByManagerID deletes all records in database with
// ManagerID the same as parameter given
func (dao *ProjectDAO) DeleteByManagerID(m uint) (error) {
	if err := dao.db.Where(&models.Project{ManagerID: m}).Delete(&models.Project{}).Error; err != nil {
		return err
	}
	return nil
}

// EditByManagerID will edit all records in database
// with the same ManagerID as parameter given
// using model given by parameter
func (dao *ProjectDAO) EditByManagerID(m uint, newVals *models.Project) (error) {
	if err := dao.db.Table("projects").Where(&models.Project{ManagerID: m}).Updates(newVals).Error; err != nil {
		return err
	}
	return nil
}

// SetManagerID will set ManagerID
// to a value given by parameter
func (dao *ProjectDAO) SetManagerID(m *models.Project, newVal uint) (*models.Project, error) {
	m.ManagerID = newVal
	record, err := dao.ReadByID((m.ID))
	if err != nil {
		return nil, err
	}

	if err := dao.db.Model(&record).Updates(m).Error; err != nil {
		return nil, err
	}

	return record, nil
}

// ReadByManagerID is a mock implementation of ReadByManagerID method
func (mock *ProjectDAOMock) ReadByManagerID(m uint) ([]models.Project, error) {
	ret := make([]models.Project, 0, len(mock.db))
	for _, val := range mock.db {
		if val.ManagerID == m {
			ret = append(ret, val)
		}
	}

	return ret, nil
}

// ReadByManagerIDT is a mock implementation of ReadByManagerIDT method
func (mock *ProjectDAOMock) ReadByManagerIDT(m uint) (*gorm.DB, error) {
	return nil, nil
}

// DeleteByManagerID is a mock implementation of DeleteByManagerID method
func (mock *ProjectDAOMock) DeleteByManagerID(m uint) (error) {
	for _, val := range mock.db {
		if val.ManagerID == m {
			delete(mock.db, val.ID)
		}
	}

	return nil
}

// EditByManagerID is a mock implementation of EditByManagerID method
func (mock *ProjectDAOMock) EditByManagerID(m uint, newVals *models.Project) (error) {
	for _, val := range mock.db {
		if val.ManagerID == m {
			id := val.ID
			val = *newVals
			val.ID = id
			val.UpdatedAt = time.Now()
		}
	}

	return nil
}

// SetManagerID is a mock implementation of SetManagerID method
func (mock *ProjectDAOMock) SetManagerID(m *models.Project, newVal uint) (*models.Project, error) {
	edit := mock.db[m.ID]
	edit.ManagerID = newVal
	edit.UpdatedAt = time.Now()

	mock.db[m.ID] = edit
	return &edit, nil
}


// ReadByID will find models.Project by ID given by parameter
func (dao *ProjectDAO) ReadByID(id uint) (*models.Project, error) {
	m := &models.Project{}
	if err := dao.db.First(&m, id).Error; err != nil {
		return nil, err
	}

	return m, nil
}


// ProjectDAOMock is a mock DAO
type ProjectDAOMock struct{
	db map[uint]models.Project
	lastID uint
}

// NewProjectDAOMock is a factory function for NewProjectDAOMock
func NewProjectDAOMock() *ProjectDAOMock{
	return &ProjectDAOMock{
		db: make(map[uint]models.Project),
	}
}

// Create will put a model into mock in-memory DB
func (mock *ProjectDAOMock) Create(m *models.Project) (error) {
	created := false
	m.CreatedAt = time.Now()
	for !created{
		mock.lastID++
		if _, exists := mock.db[mock.lastID]; !exists {
			m.ID = mock.lastID
			mock.db[mock.lastID] = *m
			created = true
		}
	}
	return nil
}

type ProjectDAOInterface interface {
	 Create(m *models.Project) (error)
	 Read(m *models.Project) ([]models.Project, error)
	 ReadT(m *models.Project) (*gorm.DB, error)
	 ReadByIDT(id uint) (*gorm.DB, error)
	 Update(m *models.Project, id uint) (*models.Project, error)
	 UpdateAllFields(m *models.Project) (*models.Project, error)
	 Delete(m *models.Project) (error)
	 GetUpdatedAfter(timestamp time.Time) ([]models.Project, error)
	 GetAll() ([]models.Project, error)
	 ExecuteCustomQueryT(query string) (*gorm.DB, error)
	 SetEnd(m *models.Project, str time.Time) (*models.Project, error)
	 AddRisksAssociation(m *models.Project, asocVal *models.Risk) (*models.Project, error)
	 RemoveRisksAssociation(m *models.Project, asocVal *models.Risk) (*models.Project, error)
	 GetAllAssociatedRisks(m *models.Project) ([]models.Risk, error)
	 ReadByName(m string) ([]models.Project, error)
	 ReadByNameT(m string) (*gorm.DB, error)
	 DeleteByName(m string) (error)
	 EditByName(m string, newVals *models.Project) (error)
	 SetName(m *models.Project, newVal string) (*models.Project, error)
	 ReadByDescription(m string) ([]models.Project, error)
	 ReadByDescriptionT(m string) (*gorm.DB, error)
	 DeleteByDescription(m string) (error)
	 EditByDescription(m string, newVals *models.Project) (error)
	 SetDescription(m *models.Project, newVal string) (*models.Project, error)
	 AddUsersAssociation(m *models.Project, asocVal *models.User) (*models.Project, error)
	 RemoveUsersAssociation(m *models.Project, asocVal *models.User) (*models.Project, error)
	 GetAllAssociatedUsers(m *models.Project) ([]models.User, error)
	 SetStart(m *models.Project, str time.Time) (*models.Project, error)
	 ReadByIsFinished(m bool) ([]models.Project, error)
	 ReadByIsFinishedT(m bool) (*gorm.DB, error)
	 DeleteByIsFinished(m bool) (error)
	 EditByIsFinished(m bool, newVals *models.Project) (error)
	 SetIsFinished(m *models.Project, newVal bool) (*models.Project, error)
	 ReadByManagerID(m uint) ([]models.Project, error)
	 ReadByManagerIDT(m uint) (*gorm.DB, error)
	 DeleteByManagerID(m uint) (error)
	 EditByManagerID(m uint, newVals *models.Project) (error)
	 SetManagerID(m *models.Project, newVal uint) (*models.Project, error)
	 ReadByID(id uint) (*models.Project, error)
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

// Read is a mock implementation of Read method
func (mock *RiskDAOMock) Read(m *models.Risk) ([]models.Risk, error) {
	ret := make([]models.Risk, 0, len(mock.db))
	for range mock.db {
		/*if val == *m {
			ret = append(ret, val)
		}*/
	}

	return ret, nil
}


// ReadByID is a mock implementation of ReadByID method
func (mock *RiskDAOMock) ReadByID(id uint) (*models.Risk, error) {
	ret := mock.db[id]
	return &ret, nil
}

// ReadByIDT is a mock implementation of ReadByIDT method
func (mock *RiskDAOMock) ReadByIDT(id uint) (*gorm.DB, error) {
	return nil, nil
}

// ReadT is a mock implementation of ReadT method
func (mock *RiskDAOMock) ReadT(m *models.Risk) (*gorm.DB, error) {
	return nil, nil
}

// Update is a mock implementation of Update method
func (mock *RiskDAOMock) Update(m *models.Risk, id uint) (*models.Risk, error) {
	m.UpdatedAt = time.Now()
	mock.db[id] = *m
	return m, nil
}

// UpdateAllFields is a mock implementation of UpdateAllFields method
func (mock *RiskDAOMock) UpdateAllFields(m *models.Risk) (*models.Risk, error) {
	m.UpdatedAt = time.Now()
	mock.db[m.ID] = *m
	return m, nil
}

// Delete is a mock implementation of Delete method
func (mock *RiskDAOMock) Delete(m *models.Risk) (error) {
	delete(mock.db, m.ID)
	return nil
}

// GetUpdatedAfter is a mock implementation of GetUpdatedAfter method
func (mock *RiskDAOMock) GetUpdatedAfter(timestamp time.Time) ([]models.Risk, error) {
	ret := make([]models.Risk, 0, len(mock.db))
	for _, val :=  range mock.db {
		if val.UpdatedAt.After(timestamp) {
			ret = append(ret, val)
		}
	}

	return ret, nil
}

// GetAll is a mock implementation of GetAll method
func (mock *RiskDAOMock) GetAll() ([]models.Risk, error) {
	ret := make([]models.Risk, 0, len(mock.db))
	for _, val := range mock.db {
		ret = append(ret, val)
	}

	return ret, nil
}

// ExecuteCustomQueryT is a mock implementation of ExecuteCustomQueryT method
func (mock *RiskDAOMock) ExecuteCustomQueryT(query string) (*gorm.DB, error) {
	return nil, nil
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

// ReadByTrigger is a mock implementation of ReadByTrigger method
func (mock *RiskDAOMock) ReadByTrigger(m string) ([]models.Risk, error) {
	ret := make([]models.Risk, 0, len(mock.db))
	for _, val := range mock.db {
		if val.Trigger == m {
			ret = append(ret, val)
		}
	}

	return ret, nil
}

// ReadByTriggerT is a mock implementation of ReadByTriggerT method
func (mock *RiskDAOMock) ReadByTriggerT(m string) (*gorm.DB, error) {
	return nil, nil
}

// DeleteByTrigger is a mock implementation of DeleteByTrigger method
func (mock *RiskDAOMock) DeleteByTrigger(m string) (error) {
	for _, val := range mock.db {
		if val.Trigger == m {
			delete(mock.db, val.ID)
		}
	}

	return nil
}

// EditByTrigger is a mock implementation of EditByTrigger method
func (mock *RiskDAOMock) EditByTrigger(m string, newVals *models.Risk) (error) {
	for _, val := range mock.db {
		if val.Trigger == m {
			id := val.ID
			val = *newVals
			val.ID = id
			val.UpdatedAt = time.Now()
		}
	}

	return nil
}

// SetTrigger is a mock implementation of SetTrigger method
func (mock *RiskDAOMock) SetTrigger(m *models.Risk, newVal string) (*models.Risk, error) {
	edit := mock.db[m.ID]
	edit.Trigger = newVal
	edit.UpdatedAt = time.Now()

	mock.db[m.ID] = edit
	return &edit, nil
}


// ReadByImpact will find all records
// matching the value given by parameter
func (dao *RiskDAO) ReadByImpact(m float64) ([]models.Risk, error) {
	retVal := []models.Risk{}
	if err := dao.db.Where(&models.Risk{Impact: m}).Find(&retVal).Error; err != nil {
		return nil, err
	}

	return retVal, nil
}

// ReadByImpactT will return a transaction that
// can be used to find all models matching the value given by parameter
func (dao *RiskDAO) ReadByImpactT(m float64) (*gorm.DB, error) {
	retVal := dao.db.Where(&models.Risk{Impact: m})

	return retVal, retVal.Error
}

// DeleteByImpact deletes all records in database with
// Impact the same as parameter given
func (dao *RiskDAO) DeleteByImpact(m float64) (error) {
	if err := dao.db.Where(&models.Risk{Impact: m}).Delete(&models.Risk{}).Error; err != nil {
		return err
	}
	return nil
}

// EditByImpact will edit all records in database
// with the same Impact as parameter given
// using model given by parameter
func (dao *RiskDAO) EditByImpact(m float64, newVals *models.Risk) (error) {
	if err := dao.db.Table("risks").Where(&models.Risk{Impact: m}).Updates(newVals).Error; err != nil {
		return err
	}
	return nil
}

// SetImpact will set Impact
// to a value given by parameter
func (dao *RiskDAO) SetImpact(m *models.Risk, newVal float64) (*models.Risk, error) {
	m.Impact = newVal
	record, err := dao.ReadByID((m.ID))
	if err != nil {
		return nil, err
	}

	if err := dao.db.Model(&record).Updates(m).Error; err != nil {
		return nil, err
	}

	return record, nil
}

// ReadByImpact is a mock implementation of ReadByImpact method
func (mock *RiskDAOMock) ReadByImpact(m float64) ([]models.Risk, error) {
	ret := make([]models.Risk, 0, len(mock.db))
	for _, val := range mock.db {
		if val.Impact == m {
			ret = append(ret, val)
		}
	}

	return ret, nil
}

// ReadByImpactT is a mock implementation of ReadByImpactT method
func (mock *RiskDAOMock) ReadByImpactT(m float64) (*gorm.DB, error) {
	return nil, nil
}

// DeleteByImpact is a mock implementation of DeleteByImpact method
func (mock *RiskDAOMock) DeleteByImpact(m float64) (error) {
	for _, val := range mock.db {
		if val.Impact == m {
			delete(mock.db, val.ID)
		}
	}

	return nil
}

// EditByImpact is a mock implementation of EditByImpact method
func (mock *RiskDAOMock) EditByImpact(m float64, newVals *models.Risk) (error) {
	for _, val := range mock.db {
		if val.Impact == m {
			id := val.ID
			val = *newVals
			val.ID = id
			val.UpdatedAt = time.Now()
		}
	}

	return nil
}

// SetImpact is a mock implementation of SetImpact method
func (mock *RiskDAOMock) SetImpact(m *models.Risk, newVal float64) (*models.Risk, error) {
	edit := mock.db[m.ID]
	edit.Impact = newVal
	edit.UpdatedAt = time.Now()

	mock.db[m.ID] = edit
	return &edit, nil
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

// AddProjectsAssociation is a mock implementation of AddProjectsAssociation method
func (mock *RiskDAOMock) AddProjectsAssociation(m *models.Risk, asocVal *models.Project) (*models.Risk, error) {
	edit := mock.db[m.ID]
	edit.Projects = append(edit.Projects, *asocVal)
	edit.UpdatedAt = time.Now()
	mock.db[m.ID] = edit

	return &edit, nil
}

// RemoveProjectsAssociation is a mock implementation of RemoveProjectsAssociation method
func (mock *RiskDAOMock) RemoveProjectsAssociation(m *models.Risk, asocVal *models.Project) (*models.Risk, error) {
	a := m.Projects
	m.UpdatedAt = time.Now()
	deletedIndex := 0
	for j, val := range a {
		if val.ID == asocVal.ID {
			deletedIndex = j
		}
	}
	a[deletedIndex] = a[len(a)-1]
	a = a[:len(a)-1]

	return m, nil
}

// GetAllAssociatedProjects is a mock implementation of GetAllAssociatedProjects method
func (mock *RiskDAOMock) GetAllAssociatedProjects(m *models.Risk) ([]models.Project, error) {
	return m.Projects, nil
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

// ReadByProbability is a mock implementation of ReadByProbability method
func (mock *RiskDAOMock) ReadByProbability(m float64) ([]models.Risk, error) {
	ret := make([]models.Risk, 0, len(mock.db))
	for _, val := range mock.db {
		if val.Probability == m {
			ret = append(ret, val)
		}
	}

	return ret, nil
}

// ReadByProbabilityT is a mock implementation of ReadByProbabilityT method
func (mock *RiskDAOMock) ReadByProbabilityT(m float64) (*gorm.DB, error) {
	return nil, nil
}

// DeleteByProbability is a mock implementation of DeleteByProbability method
func (mock *RiskDAOMock) DeleteByProbability(m float64) (error) {
	for _, val := range mock.db {
		if val.Probability == m {
			delete(mock.db, val.ID)
		}
	}

	return nil
}

// EditByProbability is a mock implementation of EditByProbability method
func (mock *RiskDAOMock) EditByProbability(m float64, newVals *models.Risk) (error) {
	for _, val := range mock.db {
		if val.Probability == m {
			id := val.ID
			val = *newVals
			val.ID = id
			val.UpdatedAt = time.Now()
		}
	}

	return nil
}

// SetProbability is a mock implementation of SetProbability method
func (mock *RiskDAOMock) SetProbability(m *models.Risk, newVal float64) (*models.Risk, error) {
	edit := mock.db[m.ID]
	edit.Probability = newVal
	edit.UpdatedAt = time.Now()

	mock.db[m.ID] = edit
	return &edit, nil
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

// ReadByCategory is a mock implementation of ReadByCategory method
func (mock *RiskDAOMock) ReadByCategory(m string) ([]models.Risk, error) {
	ret := make([]models.Risk, 0, len(mock.db))
	for _, val := range mock.db {
		if val.Category == m {
			ret = append(ret, val)
		}
	}

	return ret, nil
}

// ReadByCategoryT is a mock implementation of ReadByCategoryT method
func (mock *RiskDAOMock) ReadByCategoryT(m string) (*gorm.DB, error) {
	return nil, nil
}

// DeleteByCategory is a mock implementation of DeleteByCategory method
func (mock *RiskDAOMock) DeleteByCategory(m string) (error) {
	for _, val := range mock.db {
		if val.Category == m {
			delete(mock.db, val.ID)
		}
	}

	return nil
}

// EditByCategory is a mock implementation of EditByCategory method
func (mock *RiskDAOMock) EditByCategory(m string, newVals *models.Risk) (error) {
	for _, val := range mock.db {
		if val.Category == m {
			id := val.ID
			val = *newVals
			val.ID = id
			val.UpdatedAt = time.Now()
		}
	}

	return nil
}

// SetCategory is a mock implementation of SetCategory method
func (mock *RiskDAOMock) SetCategory(m *models.Risk, newVal string) (*models.Risk, error) {
	edit := mock.db[m.ID]
	edit.Category = newVal
	edit.UpdatedAt = time.Now()

	mock.db[m.ID] = edit
	return &edit, nil
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

// ReadByStatus is a mock implementation of ReadByStatus method
func (mock *RiskDAOMock) ReadByStatus(m string) ([]models.Risk, error) {
	ret := make([]models.Risk, 0, len(mock.db))
	for _, val := range mock.db {
		if val.Status == m {
			ret = append(ret, val)
		}
	}

	return ret, nil
}

// ReadByStatusT is a mock implementation of ReadByStatusT method
func (mock *RiskDAOMock) ReadByStatusT(m string) (*gorm.DB, error) {
	return nil, nil
}

// DeleteByStatus is a mock implementation of DeleteByStatus method
func (mock *RiskDAOMock) DeleteByStatus(m string) (error) {
	for _, val := range mock.db {
		if val.Status == m {
			delete(mock.db, val.ID)
		}
	}

	return nil
}

// EditByStatus is a mock implementation of EditByStatus method
func (mock *RiskDAOMock) EditByStatus(m string, newVals *models.Risk) (error) {
	for _, val := range mock.db {
		if val.Status == m {
			id := val.ID
			val = *newVals
			val.ID = id
			val.UpdatedAt = time.Now()
		}
	}

	return nil
}

// SetStatus is a mock implementation of SetStatus method
func (mock *RiskDAOMock) SetStatus(m *models.Risk, newVal string) (*models.Risk, error) {
	edit := mock.db[m.ID]
	edit.Status = newVal
	edit.UpdatedAt = time.Now()

	mock.db[m.ID] = edit
	return &edit, nil
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

// ReadByDescription is a mock implementation of ReadByDescription method
func (mock *RiskDAOMock) ReadByDescription(m string) ([]models.Risk, error) {
	ret := make([]models.Risk, 0, len(mock.db))
	for _, val := range mock.db {
		if val.Description == m {
			ret = append(ret, val)
		}
	}

	return ret, nil
}

// ReadByDescriptionT is a mock implementation of ReadByDescriptionT method
func (mock *RiskDAOMock) ReadByDescriptionT(m string) (*gorm.DB, error) {
	return nil, nil
}

// DeleteByDescription is a mock implementation of DeleteByDescription method
func (mock *RiskDAOMock) DeleteByDescription(m string) (error) {
	for _, val := range mock.db {
		if val.Description == m {
			delete(mock.db, val.ID)
		}
	}

	return nil
}

// EditByDescription is a mock implementation of EditByDescription method
func (mock *RiskDAOMock) EditByDescription(m string, newVals *models.Risk) (error) {
	for _, val := range mock.db {
		if val.Description == m {
			id := val.ID
			val = *newVals
			val.ID = id
			val.UpdatedAt = time.Now()
		}
	}

	return nil
}

// SetDescription is a mock implementation of SetDescription method
func (mock *RiskDAOMock) SetDescription(m *models.Risk, newVal string) (*models.Risk, error) {
	edit := mock.db[m.ID]
	edit.Description = newVal
	edit.UpdatedAt = time.Now()

	mock.db[m.ID] = edit
	return &edit, nil
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

	dao.db.Model(&m).Association("CounterMeasures").Find(&retVal)
	return retVal, nil
}

// AddCounterMeasuresAssociation is a mock implementation of AddCounterMeasuresAssociation method
func (mock *RiskDAOMock) AddCounterMeasuresAssociation(m *models.Risk, asocVal *models.CounterMeasure) (*models.Risk, error) {
	edit := mock.db[m.ID]
	edit.CounterMeasures = append(edit.CounterMeasures, *asocVal)
	edit.UpdatedAt = time.Now()
	mock.db[m.ID] = edit

	return &edit, nil
}

// RemoveCounterMeasuresAssociation is a mock implementation of RemoveCounterMeasuresAssociation method
func (mock *RiskDAOMock) RemoveCounterMeasuresAssociation(m *models.Risk, asocVal *models.CounterMeasure) (*models.Risk, error) {
	a := m.CounterMeasures
	m.UpdatedAt = time.Now()
	deletedIndex := 0
	for j, val := range a {
		if val.ID == asocVal.ID {
			deletedIndex = j
		}
	}
	a[deletedIndex] = a[len(a)-1]
	a = a[:len(a)-1]

	return m, nil
}

// GetAllAssociatedCounterMeasures is a mock implementation of GetAllAssociatedCounterMeasures method
func (mock *RiskDAOMock) GetAllAssociatedCounterMeasures(m *models.Risk) ([]models.CounterMeasure, error) {
	return m.CounterMeasures, nil
}


// ReadByValue will find all records
// matching the value given by parameter
func (dao *RiskDAO) ReadByValue(m float64) ([]models.Risk, error) {
	retVal := []models.Risk{}
	if err := dao.db.Where(&models.Risk{Value: m}).Find(&retVal).Error; err != nil {
		return nil, err
	}

	return retVal, nil
}

// ReadByValueT will return a transaction that
// can be used to find all models matching the value given by parameter
func (dao *RiskDAO) ReadByValueT(m float64) (*gorm.DB, error) {
	retVal := dao.db.Where(&models.Risk{Value: m})

	return retVal, retVal.Error
}

// DeleteByValue deletes all records in database with
// Value the same as parameter given
func (dao *RiskDAO) DeleteByValue(m float64) (error) {
	if err := dao.db.Where(&models.Risk{Value: m}).Delete(&models.Risk{}).Error; err != nil {
		return err
	}
	return nil
}

// EditByValue will edit all records in database
// with the same Value as parameter given
// using model given by parameter
func (dao *RiskDAO) EditByValue(m float64, newVals *models.Risk) (error) {
	if err := dao.db.Table("risks").Where(&models.Risk{Value: m}).Updates(newVals).Error; err != nil {
		return err
	}
	return nil
}

// SetValue will set Value
// to a value given by parameter
func (dao *RiskDAO) SetValue(m *models.Risk, newVal float64) (*models.Risk, error) {
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

// ReadByValue is a mock implementation of ReadByValue method
func (mock *RiskDAOMock) ReadByValue(m float64) ([]models.Risk, error) {
	ret := make([]models.Risk, 0, len(mock.db))
	for _, val := range mock.db {
		if val.Value == m {
			ret = append(ret, val)
		}
	}

	return ret, nil
}

// ReadByValueT is a mock implementation of ReadByValueT method
func (mock *RiskDAOMock) ReadByValueT(m float64) (*gorm.DB, error) {
	return nil, nil
}

// DeleteByValue is a mock implementation of DeleteByValue method
func (mock *RiskDAOMock) DeleteByValue(m float64) (error) {
	for _, val := range mock.db {
		if val.Value == m {
			delete(mock.db, val.ID)
		}
	}

	return nil
}

// EditByValue is a mock implementation of EditByValue method
func (mock *RiskDAOMock) EditByValue(m float64, newVals *models.Risk) (error) {
	for _, val := range mock.db {
		if val.Value == m {
			id := val.ID
			val = *newVals
			val.ID = id
			val.UpdatedAt = time.Now()
		}
	}

	return nil
}

// SetValue is a mock implementation of SetValue method
func (mock *RiskDAOMock) SetValue(m *models.Risk, newVal float64) (*models.Risk, error) {
	edit := mock.db[m.ID]
	edit.Value = newVal
	edit.UpdatedAt = time.Now()

	mock.db[m.ID] = edit
	return &edit, nil
}


// ReadByRisk will find all records
// matching the value given by parameter
func (dao *RiskDAO) ReadByRisk(m float64) ([]models.Risk, error) {
	retVal := []models.Risk{}
	if err := dao.db.Where(&models.Risk{Risk: m}).Find(&retVal).Error; err != nil {
		return nil, err
	}

	return retVal, nil
}

// ReadByRiskT will return a transaction that
// can be used to find all models matching the value given by parameter
func (dao *RiskDAO) ReadByRiskT(m float64) (*gorm.DB, error) {
	retVal := dao.db.Where(&models.Risk{Risk: m})

	return retVal, retVal.Error
}

// DeleteByRisk deletes all records in database with
// Risk the same as parameter given
func (dao *RiskDAO) DeleteByRisk(m float64) (error) {
	if err := dao.db.Where(&models.Risk{Risk: m}).Delete(&models.Risk{}).Error; err != nil {
		return err
	}
	return nil
}

// EditByRisk will edit all records in database
// with the same Risk as parameter given
// using model given by parameter
func (dao *RiskDAO) EditByRisk(m float64, newVals *models.Risk) (error) {
	if err := dao.db.Table("risks").Where(&models.Risk{Risk: m}).Updates(newVals).Error; err != nil {
		return err
	}
	return nil
}

// SetRisk will set Risk
// to a value given by parameter
func (dao *RiskDAO) SetRisk(m *models.Risk, newVal float64) (*models.Risk, error) {
	m.Risk = newVal
	record, err := dao.ReadByID((m.ID))
	if err != nil {
		return nil, err
	}

	if err := dao.db.Model(&record).Updates(m).Error; err != nil {
		return nil, err
	}

	return record, nil
}

// ReadByRisk is a mock implementation of ReadByRisk method
func (mock *RiskDAOMock) ReadByRisk(m float64) ([]models.Risk, error) {
	ret := make([]models.Risk, 0, len(mock.db))
	for _, val := range mock.db {
		if val.Risk == m {
			ret = append(ret, val)
		}
	}

	return ret, nil
}

// ReadByRiskT is a mock implementation of ReadByRiskT method
func (mock *RiskDAOMock) ReadByRiskT(m float64) (*gorm.DB, error) {
	return nil, nil
}

// DeleteByRisk is a mock implementation of DeleteByRisk method
func (mock *RiskDAOMock) DeleteByRisk(m float64) (error) {
	for _, val := range mock.db {
		if val.Risk == m {
			delete(mock.db, val.ID)
		}
	}

	return nil
}

// EditByRisk is a mock implementation of EditByRisk method
func (mock *RiskDAOMock) EditByRisk(m float64, newVals *models.Risk) (error) {
	for _, val := range mock.db {
		if val.Risk == m {
			id := val.ID
			val = *newVals
			val.ID = id
			val.UpdatedAt = time.Now()
		}
	}

	return nil
}

// SetRisk is a mock implementation of SetRisk method
func (mock *RiskDAOMock) SetRisk(m *models.Risk, newVal float64) (*models.Risk, error) {
	edit := mock.db[m.ID]
	edit.Risk = newVal
	edit.UpdatedAt = time.Now()

	mock.db[m.ID] = edit
	return &edit, nil
}


// ReadByUserID will find all records
// matching the value given by parameter
func (dao *RiskDAO) ReadByUserID(m uint) ([]models.Risk, error) {
	retVal := []models.Risk{}
	if err := dao.db.Where(&models.Risk{UserID: m}).Find(&retVal).Error; err != nil {
		return nil, err
	}

	return retVal, nil
}

// ReadByUserIDT will return a transaction that
// can be used to find all models matching the value given by parameter
func (dao *RiskDAO) ReadByUserIDT(m uint) (*gorm.DB, error) {
	retVal := dao.db.Where(&models.Risk{UserID: m})

	return retVal, retVal.Error
}

// DeleteByUserID deletes all records in database with
// UserID the same as parameter given
func (dao *RiskDAO) DeleteByUserID(m uint) (error) {
	if err := dao.db.Where(&models.Risk{UserID: m}).Delete(&models.Risk{}).Error; err != nil {
		return err
	}
	return nil
}

// EditByUserID will edit all records in database
// with the same UserID as parameter given
// using model given by parameter
func (dao *RiskDAO) EditByUserID(m uint, newVals *models.Risk) (error) {
	if err := dao.db.Table("risks").Where(&models.Risk{UserID: m}).Updates(newVals).Error; err != nil {
		return err
	}
	return nil
}

// SetUserID will set UserID
// to a value given by parameter
func (dao *RiskDAO) SetUserID(m *models.Risk, newVal uint) (*models.Risk, error) {
	m.UserID = newVal
	record, err := dao.ReadByID((m.ID))
	if err != nil {
		return nil, err
	}

	if err := dao.db.Model(&record).Updates(m).Error; err != nil {
		return nil, err
	}

	return record, nil
}

// ReadByUserID is a mock implementation of ReadByUserID method
func (mock *RiskDAOMock) ReadByUserID(m uint) ([]models.Risk, error) {
	ret := make([]models.Risk, 0, len(mock.db))
	for _, val := range mock.db {
		if val.UserID == m {
			ret = append(ret, val)
		}
	}

	return ret, nil
}

// ReadByUserIDT is a mock implementation of ReadByUserIDT method
func (mock *RiskDAOMock) ReadByUserIDT(m uint) (*gorm.DB, error) {
	return nil, nil
}

// DeleteByUserID is a mock implementation of DeleteByUserID method
func (mock *RiskDAOMock) DeleteByUserID(m uint) (error) {
	for _, val := range mock.db {
		if val.UserID == m {
			delete(mock.db, val.ID)
		}
	}

	return nil
}

// EditByUserID is a mock implementation of EditByUserID method
func (mock *RiskDAOMock) EditByUserID(m uint, newVals *models.Risk) (error) {
	for _, val := range mock.db {
		if val.UserID == m {
			id := val.ID
			val = *newVals
			val.ID = id
			val.UpdatedAt = time.Now()
		}
	}

	return nil
}

// SetUserID is a mock implementation of SetUserID method
func (mock *RiskDAOMock) SetUserID(m *models.Risk, newVal uint) (*models.Risk, error) {
	edit := mock.db[m.ID]
	edit.UserID = newVal
	edit.UpdatedAt = time.Now()

	mock.db[m.ID] = edit
	return &edit, nil
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

// ReadByCost is a mock implementation of ReadByCost method
func (mock *RiskDAOMock) ReadByCost(m int) ([]models.Risk, error) {
	ret := make([]models.Risk, 0, len(mock.db))
	for _, val := range mock.db {
		if val.Cost == m {
			ret = append(ret, val)
		}
	}

	return ret, nil
}

// ReadByCostT is a mock implementation of ReadByCostT method
func (mock *RiskDAOMock) ReadByCostT(m int) (*gorm.DB, error) {
	return nil, nil
}

// DeleteByCost is a mock implementation of DeleteByCost method
func (mock *RiskDAOMock) DeleteByCost(m int) (error) {
	for _, val := range mock.db {
		if val.Cost == m {
			delete(mock.db, val.ID)
		}
	}

	return nil
}

// EditByCost is a mock implementation of EditByCost method
func (mock *RiskDAOMock) EditByCost(m int, newVals *models.Risk) (error) {
	for _, val := range mock.db {
		if val.Cost == m {
			id := val.ID
			val = *newVals
			val.ID = id
			val.UpdatedAt = time.Now()
		}
	}

	return nil
}

// SetCost is a mock implementation of SetCost method
func (mock *RiskDAOMock) SetCost(m *models.Risk, newVal int) (*models.Risk, error) {
	edit := mock.db[m.ID]
	edit.Cost = newVal
	edit.UpdatedAt = time.Now()

	mock.db[m.ID] = edit
	return &edit, nil
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

// ReadByName is a mock implementation of ReadByName method
func (mock *RiskDAOMock) ReadByName(m string) ([]models.Risk, error) {
	ret := make([]models.Risk, 0, len(mock.db))
	for _, val := range mock.db {
		if val.Name == m {
			ret = append(ret, val)
		}
	}

	return ret, nil
}

// ReadByNameT is a mock implementation of ReadByNameT method
func (mock *RiskDAOMock) ReadByNameT(m string) (*gorm.DB, error) {
	return nil, nil
}

// DeleteByName is a mock implementation of DeleteByName method
func (mock *RiskDAOMock) DeleteByName(m string) (error) {
	for _, val := range mock.db {
		if val.Name == m {
			delete(mock.db, val.ID)
		}
	}

	return nil
}

// EditByName is a mock implementation of EditByName method
func (mock *RiskDAOMock) EditByName(m string, newVals *models.Risk) (error) {
	for _, val := range mock.db {
		if val.Name == m {
			id := val.ID
			val = *newVals
			val.ID = id
			val.UpdatedAt = time.Now()
		}
	}

	return nil
}

// SetName is a mock implementation of SetName method
func (mock *RiskDAOMock) SetName(m *models.Risk, newVal string) (*models.Risk, error) {
	edit := mock.db[m.ID]
	edit.Name = newVal
	edit.UpdatedAt = time.Now()

	mock.db[m.ID] = edit
	return &edit, nil
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

// ReadByThreat is a mock implementation of ReadByThreat method
func (mock *RiskDAOMock) ReadByThreat(m string) ([]models.Risk, error) {
	ret := make([]models.Risk, 0, len(mock.db))
	for _, val := range mock.db {
		if val.Threat == m {
			ret = append(ret, val)
		}
	}

	return ret, nil
}

// ReadByThreatT is a mock implementation of ReadByThreatT method
func (mock *RiskDAOMock) ReadByThreatT(m string) (*gorm.DB, error) {
	return nil, nil
}

// DeleteByThreat is a mock implementation of DeleteByThreat method
func (mock *RiskDAOMock) DeleteByThreat(m string) (error) {
	for _, val := range mock.db {
		if val.Threat == m {
			delete(mock.db, val.ID)
		}
	}

	return nil
}

// EditByThreat is a mock implementation of EditByThreat method
func (mock *RiskDAOMock) EditByThreat(m string, newVals *models.Risk) (error) {
	for _, val := range mock.db {
		if val.Threat == m {
			id := val.ID
			val = *newVals
			val.ID = id
			val.UpdatedAt = time.Now()
		}
	}

	return nil
}

// SetThreat is a mock implementation of SetThreat method
func (mock *RiskDAOMock) SetThreat(m *models.Risk, newVal string) (*models.Risk, error) {
	edit := mock.db[m.ID]
	edit.Threat = newVal
	edit.UpdatedAt = time.Now()

	mock.db[m.ID] = edit
	return &edit, nil
}


// ReadByID will find models.Risk by ID given by parameter
func (dao *RiskDAO) ReadByID(id uint) (*models.Risk, error) {
	m := &models.Risk{}
	if err := dao.db.First(&m, id).Error; err != nil {
		return nil, err
	}

	return m, nil
}


// RiskDAOMock is a mock DAO
type RiskDAOMock struct{
	db map[uint]models.Risk
	lastID uint
}

// NewRiskDAOMock is a factory function for NewRiskDAOMock
func NewRiskDAOMock() *RiskDAOMock{
	return &RiskDAOMock{
		db: make(map[uint]models.Risk),
	}
}

// Create will put a model into mock in-memory DB
func (mock *RiskDAOMock) Create(m *models.Risk) (error) {
	created := false
	m.CreatedAt = time.Now()
	for !created{
		mock.lastID++
		if _, exists := mock.db[mock.lastID]; !exists {
			m.ID = mock.lastID
			mock.db[mock.lastID] = *m
			created = true
		}
	}
	return nil
}

type RiskDAOInterface interface {
	 Create(m *models.Risk) (error)
	 Read(m *models.Risk) ([]models.Risk, error)
	 ReadT(m *models.Risk) (*gorm.DB, error)
	 ReadByIDT(id uint) (*gorm.DB, error)
	 Update(m *models.Risk, id uint) (*models.Risk, error)
	 UpdateAllFields(m *models.Risk) (*models.Risk, error)
	 Delete(m *models.Risk) (error)
	 GetUpdatedAfter(timestamp time.Time) ([]models.Risk, error)
	 GetAll() ([]models.Risk, error)
	 ExecuteCustomQueryT(query string) (*gorm.DB, error)
	 SetStart(m *models.Risk, str time.Time) (*models.Risk, error)
	 ReadByTrigger(m string) ([]models.Risk, error)
	 ReadByTriggerT(m string) (*gorm.DB, error)
	 DeleteByTrigger(m string) (error)
	 EditByTrigger(m string, newVals *models.Risk) (error)
	 SetTrigger(m *models.Risk, newVal string) (*models.Risk, error)
	 ReadByImpact(m float64) ([]models.Risk, error)
	 ReadByImpactT(m float64) (*gorm.DB, error)
	 DeleteByImpact(m float64) (error)
	 EditByImpact(m float64, newVals *models.Risk) (error)
	 SetImpact(m *models.Risk, newVal float64) (*models.Risk, error)
	 SetEnd(m *models.Risk, str time.Time) (*models.Risk, error)
	 AddProjectsAssociation(m *models.Risk, asocVal *models.Project) (*models.Risk, error)
	 RemoveProjectsAssociation(m *models.Risk, asocVal *models.Project) (*models.Risk, error)
	 GetAllAssociatedProjects(m *models.Risk) ([]models.Project, error)
	 ReadByProbability(m float64) ([]models.Risk, error)
	 ReadByProbabilityT(m float64) (*gorm.DB, error)
	 DeleteByProbability(m float64) (error)
	 EditByProbability(m float64, newVals *models.Risk) (error)
	 SetProbability(m *models.Risk, newVal float64) (*models.Risk, error)
	 ReadByCategory(m string) ([]models.Risk, error)
	 ReadByCategoryT(m string) (*gorm.DB, error)
	 DeleteByCategory(m string) (error)
	 EditByCategory(m string, newVals *models.Risk) (error)
	 SetCategory(m *models.Risk, newVal string) (*models.Risk, error)
	 ReadByStatus(m string) ([]models.Risk, error)
	 ReadByStatusT(m string) (*gorm.DB, error)
	 DeleteByStatus(m string) (error)
	 EditByStatus(m string, newVals *models.Risk) (error)
	 SetStatus(m *models.Risk, newVal string) (*models.Risk, error)
	 ReadByDescription(m string) ([]models.Risk, error)
	 ReadByDescriptionT(m string) (*gorm.DB, error)
	 DeleteByDescription(m string) (error)
	 EditByDescription(m string, newVals *models.Risk) (error)
	 SetDescription(m *models.Risk, newVal string) (*models.Risk, error)
	 AddCounterMeasuresAssociation(m *models.Risk, asocVal *models.CounterMeasure) (*models.Risk, error)
	 RemoveCounterMeasuresAssociation(m *models.Risk, asocVal *models.CounterMeasure) (*models.Risk, error)
	 GetAllAssociatedCounterMeasures(m *models.Risk) ([]models.CounterMeasure, error)
	 ReadByValue(m float64) ([]models.Risk, error)
	 ReadByValueT(m float64) (*gorm.DB, error)
	 DeleteByValue(m float64) (error)
	 EditByValue(m float64, newVals *models.Risk) (error)
	 SetValue(m *models.Risk, newVal float64) (*models.Risk, error)
	 ReadByRisk(m float64) ([]models.Risk, error)
	 ReadByRiskT(m float64) (*gorm.DB, error)
	 DeleteByRisk(m float64) (error)
	 EditByRisk(m float64, newVals *models.Risk) (error)
	 SetRisk(m *models.Risk, newVal float64) (*models.Risk, error)
	 ReadByUserID(m uint) ([]models.Risk, error)
	 ReadByUserIDT(m uint) (*gorm.DB, error)
	 DeleteByUserID(m uint) (error)
	 EditByUserID(m uint, newVals *models.Risk) (error)
	 SetUserID(m *models.Risk, newVal uint) (*models.Risk, error)
	 ReadByCost(m int) ([]models.Risk, error)
	 ReadByCostT(m int) (*gorm.DB, error)
	 DeleteByCost(m int) (error)
	 EditByCost(m int, newVals *models.Risk) (error)
	 SetCost(m *models.Risk, newVal int) (*models.Risk, error)
	 ReadByName(m string) ([]models.Risk, error)
	 ReadByNameT(m string) (*gorm.DB, error)
	 DeleteByName(m string) (error)
	 EditByName(m string, newVals *models.Risk) (error)
	 SetName(m *models.Risk, newVal string) (*models.Risk, error)
	 ReadByThreat(m string) ([]models.Risk, error)
	 ReadByThreatT(m string) (*gorm.DB, error)
	 DeleteByThreat(m string) (error)
	 EditByThreat(m string, newVals *models.Risk) (error)
	 SetThreat(m *models.Risk, newVal string) (*models.Risk, error)
	 ReadByID(id uint) (*models.Risk, error)
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

// Read is a mock implementation of Read method
func (mock *CounterMeasureDAOMock) Read(m *models.CounterMeasure) ([]models.CounterMeasure, error) {
	ret := make([]models.CounterMeasure, 0, len(mock.db))
	for range mock.db {
		/*if val == *m {
			ret = append(ret, val)
		}*/
	}

	return ret, nil
}


// ReadByID is a mock implementation of ReadByID method
func (mock *CounterMeasureDAOMock) ReadByID(id uint) (*models.CounterMeasure, error) {
	ret := mock.db[id]
	return &ret, nil
}

// ReadByIDT is a mock implementation of ReadByIDT method
func (mock *CounterMeasureDAOMock) ReadByIDT(id uint) (*gorm.DB, error) {
	return nil, nil
}

// ReadT is a mock implementation of ReadT method
func (mock *CounterMeasureDAOMock) ReadT(m *models.CounterMeasure) (*gorm.DB, error) {
	return nil, nil
}

// Update is a mock implementation of Update method
func (mock *CounterMeasureDAOMock) Update(m *models.CounterMeasure, id uint) (*models.CounterMeasure, error) {
	m.UpdatedAt = time.Now()
	mock.db[id] = *m
	return m, nil
}

// UpdateAllFields is a mock implementation of UpdateAllFields method
func (mock *CounterMeasureDAOMock) UpdateAllFields(m *models.CounterMeasure) (*models.CounterMeasure, error) {
	m.UpdatedAt = time.Now()
	mock.db[m.ID] = *m
	return m, nil
}

// Delete is a mock implementation of Delete method
func (mock *CounterMeasureDAOMock) Delete(m *models.CounterMeasure) (error) {
	delete(mock.db, m.ID)
	return nil
}

// GetUpdatedAfter is a mock implementation of GetUpdatedAfter method
func (mock *CounterMeasureDAOMock) GetUpdatedAfter(timestamp time.Time) ([]models.CounterMeasure, error) {
	ret := make([]models.CounterMeasure, 0, len(mock.db))
	for _, val :=  range mock.db {
		if val.UpdatedAt.After(timestamp) {
			ret = append(ret, val)
		}
	}

	return ret, nil
}

// GetAll is a mock implementation of GetAll method
func (mock *CounterMeasureDAOMock) GetAll() ([]models.CounterMeasure, error) {
	ret := make([]models.CounterMeasure, 0, len(mock.db))
	for _, val := range mock.db {
		ret = append(ret, val)
	}

	return ret, nil
}

// ExecuteCustomQueryT is a mock implementation of ExecuteCustomQueryT method
func (mock *CounterMeasureDAOMock) ExecuteCustomQueryT(query string) (*gorm.DB, error) {
	return nil, nil
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

// ReadByName is a mock implementation of ReadByName method
func (mock *CounterMeasureDAOMock) ReadByName(m string) ([]models.CounterMeasure, error) {
	ret := make([]models.CounterMeasure, 0, len(mock.db))
	for _, val := range mock.db {
		if val.Name == m {
			ret = append(ret, val)
		}
	}

	return ret, nil
}

// ReadByNameT is a mock implementation of ReadByNameT method
func (mock *CounterMeasureDAOMock) ReadByNameT(m string) (*gorm.DB, error) {
	return nil, nil
}

// DeleteByName is a mock implementation of DeleteByName method
func (mock *CounterMeasureDAOMock) DeleteByName(m string) (error) {
	for _, val := range mock.db {
		if val.Name == m {
			delete(mock.db, val.ID)
		}
	}

	return nil
}

// EditByName is a mock implementation of EditByName method
func (mock *CounterMeasureDAOMock) EditByName(m string, newVals *models.CounterMeasure) (error) {
	for _, val := range mock.db {
		if val.Name == m {
			id := val.ID
			val = *newVals
			val.ID = id
			val.UpdatedAt = time.Now()
		}
	}

	return nil
}

// SetName is a mock implementation of SetName method
func (mock *CounterMeasureDAOMock) SetName(m *models.CounterMeasure, newVal string) (*models.CounterMeasure, error) {
	edit := mock.db[m.ID]
	edit.Name = newVal
	edit.UpdatedAt = time.Now()

	mock.db[m.ID] = edit
	return &edit, nil
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

// ReadByDescription is a mock implementation of ReadByDescription method
func (mock *CounterMeasureDAOMock) ReadByDescription(m string) ([]models.CounterMeasure, error) {
	ret := make([]models.CounterMeasure, 0, len(mock.db))
	for _, val := range mock.db {
		if val.Description == m {
			ret = append(ret, val)
		}
	}

	return ret, nil
}

// ReadByDescriptionT is a mock implementation of ReadByDescriptionT method
func (mock *CounterMeasureDAOMock) ReadByDescriptionT(m string) (*gorm.DB, error) {
	return nil, nil
}

// DeleteByDescription is a mock implementation of DeleteByDescription method
func (mock *CounterMeasureDAOMock) DeleteByDescription(m string) (error) {
	for _, val := range mock.db {
		if val.Description == m {
			delete(mock.db, val.ID)
		}
	}

	return nil
}

// EditByDescription is a mock implementation of EditByDescription method
func (mock *CounterMeasureDAOMock) EditByDescription(m string, newVals *models.CounterMeasure) (error) {
	for _, val := range mock.db {
		if val.Description == m {
			id := val.ID
			val = *newVals
			val.ID = id
			val.UpdatedAt = time.Now()
		}
	}

	return nil
}

// SetDescription is a mock implementation of SetDescription method
func (mock *CounterMeasureDAOMock) SetDescription(m *models.CounterMeasure, newVal string) (*models.CounterMeasure, error) {
	edit := mock.db[m.ID]
	edit.Description = newVal
	edit.UpdatedAt = time.Now()

	mock.db[m.ID] = edit
	return &edit, nil
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

// ReadByCost is a mock implementation of ReadByCost method
func (mock *CounterMeasureDAOMock) ReadByCost(m int) ([]models.CounterMeasure, error) {
	ret := make([]models.CounterMeasure, 0, len(mock.db))
	for _, val := range mock.db {
		if val.Cost == m {
			ret = append(ret, val)
		}
	}

	return ret, nil
}

// ReadByCostT is a mock implementation of ReadByCostT method
func (mock *CounterMeasureDAOMock) ReadByCostT(m int) (*gorm.DB, error) {
	return nil, nil
}

// DeleteByCost is a mock implementation of DeleteByCost method
func (mock *CounterMeasureDAOMock) DeleteByCost(m int) (error) {
	for _, val := range mock.db {
		if val.Cost == m {
			delete(mock.db, val.ID)
		}
	}

	return nil
}

// EditByCost is a mock implementation of EditByCost method
func (mock *CounterMeasureDAOMock) EditByCost(m int, newVals *models.CounterMeasure) (error) {
	for _, val := range mock.db {
		if val.Cost == m {
			id := val.ID
			val = *newVals
			val.ID = id
			val.UpdatedAt = time.Now()
		}
	}

	return nil
}

// SetCost is a mock implementation of SetCost method
func (mock *CounterMeasureDAOMock) SetCost(m *models.CounterMeasure, newVal int) (*models.CounterMeasure, error) {
	edit := mock.db[m.ID]
	edit.Cost = newVal
	edit.UpdatedAt = time.Now()

	mock.db[m.ID] = edit
	return &edit, nil
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

// AddRisksAssociation is a mock implementation of AddRisksAssociation method
func (mock *CounterMeasureDAOMock) AddRisksAssociation(m *models.CounterMeasure, asocVal *models.Risk) (*models.CounterMeasure, error) {
	edit := mock.db[m.ID]
	edit.Risks = append(edit.Risks, *asocVal)
	edit.UpdatedAt = time.Now()
	mock.db[m.ID] = edit

	return &edit, nil
}

// RemoveRisksAssociation is a mock implementation of RemoveRisksAssociation method
func (mock *CounterMeasureDAOMock) RemoveRisksAssociation(m *models.CounterMeasure, asocVal *models.Risk) (*models.CounterMeasure, error) {
	a := m.Risks
	m.UpdatedAt = time.Now()
	deletedIndex := 0
	for j, val := range a {
		if val.ID == asocVal.ID {
			deletedIndex = j
		}
	}
	a[deletedIndex] = a[len(a)-1]
	a = a[:len(a)-1]

	return m, nil
}

// GetAllAssociatedRisks is a mock implementation of GetAllAssociatedRisks method
func (mock *CounterMeasureDAOMock) GetAllAssociatedRisks(m *models.CounterMeasure) ([]models.Risk, error) {
	return m.Risks, nil
}


// ReadByID will find models.CounterMeasure by ID given by parameter
func (dao *CounterMeasureDAO) ReadByID(id uint) (*models.CounterMeasure, error) {
	m := &models.CounterMeasure{}
	if err := dao.db.First(&m, id).Error; err != nil {
		return nil, err
	}

	return m, nil
}


// CounterMeasureDAOMock is a mock DAO
type CounterMeasureDAOMock struct{
	db map[uint]models.CounterMeasure
	lastID uint
}

// NewCounterMeasureDAOMock is a factory function for NewCounterMeasureDAOMock
func NewCounterMeasureDAOMock() *CounterMeasureDAOMock{
	return &CounterMeasureDAOMock{
		db: make(map[uint]models.CounterMeasure),
	}
}

// Create will put a model into mock in-memory DB
func (mock *CounterMeasureDAOMock) Create(m *models.CounterMeasure) (error) {
	created := false
	m.CreatedAt = time.Now()
	for !created{
		mock.lastID++
		if _, exists := mock.db[mock.lastID]; !exists {
			m.ID = mock.lastID
			mock.db[mock.lastID] = *m
			created = true
		}
	}
	return nil
}

type CounterMeasureDAOInterface interface {
	 Create(m *models.CounterMeasure) (error)
	 Read(m *models.CounterMeasure) ([]models.CounterMeasure, error)
	 ReadT(m *models.CounterMeasure) (*gorm.DB, error)
	 ReadByIDT(id uint) (*gorm.DB, error)
	 Update(m *models.CounterMeasure, id uint) (*models.CounterMeasure, error)
	 UpdateAllFields(m *models.CounterMeasure) (*models.CounterMeasure, error)
	 Delete(m *models.CounterMeasure) (error)
	 GetUpdatedAfter(timestamp time.Time) ([]models.CounterMeasure, error)
	 GetAll() ([]models.CounterMeasure, error)
	 ExecuteCustomQueryT(query string) (*gorm.DB, error)
	 ReadByName(m string) ([]models.CounterMeasure, error)
	 ReadByNameT(m string) (*gorm.DB, error)
	 DeleteByName(m string) (error)
	 EditByName(m string, newVals *models.CounterMeasure) (error)
	 SetName(m *models.CounterMeasure, newVal string) (*models.CounterMeasure, error)
	 ReadByDescription(m string) ([]models.CounterMeasure, error)
	 ReadByDescriptionT(m string) (*gorm.DB, error)
	 DeleteByDescription(m string) (error)
	 EditByDescription(m string, newVals *models.CounterMeasure) (error)
	 SetDescription(m *models.CounterMeasure, newVal string) (*models.CounterMeasure, error)
	 ReadByCost(m int) ([]models.CounterMeasure, error)
	 ReadByCostT(m int) (*gorm.DB, error)
	 DeleteByCost(m int) (error)
	 EditByCost(m int, newVals *models.CounterMeasure) (error)
	 SetCost(m *models.CounterMeasure, newVal int) (*models.CounterMeasure, error)
	 AddRisksAssociation(m *models.CounterMeasure, asocVal *models.Risk) (*models.CounterMeasure, error)
	 RemoveRisksAssociation(m *models.CounterMeasure, asocVal *models.Risk) (*models.CounterMeasure, error)
	 GetAllAssociatedRisks(m *models.CounterMeasure) ([]models.Risk, error)
	 ReadByID(id uint) (*models.CounterMeasure, error)
}