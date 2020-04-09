package global

import "context"

func NoDecodeRequestFunc(ctx context.Context, i interface{}) (request interface{}, err error) {
	return i, nil
}

func NoEncodeResponseFunc(ctx context.Context, i interface{}) (response interface{}, err error) {
	return i, nil
}

func NoEncodeRequestFunc(ctx context.Context, i interface{}) (request interface{}, err error) {
	return i, nil
}

func NoDecodeResponseFunc(ctx context.Context, i interface{}) (response interface{}, err error) {
	return i, nil
}
