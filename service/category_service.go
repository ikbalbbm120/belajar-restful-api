package service

import(
	"context"
	"belajar-golang-restful-api/model/web"
)

type CategoryService struct {
	Create(ctx context.Context, request web.CategoryCreateRequest) web.CategoryResponse
	Update(ctx context.Context, request web.CategoryUpdateRequest) web.CategoryResponse
	Delete(ctx context.Context, categoryId int)
	FindByID(ctx context.Context, categoryId int)  web.CategoryResponse
	FindAll(ctx context.Context) []web.CategoryResponse
}