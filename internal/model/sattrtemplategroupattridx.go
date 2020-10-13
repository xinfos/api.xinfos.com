package model

//STemplateGroupAttrIdx -
type STemplateGroupAttrIdx struct {
	SGroupID uint64 `json:"s_group_id"`
	SAttrID  uint64 `json:"s_attr_id"`
	IsDelete uint   `json:"-"`
}

var sTemplateGroupAttrIdx *STemplateGroupAttrIdx

//STemplateGroupAttrIdxModel 实例化 SAttr 模型
func STemplateGroupAttrIdxModel() *STemplateGroupAttrIdx {
	return sTemplateGroupAttrIdx
}

//TableName 返回对应的表名
func (t *STemplateGroupAttrIdx) TableName() string {
	return `t_sys_spu_template_group_attr_idx`
}
