package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Data struct {
	example string
}

type Node struct {
	version   float64
	data      *Data
	node_type string
	notes     []Node
}

func GetLatestVersion(node Node) float64 {
	if len(node.notes) == 0 {
		return node.version
	}
	latest_version := node.version
	for _, note := range node.notes {
		latest_version = GetLatestVersion(note)
	}
	return latest_version
}

func CreateJson(nodes []Node) string {
	var jsonParts []string

	for _, node := range nodes {
		nodeJson := "{"
		nodeJson += "\"version\": " + strconv.FormatFloat(node.version, 'f', -1, 64) + ","
		nodeJson += "\"data\": {"
		nodeJson += "\"example\": \"" + node.data.example + "\"},"
		nodeJson += "\"node_type\": \"" + node.node_type + "\","
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
	Docs = append(Docs, Node{1.0, &data, "crated", []Node{
		Node{1.0, &data, "anonymized", []Node{
			Node{1.0, &data, "signed", []Node{}},
		}},
		Node{1.1, &data, "anonymized", []Node{
			Node{1.1, &data, "signed", []Node{}},
		}},
	}})
	print(GetLatestVersion(Docs[0]))

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
