package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/milkygraph/simplebank/db/sqlc"
)

type transferRequest struct {
	FromAccountID int64  `json:"from_account_id" binding:"required"`
	ToAccountID   int64  `json:"to_account_id" binding:"required"`
	Amount        int64  `json:"amount" binding:"required"`
	Currency      string `json:"currency" binding:"required,oneof=USD EUR CAD"`
}

func (server *Server) createTransfer(ctx *gin.Context) {
	var req transferRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if !server.validCurrency(ctx, req.FromAccountID, req.ToAccountID) {
		return
	}

	arg := db.TransferTxParams{
		FromAccountID: req.FromAccountID,
		ToAccountID:   req.ToAccountID,
		Amount:        req.Amount,
	}

	transfer, err := server.store.TransferTx(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, transfer)
}

func (server *Server) validCurrency(ctx *gin.Context, fromAccountID int64, toAccountID int64) bool {
	fromAccount, err := server.store.GetAccount(ctx, fromAccountID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return false
	}

	toAccount, err := server.store.GetAccount(ctx, toAccountID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return false
	}

	if fromAccount.Currency != toAccount.Currency {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "currency mismatch"})
		return false
	}

	return true
}
