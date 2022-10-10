// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0

package db

import (
	"context"

	"github.com/google/uuid"
)

type Querier interface {
	AddAccountBalance(ctx context.Context, arg AddAccountBalanceParams) (Account, error)
	CreateAccount(ctx context.Context, arg CreateAccountParams) (Account, error)
	CreateEntry(ctx context.Context, arg CreateEntryParams) (Entry, error)
	CreateOrder(ctx context.Context, arg CreateOrderParams) (Order, error)
	CreateOrderItem(ctx context.Context, arg CreateOrderItemParams) (OrderItem, error)
	CreateProduct(ctx context.Context, arg CreateProductParams) (Product, error)
	CreateSession(ctx context.Context, arg CreateSessionParams) (Session, error)
	CreateTransfer(ctx context.Context, arg CreateTransferParams) (Transfer, error)
	CreateUser(ctx context.Context, arg CreateUserParams) (User, error)
	DeleteAccount(ctx context.Context, id int64) error
	DeleteOrder(ctx context.Context, id int64) error
	DeleteOrderItem(ctx context.Context, id int64) error
	DeleteProduct(ctx context.Context, id int64) error
	GetAccount(ctx context.Context, id int64) (Account, error)
	GetAccountForUpdate(ctx context.Context, id int64) (Account, error)
	GetEntry(ctx context.Context, id int64) (Entry, error)
	GetOrder(ctx context.Context, id int64) (Order, error)
	GetOrderForUpdate(ctx context.Context, id int64) (Order, error)
	GetOrderItem(ctx context.Context, id int64) (OrderItem, error)
	GetOrderItemForUpdate(ctx context.Context, id int64) (OrderItem, error)
	GetProduct(ctx context.Context, id int64) (Product, error)
	GetProductForUpdate(ctx context.Context, id int64) (Product, error)
	GetSession(ctx context.Context, id uuid.UUID) (Session, error)
	GetTransfer(ctx context.Context, id int64) (Transfer, error)
	GetUser(ctx context.Context, username string) (User, error)
	ListAccounts(ctx context.Context, arg ListAccountsParams) ([]Account, error)
	ListEntries(ctx context.Context, arg ListEntriesParams) ([]Entry, error)
	ListOrderItems(ctx context.Context, arg ListOrderItemsParams) ([]OrderItem, error)
	ListOrders(ctx context.Context, arg ListOrdersParams) ([]Order, error)
	ListProducts(ctx context.Context) ([]Product, error)
	ListTransfers(ctx context.Context, arg ListTransfersParams) ([]Transfer, error)
	UpdateAccount(ctx context.Context, arg UpdateAccountParams) (Account, error)
	UpdateOrder(ctx context.Context, arg UpdateOrderParams) (Order, error)
	UpdateOrderItem(ctx context.Context, arg UpdateOrderItemParams) (OrderItem, error)
	UpdateProduct(ctx context.Context, arg UpdateProductParams) (Product, error)
}

var _ Querier = (*Queries)(nil)
