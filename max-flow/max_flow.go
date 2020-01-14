package max_flow

type Node interface {
	ID() string
	Capacity()	int

	Match(p Node)	bool
	AddFlow(num int)
}

type Flow struct {
	From	Node
	To		Node
	Count	int
}

type MaxFlow struct {
	flow_table map[string]*Flow
}

const MAX_NODE = 1000000

func min(x int, y int) int {
	if x > y {
		return y
	} else {
		return x
	}
}

func (mf* MaxFlow) Match(nx []Node, ny []Node) []*Flow {

	mf.flow_table = make(map[string]*Flow)

	for _, x := range nx {

		new_flow_list := mf.searchPath(x, ny, x.Capacity())
		mf.AddFlow(new_flow_list)
	}

	var flow_list []*Flow
	for _, v := range mf.flow_table {
		flow_list = append(flow_list, v)
	}

	return flow_list
}

func (mf *MaxFlow) AddFlow(new_flow_list []*Flow) {

	for _, n := range new_flow_list {
		key := n.From.ID() + "-" + n.To.ID()
		n.From.AddFlow(n.Count)
		n.To.AddFlow(n.Count)

		flow, ok := mf.flow_table[key]
		if ok {
			flow.Count += n.Count
			if flow.Count == 0 {
				delete(mf.flow_table, key)
			}
		} else {
			mf.flow_table[key] = n
		}
	}
}

func (mf* MaxFlow) searchPath(start Node, ny []Node, flow_limit int) []*Flow {

	var flow_list []*Flow

	// 先匹配空闲
	for _, v := range ny {
		if v.Capacity() == 0 {
			continue
		}

		if start.Match(v) {
			flow_num := min(flow_limit, v.Capacity())
			flow := &Flow{start, v, flow_num}
			flow_list = append(flow_list, flow)

			return flow_list
		}
	}

	// 增广路
	for _, v := range mf.flow_table {

		if start.Match(v.To) {
			new_flow_list := mf.searchPath(v.From, ny, v.Count)
			if len(new_flow_list) > 0 {

				flow_num := new_flow_list[0].Count
				flow := &Flow{start, v.To, flow_num}
				reduce_flow := &Flow{v.From, v.To, -flow_num}

				flow_list = append(flow_list, flow)
				flow_list = append(flow_list, reduce_flow)
				flow_list = append(flow_list, new_flow_list...)

				return flow_list
			}
		}
	}

	return flow_list
}