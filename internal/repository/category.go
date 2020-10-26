package repository

import (
	"fmt"

	"api.xinfos.com/internal/model"
	"api.xinfos.com/internal/repository/cache"
	"api.xinfos.com/utils/errs"
	"github.com/go-ego/riot"
	"github.com/go-ego/riot/types"
)

var (
	// searcher is coroutine safe
	searcher = riot.Engine{}
)

//CategoryList - Category list strcut
type CategoryList struct {
	List            []*model.Category `json:"list"`
	CurrentPageNo   uint              `json:"current_page_no"`
	CurrentPageSize uint              `json:"current_page_size"`
	TotalCount      int               `json:"total_count"`
}

//CategoryRepository - The Category Repository
type CategoryRepository struct {
	c *cache.CategoryCache
}

//NewCategoryRepository - Init BrandRepository
func NewCategoryRepository() *CategoryRepository {
	return &CategoryRepository{
		c: cache.NewCategoryCache(),
	}
}

//Create - Create a category
func (repo *CategoryRepository) Create(m *model.Category) (uint64, *errs.Errs) {
	//1.Check that the pid is normal, 10000 is the top-level category ID
	var parent *model.Category
	if m.PID == 10000 {
		parent = m.GetRootCategory()
	} else {
		parent, _ = model.CategoryModel().FindByID(m.PID)
		if parent == nil || parent.CatID <= 0 {
			return 0, errs.ErrCatParentIsNotFound
		}
	}

	//2.Check if the category name already exists.
	if m.IsCategoryNameExists(m.Name) {
		return 0, errs.ErrCatNameIsExists
	}

	//3.set current category depth
	m.Depth = parent.Depth + 1

	//4.Create
	err := m.Create()
	if err != nil {
		return 0, errs.ErrCatCreateFail
	}
	return m.CatID, nil
}

//Delete - Delete a category
func (repo *CategoryRepository) Delete(id uint64) *errs.Errs {
	//1.Check if the category is exists.
	m, err := model.CategoryModel().FindByID(id)
	if err != nil || m == nil || m.CatID <= 0 {
		return errs.ErrCatDeleteIsNotFound
	}
	//2.Check is there a subcategory.
	if m.IsCategoryHadSubChild(id) {
		return errs.ErrCatDeleteHasSubCat
	}
	//3.Delete
	if err := m.Delete(); err != nil {
		return errs.ErrCatDeleteFail
	}
	return nil
}

//Update - Update a Category
func (repo *CategoryRepository) Update(m *model.Category) *errs.Errs {
	//1.Check if the category is exists.
	data, err := m.FindByID(m.CatID)
	if err != nil || data == nil {
		return errs.ErrCatUpdateIsNotFound
	}
	//2.Check if the category name already exists.
	c, _ := m.FindByName(m.Name)
	if c != nil && c.CatID > 0 && c.CatID != m.CatID {
		return errs.ErrCatNameIsExists
	}
	//3.Check the data that needs to be updated
	if data.Name != m.Name {
		data.Name = m.Name
	}
	if data.Alias != m.Alias {
		data.Alias = m.Alias
	}
	if data.Desc != m.Desc {
		data.Desc = m.Desc
	}
	if data.State != m.State {
		data.State = m.State
	}

	fmt.Println(data)
	fmt.Println(m)
	//4.Update
	if err = data.Update(); err != nil {
		return errs.ErrCatUpdateFail
	}
	repo.c.Del(data.CatID)
	return nil
}

//FindByID - Find category by cat_id
func (repo *CategoryRepository) FindByID(id uint64) (*model.Category, *errs.Errs) {
	data := repo.c.Get(id)
	if data != nil && data.CatID > 0 {
		return data, nil
	}
	data, err := model.CategoryModel().FindByID(id)
	if err != nil && data == nil && data.CatID != id {
		return nil, nil
	}
	repo.c.Set(data)
	return data, nil
}

//FindAllByPID - Find sub category list by PID
func (repo *CategoryRepository) FindAllByPID(id uint64) ([]*model.Category, *errs.Errs) {
	k := fmt.Sprintf("category:parent:%d", id)
	data := repo.c.GetAll(k)
	if data != nil && len(data) > 0 {
		return data, nil
	}
	data, _ = model.CategoryModel().FindAllByPID(id)
	if data != nil && len(data) > 0 {
		repo.c.SetAll(k, data)
	}
	return data, nil
}

//FindAll - Find category list by query
func (repo *CategoryRepository) FindAll(query string, args []interface{}, orderby string, page, pageSize uint) (*CategoryList, *errs.Errs) {
	data, count, err := model.CategoryModel().FindAllByQuery(query, args, orderby, "", page, pageSize)
	if err != nil {
		return nil, errs.ErrBrandCreateFail
	}
	l := &CategoryList{
		List:            data,
		CurrentPageNo:   page,
		CurrentPageSize: pageSize,
		TotalCount:      count,
	}
	return l, nil
}

//SearchByKeyword - Search category by keyword
func (repo *CategoryRepository) SearchByKeyword(keyword string) (string, *errs.Errs) {
	sea := searcher.Search(types.SearchReq{
		Text: "zl",
		RankOpts: &types.RankOpts{
			OutputOffset: 0,
			MaxOutputs:   100,
		}})
	fmt.Println("search response: ", sea, "; docs = ", sea.Docs)
	return keyword, nil
}
