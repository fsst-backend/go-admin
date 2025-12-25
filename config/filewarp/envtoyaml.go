package filewrap

import (
	"os"

	"github.com/go-admin-team/go-admin-core/config/source"
	"github.com/go-admin-team/go-admin-core/config/source/file"
	"gopkg.in/yaml.v3"
)

// fileWrap wraps the original file source and injects env vars into YAML
type EnvToYaml struct {
	inner source.Source
	path  string
}

func NewFileWrap(path string) source.Source {
	inner := file.NewSource(file.WithPath(path))

	return &EnvToYaml{
		inner: inner,
		path:  path,
	}
}

func (f *EnvToYaml) Read() (*source.ChangeSet, error) {
	// 读取文件
	b, err := os.ReadFile(f.path)
	if err != nil {
		return nil, err
	}

	// 替换环境变量
	content := os.ExpandEnv(string(b))

	// 如果是 YAML，可以进一步解析成 JSON bytes（ChangeSet.Data一般是JSON）
	var data interface{}
	if err := yaml.Unmarshal([]byte(content), &data); err != nil {
		return nil, err
	}

	// 再 marshal 回 YAML 或 JSON bytes
	out, err := yaml.Marshal(data)
	if err != nil {
		return nil, err
	}

	info, err := os.Stat(f.path)
	if err != nil {
		return nil, err
	}

	cs := &source.ChangeSet{
		Format:    "yaml",
		Source:    f.String(),
		Timestamp: info.ModTime(),
		Data:      out,
	}
	cs.Checksum = cs.Sum()
	return cs, nil
}

func (f *EnvToYaml) String() string {
	return f.inner.String()
}

func (f *EnvToYaml) Watch() (source.Watcher, error) {
	// 直接用原来的 inner Watch
	return f.inner.Watch()
}

func (f *EnvToYaml) Write(cs *source.ChangeSet) error {
	// 不实现写入
	return nil
}
