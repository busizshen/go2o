package parser

import (
	"github.com/ixre/gof/types"
	"go2o/core/domain/interface/item"
	"go2o/core/domain/interface/valueobject"
	"go2o/core/service/proto"
)

/**
 * Copyright (C) 2007-2020 56X.NET,All rights reserved.
 *
 * name : item_dto.go
 * author : jarrysix (jarrysix#gmail.com)
 * date : 2020-09-19 13:50
 * description :
 * history :
 */

func ParseSaleLabelDto(src *item.Label) *proto.SItemLabel {
	return &proto.SItemLabel{
		Id:                   src.Id,
		Name:                 src.TagName,
		TagCode:              src.TagCode,
		LabelImage:           src.LabelImage,
		Enabled:              src.Enabled == 1,
		XXX_NoUnkeyedLiteral: struct{}{},
		XXX_unrecognized:     nil,
		XXX_sizecache:        0,
	}
}

func FromSaleLabelDto(src *proto.SItemLabel) *item.Label {
	return &item.Label{
		Id:         src.Id,
		MerchantId: -1,
		TagCode:    src.TagCode,
		TagName:    src.Name,
		LabelImage: src.LabelImage,
		Enabled:    types.IntCond(src.Enabled, 1, 0),
	}
}

func PriceArrayDto(src []*item.MemberPrice) []*proto.SLevelPrice {
	dst := make([]*proto.SLevelPrice,len(src))
	for i,v := range src{
		dst[i] = LevelPriceDto(v)
	}
	return dst
}

func LevelPriceDto(src *item.MemberPrice) *proto.SLevelPrice {
	return &proto.SLevelPrice{
		Id:                   int64(src.Id),
		Level:                int32(src.Level),
		Price:                int64(src.Price*100),
		MaxNumber:            int32(src.MaxQuota),
		Enabled:              src.Enabled == 1,
	}
}

func ParseLevelPrice(src *proto.SLevelPrice)*item.MemberPrice{
	return &item.MemberPrice{
		Id:       int(src.Id),
		Level:    int(src.Level),
		Price:    float32(src.Price)/100,
		MaxQuota: int(src.MaxNumber),
		Enabled:  types.IntCond(src.Enabled,1,0),
	}
}

func ParseGoodsDto_(src *valueobject.Goods) *proto.SUnifiedViewItem {
	return &proto.SUnifiedViewItem{
		ItemId:               src.ItemId,
		ProductId:            src.ProductId,
		CategoryId:           src.CategoryId,
		VendorId:             int64(src.VendorId),
		BrandId:              0,
		Title:                src.Title,
		Code:                 "",
		SkuId:                src.SkuId,
		Image:                src.Image,
		Price:                float64(src.Price),
		PriceRange:           src.PriceRange,
		StockNum:             src.StockNum,
		ShelveState:          item.ShelvesOn,
		ReviewState:          0,
		UpdateTime:           0,
	}
	return nil
}