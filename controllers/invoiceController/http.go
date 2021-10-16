package invoiceController

import (
	"TiBO_API/app/middleware/auth"
	"TiBO_API/businesses/invoiceEntity"
	"TiBO_API/controllers/invoiceController/request"
	"TiBO_API/controllers/invoiceController/response"
	"TiBO_API/helpers"
	"net/http"

	"github.com/labstack/echo/v4"
)

type InvoiceController struct {
	invoiceService invoiceEntity.Service
}

func NewInvoiceController(service invoiceEntity.Service) *InvoiceController {
	return &InvoiceController{
		invoiceService: service,
	}
}

func (ctrl *InvoiceController) Create(c echo.Context) error {
	req := request.Invoices{}
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest,
			helpers.BuildErrorResponse("An error occurred while inputing data",
				err, helpers.EmptyObj{}))
	}

	userID := auth.GetUser(c)

	data, err := ctrl.invoiceService.Create(c.Request().Context(), userID.Uuid, req.ToDomain())
	if err != nil {
		return c.JSON(http.StatusInternalServerError,
			helpers.BuildErrorResponse("Something Gone Wrong,Please Contact Administrator",
				err, helpers.EmptyObj{}))
	}
	return c.JSON(http.StatusCreated,
		helpers.BuildSuccessResponse("Successfully Created invoice ",
			data))
}

func (ctrl *InvoiceController) GetByUserID(c echo.Context) error {
	userID := auth.GetUser(c).Uuid
	data, err := ctrl.invoiceService.GetByUserID(c.Request().Context(), userID)
	if err != nil {
		return c.JSON(http.StatusNotFound,
			helpers.BuildErrorResponse("No one Cinema found within your area",
				err, helpers.EmptyObj{}))
	}
	res := response.FromDomainArray(data)
	return c.JSON(http.StatusOK,
		helpers.BuildSuccessResponse("Successfully Get invoice By userID",
			res))
}
