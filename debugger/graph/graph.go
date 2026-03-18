package graph

import (
	"fmt"

	"gioui.org/layout"
	"github.com/ddkwork/HyperDbg/debugger/api"
	"github.com/ddkwork/golibrary/std/safemap"
)

type NodeType int

const (
	NodeTypeEntry NodeType = iota
	NodeTypeExit
	NodeTypeBasic
	NodeTypeCall
	NodeTypeBranch
	NodeTypeLoop
)

type EdgeType int

const (
	EdgeTypeUnconditional EdgeType = iota
	EdgeTypeConditional
	EdgeTypeCall
	EdgeTypeReturn
)

type GraphNode struct {
	ID           uint64
	Address      uint64
	Type         NodeType
	Name         string
	Instructions []string
	Edges        []*GraphEdge
}

type GraphEdge struct {
	FromID    uint64
	ToID      uint64
	Type      EdgeType
	Condition string
}

type ControlFlowGraph struct {
	Nodes *safemap.M[uint64, *GraphNode]
	Edges *safemap.M[uint64, []*GraphEdge]
}

type Manager struct {
	graphs *safemap.M[uint64, *ControlFlowGraph]
}

func New() api.Interface {
	return &Manager{
		graphs: safemap.New[uint64, *ControlFlowGraph](),
	}
}

func NewControlFlowGraph() *ControlFlowGraph {
	return &ControlFlowGraph{
		Nodes: safemap.New[uint64, *GraphNode](),
		Edges: safemap.New[uint64, []*GraphEdge](),
	}
}

func (m *Manager) CreateGraph(entryAddress uint64) *ControlFlowGraph {
	graph := NewControlFlowGraph()
	m.graphs.Update(entryAddress, graph)
	return graph
}

func (m *Manager) GetGraph(entryAddress uint64) *ControlFlowGraph {
	graph, _ := m.graphs.Get(entryAddress)
	return graph
}

func (m *Manager) DeleteGraph(entryAddress uint64) {
	m.graphs.Delete(entryAddress)
}

func (m *Manager) GetAllGraphs() []*ControlFlowGraph {
	result := make([]*ControlFlowGraph, 0)
	for _, graph := range m.graphs.Range() {
		result = append(result, graph)
	}
	return result
}

func (m *Manager) Clear() {
	m.graphs.Reset()
}

func (g *ControlFlowGraph) AddNode(id, address uint64, nodeType NodeType, name string) *GraphNode {
	node := &GraphNode{
		ID:           id,
		Address:      address,
		Type:         nodeType,
		Name:         name,
		Instructions: make([]string, 0),
		Edges:        make([]*GraphEdge, 0),
	}

	g.Nodes.Update(id, node)
	return node
}

func (g *ControlFlowGraph) GetNode(id uint64) *GraphNode {
	node, _ := g.Nodes.Get(id)
	return node
}

func (g *ControlFlowGraph) GetNodeByAddress(address uint64) *GraphNode {
	for _, node := range g.Nodes.Range() {
		if node.Address == address {
			return node
		}
	}
	return nil
}

func (g *ControlFlowGraph) DeleteNode(id uint64) {
	g.Nodes.Delete(id)
	g.Edges.Delete(id)

	for _, edges := range g.Edges.Range() {
		for i, edge := range edges {
			if edge.ToID == id {
				edges = append(edges[:i], edges[i+1:]...)
				break
			}
		}
	}
}

func (g *ControlFlowGraph) AddEdge(fromID, toID uint64, edgeType EdgeType, condition string) {
	edge := &GraphEdge{
		FromID:    fromID,
		ToID:      toID,
		Type:      edgeType,
		Condition: condition,
	}

	edges, _ := g.Edges.Get(fromID)
	if edges == nil {
		edges = make([]*GraphEdge, 0)
		g.Edges.Update(fromID, edges)
	}

	edges = append(edges, edge)
	g.Edges.Update(fromID, edges)

	if node, exists := g.Nodes.Get(fromID); exists {
		node.Edges = append(node.Edges, edge)
	}
}

func (g *ControlFlowGraph) GetEdges(fromID uint64) []*GraphEdge {
	edges, exists := g.Edges.Get(fromID)
	if exists {
		result := make([]*GraphEdge, len(edges))
		copy(result, edges)
		return result
	}
	return nil
}

func (g *ControlFlowGraph) GetIncomingEdges(toID uint64) []*GraphEdge {
	result := make([]*GraphEdge, 0)
	for _, edges := range g.Edges.Range() {
		for _, edge := range edges {
			if edge.ToID == toID {
				result = append(result, edge)
			}
		}
	}
	return result
}

func (g *ControlFlowGraph) DeleteEdge(fromID, toID uint64) {
	edges, exists := g.Edges.Get(fromID)
	if exists {
		for i, edge := range edges {
			if edge.ToID == toID {
				edges = append(edges[:i], edges[i+1:]...)
				g.Edges.Update(fromID, edges)
				break
			}
		}
	}
}

func (g *ControlFlowGraph) AddInstruction(nodeID uint64, instruction string) {
	if node, exists := g.Nodes.Get(nodeID); exists {
		node.Instructions = append(node.Instructions, instruction)
	}
}

func (g *ControlFlowGraph) GetInstructions(nodeID uint64) []string {
	if node, exists := g.Nodes.Get(nodeID); exists {
		result := make([]string, len(node.Instructions))
		copy(result, node.Instructions)
		return result
	}
	return nil
}

func (g *ControlFlowGraph) GetNodeCount() int {
	count := 0
	for range g.Nodes.Range() {
		count++
	}
	return count
}

func (g *ControlFlowGraph) GetEdgeCount() int {
	count := 0
	for _, edges := range g.Edges.Range() {
		count += len(edges)
	}
	return count
}

func (g *ControlFlowGraph) GetEntryNode() *GraphNode {
	for _, node := range g.Nodes.Range() {
		if node.Type == NodeTypeEntry {
			return node
		}
	}
	return nil
}

func (g *ControlFlowGraph) GetExitNodes() []*GraphNode {
	result := make([]*GraphNode, 0)
	for _, node := range g.Nodes.Range() {
		if node.Type == NodeTypeExit {
			result = append(result, node)
		}
	}
	return result
}

func (g *ControlFlowGraph) GetBasicBlocks() []*GraphNode {
	result := make([]*GraphNode, 0)
	for _, node := range g.Nodes.Range() {
		if node.Type == NodeTypeBasic {
			result = append(result, node)
		}
	}
	return result
}

func (g *ControlFlowGraph) GetBranchNodes() []*GraphNode {
	result := make([]*GraphNode, 0)
	for _, node := range g.Nodes.Range() {
		if node.Type == NodeTypeBranch {
			result = append(result, node)
		}
	}
	return result
}

func (g *ControlFlowGraph) GetLoopNodes() []*GraphNode {
	result := make([]*GraphNode, 0)
	for _, node := range g.Nodes.Range() {
		if node.Type == NodeTypeLoop {
			result = append(result, node)
		}
	}
	return result
}

func (g *ControlFlowGraph) SetNodeType(id uint64, nodeType NodeType) {
	if node, exists := g.Nodes.Get(id); exists {
		node.Type = nodeType
	}
}

func (g *ControlFlowGraph) SetNodeName(id uint64, name string) {
	if node, exists := g.Nodes.Get(id); exists {
		node.Name = name
	}
}

func (g *ControlFlowGraph) SetEdgeType(fromID, toID uint64, edgeType EdgeType) {
	if edges, exists := g.Edges.Get(fromID); exists {
		for _, edge := range edges {
			if edge.ToID == toID {
				edge.Type = edgeType
				break
			}
		}
	}
}

func (g *ControlFlowGraph) SetEdgeCondition(fromID, toID uint64, condition string) {
	if edges, exists := g.Edges.Get(fromID); exists {
		for _, edge := range edges {
			if edge.ToID == toID {
				edge.Condition = condition
				break
			}
		}
	}
}

func (g *ControlFlowGraph) HasNode(id uint64) bool {
	_, exists := g.Nodes.Get(id)
	return exists
}

func (g *ControlFlowGraph) HasEdge(fromID, toID uint64) bool {
	if edges, exists := g.Edges.Get(fromID); exists {
		for _, edge := range edges {
			if edge.ToID == toID {
				return true
			}
		}
	}
	return false
}

func (g *ControlFlowGraph) Clear() {
	g.Nodes.Reset()
	g.Edges.Reset()
}

func (g *ControlFlowGraph) GetSuccessors(nodeID uint64) []*GraphNode {
	result := make([]*GraphNode, 0)
	if edges, exists := g.Edges.Get(nodeID); exists {
		for _, edge := range edges {
			if node, exists := g.Nodes.Get(edge.ToID); exists {
				result = append(result, node)
			}
		}
	}
	return result
}

func (g *ControlFlowGraph) GetPredecessors(nodeID uint64) []*GraphNode {
	result := make([]*GraphNode, 0)
	for fromID, edges := range g.Edges.Range() {
		for _, edge := range edges {
			if edge.ToID == nodeID {
				if node, exists := g.Nodes.Get(fromID); exists {
					result = append(result, node)
				}
			}
		}
	}
	return result
}

func (g *ControlFlowGraph) AnalyzeCyclomaticComplexity() int {
	edges := g.GetEdgeCount()
	nodes := g.GetNodeCount()

	if nodes == 0 {
		return 0
	}

	return edges - nodes + 2
}

func (m *Manager) AnalyzeFunction(entryAddress uint64, disasmFunc func(uint64) (string, uint64, error)) (*ControlFlowGraph, error) {
	graph := m.CreateGraph(entryAddress)

	nodeID := uint64(1)
	entryNode := graph.AddNode(nodeID, entryAddress, NodeTypeEntry, "entry")
	nodeID++

	currentAddress := entryAddress
	visited := make(map[uint64]bool)
	prevNode := entryNode

	for !visited[currentAddress] {
		visited[currentAddress] = true

		instruction, nextAddr, err := disasmFunc(currentAddress)
		if err != nil {
			return nil, fmt.Errorf("disassembly error at 0x%X: %v", currentAddress, err)
		}

		if nextAddr == 0 {
			exitNode := graph.AddNode(nodeID, currentAddress, NodeTypeExit, "exit")
			nodeID++
			graph.AddEdge(prevNode.ID, exitNode.ID, EdgeTypeUnconditional, "")
			break
		}

		basicBlockNode := graph.AddNode(nodeID, currentAddress, NodeTypeBasic, fmt.Sprintf("block_%X", currentAddress))
		basicBlockNode.Instructions = append(basicBlockNode.Instructions, instruction)
		nodeID++

		graph.AddEdge(prevNode.ID, basicBlockNode.ID, EdgeTypeUnconditional, "")
		prevNode = basicBlockNode

		currentAddress = nextAddr
	}

	return graph, nil
}

func (m *Manager) Layout() layout.Widget {
	return func(gtx layout.Context) layout.Dimensions {
		return layout.Dimensions{}
	}
}

func (m *Manager) Update() error {
	return nil
}

func (m *Manager) Self() any {
	return m
}
