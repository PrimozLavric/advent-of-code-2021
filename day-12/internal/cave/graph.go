package cave

import (
	"errors"
	"fmt"

	"github.com/PrimozLavric/advent-of-code-2021/day-12/internal/util"
	mapset "github.com/deckarep/golang-set"
)

type nodeType int

const (
	StartNode nodeType = 0
	EndNode            = 1
	BigNode            = 2
	SmallNode          = 3
)

const (
	StartNodeName string = "start"
	EndNodeName          = "end"
)

// Node is composed of identifier, type and pointers to the nodes that are connected to it.
type Node struct {
	id          string
	t           nodeType
	connections []*Node
}

// NewNode creates new node and deduces node type from the name.
func NewNode(id string) (*Node, error) {
	var t nodeType

	if id == StartNodeName {
		t = StartNode
	} else if id == EndNodeName {
		t = EndNode
	} else if util.StringIsAllUppercase(id) {
		t = BigNode
	} else if util.StringIsAllLowercase(id) {
		t = SmallNode
	} else {
		return nil, errors.New(fmt.Sprintf("invalid node id '%s'", id))
	}

	return &Node{id: id, t: t}, nil
}

// Graph contains uniquely named nodes and pointer to start node.
type Graph struct {
	nodes map[string]*Node
	start *Node
}

// NewGraph initializes graph from the provided edge map.
func NewGraph(edgeMap map[string][]string) (*Graph, error) {
	g := Graph{nodes: make(map[string]*Node)}

	// Create nodes and initialize connections (edges)
	for nodeId, connectedNodesIds := range edgeMap {
		var err error

		nodeA, ok := g.nodes[nodeId]

		// Create node if it does not exist.
		if !ok {
			nodeA, err = NewNode(nodeId)

			if err != nil {
				return nil, err
			}

			g.nodes[nodeId] = nodeA
		}

		for _, connectedNodeId := range connectedNodesIds {
			connectedNode, ok := g.nodes[connectedNodeId]

			// Create node if it does not exist.
			if !ok {
				connectedNode, err = NewNode(connectedNodeId)

				if err != nil {
					return nil, err
				}

				g.nodes[connectedNodeId] = connectedNode
			}

			nodeA.connections = append(nodeA.connections, connectedNode)
			connectedNode.connections = append(connectedNode.connections, nodeA)
		}
	}

	// Store ptr to start node.
	var ok bool
	g.start, ok = g.nodes[StartNodeName]

	if !ok {
		return nil, errors.New("missing start node")
	}

	// Make sure that end node is present.
	_, ok = g.nodes[EndNodeName]

	if !ok {
		return nil, errors.New("missing end node")
	}

	return &g, nil
}

// FindNumberOfPathsToEnd finds number of paths from the start node to the end node. Small node can only be visited once
// unless visitOneSmallNodeTwice is set to true.
func (g *Graph) FindNumberOfPathsToEnd(visitOneSmallNodeTwice bool) int {
	visitedSmallNodes := mapset.NewSet()
	numPaths := 0
	canVisitTwice := visitOneSmallNodeTwice

	var nodeVisitor func(node *Node)
	nodeVisitor = func(node *Node) {
		if node.t == EndNode {
			// Reached end. Increment num paths and stop traversal here.
			numPaths++
			return
		} else if node.t == SmallNode {
			// Check if this small node was already visited.
			if visitedSmallNodes.Contains(node) {
				// Stop here if node cannot be visited additional time.
				if !canVisitTwice {
					return
				}

				// Use double visit and reset it on exit.
				canVisitTwice = false
				defer func() { canVisitTwice = true }()
			} else {
				// Add node among visited and remove it on exit.
				visitedSmallNodes.Add(node)
				defer visitedSmallNodes.Remove(node)
			}
		}

		// Continue search for end in the connecting nodes.
		for _, conn := range node.connections {
			if conn.t == StartNode {
				continue
			}

			nodeVisitor(conn)
		}
	}

	nodeVisitor(g.start)

	return numPaths
}
