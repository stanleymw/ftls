package protocol

// import "os"

type FtlsRequest struct {
	Opcode byte
}

type FtlsResponse struct {
	Body string
}

type FtlsFile struct {
	Size int64
}

const (
	GET_SERVER_INFO       byte = 0x0
	CLOSE_CONNECTION      byte = 0x1
	GET_CURRENT_DIRECTORY byte = 0x2
	GET_DIRECTORY_LIST    byte = 0x3
	RETRIEVE_FILE         byte = 0x4
)
