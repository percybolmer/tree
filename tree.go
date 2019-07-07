// package tree is used to enable tree structure for any kind of struct
// it is ment to be embedded into currently existing structs to enable treeifying
//
//
// Current tree format is in regard of Extjs way of treeing
//
// @author perbol
// @version 0.1
package tree

// Node is a tree node that is meant to be embedded
type Node struct {
	Text     string        `json:"text"`
	Expanded bool          `json:"expanded"`
	Children []interface{} `json:"children"`
}
