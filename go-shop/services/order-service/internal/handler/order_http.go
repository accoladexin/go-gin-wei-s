package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	orderpb "go-shop/order-service/proto"
	"google.golang.org/grpc"
)

// OrderHTTPHandler HTTP 层只做：参数解析 + 转 gRPC
type OrderHTTPHandler struct {
	orderClient orderpb.OrderServiceClient
}

// NewOrderHTTPHandler 创建 HTTP Handler
func NewOrderHTTPHandler(grpcAddr string) *OrderHTTPHandler {
	conn, err := grpc.Dial(
		grpcAddr,
		grpc.WithInsecure(),
	)
	if err != nil {
		panic(err)
	}

	return &OrderHTTPHandler{
		orderClient: orderpb.NewOrderServiceClient(conn),
	}
}

// CreateOrder HTTP 接口
// POST /orders?user_id=1&product_id=1
func (h *OrderHTTPHandler) CreateOrder(c *gin.Context) {
	userID, _ := strconv.ParseInt(c.Query("user_id"), 10, 64)
	productID, _ := strconv.ParseInt(c.Query("product_id"), 10, 64)

	resp, err := h.orderClient.CreateOrder(
		c.Request.Context(),
		&orderpb.CreateOrderRequest{
			UserId:    userID,
			ProductId: productID,
		},
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"order_id":     resp.OrderId,
		"user_name":    resp.UserName,
		"product_name": resp.ProductName,
		"price":        resp.Price,
	})
}
