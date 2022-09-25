package config

type PathConfig struct {
	Base  string
	Image string
	Doc   string
}

func (pc PathConfig) SetPathConfig() {
	pc.Base = "public"
	pc.Image = "images"
	pc.Doc = "others"
}

func (pc PathConfig) GetImage() string {
	return pc.Base + "/" + pc.Image
}

func (pc PathConfig) GetDoc() string {
	return pc.Base + "/" + pc.Doc
}
