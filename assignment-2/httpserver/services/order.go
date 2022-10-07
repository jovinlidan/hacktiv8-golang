package services

import (
	"assignment2/httpserver/controllers/params"
	"assignment2/httpserver/controllers/views"
	"assignment2/httpserver/repositories"
	"assignment2/httpserver/repositories/models"
	"database/sql"
	"strconv"
	"strings"
)



type OrderSvc struct {
	repo repositories.OrderRepo
}

func NewOrderSvc(repo repositories.OrderRepo) *OrderSvc {
	return &OrderSvc{
		repo: repo,
	}
}


func (s *OrderSvc) GetAllOrders() *views.Response {
	orders, err := s.repo.GetOrders()
	if err != nil {
		if err == sql.ErrNoRows {
			return views.DataNotFound(err)
		}
		return views.InternalServerError(err)
	}

	return views.SuccessFindAllResponse(parseModelToOrderGetAll(orders), "GET_ALL_ORDERS")
}

func (s *OrderSvc) CreateOrder(req *params.OrderCreateRequest) *views.Response {
	order := parseCreateRequestToModel(req)
	
	err := s.repo.CreateOrder(order)
	if err != nil {
		if strings.Contains(err.Error(), "duplicate") {
			return views.DataConflict(err)
		}
		return views.InternalServerError(err)
	}
	return views.SuccessCreateResponse(order, "CREATE_ORDER")
}

/// Create
func parseCreateRequestToItem(req *[]params.OrderItemCreateRequest) *[]models.Item{
	var o []models.Item
	for _, st := range *req{
		o = append(o, models.Item{
			ItemCode: st.ItemCode,
			Description: st.Description,
			Quantity: st.Quantity,
		} )
	}
	return &o;
}


func parseCreateRequestToModel(req *params.OrderCreateRequest) *models.Order {
	return &models.Order{
		CustomerName: req.CustomerName,
		OrderedAt: req.OrderedAt,
		Items: *parseCreateRequestToItem(&req.Items),
	}
}
///

/// Update
func parseUpdateRequestToItem(req *[]params.OrderItemUpdateRequest) *[]models.Item{
	var o []models.Item
	for _, st := range *req{
		o = append(o, models.Item{
			ItemCode: st.ItemCode,
			Description: st.Description,
			Quantity: st.Quantity,
			LineItemID: st.LineItemID,
		} )
	}
	return &o;
}


func parseUpdateRequestToModel(req *params.OrderUpdateRequest) *models.Order {
	return &models.Order{
		OrderID: req.OrderID,
		CustomerName: req.CustomerName,
		OrderedAt: req.OrderedAt,
		Items: *parseUpdateRequestToItem(&req.Items),
	}
}
///

func parseModelToOrderItemGetAll (mod *[]models.Item) *[]views.OrderItemGetAll {
var s []views.OrderItemGetAll
	for _, st := range *mod {
		s = append(s, views.OrderItemGetAll{
			LineItemID: st.LineItemID,
			ItemCode: st.ItemCode,
			Description: st.Description,
			Quantity: st.Quantity,
			OrderID: st.OrderID,
		})
	}
	return &s
}

func parseModelToOrderGetAll(mod *[]models.Order) *[]views.OrderGetAll {
	var s []views.OrderGetAll
	for _, st := range *mod {
		s = append(s, views.OrderGetAll{
			CustomerName:  st.CustomerName,
			OrderedAt: st.OrderedAt,
			Items:  *parseModelToOrderItemGetAll(&st.Items),
			OrderID: st.OrderID,
			
		})
	}
	return &s
}


func (s *OrderSvc) UpdateOrder(req *params.OrderUpdateRequest) *views.Response {
	order := parseUpdateRequestToModel(req)

	dbOrder, err := s.repo.GetOrder(strconv.FormatUint(uint64(order.OrderID), 10))
	if err != nil {
		if err == sql.ErrNoRows {
			return views.DataNotFound(err)
		}
		return views.InternalServerError(err)
	}
	dbOrder.CustomerName = order.CustomerName
	dbOrder.Items = order.Items
	dbOrder.OrderedAt = order.OrderedAt

	order, err = s.repo.UpdateOrder(dbOrder)
	if err != nil {
		if err == sql.ErrNoRows {
			return views.DataNotFound(err)
		}
		return views.InternalServerError(err)
	}

	return views.SuccessUpdateResponse(order,"UPDATE_ORDER")
}

func (s *OrderSvc) DeleteOrder(id *string) *views.Response {

	dbOrder, err := s.repo.GetOrder(*id)
	if err != nil {
		if err == sql.ErrNoRows {
			return views.DataNotFound(err)
		}
		return views.InternalServerError(err)
	}

	 order, err := s.repo.DeleteOrder(dbOrder)
	if err != nil {
		if err == sql.ErrNoRows {
			return views.DataNotFound(err)
		}
		return views.InternalServerError(err)
	}

	return views.SuccessDeleteResponse(order, "DELETE_ORDER")
}