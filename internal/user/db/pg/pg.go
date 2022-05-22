package pg

import (
	user2 "Scoltest/internal/domain/user"
	"Scoltest/pkg/loger"
	"context"
	"github.com/jackc/pgx/v4"
)

type db struct {
	DBConn *pgx.Conn
	loger  *loger.Logger
	//ctx    context.Context
}

func (b *db) Create(ctx context.Context, user *user2.User) error {
	b.loger.Traceln("begin query create user")
	err := b.DBConn.QueryRow(ctx, "create", user.Name, user.Surname, user.MiddleName, user.Gender, user.Age).Scan(&user.Uuid, &user.Fio)
	if err != nil {
		b.loger.Error("error create query.Data: %#v. Error : %s ", *user, err.Error())
		return err

	} else {
		b.loger.Tracef("Sucsess create user/ Data %#v ", *user)
	}
	return nil
}

func (b *db) Update(ctx context.Context, user *user2.User) error {

	b.loger.Traceln("begin query update user")

	_, err := b.DBConn.Exec(ctx, "update", user.Name, user.Surname, user.MiddleName, user.Gender, user.Age, user.Uuid)

	if err != nil {
		b.loger.Errorf("error create query.Data: %v. Error : %s ", *user, err.Error())
		return err
	} else {
		b.loger.Tracef("Sucsess update user/ Data %#v ", *user)
	}
	return err
}

func (b *db) Delete(ctx context.Context, uuid string) error {
	b.loger.Traceln("begin query delete user")

	_, err := b.DBConn.Exec(ctx, "delete", uuid)

	if err != nil {
		b.loger.Errorf("error delete user query.Data: %v. Error : %s ", uuid, err.Error())
		return err
	} else {
		b.loger.Tracef("Sucsess delete user. Uuid:  %#v ", uuid)
	}
	return nil
}

func New(ctx context.Context, loger *loger.Logger, DBConn *pgx.Conn) (user2.Storage, error) {

	loger.Infoln("Prepare create query begin")
	if _, err := DBConn.Prepare(ctx, "create", "insert into users (name,surname,midlename,gender,age) VALUES ($1,$2,$3,$4,$5) RETURNING uuid,fio"); err != nil {
		return nil, err
	}
	loger.Infoln("Prepare create query sucsess")

	loger.Infoln("Prepare update query begin")
	if _, err := DBConn.Prepare(ctx, "update", "update users set name=$1, surname=$2, midlename=$3, gender=$4,age=$5 where uuid=$6"); err != nil {
		return nil, err
	}

	loger.Infoln("Prepare delete query ")
	if _, err := DBConn.Prepare(ctx, "delete", "delete from users where uuid=$1"); err != nil {
		return nil, err
	}
	loger.Infoln("Prepare delete query sucsess")

	return &db{DBConn: DBConn, loger: loger}, nil
}
