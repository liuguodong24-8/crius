package config

import "github.com/BurntSushi/toml"

// Load 导入
func Load(f string) {
	if _, err := toml.DecodeFile(f, &Setting); nil != err {
		panic("配置文件读取错误")
	}
}
