package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Define the custom type
type NodeType string

// Define the allowed values as constants
const (
	Created    NodeType = "created"
	Anonymized NodeType = "anonymized"
	Signed     NodeType = "signed"
)

type Data struct {
	example string
}

type Node struct {
	version   float64
	data      *Data
	node_type NodeType
	notes     []Node
}

// GetLatestVersion returns the latest version number
func GetLatestVersion(node Node) float64 {
	if len(node.notes) == 0 {
		return node.version
	}
	latest_version := node.version
	for _, note := range node.notes {
		version := GetLatestVersion(note)
		if version > latest_version {
			latest_version = version
		}
	}
	return latest_version
}

// GetLatestNodeOfType returns the latest node of the specified type
func GetLatestNodeOfType(node Node, node_type NodeType) Node {
	var latestNode Node

	// Check if the current node matches the type and is the latest
	if node.node_type == node_type {
		latestNode = node
	}

	// Iterate through the notes to find the latest node of the specified type
	for _, note := range node.notes {
		noteOfType := GetLatestNodeOfType(note, node_type)
		if noteOfType.version > latestNode.version {
			latestNode = noteOfType
		}
	}

	return latestNode
}
	

// CreateBranch creates a new branch from the current node
func CreateBranch(current_n *Node, data *Data, node_type NodeType, new_id float64) {
	new_node := Node{new_id, data, node_type, []Node{}}
	current_n.notes = append(current_n.notes, new_node)
}

// CreateJson generates a JSON representation of the nodes
func CreateJson(nodes []Node) string {
	var jsonParts []string

	for _, node := range nodes {
		nodeJson := "{"
		nodeJson += "\"version\": " + strconv.FormatFloat(node.version, 'f', -1, 64) + ","
		nodeJson += "\"data\": {"
		nodeJson += "\"example\": \"" + node.data.example + "\"},"
		nodeJson += "\"node_type\": \"" + string(node.node_type) + "\","
		nodeJson += "\"notes\": ["
		if len(node.notes) > 0 {
			nodeJson += CreateJson(node.notes)
		}
		nodeJson += "]}"
		jsonParts = append(jsonParts, nodeJson)
	}

	return "[" + strings.Join(jsonParts, ",") + "]"
}

func main() {
	Docs := []Node{}
	data := Data{"Example"}
	Docs = append(Docs, Node{1.0, &data, Created, []Node{
		{1.0, &data, Anonymized, []Node{
			{1.0, &data, Signed, []Node{}},
		}},
		{1.1, &data, Anonymized, []Node{
			{1.1, &data, Signed, []Node{}},
		}},
	}})


	// Modify the existing node by reference
	doc1 := &Docs[0]
	CreateBranch(doc1, &data, Anonymized, GetLatestVersion(*doc1)+0.1)
	CreateBranch(doc1, &data, Anonymized, GetLatestVersion(*doc1)+0.1)
	anonym := &doc1.notes[0]
	CreateBranch(anonym, &data, Signed, GetLatestVersion(*anonym)+0.1)
	CreateBranch(anonym, &data, Signed, GetLatestVersion(*anonym)+0.1)

	// Get the latest node of Signed type
	latestSigned := GetLatestNodeOfType(Docs[0], Signed)
	fmt.Println("Latest Signed Node:", latestSigned)

	latestVersion := GetLatestVersion(Docs[0])
	fmt.Println("Latest Version:", latestVersion)

	// Generate JSON
	json := CreateJson(Docs)
	fmt.Println("Generated JSON:", json)

	// Save JSON to file
	file, err := os.Create("output.json")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	_, err = file.WriteString(json)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}

	fmt.Println("JSON saved to output.json")
}
