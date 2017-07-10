package template

import (
	"crypto/sha1"
	"encoding/hex"
	"os"
	"path/filepath"

	"github.com/hashicorp/terraform/helper/pathorcontents"
	"github.com/hashicorp/terraform/helper/schema"
)

func dataSourceFileList() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceFileListRead,

		Schema: map[string]*schema.Schema{
			"source_dir": {
				Type:        schema.TypeString,
				Description: "Path to the directory where the files reside",
				Required:    true,
				ForceNew:    true,
			},
			"files": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"content": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceFileListRead(d *schema.ResourceData, meta interface{}) error {
	sourceDir := d.Get("source_dir").(string)

	files := make([]map[string]interface{}, 0)

	err := filepath.Walk(sourceDir, func(p string, f os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if f.IsDir() {
			return nil
		}

		fileContent, _, inputErr := pathorcontents.Read(p)
		if inputErr != nil {
			return inputErr
		}

		checksum := sha1.Sum([]byte(f.Name() + fileContent))
		fileHash := hex.EncodeToString(checksum[:])

		file := make(map[string]interface{})
		file["id"] = fileHash
		file["name"] = f.Name()
		file["content"] = fileContent

		files = append(files, file)

		return nil
	})
	if err != nil {
		return err
	}

	if err := d.Set("files", files); err != nil {
		return err
	}

	hash, err := generateDirHash(sourceDir)
	if err != nil {
		return err
	}
	d.SetId(hash)

	return nil
}
