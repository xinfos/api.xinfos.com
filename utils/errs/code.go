package errs

// const (
// 	// ErrSuccess      = 200    //ErrSuccess  - 请求成功
// 	ErrInternal     = 500    //ErrInternal - 内部错误
// 	ErrDBQuery      = 1001   //ErrDBQuery  - DB 查询异常
// 	ErrParamInvalid = 100001 //ErrParamInvalid 参数提交错误
// 	ErrParamVerify  = 100002 //ErrParamVerify 参数校验错误
// )

//ErrorMsg - customization error message
var ErrorMsg = map[int]string{
	200:    "请求成功",
	500:    "抱歉，提交参数错误",
	1001:   "抱歉，服务内部错误",
	100001: "抱歉，数据内部错误",
	100002: "抱歉，提交参数不合法",
	120001: "抱歉，分类不存在.",
	120002: "抱歉，父级分类不存在，请重新确认!~",
	120003: "抱歉，当前分类名称已存在, 请重新输入!~",
	120004: "抱歉，分类创建失败.",
	120005: "抱歉，当前分类删除失败，请重试!~",
	120006: "抱歉，当前您要删除的分类不存在，请重新确认!~",
	120007: "抱歉，当前分类下有子分类不能直接删除，请先删除子分类!~",
	120008: "抱歉，分类更新失败.",
	120009: "抱歉，更新失败，当前分类下有子分类不能直接移动，请先移动子分类!~",
	120010: "抱歉，更新失败，当前分类不存在,请重新确认!~",
	120011: "抱歉，更新失败，当前父级分类不能是自己，请重新确认!~",
	130001: "抱歉，没有找到相应的品牌信息!~",
	130002: "抱歉，创建失败.",
	130003: "抱歉，创建失败，没有找到相应的分类，请先确认分类信息!~",
	130004: "抱歉，创建失败，当前品牌名称已存在，请重新确认!~",
	130005: "抱歉，删除失败.",
	130006: "抱歉，删除失败，没有找到相应的品牌信息!~",
	130007: "抱歉，删除失败，当前品牌下含有商品，请先删除商品!~",
	130008: "抱歉，更新失败.",
	130009: "抱歉，更新失败，当前品牌名称已存在，请重新确认!~",
	130010: "抱歉，品牌获取失败，没有找到相应的分类信息，请重新确认!~",
	140001: "抱歉，当前分类没有对应的商品属性!~",
}

var (
	ErrSuccess                     = NewErrs(200, "请求成功")
	ErrInternal                    = NewErrs(500, "抱歉，提交参数错误.")
	ErrDBQuery                     = NewErrs(500, "抱歉，服务内部错误.")
	ErrParamInvalid                = NewErrs(500, "抱歉，数据内部错误.")
	ErrParamVerify                 = NewErrs(500, "抱歉，提交参数不合法.")
	ErrCatIsNotFound               = NewErrs(120001, "抱歉，分类不存在.")
	ErrCatParentIsNotFound         = NewErrs(120002, "抱歉，父级分类不存在，请重新确认!~")
	ErrCatNameIsExists             = NewErrs(120003, "抱歉，当前分类名称已存在, 请重新输入!~")
	ErrCatCreateFail               = NewErrs(120004, "抱歉，分类创建失败.")
	ErrCatDeleteFail               = NewErrs(120005, "抱歉，当前分类删除失败，请重试!~")
	ErrCatDeleteIsNotFound         = NewErrs(120006, "抱歉，当前您要删除的分类不存在，请重新确认!~")
	ErrCatDeleteHasSubCat          = NewErrs(120007, "抱歉，当前分类下有子分类不能直接删除，请先删除子分类!~")
	ErrCatUpdateFail               = NewErrs(120008, "抱歉，分类更新失败.")
	ErrCatUpdateHasSubCat          = NewErrs(120009, "抱歉，更新失败，当前分类下有子分类不能直接移动，请先移动子分类!~")
	ErrCatUpdateIsNotFound         = NewErrs(120010, "抱歉，更新失败，当前分类不存在,请重新确认!~")
	ErrCatUpdateParentCanNotSelf   = NewErrs(120011, "抱歉，更新失败，当前父级分类不能是自己，请重新确认!~")
	ErrBrandNotFound               = NewErrs(130001, "抱歉，没有找到相应的品牌信息!~")
	ErrBrandCreateFail             = NewErrs(130002, "抱歉，创建失败.")
	ErrBrandCreateFailCateNotFound = NewErrs(130003, "抱歉，创建失败，没有找到相应的分类，请先确认分类信息!~")
	ErrBrandCreateFailNameIsExists = NewErrs(130004, "抱歉，创建失败，当前品牌名称已存在，请重新确认!~")
	ErrBrandDeleteFail             = NewErrs(130005, "抱歉，删除失败.")
	ErrBrandDeleteFailNotFound     = NewErrs(130006, "抱歉，删除失败，没有找到相应的品牌信息!~")
	ErrBrandDeleteFailHasProduct   = NewErrs(130007, "抱歉，删除失败，当前品牌下含有商品，请先删除商品!~")
	ErrBrandUpdateFail             = NewErrs(130008, "抱歉，更新失败.")
	ErrBrandUpdateFailNameIsExists = NewErrs(130009, "抱歉，更新失败，当前品牌名称已存在，请重新确认!~")
	ErrGetFailCateIsNotExists      = NewErrs(130010, "抱歉，品牌获取失败，没有找到相应的分类信息，请重新确认!~")
	ErrProductNoAttr               = NewErrs(140001, "抱歉，当前分类没有对应的商品属性!~")
)
