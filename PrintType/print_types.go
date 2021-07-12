package print

type PrintType string

//====================================================
const (
	TypeError   PrintType = "Error"
	TypeWarn    PrintType = "Warn"
	TypeInfo    PrintType = "Info"
	TypeSuccess PrintType = "Success"
	TypeNeutral PrintType = "Neutral"
)
