package uniqueid

import (
	"github.com/sony/sonyflake"
	"github.com/zeromicro/go-zero/core/logx"
)

var (
	userFlake         *sonyflake.Sonyflake
	userRelationFlake *sonyflake.Sonyflake
	feedsFlake        *sonyflake.Sonyflake
	commentsFlake     *sonyflake.Sonyflake
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
	feedsFlake = sonyflake.NewSonyflake(settings)
	if feedsFlake == nil {
		logx.Error("Sonyflake for feeds not created")
	}
	commentsFlake = sonyflake.NewSonyflake(settings)
	if commentsFlake == nil {
		logx.Error("Sonyflake for comments not created")
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

func GenFeedID() int64 {
	id, err := feedsFlake.NextID()
	if err != nil {
		logx.Severef("flake NextID failed with %s \n", err)
		panic(err)
	}

	return int64(id)
}

func GenCommentID() int64 {
	id, err := commentsFlake.NextID()
	if err != nil {
		logx.Severef("flake NextID failed with %s \n", err)
		panic(err)
	}

	return int64(id)
}
