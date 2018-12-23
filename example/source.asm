start	LDA	nil

	SLP	3000

	ADD	one
	ADD	one
	ADD	two
	SUB	one
	SUB	nil
	ADD	five

	JMP	skip
	SUB	five
	SUB	one
	STOP	3
skip


	LDA	two
	JNE	end

	ADD	one
	STOP	1

end	STOP	2

nil	DEFW	0
one	DEFW	1
two	DEFW	2

	DEFW	5
five	EQU	11
