package snowflake

import (
	"github.com/bwmarrin/snowflake"
	logger "github.com/sirupsen/logrus"
)

var node *snowflake.Node

func StartSnowflakeNode() {
	var err error
	node, err = snowflake.NewNode(69)
	if err != nil {
		logger.Fatalf("Error starting snowflake node\n%s\n", err)
	}
}

func GenerateUint() uint {
	return uint(node.Generate().Int64())
}
