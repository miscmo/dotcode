TEXT runtime.main.func2(SB) /usr/local/go/src/runtime/proc.go
  proc.go:208		0x1031e60		493b6610		CMPQ 0x10(R14), SP		
  proc.go:208		0x1031e64		7628			JBE 0x1031e8e			
  proc.go:208		0x1031e66		4883ec08		SUBQ $0x8, SP			
  proc.go:208		0x1031e6a		48892c24		MOVQ BP, 0(SP)			
  proc.go:208		0x1031e6e		488d2c24		LEAQ 0(SP), BP			
  proc.go:208		0x1031e72		488b4208		MOVQ 0x8(DX), AX		
  proc.go:209		0x1031e76		803800			CMPB $0x0, 0(AX)		
  proc.go:209		0x1031e79		740a			JE 0x1031e85			
  proc.go:210		0x1031e7b		0f1f440000		NOPL 0(AX)(AX*1)		
  proc.go:210		0x1031e80		e8bb900000		CALL runtime.unlockOSThread(SB)	
  proc.go:212		0x1031e85		488b2c24		MOVQ 0(SP), BP			
  proc.go:212		0x1031e89		4883c408		ADDQ $0x8, SP			
  proc.go:212		0x1031e8d		c3			RET				
  proc.go:208		0x1031e8e		e88d720200		CALL runtime.morestack.abi0(SB)	
  proc.go:208		0x1031e93		ebcb			JMP runtime.main.func2(SB)	
  :-1			0x1031e95		cc			INT $0x3			
  :-1			0x1031e96		cc			INT $0x3			
  :-1			0x1031e97		cc			INT $0x3			
  :-1			0x1031e98		cc			INT $0x3			
  :-1			0x1031e99		cc			INT $0x3			
  :-1			0x1031e9a		cc			INT $0x3			
  :-1			0x1031e9b		cc			INT $0x3			
  :-1			0x1031e9c		cc			INT $0x3			
  :-1			0x1031e9d		cc			INT $0x3			
  :-1			0x1031e9e		cc			INT $0x3			
  :-1			0x1031e9f		cc			INT $0x3			

TEXT runtime.main.func1(SB) /usr/local/go/src/runtime/proc.go
  proc.go:174		0x10564a0		493b6610		CMPQ 0x10(R14), SP			
  proc.go:174		0x10564a4		762d			JBE 0x10564d3				
  proc.go:174		0x10564a6		4883ec20		SUBQ $0x20, SP				
  proc.go:174		0x10564aa		48896c2418		MOVQ BP, 0x18(SP)			
  proc.go:174		0x10564af		488d6c2418		LEAQ 0x18(SP), BP			
  proc.go:175		0x10564b4		488d05fd330500		LEAQ go.func.*+1167(SB), AX		
  proc.go:175		0x10564bb		31db			XORL BX, BX				
  proc.go:175		0x10564bd		48c7c1ffffffff		MOVQ $-0x1, CX				
  proc.go:175		0x10564c4		e817f4fdff		CALL runtime.newm(SB)			
  proc.go:176		0x10564c9		488b6c2418		MOVQ 0x18(SP), BP			
  proc.go:176		0x10564ce		4883c420		ADDQ $0x20, SP				
  proc.go:176		0x10564d2		c3			RET					
  proc.go:174		0x10564d3		e8e82c0000		CALL runtime.morestack_noctxt.abi0(SB)	
  proc.go:174		0x10564d8		ebc6			JMP runtime.main.func1(SB)		
  :-1			0x10564da		cc			INT $0x3				
  :-1			0x10564db		cc			INT $0x3				
  :-1			0x10564dc		cc			INT $0x3				
  :-1			0x10564dd		cc			INT $0x3				
  :-1			0x10564de		cc			INT $0x3				
  :-1			0x10564df		cc			INT $0x3				

TEXT main.main(SB) /Users/haigangliu/Documents/github/dotcode/go/defer/main.go
  main.go:18		0x108a140		493b6610		CMPQ 0x10(R14), SP			
  main.go:18		0x108a144		0f86fc000000		JBE 0x108a246				
  main.go:18		0x108a14a		4883c480		ADDQ $-0x80, SP				
  main.go:18		0x108a14e		48896c2478		MOVQ BP, 0x78(SP)			
  main.go:18		0x108a153		488d6c2478		LEAQ 0x78(SP), BP			
  main.go:30		0x108a158		48c744242001000000	MOVQ $0x1, 0x20(SP)			
  main.go:30		0x108a161		48c744241802000000	MOVQ $0x2, 0x18(SP)			
  main.go:31		0x108a16a		488b442420		MOVQ 0x20(SP), AX			
  main.go:31		0x108a16f		bb02000000		MOVL $0x2, BX				
  main.go:31		0x108a174		e8e7000000		CALL main.swap1(SB)			
  main.go:32		0x108a179		488d4c2458		LEAQ 0x58(SP), CX			
  main.go:32		0x108a17e		440f1139		MOVUPS X15, 0(CX)			
  main.go:32		0x108a182		488d4c2468		LEAQ 0x68(SP), CX			
  main.go:32		0x108a187		440f1139		MOVUPS X15, 0(CX)			
  main.go:32		0x108a18b		488d4c2458		LEAQ 0x58(SP), CX			
  main.go:32		0x108a190		48894c2438		MOVQ CX, 0x38(SP)			
  main.go:32		0x108a195		488b442420		MOVQ 0x20(SP), AX			
  main.go:32		0x108a19a		e821f0f7ff		CALL runtime.convT64(SB)		
  main.go:32		0x108a19f		4889442430		MOVQ AX, 0x30(SP)			
  main.go:32		0x108a1a4		488b4c2438		MOVQ 0x38(SP), CX			
  main.go:32		0x108a1a9		8401			TESTB AL, 0(CX)				
  main.go:32		0x108a1ab		488d152e730000		LEAQ runtime.rodata+28960(SB), DX	
  main.go:32		0x108a1b2		488911			MOVQ DX, 0(CX)				
  main.go:32		0x108a1b5		488d7908		LEAQ 0x8(CX), DI			
  main.go:32		0x108a1b9		833d40fd0d0000		CMPL $0x0, runtime.writeBarrier(SB)	
  main.go:32		0x108a1c0		7402			JE 0x108a1c4				
  main.go:32		0x108a1c2		eb06			JMP 0x108a1ca				
  main.go:32		0x108a1c4		48894108		MOVQ AX, 0x8(CX)			
  main.go:32		0x108a1c8		eb07			JMP 0x108a1d1				
  main.go:32		0x108a1ca		e8d10ffdff		CALL runtime.gcWriteBarrier(SB)		
  main.go:32		0x108a1cf		eb00			JMP 0x108a1d1				
  main.go:32		0x108a1d1		488b442418		MOVQ 0x18(SP), AX			
  main.go:32		0x108a1d6		e8e5eff7ff		CALL runtime.convT64(SB)		
  main.go:32		0x108a1db		4889442428		MOVQ AX, 0x28(SP)			
  main.go:32		0x108a1e0		488b4c2438		MOVQ 0x38(SP), CX			
  main.go:32		0x108a1e5		8401			TESTB AL, 0(CX)				
  main.go:32		0x108a1e7		488d15f2720000		LEAQ runtime.rodata+28960(SB), DX	
  main.go:32		0x108a1ee		48895110		MOVQ DX, 0x10(CX)			
  main.go:32		0x108a1f2		488d7918		LEAQ 0x18(CX), DI			
  main.go:32		0x108a1f6		833d03fd0d0000		CMPL $0x0, runtime.writeBarrier(SB)	
  main.go:32		0x108a1fd		7403			JE 0x108a202				
  main.go:32		0x108a1ff		90			NOPL					
  main.go:32		0x108a200		eb06			JMP 0x108a208				
  main.go:32		0x108a202		48894118		MOVQ AX, 0x18(CX)			
  main.go:32		0x108a206		eb07			JMP 0x108a20f				
  main.go:32		0x108a208		e8930ffdff		CALL runtime.gcWriteBarrier(SB)		
  main.go:32		0x108a20d		eb00			JMP 0x108a20f				
  main.go:32		0x108a20f		488b442438		MOVQ 0x38(SP), AX			
  main.go:32		0x108a214		8400			TESTB AL, 0(AX)				
  main.go:32		0x108a216		eb00			JMP 0x108a218				
  main.go:32		0x108a218		4889442440		MOVQ AX, 0x40(SP)			
  main.go:32		0x108a21d		48c744244802000000	MOVQ $0x2, 0x48(SP)			
  main.go:32		0x108a226		48c744245002000000	MOVQ $0x2, 0x50(SP)			
  main.go:32		0x108a22f		bb02000000		MOVL $0x2, BX				
  main.go:32		0x108a234		4889d9			MOVQ BX, CX				
  main.go:32		0x108a237		e824acffff		CALL fmt.Println(SB)			
  main.go:33		0x108a23c		488b6c2478		MOVQ 0x78(SP), BP			
  main.go:33		0x108a241		4883ec80		SUBQ $-0x80, SP				
  main.go:33		0x108a245		c3			RET					
  main.go:18		0x108a246		e875effcff		CALL runtime.morestack_noctxt.abi0(SB)	
  main.go:18		0x108a24b		e9f0feffff		JMP main.main(SB)			
  :-1			0x108a250		cc			INT $0x3				
  :-1			0x108a251		cc			INT $0x3				
  :-1			0x108a252		cc			INT $0x3				
  :-1			0x108a253		cc			INT $0x3				
  :-1			0x108a254		cc			INT $0x3				
  :-1			0x108a255		cc			INT $0x3				
  :-1			0x108a256		cc			INT $0x3				
  :-1			0x108a257		cc			INT $0x3				
  :-1			0x108a258		cc			INT $0x3				
  :-1			0x108a259		cc			INT $0x3				
  :-1			0x108a25a		cc			INT $0x3				
  :-1			0x108a25b		cc			INT $0x3				
  :-1			0x108a25c		cc			INT $0x3				
  :-1			0x108a25d		cc			INT $0x3				
  :-1			0x108a25e		cc			INT $0x3				
  :-1			0x108a25f		cc			INT $0x3				

TEXT main.swap1(SB) /Users/haigangliu/Documents/github/dotcode/go/defer/main.go
  main.go:35		0x108a260		4883ec10		SUBQ $0x10, SP		
  main.go:35		0x108a264		48896c2408		MOVQ BP, 0x8(SP)	
  main.go:35		0x108a269		488d6c2408		LEAQ 0x8(SP), BP	
  main.go:35		0x108a26e		4889442418		MOVQ AX, 0x18(SP)	
  main.go:35		0x108a273		48895c2420		MOVQ BX, 0x20(SP)	
  main.go:36		0x108a278		48890424		MOVQ AX, 0(SP)		
  main.go:36		0x108a27c		488b442420		MOVQ 0x20(SP), AX	
  main.go:36		0x108a281		4889442418		MOVQ AX, 0x18(SP)	
  main.go:36		0x108a286		488b0424		MOVQ 0(SP), AX		
  main.go:36		0x108a28a		4889442420		MOVQ AX, 0x20(SP)	
  main.go:37		0x108a28f		488b6c2408		MOVQ 0x8(SP), BP	
  main.go:37		0x108a294		4883c410		ADDQ $0x10, SP		
  main.go:37		0x108a298		c3			RET			
