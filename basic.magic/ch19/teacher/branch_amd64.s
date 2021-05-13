#include "textflag.h"

TEXT ·ASMIf+0(SB), $0-16
	MOVQ	ok+0(FP),	AX	// AX=ok
	CMPQ	AX,			$0	// AX == 0 比较
	JZ		COND			// ok == 0 跳转COND	
	MOVQ	$2,		r+8(FP)
	RET
// ok ==0 return 1
COND:
	MOVQ	$1,		r+8(FP)
	RET



TEXT ·ASMFor+0(SB), NOSPLIT, $0-32
LOOP_BEFORE:
	MOVQ	startLoop+0(FP),	AX	// AX = startLoop
	MOVQ	endLoop+8(FP),		BX	// BX = endLoop
	MOVQ	step+16(FP),		CX	// CX = step
	MOVQ	$0,					DX	// ret =0

LOOP_IF:
	CMPQ	AX,					BX	// i < endLoop
	JL		LOOP_BODY				// i < endLoop; goto LOOP_BODY
	JMP		LOOP_END

LOOP_BODY:
	ADDQ	AX,					DX	// ret += i
	ADDQ	CX,					AX	// i += step
	JMP		LOOP_IF

LOOP_END:
	MOVQ	DX,			ret + 24(FP)// push ret to stack
	RET

