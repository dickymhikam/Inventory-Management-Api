package handler

import (
	"encoding/json"
	"inventory-api/internal/model"
	"inventory-api/internal/service"
	"net/http"
	"strconv"
)

type ProductHandler struct {
  productService service.ProductService
}

func NewProductHandler(productService service.ProductService) *ProductHandler  {
	return &ProductHandler{
		productService: productService,
	}
}


func (handler *ProductHandler) createProduct(
	writer http.ResponseWriter,
	request *http.Request,
) {
	var product model.Products
	
	err := json.NewDecoder(request.Body).Decode(&product)
	if err != nil {
		http.Error(writer, "invalid request body", http.StatusBadRequest)
		return
	}

	err = handler.productService.CreateProduct(
		request.Context(),
		&product,
	)

	if err != nil {
		http.Error(writer, "Failed to create product", http.StatusInternalServerError)
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusCreated)

	json.NewEncoder(writer).Encode(map[string]interface{}{
		"message": "Product created Successfully",
		"data": product,
	})
}

func (handler *ProductHandler) UpdateProduct(
	writer http.ResponseWriter, 
	request *http.Request,
) {
	var product model.Products

	productID, err := strconv.ParseUint(request.PathValue("id"), 10, 64)
	if err != nil {
		http.Error(writer, "Invalid product ID", http.StatusBadRequest)
		return
	}

	err = json.NewDecoder(request.Body).Decode(&product)
	if err != nil {
		http.Error(writer, "Invalid request body", http.StatusBadRequest)
		return
	}

	product.ID = productID
	
	err = handler.productService.UpdateProduct(
		request.Context(),
		&product,
	)

	if err != nil {
		http.Error(writer, "Failed to update product", http.StatusInternalServerError)
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)

	json.NewEncoder(writer).Encode(map[string]interface{}{
		"message": "Update Product Successfully",
		"data": product,
	})
}

func (handler *ProductHandler) DeleteProduct(
	writer http.ResponseWriter,
	request *http.Request,
) {
	var product *model.Products	

	productID, err := strconv.ParseUint(request.PathValue("id"),10, 64,)
	if err !=  nil {
		http.Error(writer, "Invalid Product ID", http.StatusBadRequest)
	}

	err = json.NewDecoder(request.Body).Decode(&product)
	if err != nil {
		http.Error(writer, "Invalid request body", http.StatusBadRequest)
	}

	product.ID = productID

	err = handler.productService.DeleteProduct(
		request.Context(),
		product,
	)
	if err != nil {
		http.Error(writer, "Failed to Delete Product", http.StatusInternalServerError)
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)

	json.NewEncoder(writer).Encode(map[string]interface{}{
		"message": "Delete Product Successfully",
		"data": product,
	})
}
