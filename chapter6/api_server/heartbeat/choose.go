package heartbeat

import "math/rand"

// 随机选择多个数据服务
// @n:要选择n个可用的数据服务
// @exclude:需要排除的项
func ChooseRandomDataServers(n int, exclude map[int]string) (ds []string) {
	// 候选
	candidates := make([]string, 0)
	// 将IP作为键方便匹配
	reverseExcludeMap := make(map[string]int)
	for id, addr := range exclude {
		reverseExcludeMap[addr] = id
	}
	// 当前可用的数据服务列表
	servers := GetDataServers()
	for i := range servers {
		// 可用的数据服务IP
		s := servers[i]
		// 该IP不存在于排除项时才添加到候选
		_, excluded := reverseExcludeMap[s]
		if !excluded {
			candidates = append(candidates, s)
		}
	}

	// 候选人数小于需要的人数则视为失败
	length := len(candidates)
	if length < n {
		return
	}

	// 从候选人中随机选取
	p := rand.Perm(length)
	for i := 0; i < n; i++ {
		ds = append(ds, candidates[p[i]])
	}
	return
}
