package hwinfo

/*
#cgo pkg-config: hwinfo
#include <hd.h>

// CGO cannot access union type fields, so we do this as a workaround
res_cache_t hd_res_get_cache(hd_res_t *res) { return res->cache; }
*/
import "C"
import "fmt"

type ResourceCache struct {
	Type ResourceType `json:"type"`
	Size uint         `json:"size"`
}

func (r ResourceCache) ResourceType() ResourceType {
	return ResourceTypeCache
}

func NewResourceCache(res *C.hd_res_t, resType ResourceType) (*ResourceCache, error) {
	if res == nil {
		return nil, nil
	}

	if resType != ResourceTypeCache {
		return nil, fmt.Errorf("expected resource type '%s', found '%s'", ResourceTypeCache, resType)
	}

	cache := C.hd_res_get_cache(res)

	return &ResourceCache{
		Type: resType,
		Size: uint(cache.size),
	}, nil
}
