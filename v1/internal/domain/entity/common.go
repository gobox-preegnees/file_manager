package entity

type Path string
type ModTime int64

type Username string
type FolderID string
type ClientID string
type Identifier struct {
	Username Username
	FolderID FolderID
	ClientID ClientID
}
