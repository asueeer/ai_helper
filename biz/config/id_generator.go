package config

import "github.com/bwmarrin/snowflake"

var (
	WorkID      = int64(1)
	iDGenerator *snowflake.Node
)

func init() {
	var err error
	iDGenerator, err = snowflake.NewNode(WorkID)
	if err != nil {
		panic(err)
	}
}

func GenerateIDInt64() int64 {
	return iDGenerator.Generate().Int64()
}

func GenerateIDString() string {
	return iDGenerator.Generate().String()
}
