package main

import (
	"common/api"
	"errors"
	"fmt"
	"gateway/cmd/utils"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func validateItems(items []*api.Item) error {
	if len(items) == 0 {
		return errors.New("you must have at least one item")
	}

	for _, item := range(items) {
		if item.ID == "" {
			return errors.New("invalid item id")
		}

		if item.Quantity < 1 {
			return errors.New("item quantity must be greater than 0")
		}
	}

	return nil
}

func (app *Http) createOrderHandler(c *fiber.Ctx) error {
	// Get customer Id from url params
	customerId := c.Params("customerId")

	if customerId == "" {
		utils.RespondWithError(c, http.StatusBadRequest, "Invalid Parameters!")

		return fmt.Errorf("customer ID is required")

	}

	// Get and validate items from req body
	var items []*api.Item

	err := c.BodyParser(&items)

	if err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, "Error parsing JSON")

		return err
	}

	validateErr := validateItems(items)

	if validateErr != nil {
		utils.RespondWithError(c, http.StatusBadRequest, validateErr.Error())

		return validateErr
	}


	order, taskErr := app.client.CreateOrder(c.Context(), &api.CreateOrderRequest{
		CustomerID: customerId,
		Items: items,
	})

	if taskErr != nil {
		rpcStatus := status.Convert(taskErr)

		if rpcStatus.Code() == codes.InvalidArgument {
			utils.RespondWithError(c, http.StatusBadRequest, rpcStatus.Message())

		    return taskErr
		}

		utils.RespondWithError(c, http.StatusInternalServerError, taskErr.Error())

		return taskErr
	}

	utils.RespondWithJSON(c, http.StatusCreated, order)

	return nil
}