package services

import (
	"fmt"
	"golang-template/dto"
	"golang-template/middleware"
	"golang-template/models"
	"golang-template/repository"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type ProductService interface {
	UploadImage(c echo.Context) (string, error)
	CreateProduct(c echo.Context) error
	GetProductById(id string) (*models.Product, error)
	GetProductByStoreId(id string, desc, page, pageSize int, search, sort string) (*[]models.Product, *dto.Pagination, error)
	GetAllProduct(desc, page, pageSize int, search, sort string) (*[]models.Product, *dto.Pagination, error)
	UpdateProduct(c echo.Context, id string) error
	DeleteProduct(id string) error
	CheckProductStore(id string, c echo.Context) error
	AdvertiseProduct(c echo.Context, id string) error
	UnadvertiseProduct(c echo.Context, id string) error
}

type productService struct {
	productRepo  repository.ProductRepository
	uploader     *ClientUploader
	storeService StoreService
	ptService    ProductTypeService
	tokenRepo    repository.TokenRepository
}

func NewProductService() ProductService {
	return &productService{
		productRepo:  repository.NewProductRepositoryGORM(),
		uploader:     NewClientUploader(),
		storeService: NewStoreService(),
		tokenRepo:    repository.NewTokenRepositoryGORM(),
		ptService:    NewProductTypeService(),
	}
}

func (service *productService) UploadImage(c echo.Context) (string, error) {
	file, err := c.FormFile("file")
	if err != nil {
		return "", err
	}

	src, err := file.Open()
	if err != nil {
		return "", err
	}
	defer src.Close()

	storeUsername, err := service.storeService.GetStoreUsernameByToken(c)
	if err != nil {
		return "", err
	}

	randomNumber := strconv.Itoa(rand.Intn(1000))
	fileName := fmt.Sprintf("%s-%s-%s", storeUsername, file.Filename, randomNumber)
	err = service.uploader.UploadImage(src, fileName)
	if err != nil {
		return "", err
	}

	uploadedImagePath := service.uploader.uploadPath + fileName
	return uploadedImagePath, nil
}

func ParseProductForm(c echo.Context) (*dto.ProductRegisterForm, error) {
	productName := c.FormValue("product_name")
	productPriceStr := c.FormValue("product_price")
	productDesc := c.FormValue("product_desc")
	productIsShowStr := c.FormValue("product_isshow")
	productLemakTotalStr := c.FormValue("product_lemaktotal")
	productProteinStr := c.FormValue("product_protein")
	productKarbohidratStr := c.FormValue("product_karbohidrat")
	productGaramStr := c.FormValue("product_garam")
	productGrade := c.FormValue("product_grade")
	productServingSizeStr := c.FormValue("product_servingsize")
	ptName := c.FormValue("pt_name")

	productPrice, err := strconv.ParseFloat(productPriceStr, 64)
	if err != nil {
		return nil, echo.NewHTTPError(http.StatusBadRequest, "Invalid product price")
	}
	productIsShow, err := strconv.ParseBool(productIsShowStr)
	if err != nil {
		return nil, echo.NewHTTPError(http.StatusBadRequest, "Invalid product isShow value")
	}
	productLemakTotal, err := strconv.ParseFloat(productLemakTotalStr, 64)
	if err != nil {
		return nil, echo.NewHTTPError(http.StatusBadRequest, "Invalid product lemak total")
	}
	productProtein, err := strconv.ParseFloat(productProteinStr, 64)
	if err != nil {
		return nil, echo.NewHTTPError(http.StatusBadRequest, "Invalid product protein")
	}
	productKarbohidrat, err := strconv.ParseFloat(productKarbohidratStr, 64)
	if err != nil {
		return nil, echo.NewHTTPError(http.StatusBadRequest, "Invalid product karbohidrat")
	}
	productGaram, err := strconv.ParseFloat(productGaramStr, 64)
	if err != nil {
		return nil, echo.NewHTTPError(http.StatusBadRequest, "Invalid product garam")
	}
	productServingSize, err := strconv.ParseFloat(productServingSizeStr, 64)
	if err != nil {
		return nil, echo.NewHTTPError(http.StatusBadRequest, "Invalid product serving size")
	}

	return &dto.ProductRegisterForm{
		ProductName:        productName,
		ProductPrice:       productPrice,
		ProductDesc:        productDesc,
		ProductIsShow:      productIsShow,
		ProductLemakTotal:  productLemakTotal,
		ProductProtein:     productProtein,
		ProductKarbohidrat: productKarbohidrat,
		ProductGaram:       productGaram,
		ProductGrade:       productGrade,
		ProductServingSize: productServingSize,
		ProductExpShow:     "",
		PT_Name:            ptName,
	}, nil
}

func (service *productService) CreateProduct(c echo.Context) error {
	productForm, err := ParseProductForm(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	imagePath, err := service.UploadImage(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	UserToken, err := service.tokenRepo.UserToken(middleware.GetToken(c))
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, "Unauthorized")
	}

	store, err := service.storeService.GetStoreByUserId(UserToken.ID)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "Store not found")
	}

	pt_id, err := service.ptService.GetProductTypeIdByName(productForm.PT_Name)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "Product type not found")
	}

	product := models.Product{
		PRODUCT_ID:          uuid.New().String(),
		PRODUCT_NAME:        productForm.ProductName,
		PRODUCT_PRICE:       productForm.ProductPrice,
		PRODUCT_DESC:        productForm.ProductDesc,
		PRODUCT_ISSHOW:      productForm.ProductIsShow,
		PRODUCT_LEMAKTOTAL:  productForm.ProductLemakTotal,
		PRODUCT_PROTEIN:     productForm.ProductProtein,
		PRODUCT_KARBOHIDRAT: productForm.ProductKarbohidrat,
		PRODUCT_GARAM:       productForm.ProductGaram,
		PRODUCT_SERVINGSIZE: productForm.ProductServingSize,
		PRODUCT_PICTURE:     imagePath,
		PRODUCT_GRADING:     productForm.ProductGrade,
		PRODUCT_EXPSHOW:     time.Now(),
		CreatedAt:           time.Now(),
		UpdatedAt:           time.Now(),
		STORE_ID:            store.STORE_ID,
		PT_ID:               pt_id,
	}

	err = service.productRepo.CreateProduct(&product)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusCreated, product)
}

func (service *productService) GetProductById(id string) (*models.Product, error) {
	return service.productRepo.GetProductById(id)
}

func (service *productService) GetProductByStoreId(id string, desc, page, pageSize int, search, sort string) (*[]models.Product, *dto.Pagination, error) {
	return service.productRepo.GetProductByStoreId(id, desc, page, pageSize, search, sort)
}

func (service *productService) GetAllProduct(desc, page, pageSize int, search, sort string) (*[]models.Product, *dto.Pagination, error) {
	return service.productRepo.GetAllProduct(desc, page, pageSize, search, sort)
}

func (service *productService) CheckProductStore(id string, c echo.Context) error {
	UserToken, err := service.tokenRepo.UserToken(middleware.GetToken(c))
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, "Unauthorized")
	}

	store, err := service.storeService.GetStoreByUserId(UserToken.ID)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "Store not found")
	}

	product, err := service.GetProductById(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "Product not found")
	}

	if product.STORE_ID != store.STORE_ID {
		return echo.NewHTTPError(http.StatusForbidden, "Forbidden")
	}

	return nil

}

func (service *productService) UpdateProduct(c echo.Context, id string) error {
	productForm, err := ParseProductForm(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	product, err := service.GetProductById(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "Product not found")
	}

	file, _ := c.FormFile("file")
	if file != nil {
		imagePath, err := service.UploadImage(c)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		service.uploader.DeleteImage(product.PRODUCT_PICTURE)
		product.PRODUCT_PICTURE = imagePath
	}

	if productForm.ProductName != "" {
		product.PRODUCT_NAME = productForm.ProductName
	}
	if productForm.ProductPrice != 0 {
		product.PRODUCT_PRICE = productForm.ProductPrice
	}
	if productForm.ProductDesc != "" {
		product.PRODUCT_DESC = productForm.ProductDesc
	}
	if productForm.ProductLemakTotal != 0 {
		product.PRODUCT_LEMAKTOTAL = productForm.ProductLemakTotal
	}
	if productForm.ProductProtein != 0 {
		product.PRODUCT_PROTEIN = productForm.ProductProtein
	}
	if productForm.ProductKarbohidrat != 0 {
		product.PRODUCT_KARBOHIDRAT = productForm.ProductKarbohidrat
	}
	if productForm.ProductGaram != 0 {
		product.PRODUCT_GARAM = productForm.ProductGaram
	}
	if productForm.ProductGrade != "" {
		product.PRODUCT_GRADING = productForm.ProductGrade
	}
	if productForm.ProductServingSize != 0 {
		product.PRODUCT_SERVINGSIZE = productForm.ProductServingSize
	}
	if productForm.PT_Name != "" {
		pt_id, err := service.ptService.GetProductTypeIdByName(productForm.PT_Name)
		if err != nil {
			return echo.NewHTTPError(http.StatusNotFound, "Product type not found")
		}
		product.PT_ID = pt_id
	}

	product.UpdatedAt = time.Now()

	return service.productRepo.UpdateProduct(product)
}

func (service *productService) DeleteProduct(id string) error {
	product, err := service.GetProductById(id)
	if err != nil {
		return err
	}
	service.uploader.DeleteImage(product.PRODUCT_PICTURE)
	return service.productRepo.DeleteProduct(id)
}

func (service *productService) AdvertiseProduct(c echo.Context, id string) error {
	product, err := service.GetProductById(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "Product not found")
	}

	product.PRODUCT_ISSHOW = true
	product.PRODUCT_EXPSHOW = time.Now().AddDate(0, 1, 0)
	err = service.productRepo.UpdateProduct(product)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal server error")
	}

	return c.JSON(http.StatusOK, product)
}

func (service *productService) UnadvertiseProduct(c echo.Context, id string) error {
	product, err := service.GetProductById(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "Product not found")
	}

	product.PRODUCT_ISSHOW = false
	err = service.productRepo.UpdateProduct(product)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal server error")
	}

	return c.JSON(http.StatusOK, product)
}
