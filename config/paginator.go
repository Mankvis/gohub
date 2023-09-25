package config

import "gohub/pkg/config"

func init() {
	config.Add("paging", func() map[string]interface{} {
		return map[string]interface{}{

			// 默认每页条数
			"perpage": 20,

			// URL 中用以分别多少页的参数
			// 此值若修改需一并修改该请求验证规则
			"url_query_page": "page",

			// URL 中用以分辨排序的参数（使用 id 或者其他）
			// 此值若修改需一并修改该请求验证规则
			"url_query_sort": "sort",

			// URL 中用以分别排序规则的参数（辨别是正序还是倒叙）
			// 此值若修改需一并修改该请求验证规则
			"url_query_order": "order",

			// URL 中用以分别每页条数的参数
			// 此值若修改需一并修改该请求验证规则
			"url_query_per_page": "per_page",
		}
	})
}
