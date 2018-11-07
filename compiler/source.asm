;// 	|-------------------------------------|
;// 	|  6 DISPLAY RUNNER - A MU0 GAME      |
;// 	|  by Alexandru-Paul Copil, mbaxaac3  |
;// 	|  (thee-engineer)                    |
;// 	|                                     |
;// 	|  LICENSE: MIT                       |
;// 	|                                     |
;// 	|  STARTED ON : 28/11/2016            |
;// 	|  LAST EDIT  : 07/12/2016            |
;// 	|                                     |
;// 	|-------------------------------------|

;//		|-------------------------------------|
;//		| SOURCE CODE BEGINS BELOW            |
;// 	|-------------------------------------|

;// 	|-------------------------------------|
;// 	| PROGRAM RUNTIME BEGINS BELOW        |
;// 	|-------------------------------------|

init	ORG	0000		;// RESET MEMORY ADRESS

			JMP	mrst		;// RESET BOARD COMPONENTS

runt							;// START PROGRAM RUNTIME

;// 	|-------------------------------------|
;// 	| WAIT FOR START INPUT (C)		        |
;// 	|-------------------------------------|

menu							;// LOOP FOR USER START (DIF 1,2,X)

			LDA	kr3			;// CHECK KEYROW 3
			SUB	df1			;// CHECK DIF 1
			JNE	me1			;// NOT
			LDA kr3			;// STORE IT
			STA dff
			JGE	load		;// LOAD

me1		LDA	kr3			;// CHECK KEYROW 3
			SUB df2			;// CHECK DIF 2
			JNE me2			;// NOT
			LDA kr3			;// STORE IT
			STA dff
			JGE load		;// LOAD

me2		LDA	kr3			;// CHECK KEYROW 3
			SUB dfx			;// CHECK DIF X
			JNE me3			;// NOT
			LDA kr3			;// STORE IT
			STA dff
			JGE load		;// LOAD

me3		LDA	kr3			;// CHECK KEYROW 3
			SUB df3			;// CHECK DIF 3
			JNE menu		;// NOT, LOOP AGAIN
			LDA kr3			;// STORE IT
			STA dff
			JGE load		;// LOAD

			JMP menu

;// 	|-------------------------------------|
;// 	| PASS LOADING TIME						        |
;// 	|-------------------------------------|

load							;// START WAITING TIME

			LDA	dlc			;// LOAD GLOBAL COUNT
			STA	tmp			;// STORE LOCAL COUNT
ldc
			LDA	dly			;// LOAD DELAY TIME
ldl0	SUB	one			;// COUNT DOWN
			JNE	ldl0		;// LOOP
			LDA	tmp			;// LOAD DELAY COUNT
			SUB	one			;// COUNT DOWN
			STA	tmp			;// STORE COUNT

;//		| LOADING BAR													|

			LDA	tbg			;// LOAD TEMP BAR GRAPH
			STA	dbg			;// STORE IT TO BAR GRAPH
			ADD tbg			;// COUNT UP
			STA tbg			;// STORE TO TEMP BAR GRAPH

			JNE	ldc			;// DELAY MORE

;//		| INIT SETUP 													|

			LDA ph3			;// LOAD HP
			STA php			;// STORE HP
			STA dbg			;// STORE HP

			LDA mid			;// LOAD PP
			STA dp5			;// SET PLAYER

;// 	|-------------------------------------|
;// 	| SKIP, LOADING TIME					        |
;// 	|-------------------------------------|

			JMP input		;// SKIP SKIP, TAKE INPUT

skip							;// IF NO INPUT, DELAY, SKIP

			LDA	dlp			;// LOAD GLOBAL COUNT
			STA	tmp			;// STORE LOCAL COUNT
ldp
			LDA	dly			;// LOAD DELAY TIME
ldl1	SUB	one			;// COUNT DOWN
			JNE	ldl1		;// LOOP
			LDA	tmp			;// LOAD DELAY COUNT
			SUB	one			;// COUNT DOWN
			STA	tmp			;// STORE COUNT
			JNE	ldp			;// DELAY MORE

			JMP	shift		;// SHIFT DISPLAY TO THE LEFT
s5								;// SHIFT IS DONE, CONTINUE

;// 	|-------------------------------------|
;// 	| SCAN FOR USER INPUT					        |
;// 	|-------------------------------------|

input							;// LOOP FOR USER INPUT

;//		| CHECK FOR STOP SIGNAL               |

			LDA	kr1			;// CHECK KEYROW 1
			SUB	kst			;// CHECK FOR AC
			JNE	s0			;// CHECK NON ZERO
			JGE	halt		;// CALL HALT
s0

;//		| CHECK FOR RESET SIGNAL              |

			LDA	kr4			;// CHECK KEYROW 4
			SUB	krt			;// CHECK RESET
			JNE	s1			;// CHECK NOT ZERO
			JGE	mrst		;// RESET BOARD
s1

;//		| COLLISION CHECK                			|

			LDA	dp5			;// CHECK PLAYER
			JNE s9			;// OK PLAYER
			JMP s6			;// PLAYER IS NULL
s9
			SUB	dp4			;// CHECK NEXT
			JNE	s6			;// DODGE
			JGE	mhit		;// GET HIT
s6

;//		| PLAYER MOVE - BOT                   |

			LDA		ksw		;// CHECK SWITCHES
			SUB		bmv		;// CHECK SWITCH 2
			JNE		s3		;// CHECK NON ZERO
			JGE		mbot	;// MOVE bot
s3

;//		| PLAYER MOVE - TOP                   |

			LDA	ksw			;// CHECK SWITCHES
			SUB	tmv			;// CHECK FOR SWITCH 1
			JNE	s4			;// CHECK NON ZERO
			JGE	mtop		;// MOVE top
s4

;//		| SPAWN TEST MID											|

			LDA kr4			;// CHECK KEYROW 4
			SUB ksf			;// CHECK FOR SHIFT
			JNE s7			;// NOT SHIFT
			LDA mid			;// LOAD MID SEGMENT
			STA dp0			;// STORE MID SEGMENT
			;JGE input	;// CONTINUE
s7

;//		| NO INPUT, SKIP AND DELAY						|

			LDA sbm			;// LOAD BACKGROUND MUSIC
			STA bzz			;// PLAY

			JMP skip

;// 	|-------------------------------------|
;// 	| HALT												        |
;// 	|-------------------------------------|

halt
			LDA	nul			;// SET ACC TO ONE
			STA hlt			;// SET STOP SIGNAL TO ONE

			JMP mrst		;// RESET THEN STOP

;// 	|-------------------------------------|
;// 	| PROGRAM RUNTIME STOPS HERE          |
;// 	|-------------------------------------|

;// 	|-------------------------------------|
;// 	| PROGRAM METHODS                     |
;//		|-------------------------------------|

STP
mrst							;// RESET/STOP THE PROGRAM

			LDA nul			;// LOAD NULL

			STA	dp0			;// RESET DISPLAY 0
			STA	dp1			;// RESET DISPLAY 1
			STA	dp2			;// RESET DISPLAY 2
			STA	dp3			;// RESET DISPLAY 3
			STA	dp4			;// RESET DISPLAY 4
			STA	dp5			;// RESET DISPLAY 5

			LDA fff			;// LOAD FULL BG
			STA	dbg			;// RESET BAR GRAPH

			LDA	spr			;// LOAD POS SOUND
			STA	bzz			;// STORE IN BUZZER

			LDA	hlt			;//	LOAD STOP SIGNAL
			JNE	runt		;//	START THE PROGRAM

			LDA	sht			;// LOAD HALT SOUND
			STA	bzz			;// STORE IN BUZZER

			LDA nul			;// LOAD NULL
			STA	dbg			;// RESET BAR GRAPH

			STP					;// STOP THE PROGRAM

STP
mlal							;// SET ALL DISPLAYS TO ACC
			STA	dp0			;// SET DISPLAY 0
			STA	dp1			;// SET DISPLAY 1
			STA dp2			;// SET DISPLAY 2
			STA	dp3			;// SET DISPLAY 3
			STA	dp4			;// SET DISPLAY 4
			STA	dp5			;// SET DISPLAY 5
			STA	dbg			;// SET BAR GRAPH

			JMP skip		;// CONTINUE

STP
mtop							;// MOVE PLAYER TO TOP

			LDA sdu			;// LOAD DIFF SOUND
			STA	bzz			;// PLAY DIFF SOUND

			LDA dp5			;// CHECK PLAYER POS
			SUB bot			;// CHECK BOT
			JNE mt1			;// NOT   BOT
			LDA mid			;// GO TO MID
			STA dp5
			JMP skip		;// CONTINUE

mt1		LDA dp5			;// CHECK PLAYER POS
			SUB mid			;// CHECK MID
			JNE skip		;// NOT   MID
			LDA top			;// GO TO TOP
			STA dp5
			JMP skip		;// CONTINUE

STP
mbot							;// MOVE PLAYER TO BOTTOM

			LDA smd			;// LOAD DIFF SOUND
			STA	bzz			;// PLAY DIFF SOUND

			LDA dp5			;// CHECK PLAYER POS
			SUB top			;// CHECK TOP
			JNE mt2			;// NOT   TOP
			LDA mid			;// GO TO MID
			STA dp5
			JMP skip		;// CONTINUE

mt2		LDA dp5			;// CHECK PLAYER POS
			SUB mid			;// CHECK MID
			JNE skip		;// NOT   MID
			LDA bot			;// GO TO BOT
			STA dp5
			JMP skip		;// CONTINUE

STP
shift							;// MOVE dp3-0 TO THE LEFT

			LDA	dp3			;// MOVE dp3 TO THE LEFT
			STA	dp4
			LDA	dp2			;// MOVE dp2 TO THE LEFT
			STA	dp3
			LDA	dp1			;// MOVE dp1 TO THE LEFT
			STA	dp2
			LDA	dp0			;// MOVE dp0 TO THE LEFT
			STA	dp1

			LDA nul			;// EMPTY FIRST
			STA dp0

			JMP mcemp		;// SPAWN NEXT
nem		JMP	s5			;// CONTINUE

STP
mcemp							;// CHECK FOR EMPTY SCREEN


			LDA dff			;// CHECK DIFFICULTY X
			SUB dfx			;// CHECK IF X
			JNE ndx			;// NOT X
			JMP sequence;// DIF X

ndx		LDA dff			;// CHECK DIFFICULTY 3
			SUB df3
			JNE nd3

			LDA dp0			;// CHECK DISPLAY
			JNE nem
			LDA dp1
			JNE nem
			LDA dp3
			JNE nem

			JMP sequence

nd3		LDA dff			;// CHECK DIFFICULTY 2
			SUB df2
			JNE nd2

			LDA dp0			;// CHECK DISPLAY
			JNE nem
			LDA dp1
			JNE nem
			LDA dp2
			JNE nem

			JMP sequence

nd2		LDA dff			;// CHECK DIFFICULTY 1
			SUB df1
			JNE sequence

			LDA dp0
			JNE nem
			LDA dp1
			JNE nem
			LDA dp2
			JNE nem
			LDA dp3
			JNE nem
			LDA dp4
			JNE nem

			JMP	sequence	;// NEXT IN SEQUENCE

STP
mhit								;// PLAYER GOT HIT

			LDA	sgo
			STA	bzz

			LDA ph3
			SUB php
			JNE	alv1
			LDA ph2
			STA php
			STA dbg
			JMP skip

alv1	LDA ph2
			SUB php
			JNE mrst
			LDA ph1
			STA php
			STA dbg
			JMP skip

			JGE	mrst		;// GAME OVER

;// 	|-------------------------------------|
;// 	| PROGRAM MEMORY ALOCATION            |
;//		|-------------------------------------|

;//		| DISPLAY POSITIONS                   |

top		DEFW	&0001	;//	DISPLAY: 0000_0001
mid		DEFW	&0002	;//	DISPLAY: 0100_0000
bot		DEFW	&0003	;// DISPLAY: 0000_1000

fff		DEFW	&FF		;// BAR GRAPH FULL

php		DEFW	0b00000111	;// PLAYER HEALTH
ph3		DEFW	0b00000111	;// FULL HP
ph2		DEFW	0b00000011	;// HP -1
ph1		DEFW	0b00000001	;// HP -2

;//		| DECIMALS														|

nul		DEFW	&0000	;// CONSTANT ZERO VALUE
one		DEFW	&0001	;// CONSTANT ONE  VALUE
six		DEFW	&0006	;//	CONSTANT FOUR VALUE

;//		| SIGNALS															|

hlt		DEFW	&0001	;// PROGRAM HALT SIGNAL

;//		| KEYROWS															|

krt		DEFW	&0080	;// KEYROW 4, RESET
kst		DEFW	&0002	;// KEYROW 1, AC
ksa		DEFW	&0004	;// KEYROW 1, C
ksf		DEFW	&0040	;// KEYROW 4, SHIFT

df1		DEFW	&0020	;// DIFFICULTY 1
df2		DEFW	&0010	;// DIFFICULTY 2
df3		DEFW	&0008	;// DIFFICULTY 3
dfx		DEFW	&0004 ;// DIFFICULTY X

;//		|	SWITCHES														|

tmv		DEFW	&0001	;// SWITCH 1
bmv		DEFW	&0002	;// SWITCH 2
mmv		DEFW	&0003	;// SWITCH 1 & SWITCH 2

;//		| DELAYS															|

dly		DEFW	50001	;// INNER DELAY TIME
dlc		DEFW	00007	;// WAIT FOR N SECONDS
dlp		DEFW	00002	;// WAIT FOR INPUT

;//		|	TEMPORARY STORAGE	VARIABLES					|

tbg		DEFW	&0001	;// TEMPORARY BAR GRAPH
tmp		DEFW	&0000	;// TEMPROARY VARIABLE

;//		|	PROGRAM COUNTERS										|

sqc		DEFW	&0000	;// SEQUENCE COUNTER
sma		DEFW	&0006	;// SEQUENCE SIZE
dff		DEFW	&0000	;// DIFFICULTY

;//		|	OP CODES														|

jop		DEFW	&4000	;// JMP INSTRUCTION

;//		|	AUDIO AND BUZZER 										|

spr		DEFW	0b1000010001001000	;// RESET
sht		DEFW	0b1000010000010001	;// HALT

sdu		DEFW	0b1000001001010110	;// UP
smd		DEFW	0b1000001000110110	;// DOWN

sgo		DEFW	0b1000001000010000	;// HIT

sbm		DEFW	0b1000001000100011	;// BACKGROUND
;//						m___ddddoooonnnn
;//
;//						m - mode
;//						d - duration
;//						o - octave
;//						n - note

;// 	|-------------------------------------|
;// 	| COMPILER DEFINED CONSTANTS          |
;//		|-------------------------------------|

;//		|	OUTPUTS															|

dp0		EQU		&FF5	;// CONSTANT DISPLAY 0
dp1		EQU		&FF6	;// CONSTANT DISPLAY 1
dp2		EQU		&FF7	;// CONSTANT DISPLAY 2
dp3		EQU		&FF8	;// CONSTANT DISPLAY 3
dp4		EQU		&FF9	;// CONSTANT DISPLAY 4
dp5		EQU		&FFA	;// CONSTANT DISPLAY 5

dbg		EQU		&FFE	;// CONSTANT BAR GRAPH

bzz		EQU		&FFD	;// BUZZER INPUT SOUND

;//		| INPUTS															|

kr1		EQU		&FEF	;// KEY ROW 1
kr3		EQU		&FF1	;// KEY ROW 3
kr4		EQU		&FF2	;// KEY ROW 4
ksw		EQU		&FEE	;// SWITCHES

bzb		EQU		&FF3	;// BUZZER BUSY

STP								;// SAFTEY STOP

;//		|-------------------------------------|
;//		| SOURCE CODE STOPS HERE              |
;// 	|-------------------------------------|

;//		|-------------------------------------|
;//		| SEQUENCE CODE BELOW                 |
;// 	|-------------------------------------|

nsq								;// SELECT NEXT SEQUENCE
			LDA 	sqc		;// LOAD COUNT
			ADD		one		;// INCREMENT
			STA 	sqc		;// STORE COUNT
			SUB		sma		;// SUB SEQUENCE MAx
			JNE		nem		;// CHECK FOR MAX
			LDA		nul		;// LOAD ONE
			STA		sqc		;// RESET COUNT
			JMP		nem		;// CONTINUE

;#pySEQ						;// PYTHON GENERATED SEQUENCE

car1
			LDA 	bot
			STA 	dp0
			JMP		nsq

car2
			LDA 	mid
			STA 	dp0
			JMP		nsq

car3
			LDA 	top
			STA 	dp0
			JMP		nsq

car4
			LDA 	mid
			STA 	dp0
			JMP		nsq

car5
			LDA 	bot
			STA 	dp0
			JMP		nsq

car6
			LDA 	top
			STA 	dp0
			JMP		nsq

;#pyEND						;// PYTHON GEN SEQ END

sequence					;// START SEQUENCE

spc		LDA 	spc		;// LOAD SEQ PC
			ADD 	jop		;//	ADD JUMP OPERATION
			ADD		sqc		;// POINT TO SEQUENCE
			ADD		six		;// SKIP INSTRUCTIONS
			STA 	jsq		;// STORE IT

jsq		DEFW	jsq		;// DEFINE JUMP POSITION

;#pyLNK						;// PYTHON LINK SEQUENCE

			JMP 	car1
			JMP 	car2
			JMP 	car3
			JMP 	car4
			JMP		car5
			JMP		car6

;#pyEND						;// PYTHON LINK SEQ END

			JMP		nem		;// GO BACK

STP								;// SAFTEY STOP

;//		|-------------------------------------|
;//		| SEQUENCE CODE ABOVE                 |
;// 	|-------------------------------------|