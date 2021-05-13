/** 注释
 * 接口.
 * module:github.com/aixgl/ch09/main
 * source: https://github.com/aixgl/gobook/tree/master/basic.magic/ch09/code_main
 */
package main

import (
//"fmt"
)

func main() {
    forRountine()
    channelWaitGoroutine()
    channelSyncGoroutines()
    channelSyncGoroutines2()
    
    concurrencyChannel()
    concurrencyWaitGroup()
    
    timeoutSelect()
    randSelectDemo()
}
