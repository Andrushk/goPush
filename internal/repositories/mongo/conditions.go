package mongo

import (
	"github.com/Andrushk/goPush/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func primitiveId(Id entity.ID) primitive.ObjectID {
	pid, err := primitive.ObjectIDFromHex(Id.String())
	if err != nil {
		panic("wrong primitive id")
	}
	return pid
}

func CEntityId(id entity.ID) bson.D {
	return CId(primitiveId(id))
}

func CId(id primitive.ObjectID) bson.D {
	return bson.D{
		{"_id", id},
	}
}
