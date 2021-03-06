//go:build gc
// +build gc

#include "textflag.h"

// func decl_array() (uint64, uint64)
TEXT ·decl_array(SB),NOSPLIT,$16-16
	MOVD	$0x3, R0
	MOVD	$0x2, R1

	// move the immediate $0x3, and $0x2 to frame stack, which is declaring
	// local variable in function. The way we conduting here is we declaring
	// an array which named as `arr` with 2 element.
	MOVD	R0, arr-8(SP)
	MOVD	R1, arr-16(SP)

	// Take the addr from the lowest addr of the array
	MOVD	$arr-16(SP), R3
	VLD1	(R3), [V0.D2]
	VSHL	$1, V0.D2, V0.D2

	// move the first element of the vector to register
	VMOV	V0.D[0], R4
	VMOV	V0.D[1], R5
	MOVD	R4, ret+0(FP)
	MOVD	R5, ret+8(FP)

	RET
