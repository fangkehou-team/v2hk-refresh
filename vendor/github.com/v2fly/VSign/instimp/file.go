package instimp

import (
	"archive/zip"
	"crypto/sha256"
	"encoding/hex"
	"github.com/v2fly/VSign/common"
	"github.com/v2fly/VSign/insmgr"
	"io"
	"os"
	"strings"
)

type FileBasedInsYield struct {
	source string
}

type Fileinsbasic struct {
	hash     string
	filename string
}

func (u Fileinsbasic) Instruction() {
	panic("implement me")
}

func (u Fileinsbasic) Hash() string {
	return u.hash
}

func (u Fileinsbasic) Filename() string {
	return u.filename
}

func (u Fileinsbasic) File() {
	panic("implement me")
}

func (fby FileBasedInsYield) InstructionYield(instMgr insmgr.InstructionMgr) {
	//Generate Hash for file

	InstructionYieldFile(instMgr, fileOpener{filename: fby.source}, fby.source)

	//Now if that is a zip
	if strings.HasSuffix(fby.source, ".zip") {
		file, err := os.Open(fby.source)
		if err != nil {
			panic(err)
		}
		fstat, _ := file.Stat()
		InstructionYieldZip(instMgr, file, fby.source, fstat.Size())
		file.Close()
	}
}
func InstructionYieldFile(instMgr insmgr.InstructionMgr, opener Opener, filename string) {
	//SHA256 Basic
	file, err := opener.Open()
	if err != nil {
		panic(err)
	}
	hashv := sha256.New()
	common.Must2(io.Copy(hashv, file))
	hashres := hashv.Sum(nil)
	basic := Fileinsbasic{
		hash:     hex.EncodeToString(hashres[:]),
		filename: filename,
	}
	instMgr.SubmitIns(basic)
	file.Close()

}
func InstructionYieldZip(instMgr insmgr.InstructionMgr, readerAt io.ReaderAt, filename string, size int64) {
	zipf, err := zip.NewReader(readerAt, size)
	if err != nil {
		panic(err)
	}
	filenameOriginal := filename
	filename = filename[:len(filename)-len(".zip")]
	for _, f := range zipf.File {
		if !strings.HasSuffix(f.Name, "/") {
			thisfilename := filename + "/" + f.Name
			InstructionYieldFile(instMgr, f, thisfilename)
			insmgr.NewYieldSingle(
				NewSimpleFilenameKeyValueInst5(thisfilename, "obtain", "local.zip", filenameOriginal, false)).
				InstructionYield(instMgr)
		}
	}
}
func (fby *FileBasedInsYield) AsYield() insmgr.InstructionYield {
	return fby
}

type Opener interface {
	Open() (io.ReadCloser, error)
}

type fileOpener struct {
	filename string
}

func (f fileOpener) Open() (io.ReadCloser, error) {
	return os.Open(f.filename)
}

func NewFileBasedInsYield(path string) *FileBasedInsYield {
	return &FileBasedInsYield{source: path}
}
