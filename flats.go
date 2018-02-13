package main

type flats map[int]flat

func (f flats) add(flatInstance flat) {
	f[flatInstance.id] = flatInstance
	return
}

func selectFromMap(f *flats, chooser func(*flat) (bool, error)) (*flats, error) {
	ret := flats{}
	for k, v := range *f {
		good, err := chooser(&v)
		if err != nil {
			return nil, err
		}
		if good {
			ret[k] = v
		}
	}
	return &ret, nil
}
