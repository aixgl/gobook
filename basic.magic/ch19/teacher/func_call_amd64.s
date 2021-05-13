
TEXT ·ASMAddSub+0(SB), $0-32
	//MOVQ	a+8(SP),	AX		// AX = 参数1
	MOVQ	8(SP),	AX		// AX = 参数1
	//MOVQ	b+16(SP),	BX		// BX = 参数2
	MOVQ	16(SP),	BX		// BX = 参数2
	MOVQ	AX,		c+24(SP)	// c = AX
	ADDQ	BX,		c+24(SP)	// c = AX + BX
	MOVQ	AX,		d+32(SP)		// d = BX
	SUBQ	BX,		d+32(SP)		// d = AX - BX
	RET

TEXT ·ASMCaller+0(SB), $32-16
	MOVQ	$90,	0(SP)
	MOVQ	$70,	8(SP)
	CALL	·ASMAddSub(SB)
	//MOVQ	c-16(SP),	AX
	MOVQ	16(SP),	AX
	//MOVQ	d-8(SP),	BX
	MOVQ	24(SP),	BX
	MOVQ	AX,	fc+0(FP)	
	MOVQ	BX,	fd+8(FP)	
	//CALL	runtime·printnl(SB)
	RET
