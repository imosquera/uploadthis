package hooks

import (
	"github.com/imosquera/uploadthis/commands"
)

//The archive command moves files into the archive folder for later deletion
//from some other process
type ArchiveCommand struct {
	commands.Command
}

//empty archive command.  There is nothing to do in this step
//since the prepare command already moves the files to the right place
func (self *ArchiveCommand) Run() ([]string, error) {
	return self.UploadFiles, nil
}
