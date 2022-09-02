package models

import "time"

type Product struct {
	Id                   int       `xorm:"not null pk autoincr INT(11)"`
	CategoryId           int       `xorm:"default NULL comment('分类id') INT(11)"`
	CategoryName         string    `xorm:"default 'NULL' comment('分类名称') VARCHAR(100)"`
	BrandId              int       `xorm:"default NULL comment('品牌id') INT(11)"`
	Name                 string    `xorm:"default 'NULL' comment('商品名称') VARCHAR(255)"`
	ShortName            string    `xorm:"default 'NULL' comment('商品短标题') VARCHAR(100)"`
	Code                 string    `xorm:"default 'NULL' comment('商品编号') VARCHAR(36)"`
	BarCode              string    `xorm:"default 'NULL' comment('商品条码') VARCHAR(36)"`
	CreateDate           time.Time `xorm:"default 'current_timestamp()' comment('商品添加日期') DATETIME"`
	UpdateDate           time.Time `xorm:"default 'NULL' comment('商品最后更新日期') DATETIME"`
	IsDel                int       `xorm:"default 0 comment('是否删除') INT(11)"`
	IsGroup              int       `xorm:"default 0 comment('已经弃用此字段，用product_type判断商品是否为组合商品') INT(11)"`
	Pic                  string    `xorm:"default 'NULL' comment('商品图片（多图）') VARCHAR(1000)"`
	SellingPoint         string    `xorm:"default 'NULL' comment('商品卖点') VARCHAR(200)"`
	Video                string    `xorm:"default 'NULL' comment('商品视频地址') VARCHAR(500)"`
	ContentPc            string    `xorm:"default 'NULL' comment('电脑端详情内容') MEDIUMTEXT"`
	ContentMobile        string    `xorm:"default 'NULL' comment('手机端详情内容') MEDIUMTEXT"`
	IsDiscount           int       `xorm:"default NULL comment('是否参与优惠折扣') INT(11)"`
	ProductType          int       `xorm:"default NULL comment('商品类别（1-实物商品，2-虚拟商品，3-组合商品）') INT(11)"`
	IsUse                int       `xorm:"default 0 comment('商品是否被使用过（认领或者其它第三方使用，否则不能被删除商品本身及SKU）') INT(11)"`
	DelDate              time.Time `xorm:"default 'NULL' comment('删除时间') DATETIME"`
	ChannelId            string    `xorm:"default 'NULL' comment('渠道id，多渠道用逗号分隔') VARCHAR(100)"`
	IsDrugs              int       `xorm:"default 0 comment('是否为药品 0否，1是') INT(11)"`
	UseRange             string    `xorm:"default 'NULL' comment('商品应用范围（1电商，2前置仓，3门店仓）') VARCHAR(100)"`
	GroupType            int       `xorm:"default NULL comment('组合类型(1:实实组合,2:虚虚组合,3.虚实组合)只有是组合商品才有值') INT(11)"`
	TermType             int       `xorm:"default NULL comment('只有虚拟商品才有值(1.有效期至多少  2.有效期天数)') INT(11)"`
	TermValue            int       `xorm:"default NULL comment('如果term_type=1 存：时间戳  如果term_type=2 存多少天') INT(11)"`
	VirtualInvalidRefund int       `xorm:"default NULL comment('是否支持过期退款 1：是 0：否') INT(11)"`
	WarehouseType        int       `xorm:"default 0 comment('药品仓 1：巨星药品仓 0：否') INT(11)"`
	IsIntelGoods         int       `xorm:"default 0 comment('是否互联网医疗商品（1是，0否）') INT(11)"`
	Disabled             int       `xorm:"default 0 comment('oms商品同步启用禁用装填（1禁用，0启用）') INT(11)"`
	FromOms              int       `xorm:"default 0 comment('来自oms的商品（1是，0否）') INT(11)"`
	SourceType           int       `xorm:"default 0 comment('1：OMS同步 2：后台新增 3：后台导入') INT(11)"`
}
