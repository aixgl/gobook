package singleton

import "sync"

// 或者用接口
type singleton struct {
}
 
var instance *singleton

// singleton.GetInstance()
func GetInstance() *singleton {
    if instance == nil {
        instance = &singleton{}   
    }
    return instance
}

func GetInstanceSafe() *singleton {
    if instance == nil {
        sync.Do(func(){
            instance = &singleton{}   
        })
    }
    return instance
}
