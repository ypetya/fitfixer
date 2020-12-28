package fitfixer

type IEnhancer interface {
	Enhance(targetFile string, toEnhance string, with string)
}
