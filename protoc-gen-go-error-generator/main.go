package main

import (
	"flag"
	"fmt"
	"github.com/classtorch/go-error-generator/protoc-gen-go-error-generator/internal"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/types/pluginpb"
)

var (
	showVersion        = flag.Bool("version", false, "print the Version and exit")
	descriptorFileFlag = flag.String("descriptor_file", "", "custom error descriptor file path")
	mergeErrorFlag     = flag.Bool("merge_error", false, "merge error or not")
	mergeErrorPathFlag = flag.String("merge_error_path", "", "merge error file path")
)

func main() {
	flag.Parse()
	if *showVersion {
		fmt.Printf("protoc-gen-go-error-generator %s\n", internal.Version)
		return
	}
	protogen.Options{
		ParamFunc: flag.CommandLine.Set,
	}.Run(func(gen *protogen.Plugin) error {
		gen.SupportedFeatures = uint64(pluginpb.CodeGeneratorResponse_FEATURE_PROTO3_OPTIONAL)
		for i, f := range gen.Files {
			if !f.Generate {
				continue
			}
			internal.GenerateFile(gen, f, len(gen.Files)-1 == i, descriptorFileFlag, mergeErrorPathFlag, mergeErrorFlag)
		}
		return nil
	})
}
