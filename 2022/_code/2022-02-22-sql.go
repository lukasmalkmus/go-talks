//nolint
package code

import (
	"context"
	"fmt"
)

func sqlDrawbacks() {
	// START SQL_DRAWBACKS_1 OMIT
	var name string
	err = db.QueryRow(ctx, "SELECT name FROM users WHERE id = $1", 42).Scan(&name)
	// END SQL_DRAWBACKS_1 OMIT

	// START SQL_DRAWBACKS_2 OMIT
	_ = db.QueryRow(ctx, "SELECT "+scanUserFields+" FROM users WHERE id = $1", 42)
	// END SQL_DRAWBACKS_2 OMIT

	// START SQL_DRAWBACKS_3 OMIT
	userQuery := sq.Select("name").From("users").Where(sq.Eq{"id": 42})
	qs := userQuery.ToSql()
	// END SQL_DRAWBACKS_3 OMIT

	return
}

func ormsForTheRescue() {
	// START ORMS_FOR_THE_RESCUE OMIT
	dataset := new(Dataset)
	err = db.Model(dataset).
		Relation("Owner").
		Where("dataset.id = ?", 42).
		Select()
	if err != nil {
		return err
	}
	// END ORMS_FOR_THE_RESCUE OMIT
}

// START ENT OMIT
func CreateUser(ctx context.Context, client *ent.Client) (*ent.User, error) {
	return client.User.
		Create().
		SetName("lukasmalkmus").
		Save(ctx)
}

// END ENT OMIT

func sqlc() {
	// START SQLC OMIT
	queries := dbsqlc.New(db)
	user, err := queries.CreateUser(ctx, dbsqlc.CreateUser{
		Name: "Lukas Malkmus",
	})
	fmt.Println(user.Name) // "Lukas Malkmus"
	// END SQLC OMIT
}
