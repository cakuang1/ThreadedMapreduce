package MultiThreaded

func reduceFunction(key string, values []int) KeyValue {
	count := 0
	for _, v := range values {
		count += v
	}
	return KeyValue{Key: key, Value: count}
}


