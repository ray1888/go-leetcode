package undergroundsystem

type EntryInfo struct {
	Id      int
	Station string
	T       int
}

type UndergroundSystem struct {
	UsedTimeMap       map[string]map[string][]int
	InStationID2STMap map[int]EntryInfo
}

func Constructor() UndergroundSystem {
	us := UndergroundSystem{}
	us.UsedTimeMap = make(map[string]map[string][]int)
	us.InStationID2STMap = make(map[int]EntryInfo)
	return us
}

func (this *UndergroundSystem) CheckIn(id int, stationName string, t int) {
	if _, ok := this.InStationID2STMap[id]; ok {
		return
	}
	this.InStationID2STMap[id] = EntryInfo{Id: id, Station: stationName, T: t}
	_, ok := this.UsedTimeMap[stationName]
	if !ok {
		this.UsedTimeMap[stationName] = make(map[string][]int)
	}
	return
}

func (this *UndergroundSystem) CheckOut(id int, stationName string, t int) {
	data, ok := this.InStationID2STMap[id]
	if !ok {
		return
	}
	val, ok := this.UsedTimeMap[data.Station][stationName]
	if !ok {
		this.UsedTimeMap[data.Station][stationName] = make([]int, 0)
	}
	length := len(val)
	var sum int
	if length >= 1 {
		sum = t - data.T + val[length-1]
	} else {
		sum = t - data.T
	}
	this.UsedTimeMap[data.Station][stationName] = append(this.UsedTimeMap[data.Station][stationName], sum)
	delete(this.InStationID2STMap, id)
	return
}

func (this *UndergroundSystem) GetAverageTime(startStation string, endStation string) float64 {
	val, ok := this.UsedTimeMap[startStation]
	if !ok {
		return 0
	}
	end, ok := val[endStation]
	if !ok {
		return 0
	}
	return float64(end[len(end)-1]) / float64(len(end))
}

/**
 * Your UndergroundSystem object will be instantiated and called as such:
 * obj := Constructor();
 * obj.CheckIn(id,stationName,t);
 * obj.CheckOut(id,stationName,t);
 * param_3 := obj.GetAverageTime(startStation,endStation);
 */
