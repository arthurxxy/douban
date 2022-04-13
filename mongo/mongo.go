package mongo

import (
	"context"
	"strconv"

	"github.com/qiniu/qmgo"
	"go.mongodb.org/mongo-driver/bson"
)

type QueryItem struct {
	ID string `bson:"ID"`
}

func ConnectMongo(s_uri, s_db, s_coll string) (*qmgo.Collection, error) {
	ctx := context.Background()
	client, err := qmgo.NewClient(ctx, &qmgo.Config{Uri: s_uri})
	db := client.Database(s_db)
	coll := db.Collection(s_coll)

	return coll, err
}

func GetMaxID(s_uri, s_db, s_coll string) int {
	ctx := context.Background()
	client, _ := qmgo.NewClient(ctx, &qmgo.Config{Uri: s_uri})
	db := client.Database(s_db)
	coll := db.Collection(s_coll)

	one := QueryItem{}
	coll.Find(ctx, bson.M{}).Sort("-ID").One(&one)

	intvar, _ := strconv.Atoi(one.ID)

	return intvar
}
