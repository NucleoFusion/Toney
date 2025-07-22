package filetree

import (
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"github.com/SourcewareLab/Toney/internal/colors"
	"github.com/SourcewareLab/Toney/internal/config"
	"github.com/charmbracelet/lipgloss"
)

func CreateTree() (*Node, error) {
	home, _ := os.UserHomeDir()

	root, err := buildTree(nil, filepath.Join(home, config.AppConfig.General.NotesDir), 0)
	if err != nil {
		return nil, err
	}

	root.IsExpanded = true

	return root, nil
}

func buildTree(parent *Node, path string, depth int) (*Node, error) {
	info, err := os.Stat(path)
	if err != nil {
		return nil, err
	}

	node := Node{
		Name:        info.Name(),
		Parent:      parent,
		Depth:       depth,
		IsDirectory: info.IsDir(),
		IsExpanded:  false,
	}

	if node.IsDirectory {
		files, _ := os.ReadDir(path)
		for _, file := range files {
			childPath := filepath.Join(path, file.Name())
			child, _ := buildTree(&node, childPath, depth+1)
			node.Children = append(node.Children, child)
		}
	}

	return &node, nil
}

func BuildNodeTree(n *Node, prefix string, isLast bool, curr *Node) string {
	var sb strings.Builder

	branch := "├─ "
	if isLast {
		branch = "└─ "
	}

	icon := config.AppConfig.Styles.Icons.FileIcon
	if n.IsDirectory {
		icon = config.AppConfig.Styles.Icons.FolderIcon
	}

	newPrefix := prefix

	// Build the line for this node
	line := ""
	if n.Parent == nil {
		line = icon + " " + n.Name
	} else {
		line = prefix + branch + icon + " " + n.Name
	}

	// line += "\n\n\n\n"

	if n == curr {
		line = lipgloss.NewStyle().Background(colors.ColorPalette().MenuSelectedBg).
			Foreground(colors.ColorPalette().MenuSelectedText).Render(line)
	} else {
		line = lipgloss.NewStyle().Foreground(colors.ColorPalette().Text).Render(line)
	}

	sb.WriteString(line + "\n")

	// Update prefix for children
	if n.Parent != nil {
		if isLast {
			newPrefix += "   "
		} else {
			newPrefix += "│  "
		}
	}

	if n.IsExpanded {
		for i, child := range n.Children {
			sb.WriteString(BuildNodeTree(child, newPrefix, i == len(n.Children)-1, curr))
		}
	}

	return sb.String()
}

func ListFilesRel(baseDir string) ([]string, error) {
	var relPaths []string

	err := filepath.WalkDir(baseDir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			return nil
		}
		rel, err := filepath.Rel(baseDir, path)
		if err != nil {
			return err
		}
		relPaths = append(relPaths, rel)
		return nil
	})
	if err != nil {
		return nil, err
	}
	return relPaths, nil
}
