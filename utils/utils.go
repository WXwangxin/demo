package utils

import (
	"github.com/kataras/iris/v12"
	"gorm.io/gorm"
)

func WrapPage(page, pageSize int, data interface{}, query *gorm.DB) iris.Map {
	var count int64
	query.Count(&count)
	offset := pageSize * (page - 1)
	if count > int64(offset) {
		query.Limit(pageSize).Offset(offset).Find(data)
	}
	/*result := map[string]interface{}{
		"total": count,
		"data":  data,
	}*/
	result := iris.Map{
		"total": count,
		"data":  data,
	}
	return result
}
