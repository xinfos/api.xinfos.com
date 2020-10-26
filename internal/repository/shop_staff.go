package repository

import (
	"api.xinfos.com/internal/model"
	"api.xinfos.com/internal/repository/cache"
	"api.xinfos.com/utils/errs"
)

//ShopStaffList 员工列表
type ShopStaffList struct {
	List            []*model.ShopStaff `json:"list"`
	CurrentPageNo   uint               `json:"current_page_no"`
	CurrentPageSize uint               `json:"current_page_size"`
	TotalCount      int                `json:"total_count"`
}

//ShopStaffRepository - Shop Staff Repository
type ShopStaffRepository struct {
	c *cache.ShopStaffCache
}

//NewShopStaffRepository - Init ShopStaffRepository
func NewShopStaffRepository() *ShopStaffRepository {
	return &ShopStaffRepository{
		c: cache.NewShopStaffCache(),
	}
}

//Create - Create a new shop employee
func (repo *ShopStaffRepository) Create(m *model.ShopStaff) (uint64, *errs.Errs) {
	//1.Check if the shop ID is legal
	if m.ShopID <= 0 {
		return 0, errs.ErrCatParentIsNotFound
	}
	//2.Check the employee number exists
	if len(m.StaffNo) >= 0 {
		if m.IsStaffNoExists(m.StaffNo, m.ShopID) {
			return 0, errs.ErrCatNameIsExists
		}
	}
	//3.Create
	err := m.Create()
	if err != nil {
		return 0, errs.ErrCatCreateFail
	}
	return m.ID, nil
}

//Delete - Delete a shop employee
func (repo *ShopStaffRepository) Delete(id uint64) *errs.Errs {
	//1.Check if the staff is exists.
	m, err := model.ShopStaffModel().FindByID(id)
	if err != nil || m == nil || m.ID <= 0 {
		return errs.ErrCatDeleteIsNotFound
	}
	//2.Delete
	if err := m.Delete(); err != nil {
		return errs.ErrCatDeleteFail
	}
	return nil
}

//Update - Update a shop staff
func (repo *ShopStaffRepository) Update(m *model.ShopStaff) *errs.Errs {
	//1.Check if the category is exists.
	data, err := m.FindByID(m.ID)
	if err != nil || data == nil {
		return errs.ErrCatUpdateIsNotFound
	}
	//3.Update
	if err = m.Update(); err != nil {
		return errs.ErrCatUpdateFail
	}
	return nil
}

//FindByID - Query shop employee by ID
func (repo *ShopStaffRepository) FindByID(id uint64) (*model.ShopStaff, *errs.Errs) {
	data := repo.c.Get(id)
	if data != nil && data.ID > 0 {
		return data, nil
	}
	data, _ = model.ShopStaffModel().FindByID(id)
	if data != nil && data.ID == id {
		repo.c.Set(data)
	}
	return data, nil
}

//FindAll - Query shop employee
func (repo *ShopStaffRepository) FindAll(query string, args []interface{}, orderby string, page, pageSize uint) (*ShopStaffList, *errs.Errs) {
	data, count, err := model.ShopStaffModel().FindAllByQuery(query, args, orderby, "", page, pageSize)
	if err != nil {
		return nil, errs.ErrBrandCreateFail
	}
	l := &ShopStaffList{
		List:            data,
		CurrentPageNo:   page,
		CurrentPageSize: pageSize,
		TotalCount:      count,
	}
	return l, nil
}
