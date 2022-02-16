# filesystem

```golang
type Adapter interface {
	OpenFile() (*os.File, error)
	Write([]byte) error
	WriteString(string) error
	Read()
	ReadString()
	Del()
	DelDir()
	List()
	FileExist()
	DirExist()
}
```