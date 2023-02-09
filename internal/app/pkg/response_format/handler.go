package responseFormat

import "github.com/cahyacaa/test-julo/internal/app/domain"

func HandleSuccess[responseData any](data responseData) *domain.Response[responseData] {
	return &domain.Response[responseData]{StatusCode: 200, Status: "success", Data: data}
}

func HandleError[responseData any](data responseData, statusCode int) *domain.Response[responseData] {
	return &domain.Response[responseData]{StatusCode: statusCode, Status: "fail", Data: data}
}
