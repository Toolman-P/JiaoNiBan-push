package push

import "JiaoNiBan-push/services/tpns"

func WrapTagRules(tags []string) []tpns.TagRule {
	tagItem := tpns.TagItem{
		Tags: tags,
	}
	var tagRule = tpns.TagRule{
		TagItems: []tpns.TagItem{tagItem},
		Operator: tpns.TagOperationOr,
	}
	return []tpns.TagRule{tagRule}
}
