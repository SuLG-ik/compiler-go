package main

import (
	"strconv"
	"strings"
)

type AstField struct {
	Label string
	Value string
	Node  *AstNode
}

type AstNode struct {
	Kind   string
	Fields []AstField
	Line   int
	Col    int
}

func newNode(kind string, line, col int) *AstNode {
	return &AstNode{Kind: kind, Line: line, Col: col}
}

func (n *AstNode) addAttr(label, value string) *AstNode {
	n.Fields = append(n.Fields, AstField{Label: label, Value: value})
	return n
}

func (n *AstNode) addChild(label string, child *AstNode) *AstNode {
	if child == nil {
		return n
	}
	n.Fields = append(n.Fields, AstField{Label: label, Node: child})
	return n
}

func (n *AstNode) addItem(child *AstNode) *AstNode {
	if child == nil {
		return n
	}
	n.Fields = append(n.Fields, AstField{Node: child})
	return n
}

func PrintAst(node *AstNode) string {
	if node == nil {
		return ""
	}
	var sb strings.Builder
	sb.WriteString(node.Kind)
	sb.WriteByte('\n')
	printFields(&sb, node, "")
	return sb.String()
}

func printFields(sb *strings.Builder, node *AstNode, prefix string) {
	for i, field := range node.Fields {
		isLast := i == len(node.Fields)-1
		branch := "├── "
		nextPrefix := prefix + "│   "
		if isLast {
			branch = "└── "
			nextPrefix = prefix + "    "
		}

		if field.Node == nil {
			sb.WriteString(prefix)
			sb.WriteString(branch)
			if field.Label != "" {
				sb.WriteString(field.Label)
				sb.WriteString(": ")
			}
			sb.WriteString(field.Value)
			sb.WriteByte('\n')
			continue
		}

		sb.WriteString(prefix)
		sb.WriteString(branch)
		if field.Label != "" {
			sb.WriteString(field.Label)
			sb.WriteString(": ")
		}
		sb.WriteString(field.Node.Kind)
		sb.WriteByte('\n')
		printFields(sb, field.Node, nextPrefix)
	}
}

func quoted(s string) string {
	return strconv.Quote(s)
}
