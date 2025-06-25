package container

import (
	"api_techstore/internal/database"
	"api_techstore/internal/services"
	"api_techstore/pkg/jwt"
	"api_techstore/pkg/logger"
	"log"

	"github.com/go-redis/redis/v8"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type Container struct {
	DB        *gorm.DB
	Redis     *redis.Client
	JWTConfig *jwt.JWTConfig
	Logger    *logrus.Logger

	// Khai báo các service để sử dụng DI
	CategoryService services.CategoryService
	BrandService    services.BrandService
	ProductService  services.ProductService
	OrderService    services.OrderService
	AddressService  services.AddressService
	UserService     services.UserService
	PaymentService  services.PaymentService
	CartService     services.CartService
	CartItemService services.CartItemService
}

func NewContainer() *Container {
	dbConn, err := database.InitDB()
	if err != nil {
		log.Fatalf("failed to init database: %v", err)
	}

	redisClient, err := database.InitRedis()
	if err != nil {
		log.Fatalf("failed to init redis: %v", err)
	}

	jwtCfg := jwt.NewJWTConfig()

	logger.InitLogger()

	categoryService := services.NewCategoryService(dbConn.DB)
	brandService := services.NewBrandService(dbConn.DB)
	productService := services.NewProductService(dbConn.DB)
	orderService := services.NewOrderService(dbConn.DB)
	addressService := services.NewAddressService(dbConn.DB)
	userService := services.NewUserService(dbConn.DB)
	paymentService := services.NewPaymentService(dbConn.DB)
	cartService := services.NewCartService(dbConn.DB)
	cartItemService := services.NewCartItemService(dbConn.DB)

	return &Container{
		DB:        dbConn.DB,
		Redis:     redisClient,
		JWTConfig: jwtCfg,
		Logger:    logger.Log,

		CategoryService: categoryService,
		BrandService:    brandService,
		ProductService:  productService,
		OrderService:    orderService,
		AddressService:  addressService,
		UserService:     userService,
		PaymentService:  paymentService,
		CartService:     cartService,
		CartItemService: cartItemService,
	}
}
