package fileworker

type Batch struct {
	Username   string
	FolderID   string
	ClientID   string
	Path       string
	Hash       string
	ModTime    string
	FullSize   int
	Part       int
	CountParts int
	Content    []byte
	PartSize   int
	Offset     int
	SizeFile   int
}
