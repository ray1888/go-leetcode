package backtracing

import (
	"sort"
)

type Des struct {
	Name    string
	Visited bool
}

type DesList []*Des

func (d DesList) Less(i, j int) bool {
	return d[i].Name < d[j].Name
}

func (d DesList) Len() int {
	return len(d)
}

func (d DesList) Swap(i, j int) {
	tmp := d[j].Name
	d[j].Name = d[i].Name
	d[i].Name = tmp
	return
}

func dfs(start string, total int, rMap map[string]DesList, result *[]string) bool {
	if len(*result) == total {
		return true
	}
	val, _ := rMap[start]
	if len(val) == 0 {
		return false
	}
	for _, item := range val {
		if item.Visited {
			continue
		}
		item.Visited = true
		*result = append(*result, item.Name)
		if dfs(item.Name, total, rMap, result) {
			return true
		}
		item.Visited = false
		*result = (*result)[:len(*result)-1]
	}

	return false
}

func findItinerary(tickets [][]string) []string {
	if len(tickets) == 0 {
		return []string{}
	}
	routeMap := make(map[string]DesList)
	for _, item := range tickets {
		if _, ok := routeMap[item[0]]; !ok {
			routeMap[item[0]] = make(DesList, 0)
		}
		newItem := Des{Name: item[1]}
		routeMap[item[0]] = append(routeMap[item[0]], &newItem)
	}

	for _, value := range routeMap {
		sort.Sort(value)
	}
	total := len(tickets) + 1
	result := make([]string, 0)
	result = append(result, "JFK")
	dfs("JFK", total, routeMap, &result)
	return result
}
