package model

//SAttrGroupIdx -
type SAttrGroupIdx struct {
	SGroupID uint64 `json:"s_group_id"`
	SAttrID  uint64 `json:"s_attr_id"`
	IsDelete uint   `json:"-"`
}

var sAttrGroupIdx *SAttrGroupIdx

//SAttrGroupIdxModel 实例化 SAttrGroupIdx 模型
func SAttrGroupIdxModel() *SAttrGroupIdx {
	return sAttrGroupIdx
}

//TableName 返回对应的表名
func (t *SAttrGroupIdx) TableName() string {
	return `t_system_spu_attr_group_idx`
}
