package lib

import "path"

func GetObjFilePath(file string) string {
    return path.Join(".got","obj",file)
}

func GetComFilePath(file string,dir string) string {
    return path.Join(".got","com",dir,file)
}
