package factories

import (
	"github.com/bxcodec/faker/v3"
	"gohub/app/models/user"
	"gohub/pkg/helpers"
)

func MakeUsers(times int) []user.User {

	var objs []user.User

	// 设置唯一值
	faker.SetGenerateUniqueValues(true)

	for i := 0; i < times; i++ {
		model := user.User{
			Name:     faker.Username(),
			Email:    faker.Email(),
			Phone:    helpers.RandomNumber(11),
			Password: "$2a$14$c5cf765SLcLouABpx/8sbuLPRyHuH4HcdDfHcawXoh0V5yRgIW3Oq",
		}
		objs = append(objs, model)
	}
	return objs
}
