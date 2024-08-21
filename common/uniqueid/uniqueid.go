package uniqueid

import (
	"sync"

	"github.com/sony/sonyflake"
	"github.com/zeromicro/go-zero/core/logx"
)

var (
	userFlake         *sonyflake.Sonyflake
	userRelationFlake *sonyflake.Sonyflake
	once              sync.Once
)

func init() {
	settings := sonyflake.Settings{}
	userFlake = sonyflake.NewSonyflake(settings)
	if userFlake == nil {
		logx.Error("Sonyflake for user not created")
	}

	userRelationFlake = sonyflake.NewSonyflake(settings)
	if userRelationFlake == nil {
		logx.Error("Sonyflake for user_relation not created")
	}
}

func GenUserID() int64 {
	id, err := userFlake.NextID()
	if err != nil {
		logx.Severef("flake NextID failed with %s \n", err)
		panic(err)
	}

	return int64(id)
}

func GenUserRelationID() int64 {
	id, err := userRelationFlake.NextID()
	if err != nil {
		logx.Severef("flake NextID failed with %s \n", err)
		panic(err)
	}

	return int64(id)
}
