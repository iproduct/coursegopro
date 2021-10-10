package copymap

type GenericMap map[interface{}]interface{}

func CopyMap(dest GenericMap, src GenericMap) {
	for k, v := range src {
		dest[k] = v
	}
}
