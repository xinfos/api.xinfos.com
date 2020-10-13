package model

//SAttrTemplateGroupIdx -
type SAttrTemplateGroupIdx struct {
	STemplateID uint64 `json:"s_template_id"`
	SGroupID    uint64 `json:"s_group_id"`
	IsDelete    uint   `json:"-"`
}

var sAttrTemplateGroupIdx *SAttrTemplateGroupIdx

//SAttrTemplateGroupIdxModel 实例化 SAttrTemplateSpecIdx 模型
func SAttrTemplateGroupIdxModel() *SAttrTemplateGroupIdx {
	return sAttrTemplateGroupIdx
}

//TableName 返回对应的表名
func (t *SAttrTemplateGroupIdx) TableName() string {
	return `t_sys_spu_template_group_idx`
}
