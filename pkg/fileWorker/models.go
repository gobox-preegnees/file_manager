package fileworker

type BatchForSave struct {
	Path       string
	Hash       string
	ModTime    int
	FullSize   int
	Part       int
	CountParts int
	Content    int
}
