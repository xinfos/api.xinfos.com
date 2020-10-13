package model

//SAttrTemplateSpecIdx -
type SAttrTemplateSpecIdx struct {
	STemplateID uint64 `json:"s_template_id"`
	SAttrID     uint64 `json:"s_attr_id"`
	IsDelete    uint   `json:"-"`
}

var sAttrTemplateSpecIdx *SAttrTemplateSpecIdx

//SAttrTemplateSpecIdxModel 实例化 SAttrTemplateSpecIdx 模型
func SAttrTemplateSpecIdxModel() *SAttrTemplateSpecIdx {
	return sAttrTemplateSpecIdx
}

//TableName 返回对应的表名
func (t *SAttrTemplateSpecIdx) TableName() string {
	return `t_sys_spu_template_spec_idx`
}
