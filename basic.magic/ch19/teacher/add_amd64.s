#include "textflag.h"
TEXT ·Add+0(SB), $0-24
MOVQ	a+0(FP), AX
MOVQ	b+8(FP), BX
ADDQ	AX, BX
MOVQ	BX, ret+16(FP)
RET

GLOBL ·MyNum(SB), $8
DATA ·MyNum+0(SB)/2, $0x22B8

/*
// 因内存常量无法被引用，获取地址，先放在一个变量中
GLOBL ·jobData(SB),NOPTR,$8
DATA  ·jobData(SB)/8,$"gopher"

// 在Job上引用 jobData
GLOBL ·Job(SB),$16
DATA  ·Job+0(SB)/8,$·jobData(SB)
DATA  ·Job+8(SB)/8,$6
*/
GLOBL ·Job(SB),$32

DATA ·Job+0(SB)/8,$·Job+16(SB)
DATA ·Job+8(SB)/8,$16
DATA ·Job+16(SB)/16,$"gopher and phper"


GLOBL ·intMov2(SB), $8
DATA ·intMov2+0(SB)/2, $1<<2

// 布尔bool
GLOBL ·trueBool(SB), $1
DATA  ·trueBool(SB)/1, $3	// ==true

GLOBL ·falseBool(SB), $1
DATA  ·falseBool(SB)/1, $0	// ==false

// 模拟slice
GLOBL ·helloSlice(SB), $48
DATA  ·helloSlice+0(SB)/8, $·helloSlice+24(SB)  //指针 
DATA  ·helloSlice+8(SB)/8, $4	// 内容长度
DATA  ·helloSlice+16(SB)/8, $6	// 容器大小
DATA  ·helloSlice+24(SB)/4, $3  // 第一个元素是int32 value=3
DATA  ·helloSlice+28(SB)/4, $4  // 第一个元素是int32 value=4
DATA  ·helloSlice+32(SB)/4, $64	// 第一个元素是int32 value=64
DATA  ·helloSlice+36(SB)/4, $21	// 第一个元素是int32 value=64
//DATA  ·helloSlice+40(SB)/4, $0	// 第一个元素是int32 value=0
//DATA  ·helloSlice+44(SB)/4, $0	// 第一个元素是int32 value=0

// 模拟 a, b = b, a 交换
TEXT ·Swap+0(SB), $0-32
	// AX = arg0,第一个参数存储AX寄存器
	MOVQ	arg0+0(FP),  AX	
	MOVQ	arg1+8(FP),  BX	// BX = arg2
	MOVQ	BX,	ret0+16(FP) // ret0=BX
	MOVQ	AX, ret1+24(FP) // ret1=AX
	RET

