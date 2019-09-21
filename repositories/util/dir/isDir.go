package dir

import "os"

/**
path 路径
create 如果为true 不存在则创建
*/
func IsDir(path string, create bool) (bool, error) {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			if create == true {
				if err := os.MkdirAll(path, 0777); err != nil {
					return false, err
				}
			}
			return false, nil
		}
		return false, nil
	}

	return false, err
}
