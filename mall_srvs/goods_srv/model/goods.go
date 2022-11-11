package model

//Category 类型
// 实际开发过程中,尽量设置为不为null
// https://zhuanlan.zhihu.com/p/73997266
type Category struct {
	BaseModel
	Name             string      `gorm:"type:varchar(20);not null" json:"name"`
	ParentCategoryID int32       `gorm:"int(11);default:null" json:"parent_category_id"`
	ParentCategory   *Category   `json:"-"`
	SubCategory      []*Category `gorm:"foreignKey:ParentCategoryID;references:ID" json:"sub_category"`
	Level            int32       `gorm:"type:int;not null;default:1" json:"level"`
	IsTab            bool        `gorm:"default:false;not null" json:"is_tab"`
	Url              string      `gorm:"type:varchar(300);not null" json:"url"`
}

// Brands 品牌
type Brands struct {
	BaseModel

	Name string `gorm:"type:varchar(20);not null"`
	Logo string `gorm:"type:varchar(200);default:'';not null"`
}

// GoodsCategoryBrand 商品类别品牌
type GoodsCategoryBrand struct {
	BaseModel
	CategoryID int32 `gorm:"type:int;index:goodscategorybrand_category_id_brand_id,unique"`
	Category   Category

	BrandsID int32 `gorm:"type:int;index:goodscategorybrand_category_id_brand_id,unique"`
	Brands   Brands
}

func (GoodsCategoryBrand) TableName() string {
	return "goodscategorybrand"
}

// Banner 轮播图管理
type Banner struct {
	BaseModel
	Image string `gorm:"type:varchar(200);not null"`
	Url   string `gorm:"type:varchar(200);not null"`
	Index int32  `gorm:"type:int;default:1;not null"`
}

type Goods struct {
	BaseModel

	CategoryID int32 `gorm:"type:int;index:idx_category_brand,unique"`
	Category   Category
	BrandsID   int32 `gorm:"type:int;index:idx_category_brand,unique"`
	Brands     Brands

	OnSale   bool `gorm:"default:false;not null"` //是否上架
	ShipFree bool `gorm:"default:false;not null"` //是否免运费
	IsNew    bool `gorm:"default:false;not null"` //是否是新商品
	IsHot    bool `gorm:"default:false;not null"` //是否是热度商品

	Name             string   `gorm:"type:varchar(50);not null"`   //商品名
	GoodsSn          string   `gorm:"type:varchar(50);not null"`   //商品标签
	ClickNum         int32    `gorm:"type:int;default:0;not null"` //点击次数
	SoldNum          int32    `gorm:"type:int;default:0;not null"` //销售数目
	FavNum           int32    `gorm:"type:int;default:0;not null"` //喜欢人数
	MarketPrice      float32  `gorm:"not null"`                    // 原价格
	ShopPrice        float32  `gorm:"not null"`                    //销售价格(优惠价格)
	GoodsBrief       string   `gorm:"type:varchar(100);not null"`  // 商品简介
	Images           GormList `gorm:"type:varchar(1000);not null"` //照片集合
	DescImages       GormList `gorm:"type:varchar(1000);not null"` //下拉照片详情集合
	GoodsFrontImages string   `gorm:"type:varchar(1000);not null"` //商品展示得图片
}
