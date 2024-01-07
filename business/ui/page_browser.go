package ui

import (
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

// Show a navigable tree view of the current directory.
func (a *CBAToolApp) callBrowserPage(rootDir string) *tview.TreeView {
	//rootDir := "."
	root := tview.NewTreeNode(rootDir).SetColor(tcell.ColorRed)
	tree := tview.NewTreeView().SetRoot(root).SetCurrentNode(root)

	// A helper function which adds the files and directories of the given path 
	// to the given target node.
	add := func(target *tview.TreeNode, path string) {
		files, err := ioutil.ReadDir(path)
		if err != nil {
			panic(err)
		}
		for _, file := range files {
			node := tview.NewTreeNode(file.Name()).SetReference(filepath.Join(path, file.Name())).SetSelectable(true)
			if file.IsDir() {
				node.SetColor(tcell.ColorGreen)
			}
			target.AddChild(node)
		}
	}

	// Add the current directory to the root node.
	add(root, rootDir)

	// If a directory was selected, open it.
	tree.SetSelectedFunc(func(node *tview.TreeNode) {
		reference := node.GetReference()
		if reference == nil {
			return // Selecting the root node does nothing.
		}
		path := reference.(string)
		info, err := os.Stat(path)
		if err != nil {
			return 
		}
		if info.IsDir() {
			children := node.GetChildren()
			if len(children) == 0 {
				// Load and show files in this directory.
				add(node, path)
			} else {
				// Collapse if visible, expand if collapsed.
				node.SetExpanded(!node.IsExpanded())
			}
		} else {
			a.app.SetRoot(a.layout, true)
			a.app.SetFocus(a.mainMenu)
		}
	})
	return tree
}