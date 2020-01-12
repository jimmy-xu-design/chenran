package a_star

type Node interface {
	Location() int
	SurNodeList() []Node
	Distance(dst Node) float64
	Value() float64
}

type Link struct {
	curr	Node
	prev	*Link
	G		float64
	F		float64
}

func (l *Link) Init(curr Node, prev *Link, dst Node) {
	l.curr = curr
	l.prev = prev
	if prev != nil {
		l.G = curr.Value() + prev.G
	} else {
		l.G = curr.Value()
	}

	l.F = l.G + curr.Distance(dst)
}

type AStar struct {
	open_list	map[int]*Link
	close_list	map[int]*Link
}

func (p *AStar) Road(l *Link) []Node {
	var road []Node

	curr_link := l
	for curr_link != nil {
		road = append(road, curr_link.curr)
		curr_link = curr_link.prev
	}

	return road
}

func (p *AStar) BFS(src Node, dst Node) []Node {

	p.open_list = make(map[int]*Link)
	p.close_list = make(map[int]*Link)

	link := &Link{}
	link.Init(src, nil, dst)

	p.open_list[link.curr.Location()] = link

	for len(p.open_list) > 0 {
		cur_link := p.retriveBestLink(dst)

		if cur_link.curr == dst {
			return p.Road(cur_link)
		}

		node_list := cur_link.curr.SurNodeList()
		for _, v := range node_list {
			// 检查是否已经在Close表中
			_, ok := p.close_list[v.Location()]
			if ok {
				continue
			}

			sub_link := &Link{}
			sub_link.Init(v, cur_link, dst)

			// 检查是否已在Open表中
			find_link, ok := p.open_list[v.Location()]
			if ok {
				if find_link.F > sub_link.F {
					p.open_list[v.Location()] = sub_link
				}
			} else {
				p.open_list[v.Location()] = sub_link
			}
		}

		p.close_list[cur_link.curr.Location()] = cur_link
	}

	return nil
}

func (p *AStar) retriveBestLink(dst Node) *Link {

	if len(p.open_list) == 0 {
		return nil
	}

	var best_link *Link = nil

	for _, v := range p.open_list {
		if best_link == nil {
			best_link = v
		} else if best_link.F > v.F {
			best_link = v
		} else {
			// no nothing
		}
	}

	delete(p.open_list, best_link.curr.Location())

	return best_link
}