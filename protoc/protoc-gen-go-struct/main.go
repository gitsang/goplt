package main

import (
	"log"

	"google.golang.org/protobuf/compiler/protogen"
)

func generateFile(p *protogen.Plugin, file *protogen.File) {

	f := p.NewGeneratedFile(file.GeneratedFilenamePrefix+".go", file.GoImportPath)

	f.P("// Code generated by protoc-gen-go-struct. DO NOT EDIT.")
	f.P()

	log.Println("filename_prefix: ", file.GeneratedFilenamePrefix)
	log.Println("import: ", file.GoImportPath)
	log.Println("package: ", file.GoPackageName)
	for _, m := range file.Messages {
		log.Println("message: ", m.GoIdent)
		for _, field := range m.Fields {
			log.Println("name: ", field.GoName)
			log.Println("jsonName: ", field.Desc.JSONName())
			log.Println("type: ", field.Desc.Kind())
			log.Println("leadingComment: ", field.Comments.Leading.String())
			log.Println("trailingComment: ", field.Comments.Trailing.String())
		}
	}
}

func main() {
	protogen.Options{}.Run(func(p *protogen.Plugin) error {
		for _, f := range p.Files {
			if !f.Generate {
				continue
			}
			generateFile(p, f)
		}
		return nil
	})
}
