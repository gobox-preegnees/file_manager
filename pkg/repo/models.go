package repo

type Batch struct {
	Username   string
	FolderID   string
	ClientID   string
	Path       string
	Hash       string
	ModTime    int
	Part       int
	CountParts int
	PartSize   int
	Offset     int
	SizeFile   int
	Content    []byte
}
