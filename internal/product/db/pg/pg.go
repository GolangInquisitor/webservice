package pg

import (
	"Scoltest/internal/domain/product"
	"Scoltest/pkg/loger"
	"context"
	"github.com/jackc/pgx/v4"
)

type db struct {
	DBConn *pgx.Conn
	loger  *loger.Logger
	//ctx    context.Context
}

func (b *db) Create(ctx context.Context, product *product.Product) error {
	b.loger.Traceln("begin query create product")

	tx, err := b.DBConn.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		b.loger.Errorf("error create transaction for create product. Error : %s ", err.Error())
		return err
	}

	defer func() {
		if err != nil {
			b.loger.Infoln("rollback transaction create product.")
			if errTx := tx.Rollback(context.TODO()); errTx != nil {
				b.loger.Errorf("error rollback transaction . Error : %s ", errTx.Error())
			}
		} else {
			b.loger.Infoln("commit transaction create product.")

			if errTx := tx.Commit(context.TODO()); errTx != nil {
				b.loger.Errorf("error rollback transaction . Error : %s ", errTx.Error())
			}
		}
	}()

	err = tx.QueryRow(ctx, "create_product", product.Description, product.LeftInStock).Scan(&product.Uuid)
	if err != nil {
		b.loger.Errorf("error create product query. Data: %#v. Error : %s ", *product, err.Error())
		return err
	}

	var id int64
	err = tx.QueryRow(ctx, "create_price", product.Currency, product.Price, product.Uuid).Scan(&id)
	if err != nil {
		b.loger.Errorf("error create price query. Data: %#v. Error : %s ", *product, err.Error())
	}

	return err
}
func (b *db) Update(ctx context.Context, product *product.Product) error {

	tx, err := b.DBConn.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		b.loger.Errorf("error create transaction for update product. Error : %s ", err.Error())
		return err
	}

	defer func() {
		if err != nil {
			b.loger.Infoln("rollback transaction update product.")
			if errTx := tx.Rollback(context.TODO()); errTx != nil {
				b.loger.Errorf("error rollback transaction . Error : %s ", errTx.Error())
			}
		} else {
			b.loger.Infoln("commit transaction update product.")
			if errTx := tx.Commit(context.TODO()); errTx != nil {
				b.loger.Errorf("error rollback transaction . Error : %s ", errTx.Error())
			}
		}
	}()

	_, err = tx.Exec(ctx, "update_product", product.Description, product.LeftInStock, product.Uuid)
	if err != nil {
		b.loger.Errorf("error update product query. Data: %#v. Error : %s ", *product, err.Error())
		return err
	}

	_, err = tx.Exec(ctx, "update_price", product.Price, product.Uuid, product.Currency)
	if err != nil {
		b.loger.Errorf("error update price query. Data: %#v. Error : %s ", *product, err.Error())
		return err
	}

	return err
}

func (b *db) Delete(ctx context.Context, uuid string) error {
	b.loger.Traceln("begin query delete product")

	_, err := b.DBConn.Exec(ctx, "delete_product", uuid)

	if err != nil {
		b.loger.Errorf("error delete product query.Data: %v. Error : %s ", uuid, err.Error())
		return err
	} else {
		b.loger.Tracef("Sucsess delete product. Uuid: %#v ", uuid)
	}
	return nil
}

func New(ctx context.Context, loger *loger.Logger, DBConn *pgx.Conn) (product.Storage, error) {

	loger.Infoln("Prepare create product query begin")
	if _, err := DBConn.Prepare(ctx, "create_product", "insert into product (description, left_in_stock) values ($1,$2) returning uuid"); err != nil {
		return nil, err
	}
	loger.Infoln("Prepare create product query sucsess")

	loger.Infoln("Prepare create price query begin")
	if _, err := DBConn.Prepare(ctx, "create_price", "insert into price  (currency, amount, product) values ($1,$2,$3) returning price.id"); err != nil {
		return nil, err
	}
	loger.Infoln("Prepare create price query sucsess")

	loger.Infoln("Prepare update product query begin")
	if _, err := DBConn.Prepare(ctx, "update_product", "update  product set description=$1, left_in_stock=$2 where uuid=$3"); err != nil {
		return nil, err
	}
	loger.Infoln("Prepare update product query sucsess")

	loger.Infoln("Prepare update price query begin")
	if _, err := DBConn.Prepare(ctx, "update_price", "update price set amount=$1 where product=$2 and currency=$3"); err != nil {
		return nil, err
	}
	loger.Infoln("Prepare update price query sucsess")

	loger.Infoln("Prepare delete product query begin")
	if _, err := DBConn.Prepare(ctx, "delete_product", "delete from product where uuid=$1"); err != nil {
		return nil, err
	}
	loger.Infoln("Prepare delete product query sucsess")

	return &db{DBConn: DBConn, loger: loger}, nil
}
