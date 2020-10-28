package template

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	r "github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

type testTemplateFilelist struct {
	name    string
	content string
}

func testTemplateFilelistWriteFiles(testCase string, files map[string]testTemplateFilelist) (in string, err error) {
	in, err = ioutil.TempDir(os.TempDir(), testCase)
	if err != nil {
		return
	}

	for name, file := range files {
		path := filepath.Join(in, name)

		err = os.MkdirAll(filepath.Dir(path), 0777)
		if err != nil {
			return
		}

		err = ioutil.WriteFile(path, []byte(file.content), 0777)
		if err != nil {
			return
		}
	}
	return
}

func TestTemplateFileListOutput(t *testing.T) {
	var cases = []struct {
		testCase string
		files    map[string]testTemplateFilelist
	}{
		{
			testCase: "terraform_template_file_list_1",
			files: map[string]testTemplateFilelist{
				"foo.txt": {"foo.txt", "bar"},
			},
		},
		{
			testCase: "terraform_template_file_list_2",
			files: map[string]testTemplateFilelist{
				"foo.txt":           {"foo.txt", "bar"},
				"nested/monkey.txt": {"monkey.txt", "ooh-ooh-ooh-eee-eee"},
				"cheese.txt":        {"cheese.txt", "cheddar"},
			},
		},
	}

	for _, tt := range cases {
		// Write the desired templates in a temporary directory.
		in, err := testTemplateFilelistWriteFiles(tt.testCase, tt.files)
		if err != nil {
			t.Skipf("could not write templates to temporary directory: %s", err)
			continue
		}
		defer os.RemoveAll(in)

		// Run test case.
		r.UnitTest(t, r.TestCase{
			Providers: testProviders,
			Steps: []r.TestStep{
				{
					Config: testTemplateFileListConfig(in),
					Check: func(s *terraform.State) error {
						got := s.RootModule().Outputs["files"].Value.([]interface{})

						if len(got) != len(tt.files) {
							return fmt.Errorf("\ntest case:\n%s\ngot:\n%d files\nwant:\n%d files\n", tt.testCase, len(got), len(tt.files))
						}

						output := make(map[string]string)
						for _, v := range got {
							f := v.(map[string]interface{})
							output[f["name"].(string)] = f["content"].(string)
						}

						for _, file := range tt.files {
							content, ok := output[file.name]
							if !ok {
								return fmt.Errorf("\ntest case:\n%s\nfile missing:\n%s\n", tt.testCase, file.name)
							}

							if content != file.content {
								return fmt.Errorf("\ntest case:\n%s\ncontent mis-match for file:\n%s\ngot:\n%s\nwant:\n%s\n", tt.testCase, file.name, content, file.content)
							}
						}
						return nil
					},
				},
			},
		})
	}
}

func testTemplateFileListConfig(sourceDir string) string {
	return fmt.Sprintf(`
		data "template_file_list" "list" {
			source_dir = "%s"
		}
		output "files" {
			value = "${data.template_file_list.list.files}"
		}`, sourceDir)
}
