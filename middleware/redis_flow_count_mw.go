package middleware

import (
	"computer_network/public"
	"fmt"
)

func RedisFlowCountMiddleWare(counter *public.RedisFlowCountService) func(c *SliceRouterContext) {
	return func(c *SliceRouterContext) {
		counter.Increase()
		fmt.Println("QPS:", counter.QPS)
		fmt.Println("TotalCount:", counter.TotalCount)
		c.Next()
	}
}
