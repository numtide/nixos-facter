package hwinfo

/*
#cgo pkg-config: hwinfo
#include <hd.h>

// CGO cannot access union type fields, so we do this as a workaround
hd_res_t *hd_res_next(hd_res_t *res) { return res->next; }
hd_resource_types_t hd_res_get_type(hd_res_t *res) { return res->any.type; }
res_any_t hd_res_get_any(hd_res_t *res) { return res->any; }
*/
import "C"

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"slices"
)

//go:generate enumer -type=ResourceType -json -transform=snake -trimprefix ResourceType -output=./resource_enum_type.go
type ResourceType uint

const (
	ResourceTypeAny ResourceType = iota
	ResourceTypePhysMem
	ResourceTypeMem
	ResourceTypeIo
	ResourceTypeIrq
	ResourceTypeDma
	ResourceTypeMonitor

	ResourceTypeSize
	ResourceTypeDiskGeo
	ResourceTypeCache
	ResourceTypeBaud
	ResourceTypeInitStrings
	ResourceTypePppdOption

	ResourceTypeFramebuffer
	ResourceTypeHwaddr
	ResourceTypeLink
	ResourceTypeWlan
	ResourceTypeFc
	ResourceTypePhwaddr
)

func NewResource(res *C.hd_res_t) (Resource, error) {
	if res == nil {
		return nil, nil
	}

	resourceType := ResourceType(C.hd_res_get_type(res))

	switch resourceType {
	case ResourceTypeFc:
		return NewResourceFc(res, resourceType)
	case ResourceTypePhysMem:
		return NewResourcePhysicalMemory(res, resourceType)
	case ResourceTypeMem:
		return NewResourceMemory(res, resourceType)
	case ResourceTypeIo:
		return NewResourceIO(res, resourceType)
	case ResourceTypeIrq:
		return NewResourceIrq(res, resourceType)
	case ResourceTypeDma:
		return NewResourceDma(res, resourceType)
	case ResourceTypeMonitor:
		return NewResourceMonitor(res, resourceType)
	case ResourceTypeSize:
		return NewResourceSize(res, resourceType)
	case ResourceTypeDiskGeo:
		return NewResourceDiskGeo(res, resourceType)
	case ResourceTypeCache:
		return NewResourceCache(res, resourceType)
	case ResourceTypeBaud:
		return NewResourceBaud(res, resourceType)
	case ResourceTypeInitStrings:
		return NewResourceInitStrings(res, resourceType)
	case ResourceTypePppdOption:
		return NewResourcePppdOption(res, resourceType)
	case ResourceTypeFramebuffer:
		return NewResourceFrameBuffer(res, resourceType)
	case ResourceTypeHwaddr, ResourceTypePhwaddr:
		return NewResourceHardwareAddress(res, resourceType)
	case ResourceTypeLink:
		// this is the link status of a network interface and can change when we plug/unplug a cable
		return nil, nil
	case ResourceTypeWlan:
		return NewResourceWlan(res, resourceType)
	default:
		return nil, fmt.Errorf("unexpected resource type: %v", resourceType)
	}
}

func NewResources(hd *C.hd_t) ([]Resource, error) {
	var result []Resource
	for res := hd.res; res != nil; res = C.hd_res_next(res) {
		resource, err := NewResource(res)
		if err != nil {
			return nil, err
		}
		if resource == nil {
			continue
		}
		result = append(result, resource)
	}

	slices.SortFunc(result, func(a, b Resource) int {
		// We don't really care about a proper ordering for resources, just a stable sort that is reasonably quick.
		var err error
		jsonA, err := json.Marshal(a)
		if err != nil {
			log.Panicf("failed to marshal resource: %s", err)
		}
		jsonB, err := json.Marshal(b)
		if err != nil {
			log.Panicf("failed to marshal resource: %s", err)
		}

		return bytes.Compare(jsonA, jsonB)
	})

	return result, nil
}
