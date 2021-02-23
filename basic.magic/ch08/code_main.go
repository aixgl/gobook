/** 注释
 * 接口.
 * module:github.com/aixgl/ch08/main
 * source: https://github.com/aixgl/gobook/tree/master/basic.magic/ch08/code_main
 */
package main

import (
//"fmt"
)

func main() {
	jsonSimpleMarshal()
	jsonTagEmptyMarshal()

	jsonSimpleUnmarshal()
	jsonDefine()

	encodeProto()
    
    XmlMarshal()
    XmlUnmarshal()
}
