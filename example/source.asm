start	LDA	nil	; 0

	ADD	two	; 2
	ADD	one	; 3

	SUB	nil	; 3
	ADD	nil	; 3

	JMP	skip	; -
	ADD	one	; -
skip			; 3

	SUB	one	; 2

	JML	addtwo	; 4

end	STOP	3	; 4


nil	DEFW	&0
one	DEFW	0x1
two	DEFW	0b10


addtwo
	ADD	one	; +1
	ADD	two	; +2
	SUB	one	; -1

	RET	0
