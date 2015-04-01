package core

// FileDataNodeProvider is for providing a file specific data node
type FileDataNodeProvider interface {
	// Provide tries to resolve and return a DataNode for the given file.
	Provide(filePath, fileName string, parent DataNode) DataNode
}
