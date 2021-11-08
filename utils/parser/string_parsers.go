package parser

import (
	"github.com/Gunnsteinn/cryptoGuild/utils/errors"
	"go.mongodb.org/mongo-driver/bson"
)

func StringToBson(aux string) (interface{}, *errors.RestErr) {
	var doc interface{}
	err := bson.UnmarshalExtJSON([]byte(aux), true, &doc)
	if err != nil {
		return nil, errors.NewInternalServerError(err.Error())
	}
	return doc, nil
}
