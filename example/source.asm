start	LDA	nil	; 0

	SWAP	oneswap	; 1

	ADD	two	; 3
	ADD	one	; 4

	SUB	nil	; 4
	ADD	nil	; 4

	JMP	skip	; -
	ADD	one	; -
skip			; 4

	SUB	one	; 3

	JML	addtwo	; 5

end	STOP	3	; 5


nil	DEFW	&0
one	DEFW	0x1
two	DEFW	0b10
oneswap	DEFW	1


addtwo
	ADD	one	; +1
	ADD	two	; +2
	SUB	one	; -1

	RET	0
