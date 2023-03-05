package discovery

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"google.golang.org/grpc/resolver"
)

/*
	由于服务注册到etcd是以 key-value 的形式, 因此此处用于格式化 key-value
*/

// Server 存储在etcd的value格式
type Server struct {
	Name    string `json:"name"`
	Addr    string `json:"addr"`
	Version string `json:"version"`
	Weight  int64  `json:"weight"` // 权重
}

// BuildPrefix 构造key的前缀: 服务名 + 版本号
func BuildPrefix(server *Server) string {
	if server.Version == "" {
		return fmt.Sprintf("/%s", server.Name)
	}
	return fmt.Sprintf("/%s/%s", server.Name, server.Version)
}

// BuildRegisterPath 构造key的完整路径: 前缀 + 地址
func BuildRegisterPath(server *Server) string {
	return fmt.Sprintf("%s/%s", BuildPrefix(server), server.Addr)
}

// ParseValue 将value值反序列化到一个Server实例
func ParseValue(value []byte) (*Server, error) {
	server := &Server{}
	if err := json.Unmarshal(value, server); err != nil {
		return server, err
	}
	return server, nil
}

// SplitPath 切割 key, 获得服务地址
func SplitPath(path string) (*Server, error) {
	server := &Server{}
	strs := strings.Split(path, "/")
	if len(strs) == 0 {
		return server, errors.New("invalid path")
	}
	server.Addr = strs[len(strs)-1]
	return server, nil
}

// Exist helper function
func Exist(l []resolver.Address, addr resolver.Address) bool {
	for i := range l {
		if l[i].Addr == addr.Addr {
			return true
		}
	}
	return false
}

// Remove helper function
func Remove(s []resolver.Address, addr resolver.Address) ([]resolver.Address, bool) {
	for i := range s {
		if s[i].Addr == addr.Addr {
			s[i] = s[len(s)-1]
			return s[:len(s)-1], true
		}
	}
	return nil, false
}

func BuildResolverUrl(app string) string {
	return scheme + ":///" + app
}
