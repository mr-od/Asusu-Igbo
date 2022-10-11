package api

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

type searchProductRequest struct {
	Q string `uri:"q" binding:"required"`
}

func (server *Server) searchProduct(ctx *gin.Context) {
	var req searchProductRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	product, err := server.store.SearchProduct(ctx, req.Q)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}

		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, product)
}

// func (h *HTTP) SearchAddress() Handler {
// 	return func(w http.ResponseWriter, r *http.Request) error {
// 		search := r.FormValue("search")
// 		if search == "" {
// 			return nil
// 		}

// 		addr, err := h.service.GetAddressBySearch(context.Background(), search)
// 		if err != nil {
// 			return service.NewInternalServerError("searching for address").Wrap(err)
// 		}

// 		out, err := json.Marshal(addr)
// 		if err != nil {
// 			return service.NewInternalServerError("marshaling response").Wrap(err)
// 		}

// 		h.handleRequest(w, http.StatusOK, string(out))
// 		return nil
// 	}
// }
