package server

type FileState struct {
	FileName string
	FileCode string
}

type DirectoryState struct {
	AppJSCode  string
	AppCSSCode string
	OtherFiles []FileState
}

func (cd DirectoryState) CreateSysMsgState() (sysMsg string) {
	sysMsg = "The current file contents are as follows:\n\n"
	sysMsg += "`App.js`:\n\n```\n"
	sysMsg += cd.AppJSCode
	sysMsg += "\n```\n\n`App.css`:\n\n```\n"
	sysMsg += cd.AppCSSCode
	sysMsg += "\n```"
	for _, file := range cd.OtherFiles {
		sysMsg += "\n\n`" + file.FileName + ".js`:\n\n```\n" + file.FileCode + "\n```"
	}
	return
}
