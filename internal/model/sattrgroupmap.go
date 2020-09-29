package model

import "errors"

//SAttrGroupMap -
type SAttrGroupMap struct {
	CatID uint64 `json:"cat_id"`
}

type SysSPUAttrGroup struct {
	GroupID uint64        `json:"group_id"`
	Name    string        `json:"group_name"`
	Attrs   []*SysSPUAttr `json:"attrs"`
}

type SysSPUAttr struct {
	AttrID      uint64           `json:"attr_id"`
	Name        string           `json:"attr_name"`
	FillType    uint             `json:"fill_type"`
	IsRequired  uint             `json:"is_required"`
	IsNumeric   uint             `json:"is_numeric"`
	Unit        string           `json:"unit"`
	IsGeneric   uint             `json:"is_generic"`
	IsSearching uint             `json:"is_searching"`
	Segments    string           `json:"segments"`
	Values      []*SysSPUAttrVal `json:"values"`
}

type SysSPUAttrVal struct {
	ValID uint64 `json:"val_id"`
	Val   string `json:"val"`
}

var sAttrGroupMap *SAttrGroupMap

func SAttrGroupMapModel() *SAttrGroupMap {
	return sAttrGroupMap
}

//FindByID -
func (t *SAttrGroupMap) FindByID(id uint64) (*SAttrGroupMap, error) {
	return nil, nil
}

//FindAllByCatID - 根据商品系统分类ID查询SPU系统关联属性
func (t *SAttrGroupMap) FindAllByCatID(id uint64) ([]*SysSPUAttrGroup, error) {
	//1、获取系统属性组
	sysAttrGroups, err := SAttrGroupModel().FindAllByCatID(id)
	if err != nil || len(sysAttrGroups) <= 0 {
		return nil, errors.New("1")
	}

	var sAttrGroupMaps []*SysSPUAttrGroup
	for _, v := range sysAttrGroups {
		sAttrGroupMaps = append(sAttrGroupMaps, &SysSPUAttrGroup{
			GroupID: v.ID,
			Name:    v.Name,
		})
	}

	//2、获取系统属性
	sysAttrs, err := SAttrModel().FindAllByCatID(id)
	if err != nil || len(sysAttrs) <= 0 {
		return nil, errors.New("2")
	}

	for _, v1 := range sAttrGroupMaps {
		for _, v2 := range sysAttrs {
			if v1.GroupID == v2.SGroupID {
				v1.Attrs = append(v1.Attrs, &SysSPUAttr{
					AttrID:      v2.SAttrID,
					Name:        v2.Name,
					FillType:    v2.FillType,
					IsRequired:  v2.IsRequired,
					IsNumeric:   v2.IsNumeric,
					Unit:        v2.Unit,
					IsGeneric:   v2.IsGeneric,
					IsSearching: v2.IsSearching,
					Segments:    v2.Segments,
				})
			}
		}
	}

	//3、获取系统属性值
	sysAttrValues, err := SAttrValModel().FindAllByCatID(id)
	if err != nil || len(sysAttrValues) <= 0 {
		return nil, errors.New("3")
	}

	for _, v1 := range sAttrGroupMaps {
		for _, v2 := range v1.Attrs {
			for _, v3 := range sysAttrValues {
				if v2.AttrID == v3.SAttrID {
					v2.Values = append(v2.Values, &SysSPUAttrVal{
						ValID: v3.SAttrValID,
						Val:   v3.Value,
					})
				}
			}
		}
	}
	return sAttrGroupMaps, nil
}
