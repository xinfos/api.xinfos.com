package repository

import (
	"fmt"

	"api.xinfos.com/internal/model"
	"api.xinfos.com/internal/repository/cache"
	"api.xinfos.com/utils/errs"
)

//SAttrTemplateRepository - System defined standard classification attribute template
type SAttrTemplateRepository struct {
	c *cache.SAttrGroupCache
}

//NewSAttrTemplateRepository - 初始化 NewSAttrTemplateRepository
func NewSAttrTemplateRepository() *SAttrTemplateRepository {
	return &SAttrTemplateRepository{
		c: cache.NewSAttrGroupCache(),
	}
}

//Create - 根据model模型创建记录
func (repo *SAttrTemplateRepository) Create(m *model.SAttrTempalte, generalAttrGroupIDs, specAttrIDs []uint64, generalAttrIDs map[uint64][]uint64) (uint64, *errs.Errs) {
	//1.判断当前分类信息是否存在
	if m.CatID <= 0 {
		return 0, errs.ErrAttrGroupCreateFailCateNotFound
	}
	cate, err := model.CategoryModel().FindByID(m.CatID)
	if err != nil || cate.CatID != m.CatID {
		return 0, errs.ErrAttrGroupCreateFailCateNotFound
	}

	//2.判断当前模板名称是否存在
	existsSAttrTemplate, err := m.FindByCatIDAndName(m.CatID, m.Name)
	if err != nil {
		return 0, errs.ErrAttrGroupCreateFailCateNotFound
	}
	if existsSAttrTemplate != nil && existsSAttrTemplate.ID > 0 {
		return existsSAttrTemplate.ID, nil
	}

	//3.判断规格属性是否存在
	toBebindSpecAttrIDs := []uint64{}
	if len(specAttrIDs) > 0 {
		existsSpecAttrs, err := model.SAttrModel().FindAllBySAttrIDs(specAttrIDs)
		if err != nil {

		}
		//将已存在的属性组进行绑定
		if len(existsSpecAttrs) > 0 {
			for _, v := range existsSpecAttrs {
				toBebindSpecAttrIDs = append(toBebindSpecAttrIDs, v.ID)
			}
		}
	}

	//4.判断当前的属性组是否存在, 如果不存在则忽略
	toBeBindAttrGroupIDs := []uint64{}
	toBeBindGeneralAttrIDs := map[uint64][]uint64{}
	if len(generalAttrGroupIDs) > 0 {
		existsSAttrGroup, err := model.SAttrGroupModel().FindBySGroupIDs(generalAttrGroupIDs)
		if err != nil {
			fmt.Println(err.Error())
			return 0, nil
		}
		//将已存在的属性组进行绑定, 并且将不存在的属性组公共属性从公共数组中剔除
		if len(existsSAttrGroup) > 0 {
			for _, v := range existsSAttrGroup {
				toBeBindAttrGroupIDs = append(toBeBindAttrGroupIDs, v.ID)
				if _, isOk := generalAttrIDs[v.ID]; isOk {
					toBeBindGeneralAttrIDs[v.ID] = generalAttrIDs[v.ID]
				}
			}
		}
	}
	m.IsDelete = 2
	err = m.Create(toBeBindAttrGroupIDs, toBebindSpecAttrIDs, toBeBindGeneralAttrIDs)
	if err != nil {
		fmt.Println(err.Error())
		return 0, errs.ErrAttrGroupCreateFail
	}
	if m == nil || m.ID <= 0 {
		return 0, errs.ErrAttrGroupCreateFail
	}
	return m.ID, nil

	return 0, nil
}
