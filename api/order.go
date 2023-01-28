package api

import (
	"crypto/rand"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"suning/model"
	"suning/service"
	"suning/util"
	"time"
)

// GenOrder 实现生成订单接口（用户选择需要购买的物品）
func GenOrder(c *gin.Context) {
	userId := c.Param("user_id")
	if userId == "" {
		util.RespParamErr(c)
		return
	}
	//获取一系列商品
	payProducts := c.QueryArray("pay_products")
	if payProducts == nil {
		util.RespParamErr(c)
		return
	}
	//获取购物车信息
	carts, err := service.SearchCartInOrder(userId, payProducts)
	if err != nil {
		fmt.Printf("search cart err:%v", err)
		util.RespInternalErr(c)
		return
	}
	//删除购物车相应信息，获得总价
	totalPrice, err := service.GenOrder(userId, payProducts)
	if err != nil {
		fmt.Printf("generate order err:%v", err)
		util.RespInternalErr(c)
		return
	}
	//获取地址
	addressId := c.PostForm("address_id")
	address, err := service.SearchAddressById(addressId)
	if err != nil {
		fmt.Printf("search address err:%v", err)
		util.RespInternalErr(c)
		return
	}
	//生成订单编号
	timestamp := time.Now().Unix()
	n := 10
	b := make([]byte, n)
	_, err = rand.Read(b)
	if err != nil {
		util.RespInternalErr(c)
		return
	}
	orderNumber := fmt.Sprintf("%d%x", timestamp, b)
	//写入order信息
	order := model.Order{
		OrderNumber:      orderNumber,
		OrderTime:        time.Now(),
		Status:           "待支付",
		PaymentMethod:    "",
		PaymentAmount:    totalPrice,
		PaymentTime:      "",
		RecipientName:    address.RecipientName,
		RecipientAddress: address.Province + address.City + address.StateOrCommunity,
		RecipientPhone:   address.RecipientPhone,
		UserId:           userId,
	}
	err = service.InsertOrder(order)
	if err != nil {
		fmt.Printf("insert order err:%v", err)
		util.RespInternalErr(c)
		return
	}
	//查询此order_id
	orderId, err := service.SearchOrderIdByNumber(orderNumber)
	if err != nil {
		fmt.Printf("search order_id err:%v", err)
		util.RespInternalErr(c)
		return
	}
	//添加订单明细
	for _, cart := range carts {
		err = service.InsertOrderDetail(orderId, cart)
		if err != nil {
			fmt.Printf("insert order detail err:%v", err)
			util.RespInternalErr(c)
			return
		}
	}
}

// SettleBill 实现商品结算接口（付款）
func SettleBill(c *gin.Context) {
	userId := c.Param("user_id")
	if userId == "" {
		util.RespParamErr(c)
		return
	}
	orderId := c.Query("order_id")
	if userId == "" {
		util.RespParamErr(c)
		return
	}
	//支付
	totalPrice, err := service.SearchTotalPriceById(orderId)
	if err != nil {
		fmt.Printf("search total price err:%v", err)
		util.RespInternalErr(c)
		return
	}
	payment := c.Query("payment") //支付金额
	if payment != totalPrice {
		util.NormErr(c, 400, "wrong payment")
		return
	}
	price, err := strconv.ParseFloat(totalPrice, 64)
	if err != nil {
		util.NormErr(c, 400, "wrong total price")
		return
	}
	//扣除余额
	err = service.DecreaseBalance(userId, price)
	if err != nil {
		fmt.Printf("decrease balance err:%v", err)
		util.RespInternalErr(c)
		return
	}
	//修改订单信息
	paymentMethod := c.Query("payment_method") //支付方式
	err = service.UpdateOrderStatus(orderId, "待收货")
	if err != nil {
		fmt.Printf("update status err:%v", err)
		util.RespInternalErr(c)
		return
	}
	err = service.UpdatePaymentMethodTime(orderId, paymentMethod, time.Now())
	if err != nil {
		fmt.Printf("update payment method and time err:%v", err)
		util.RespInternalErr(c)
		return
	}
	util.RespOK(c, "settle bill success")
}

// SearchOrder 实现按状态查看订单接口
func SearchOrder(c *gin.Context) {
	//获取信息
	userId := c.Param("user_id")
	if userId == "" {
		util.RespParamErr(c)
		return
	}
	status := c.Query("status")
	if status == "" {
		util.RespParamErr(c)
		return
	}
	//查找订单
	orders, err := service.SearchOrder(userId, status)
	if err != nil {
		fmt.Printf("search order err:%v", err)
		util.RespInternalErr(c)
		return
	}
	//查找订单明细
	orderDetailsS, err := service.SearchOrderDetailByOrderId(orders)
	if err != nil {
		fmt.Printf("search order detail err:%v", err)
		util.RespInternalErr(c)
		return
	}
	c.JSON(http.StatusOK, model.RespOrder{
		Status:      200,
		Info:        "search order success",
		Order:       orders,
		OrderDetail: orderDetailsS,
	})
}

// UpdateOrderStatus 实现改变订单状态接口
func UpdateOrderStatus(c *gin.Context) {
	//获取信息
	userId := c.Param("user_id")
	if userId == "" {
		util.RespParamErr(c)
		return
	}
	orderId := c.Query("order_id") //
	status := c.Query("status")    //
	if status == "" || orderId == "" {
		util.RespParamErr(c)
		return
	}
	//修改状态
	err := service.UpdateOrderStatus(orderId, status)
	if err != nil {
		fmt.Printf("udpate order status err:%v", err)
		util.RespInternalErr(c)
		return
	}
	util.RespOK(c, "update order status success")
}

// ViewOrder 实现查看所有订单接口
func ViewOrder(c *gin.Context) {
	userId := c.Param("user_id")
	if userId == "" {
		util.RespParamErr(c)
		return
	}
	//获取orders
	orders, err := service.ViewOrder(userId)
	if err != nil {
		fmt.Printf("view order err:%v", err)
		util.RespInternalErr(c)
		return
	}
	//获取order detail
	orderDetailsS, err := service.SearchOrderDetailByOrderId(orders)
	if err != nil {
		fmt.Printf("search order detail err:%v", err)
		util.RespInternalErr(c)
		return
	}
	c.JSON(http.StatusOK, model.RespOrder{
		Status:      200,
		Info:        "view order success",
		Order:       orders,
		OrderDetail: orderDetailsS,
	})
}
