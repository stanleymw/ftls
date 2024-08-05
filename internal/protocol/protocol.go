package protocol

type Request struct {
	Opcode byte
}

type DirList []File

type Response struct {
	Body string
}

type File struct {
	Name  string
	IsDir bool
	Size  int64
}

const (
	GET_SERVER_INFO       byte = 0x0
	CLOSE_CONNECTION      byte = 0x1
	GET_CURRENT_DIRECTORY byte = 0x2
	GET_DIRECTORY_LIST    byte = 0x3
	RETRIEVE_FILE         byte = 0x4
)
