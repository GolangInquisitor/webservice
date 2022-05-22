package pg

import (
	"Scoltest/internal/domain/order"
	"Scoltest/pkg/loger"
	"context"
	"encoding/json"
	"fmt"
	"github.com/jackc/pgx/v4"
	"strconv"
	"strings"
)

type db struct {
	DBConn *pgx.Conn
	loger  *loger.Logger
	//ctx    context.Context
}

func (b *db) Create(ctx context.Context, userUuid string, order *order.Order) error {
	b.loger.Traceln("begin query create order")

	tx, err := b.DBConn.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		b.loger.Errorf("error create transaction for create order. Error : %s ", err.Error())
		return err
	}

	defer func() {
		if err != nil {
			b.loger.Infoln("rollback transaction create order.")
			if errTx := tx.Rollback(context.TODO()); errTx != nil {
				b.loger.Errorf("error rollback transaction . Error : %s ", errTx.Error())
			}
		} else {
			b.loger.Infoln("commit transaction create order.")

			if errTx := tx.Commit(context.TODO()); errTx != nil {
				b.loger.Errorf("error rollback transaction . Error : %s ", errTx.Error())
			}
		}
	}()
	var count int64

	err = tx.QueryRow(ctx, "max_user_order_number", userUuid).Scan(&count)
	if err != nil {
		b.loger.Errorf("error create product query. Data: %#v. Error : %s ", *order, err.Error())
		return err
	}
	count++
	err = b.createOrder(ctx, tx, userUuid, count, order)
	order.Id = strconv.FormatInt(count, 10)

	return err
}
func (b *db) createOrder(ctx context.Context, tx pgx.Tx, userUuid string, odrerNumber int64, order *order.Order) error {
	var products []string
	var err error
	if err = json.Unmarshal([]byte(order.Product), &products); err != nil {
		b.loger.Errorf("error unmarshal product slice. Data: %#v. Error : %s ", order.Product, err.Error())
		return err
	}
	orders := make([]string, len(products))

	for i, productUuid := range products {
		orders[i], err = b.newOrder(ctx, tx, userUuid, odrerNumber, productUuid)
		if err != nil {
			return err
		}

	}

	var a []byte
	if a, err = json.Marshal(&orders); err != nil {
		b.loger.Errorf("error unmarshal product slice. Data: %#v. Error : %s ", order.Product, err.Error())
		return err
	}
	order.Uuid = string(a)
	return nil
}
func (b *db) newOrder(ctx context.Context, tx pgx.Tx, userUuid string, odrerNumber int64, productUuid string) (string, error) {
	var count int64
	var orderUuid string
	err := tx.QueryRow(ctx, "product_quantity", productUuid).Scan(&count)
	if err != nil {
		b.loger.Errorf("error create product query. Data: %#v. Error : %s ", productUuid, err.Error())
		return "", err
	}
	if count <= 0 {
		b.loger.Warningf("the product is out of stock. Data: %#v. Error : %s ", count, err.Error())
	}
	err = tx.QueryRow(ctx, "create_order", userUuid, productUuid, odrerNumber).Scan(&orderUuid)
	if err != nil {
		b.loger.Errorf("error create product query. Error : %s ", err.Error())
		return "", err
	}
	_, err = tx.Exec(ctx, "decrement_quantity", productUuid)
	if err != nil {
		b.loger.Errorf("error decrement_quantity product. Error : %s ", err.Error())
		return "", err
	}
	return orderUuid, nil
}

func (b *db) Update(ctx context.Context, userUuid string, order *order.Order) error {
	b.loger.Traceln("begin query create order")

	tx, err := b.DBConn.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		b.loger.Errorf("error create transaction for create order. Error : %s ", err.Error())
		return err
	}

	defer func() {
		if err != nil {
			b.loger.Infoln("rollback transaction create order.")
			if errTx := tx.Rollback(context.TODO()); errTx != nil {
				b.loger.Errorf("error rollback transaction . Error : %s ", errTx.Error())
			}
		} else {
			b.loger.Infoln("commit transaction create order.")

			if errTx := tx.Commit(context.TODO()); errTx != nil {
				b.loger.Errorf("error rollback transaction . Error : %s ", errTx.Error())
			}
		}
	}()
	if err = b.updateOrder(ctx, tx, userUuid, order); err != nil {
		return err
	}
	err = b.deleteOrderItems(ctx, tx, userUuid, order)

	return err
}
func (b *db) updateOrder(ctx context.Context, tx pgx.Tx, userUuid string, order *order.Order) error {
	var products []string
	var err error
	if err = json.Unmarshal([]byte(order.Product), &products); err != nil {
		b.loger.Errorf("error unmarshal product slice. Data: %#v. Error : %s ", order.Product, err.Error())
		return err
	}

	odrerNumber, err := strconv.ParseInt(order.Id, 10, 64)
	if err != nil {
		b.loger.Errorf("error parse order.Id . Data: %#v. Error : %s ", order.Id, err.Error())
		return err
	}
	var orders []string
	if err = json.Unmarshal([]byte(order.Uuid), &orders); err != nil {
		b.loger.Errorf("error unmarshal uuid slice. Data: %#v. Error : %s ", order.Product, err.Error())
		return err
	}

	if len(products) != len(orders) {
		err := fmt.Errorf("length order slice not equal length uuid")
		b.loger.Errorf("error unmarshal uuid slice. Data: %#v. Error : %s ", order.Product, err.Error())
		return err
	}

	for i, productUuid := range products {
		if orders[i] == "" {
			orders[i], err = b.newOrder(ctx, tx, userUuid, odrerNumber, productUuid)
			if err != nil {
				return err
			}
		}

	}
	var a []byte
	if a, err = json.Marshal(&orders); err != nil {
		b.loger.Errorf("error unmarshal product slice. Data: %#v. Error : %s ", order.Product, err.Error())
		return err
	}
	order.Uuid = string(a)
	return nil
}

func (b *db) deleteOrderItems(ctx context.Context, tx pgx.Tx, userUuid string, order *order.Order) error {

	orderItems := strings.Replace(order.Uuid, "[", "{", 1)
	orderItems = strings.Replace(orderItems, "]", "}", 1)

	_, err := tx.Exec(ctx, "delete_order_items", userUuid, order.Id, orderItems)
	if err != nil {
		b.loger.Errorf("error decrement_quantity product. Error : %s ", err.Error())
		return err
	}
	return nil
}
func (b *db) Delete(ctx context.Context, uuid, orderId string) error {
	odrerNumber, err := strconv.ParseInt(orderId, 10, 64)
	if err != nil {
		b.loger.Errorf("error parse order.Id . Data: %#v. Error : %s ", odrerNumber, err.Error())
		return err
	}
	b.loger.Traceln("begin query delete order")
	_, err = b.DBConn.Exec(ctx, "delete_order", uuid, odrerNumber)
	if err != nil {
		b.loger.Errorf("error delete order query. uuid: %v. orderId: %s. Error : %s ", uuid, orderId, err.Error())
		return err
	} else {
		b.loger.Tracef("Sucsess delete product. Uuid: %#v ", uuid)
	}
	return nil
}

func New(ctx context.Context, loger *loger.Logger, DBConn *pgx.Conn) (order.Storage, error) {

	loger.Infoln("Prepare get product quantity  query begin")
	if _, err := DBConn.Prepare(ctx, "product_quantity", "select left_in_stock from product where  uuid=$1"); err != nil {
		return nil, err
	}
	loger.Infoln("Prepare product quantity  query sucsess")

	loger.Infoln("Prepare get user max order id  query begin")
	if _, err := DBConn.Prepare(ctx, "max_user_order_number", "select COALESCE(max(id),0) from orders where  user_id=$1"); err != nil {
		return nil, err
	}
	loger.Infoln("Prepare product quantity  query sucsess")

	loger.Infoln("Prepare create order  query begin")
	if _, err := DBConn.Prepare(ctx, "create_order",
		"insert into orders (user_id, product,id, price, description) "+
			"select COALESCE($1::uuid),COALESCE($2::uuid),COALESCE($3::bigint),amount,description "+
			"from product left JOIN price on product.uuid=price.product where  product.uuid=$2 returning uuid"); err != nil {
		return nil, err
	}
	loger.Infoln("Prepare create order  query sucsess")

	loger.Infoln("Prepare decrement  product quantity query begin")
	if _, err := DBConn.Prepare(ctx, "decrement_quantity", "update product set left_in_stock = left_in_stock - 1  where  uuid=$1"); err != nil {
		return nil, err
	}
	loger.Infoln("Prepare decrement  product quantity query sucsess")

	loger.Infoln("Prepare delete order items query begin")
	if _, err := DBConn.Prepare(ctx, "delete_order_items", "delete from orders where orders.user_id=$1 and id=$2 and not uuid = ANY($3::uuid[]) "); err != nil {
		return nil, err
	}
	loger.Infoln("Prepare delete order items query sucsess")

	loger.Infoln("Prepare delete order items  product quantity query begin")
	if _, err := DBConn.Prepare(ctx, "delete_order", "delete from orders where orders.user_id=$1 and id=$2"); err != nil {
		return nil, err
	}
	loger.Infoln("Prepare delete order items query sucsess")
	return &db{DBConn: DBConn, loger: loger}, nil
}
