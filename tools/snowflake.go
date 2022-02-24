package tools

import (
	"fmt"
	"os"
	"strconv"

	"github.com/bwmarrin/snowflake"
	_ "github.com/joho/godotenv/autoload"
)

func GenerateSnowflakeId() int64 {
	snowflake.Epoch = 1604971882247
	initNode, _ := strconv.ParseInt(os.Getenv("NODE"), 10, 64)
	node, err := snowflake.NewNode(initNode)
	if err != nil {
		fmt.Println(err)
		return 0
	}
	id := node.Generate().Int64()
	return id
}
