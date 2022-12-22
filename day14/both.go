package main

import (
	"fmt"
	"strconv"
	"strings"
)

const example = `498,4 -> 498,6 -> 496,6
503,4 -> 502,4 -> 502,9 -> 494,9
`

const input = `461,87 -> 461,80 -> 461,87 -> 463,87 -> 463,80 -> 463,87 -> 465,87 -> 465,80 -> 465,87 -> 467,87 -> 467,86 -> 467,87
468,137 -> 468,132 -> 468,137 -> 470,137 -> 470,130 -> 470,137 -> 472,137 -> 472,136 -> 472,137 -> 474,137 -> 474,127 -> 474,137 -> 476,137 -> 476,131 -> 476,137 -> 478,137 -> 478,130 -> 478,137 -> 480,137 -> 480,128 -> 480,137 -> 482,137 -> 482,128 -> 482,137 -> 484,137 -> 484,133 -> 484,137
460,159 -> 460,154 -> 460,159 -> 462,159 -> 462,151 -> 462,159 -> 464,159 -> 464,152 -> 464,159
464,109 -> 464,111 -> 463,111 -> 463,118 -> 473,118 -> 473,111 -> 468,111 -> 468,109
468,137 -> 468,132 -> 468,137 -> 470,137 -> 470,130 -> 470,137 -> 472,137 -> 472,136 -> 472,137 -> 474,137 -> 474,127 -> 474,137 -> 476,137 -> 476,131 -> 476,137 -> 478,137 -> 478,130 -> 478,137 -> 480,137 -> 480,128 -> 480,137 -> 482,137 -> 482,128 -> 482,137 -> 484,137 -> 484,133 -> 484,137
478,33 -> 478,37 -> 471,37 -> 471,41 -> 483,41 -> 483,37 -> 482,37 -> 482,33
466,74 -> 466,71 -> 466,74 -> 468,74 -> 468,64 -> 468,74 -> 470,74 -> 470,65 -> 470,74 -> 472,74 -> 472,68 -> 472,74 -> 474,74 -> 474,68 -> 474,74 -> 476,74 -> 476,71 -> 476,74 -> 478,74 -> 478,65 -> 478,74 -> 480,74 -> 480,73 -> 480,74 -> 482,74 -> 482,64 -> 482,74
466,96 -> 470,96
494,22 -> 499,22
499,13 -> 499,16 -> 498,16 -> 498,19 -> 507,19 -> 507,16 -> 501,16 -> 501,13
466,74 -> 466,71 -> 466,74 -> 468,74 -> 468,64 -> 468,74 -> 470,74 -> 470,65 -> 470,74 -> 472,74 -> 472,68 -> 472,74 -> 474,74 -> 474,68 -> 474,74 -> 476,74 -> 476,71 -> 476,74 -> 478,74 -> 478,65 -> 478,74 -> 480,74 -> 480,73 -> 480,74 -> 482,74 -> 482,64 -> 482,74
460,145 -> 460,146 -> 462,146
468,137 -> 468,132 -> 468,137 -> 470,137 -> 470,130 -> 470,137 -> 472,137 -> 472,136 -> 472,137 -> 474,137 -> 474,127 -> 474,137 -> 476,137 -> 476,131 -> 476,137 -> 478,137 -> 478,130 -> 478,137 -> 480,137 -> 480,128 -> 480,137 -> 482,137 -> 482,128 -> 482,137 -> 484,137 -> 484,133 -> 484,137
466,74 -> 466,71 -> 466,74 -> 468,74 -> 468,64 -> 468,74 -> 470,74 -> 470,65 -> 470,74 -> 472,74 -> 472,68 -> 472,74 -> 474,74 -> 474,68 -> 474,74 -> 476,74 -> 476,71 -> 476,74 -> 478,74 -> 478,65 -> 478,74 -> 480,74 -> 480,73 -> 480,74 -> 482,74 -> 482,64 -> 482,74
468,137 -> 468,132 -> 468,137 -> 470,137 -> 470,130 -> 470,137 -> 472,137 -> 472,136 -> 472,137 -> 474,137 -> 474,127 -> 474,137 -> 476,137 -> 476,131 -> 476,137 -> 478,137 -> 478,130 -> 478,137 -> 480,137 -> 480,128 -> 480,137 -> 482,137 -> 482,128 -> 482,137 -> 484,137 -> 484,133 -> 484,137
460,159 -> 460,154 -> 460,159 -> 462,159 -> 462,151 -> 462,159 -> 464,159 -> 464,152 -> 464,159
454,168 -> 458,168
476,60 -> 476,61 -> 489,61
502,26 -> 507,26
447,106 -> 452,106
468,137 -> 468,132 -> 468,137 -> 470,137 -> 470,130 -> 470,137 -> 472,137 -> 472,136 -> 472,137 -> 474,137 -> 474,127 -> 474,137 -> 476,137 -> 476,131 -> 476,137 -> 478,137 -> 478,130 -> 478,137 -> 480,137 -> 480,128 -> 480,137 -> 482,137 -> 482,128 -> 482,137 -> 484,137 -> 484,133 -> 484,137
498,24 -> 503,24
478,33 -> 478,37 -> 471,37 -> 471,41 -> 483,41 -> 483,37 -> 482,37 -> 482,33
466,74 -> 466,71 -> 466,74 -> 468,74 -> 468,64 -> 468,74 -> 470,74 -> 470,65 -> 470,74 -> 472,74 -> 472,68 -> 472,74 -> 474,74 -> 474,68 -> 474,74 -> 476,74 -> 476,71 -> 476,74 -> 478,74 -> 478,65 -> 478,74 -> 480,74 -> 480,73 -> 480,74 -> 482,74 -> 482,64 -> 482,74
499,13 -> 499,16 -> 498,16 -> 498,19 -> 507,19 -> 507,16 -> 501,16 -> 501,13
457,166 -> 461,166
478,33 -> 478,37 -> 471,37 -> 471,41 -> 483,41 -> 483,37 -> 482,37 -> 482,33
464,109 -> 464,111 -> 463,111 -> 463,118 -> 473,118 -> 473,111 -> 468,111 -> 468,109
499,13 -> 499,16 -> 498,16 -> 498,19 -> 507,19 -> 507,16 -> 501,16 -> 501,13
466,74 -> 466,71 -> 466,74 -> 468,74 -> 468,64 -> 468,74 -> 470,74 -> 470,65 -> 470,74 -> 472,74 -> 472,68 -> 472,74 -> 474,74 -> 474,68 -> 474,74 -> 476,74 -> 476,71 -> 476,74 -> 478,74 -> 478,65 -> 478,74 -> 480,74 -> 480,73 -> 480,74 -> 482,74 -> 482,64 -> 482,74
466,74 -> 466,71 -> 466,74 -> 468,74 -> 468,64 -> 468,74 -> 470,74 -> 470,65 -> 470,74 -> 472,74 -> 472,68 -> 472,74 -> 474,74 -> 474,68 -> 474,74 -> 476,74 -> 476,71 -> 476,74 -> 478,74 -> 478,65 -> 478,74 -> 480,74 -> 480,73 -> 480,74 -> 482,74 -> 482,64 -> 482,74
478,33 -> 478,37 -> 471,37 -> 471,41 -> 483,41 -> 483,37 -> 482,37 -> 482,33
466,74 -> 466,71 -> 466,74 -> 468,74 -> 468,64 -> 468,74 -> 470,74 -> 470,65 -> 470,74 -> 472,74 -> 472,68 -> 472,74 -> 474,74 -> 474,68 -> 474,74 -> 476,74 -> 476,71 -> 476,74 -> 478,74 -> 478,65 -> 478,74 -> 480,74 -> 480,73 -> 480,74 -> 482,74 -> 482,64 -> 482,74
457,99 -> 461,99
499,13 -> 499,16 -> 498,16 -> 498,19 -> 507,19 -> 507,16 -> 501,16 -> 501,13
463,162 -> 467,162
466,74 -> 466,71 -> 466,74 -> 468,74 -> 468,64 -> 468,74 -> 470,74 -> 470,65 -> 470,74 -> 472,74 -> 472,68 -> 472,74 -> 474,74 -> 474,68 -> 474,74 -> 476,74 -> 476,71 -> 476,74 -> 478,74 -> 478,65 -> 478,74 -> 480,74 -> 480,73 -> 480,74 -> 482,74 -> 482,64 -> 482,74
457,104 -> 462,104
450,104 -> 455,104
476,60 -> 476,61 -> 489,61
468,137 -> 468,132 -> 468,137 -> 470,137 -> 470,130 -> 470,137 -> 472,137 -> 472,136 -> 472,137 -> 474,137 -> 474,127 -> 474,137 -> 476,137 -> 476,131 -> 476,137 -> 478,137 -> 478,130 -> 478,137 -> 480,137 -> 480,128 -> 480,137 -> 482,137 -> 482,128 -> 482,137 -> 484,137 -> 484,133 -> 484,137
462,143 -> 473,143 -> 473,142
483,44 -> 483,48 -> 481,48 -> 481,55 -> 489,55 -> 489,48 -> 487,48 -> 487,44
463,99 -> 467,99
461,87 -> 461,80 -> 461,87 -> 463,87 -> 463,80 -> 463,87 -> 465,87 -> 465,80 -> 465,87 -> 467,87 -> 467,86 -> 467,87
478,33 -> 478,37 -> 471,37 -> 471,41 -> 483,41 -> 483,37 -> 482,37 -> 482,33
472,168 -> 476,168
472,123 -> 472,124 -> 476,124
468,137 -> 468,132 -> 468,137 -> 470,137 -> 470,130 -> 470,137 -> 472,137 -> 472,136 -> 472,137 -> 474,137 -> 474,127 -> 474,137 -> 476,137 -> 476,131 -> 476,137 -> 478,137 -> 478,130 -> 478,137 -> 480,137 -> 480,128 -> 480,137 -> 482,137 -> 482,128 -> 482,137 -> 484,137 -> 484,133 -> 484,137
468,137 -> 468,132 -> 468,137 -> 470,137 -> 470,130 -> 470,137 -> 472,137 -> 472,136 -> 472,137 -> 474,137 -> 474,127 -> 474,137 -> 476,137 -> 476,131 -> 476,137 -> 478,137 -> 478,130 -> 478,137 -> 480,137 -> 480,128 -> 480,137 -> 482,137 -> 482,128 -> 482,137 -> 484,137 -> 484,133 -> 484,137
483,44 -> 483,48 -> 481,48 -> 481,55 -> 489,55 -> 489,48 -> 487,48 -> 487,44
489,30 -> 494,30
462,143 -> 473,143 -> 473,142
466,74 -> 466,71 -> 466,74 -> 468,74 -> 468,64 -> 468,74 -> 470,74 -> 470,65 -> 470,74 -> 472,74 -> 472,68 -> 472,74 -> 474,74 -> 474,68 -> 474,74 -> 476,74 -> 476,71 -> 476,74 -> 478,74 -> 478,65 -> 478,74 -> 480,74 -> 480,73 -> 480,74 -> 482,74 -> 482,64 -> 482,74
461,87 -> 461,80 -> 461,87 -> 463,87 -> 463,80 -> 463,87 -> 465,87 -> 465,80 -> 465,87 -> 467,87 -> 467,86 -> 467,87
466,74 -> 466,71 -> 466,74 -> 468,74 -> 468,64 -> 468,74 -> 470,74 -> 470,65 -> 470,74 -> 472,74 -> 472,68 -> 472,74 -> 474,74 -> 474,68 -> 474,74 -> 476,74 -> 476,71 -> 476,74 -> 478,74 -> 478,65 -> 478,74 -> 480,74 -> 480,73 -> 480,74 -> 482,74 -> 482,64 -> 482,74
468,137 -> 468,132 -> 468,137 -> 470,137 -> 470,130 -> 470,137 -> 472,137 -> 472,136 -> 472,137 -> 474,137 -> 474,127 -> 474,137 -> 476,137 -> 476,131 -> 476,137 -> 478,137 -> 478,130 -> 478,137 -> 480,137 -> 480,128 -> 480,137 -> 482,137 -> 482,128 -> 482,137 -> 484,137 -> 484,133 -> 484,137
482,30 -> 487,30
461,87 -> 461,80 -> 461,87 -> 463,87 -> 463,80 -> 463,87 -> 465,87 -> 465,80 -> 465,87 -> 467,87 -> 467,86 -> 467,87
468,137 -> 468,132 -> 468,137 -> 470,137 -> 470,130 -> 470,137 -> 472,137 -> 472,136 -> 472,137 -> 474,137 -> 474,127 -> 474,137 -> 476,137 -> 476,131 -> 476,137 -> 478,137 -> 478,130 -> 478,137 -> 480,137 -> 480,128 -> 480,137 -> 482,137 -> 482,128 -> 482,137 -> 484,137 -> 484,133 -> 484,137
460,145 -> 460,146 -> 462,146
460,168 -> 464,168
460,159 -> 460,154 -> 460,159 -> 462,159 -> 462,151 -> 462,159 -> 464,159 -> 464,152 -> 464,159
466,74 -> 466,71 -> 466,74 -> 468,74 -> 468,64 -> 468,74 -> 470,74 -> 470,65 -> 470,74 -> 472,74 -> 472,68 -> 472,74 -> 474,74 -> 474,68 -> 474,74 -> 476,74 -> 476,71 -> 476,74 -> 478,74 -> 478,65 -> 478,74 -> 480,74 -> 480,73 -> 480,74 -> 482,74 -> 482,64 -> 482,74
468,137 -> 468,132 -> 468,137 -> 470,137 -> 470,130 -> 470,137 -> 472,137 -> 472,136 -> 472,137 -> 474,137 -> 474,127 -> 474,137 -> 476,137 -> 476,131 -> 476,137 -> 478,137 -> 478,130 -> 478,137 -> 480,137 -> 480,128 -> 480,137 -> 482,137 -> 482,128 -> 482,137 -> 484,137 -> 484,133 -> 484,137
492,28 -> 497,28
468,137 -> 468,132 -> 468,137 -> 470,137 -> 470,130 -> 470,137 -> 472,137 -> 472,136 -> 472,137 -> 474,137 -> 474,127 -> 474,137 -> 476,137 -> 476,131 -> 476,137 -> 478,137 -> 478,130 -> 478,137 -> 480,137 -> 480,128 -> 480,137 -> 482,137 -> 482,128 -> 482,137 -> 484,137 -> 484,133 -> 484,137
464,109 -> 464,111 -> 463,111 -> 463,118 -> 473,118 -> 473,111 -> 468,111 -> 468,109
466,74 -> 466,71 -> 466,74 -> 468,74 -> 468,64 -> 468,74 -> 470,74 -> 470,65 -> 470,74 -> 472,74 -> 472,68 -> 472,74 -> 474,74 -> 474,68 -> 474,74 -> 476,74 -> 476,71 -> 476,74 -> 478,74 -> 478,65 -> 478,74 -> 480,74 -> 480,73 -> 480,74 -> 482,74 -> 482,64 -> 482,74
466,164 -> 470,164
466,74 -> 466,71 -> 466,74 -> 468,74 -> 468,64 -> 468,74 -> 470,74 -> 470,65 -> 470,74 -> 472,74 -> 472,68 -> 472,74 -> 474,74 -> 474,68 -> 474,74 -> 476,74 -> 476,71 -> 476,74 -> 478,74 -> 478,65 -> 478,74 -> 480,74 -> 480,73 -> 480,74 -> 482,74 -> 482,64 -> 482,74
464,109 -> 464,111 -> 463,111 -> 463,118 -> 473,118 -> 473,111 -> 468,111 -> 468,109
461,87 -> 461,80 -> 461,87 -> 463,87 -> 463,80 -> 463,87 -> 465,87 -> 465,80 -> 465,87 -> 467,87 -> 467,86 -> 467,87
478,33 -> 478,37 -> 471,37 -> 471,41 -> 483,41 -> 483,37 -> 482,37 -> 482,33
460,96 -> 464,96
468,137 -> 468,132 -> 468,137 -> 470,137 -> 470,130 -> 470,137 -> 472,137 -> 472,136 -> 472,137 -> 474,137 -> 474,127 -> 474,137 -> 476,137 -> 476,131 -> 476,137 -> 478,137 -> 478,130 -> 478,137 -> 480,137 -> 480,128 -> 480,137 -> 482,137 -> 482,128 -> 482,137 -> 484,137 -> 484,133 -> 484,137
472,123 -> 472,124 -> 476,124
468,137 -> 468,132 -> 468,137 -> 470,137 -> 470,130 -> 470,137 -> 472,137 -> 472,136 -> 472,137 -> 474,137 -> 474,127 -> 474,137 -> 476,137 -> 476,131 -> 476,137 -> 478,137 -> 478,130 -> 478,137 -> 480,137 -> 480,128 -> 480,137 -> 482,137 -> 482,128 -> 482,137 -> 484,137 -> 484,133 -> 484,137
460,159 -> 460,154 -> 460,159 -> 462,159 -> 462,151 -> 462,159 -> 464,159 -> 464,152 -> 464,159
461,87 -> 461,80 -> 461,87 -> 463,87 -> 463,80 -> 463,87 -> 465,87 -> 465,80 -> 465,87 -> 467,87 -> 467,86 -> 467,87
466,74 -> 466,71 -> 466,74 -> 468,74 -> 468,64 -> 468,74 -> 470,74 -> 470,65 -> 470,74 -> 472,74 -> 472,68 -> 472,74 -> 474,74 -> 474,68 -> 474,74 -> 476,74 -> 476,71 -> 476,74 -> 478,74 -> 478,65 -> 478,74 -> 480,74 -> 480,73 -> 480,74 -> 482,74 -> 482,64 -> 482,74
499,13 -> 499,16 -> 498,16 -> 498,19 -> 507,19 -> 507,16 -> 501,16 -> 501,13
464,109 -> 464,111 -> 463,111 -> 463,118 -> 473,118 -> 473,111 -> 468,111 -> 468,109
463,166 -> 467,166
483,44 -> 483,48 -> 481,48 -> 481,55 -> 489,55 -> 489,48 -> 487,48 -> 487,44
491,24 -> 496,24
466,74 -> 466,71 -> 466,74 -> 468,74 -> 468,64 -> 468,74 -> 470,74 -> 470,65 -> 470,74 -> 472,74 -> 472,68 -> 472,74 -> 474,74 -> 474,68 -> 474,74 -> 476,74 -> 476,71 -> 476,74 -> 478,74 -> 478,65 -> 478,74 -> 480,74 -> 480,73 -> 480,74 -> 482,74 -> 482,64 -> 482,74
488,26 -> 493,26
468,137 -> 468,132 -> 468,137 -> 470,137 -> 470,130 -> 470,137 -> 472,137 -> 472,136 -> 472,137 -> 474,137 -> 474,127 -> 474,137 -> 476,137 -> 476,131 -> 476,137 -> 478,137 -> 478,130 -> 478,137 -> 480,137 -> 480,128 -> 480,137 -> 482,137 -> 482,128 -> 482,137 -> 484,137 -> 484,133 -> 484,137
468,137 -> 468,132 -> 468,137 -> 470,137 -> 470,130 -> 470,137 -> 472,137 -> 472,136 -> 472,137 -> 474,137 -> 474,127 -> 474,137 -> 476,137 -> 476,131 -> 476,137 -> 478,137 -> 478,130 -> 478,137 -> 480,137 -> 480,128 -> 480,137 -> 482,137 -> 482,128 -> 482,137 -> 484,137 -> 484,133 -> 484,137
466,74 -> 466,71 -> 466,74 -> 468,74 -> 468,64 -> 468,74 -> 470,74 -> 470,65 -> 470,74 -> 472,74 -> 472,68 -> 472,74 -> 474,74 -> 474,68 -> 474,74 -> 476,74 -> 476,71 -> 476,74 -> 478,74 -> 478,65 -> 478,74 -> 480,74 -> 480,73 -> 480,74 -> 482,74 -> 482,64 -> 482,74
483,44 -> 483,48 -> 481,48 -> 481,55 -> 489,55 -> 489,48 -> 487,48 -> 487,44
460,159 -> 460,154 -> 460,159 -> 462,159 -> 462,151 -> 462,159 -> 464,159 -> 464,152 -> 464,159
499,13 -> 499,16 -> 498,16 -> 498,19 -> 507,19 -> 507,16 -> 501,16 -> 501,13
468,137 -> 468,132 -> 468,137 -> 470,137 -> 470,130 -> 470,137 -> 472,137 -> 472,136 -> 472,137 -> 474,137 -> 474,127 -> 474,137 -> 476,137 -> 476,131 -> 476,137 -> 478,137 -> 478,130 -> 478,137 -> 480,137 -> 480,128 -> 480,137 -> 482,137 -> 482,128 -> 482,137 -> 484,137 -> 484,133 -> 484,137
453,102 -> 458,102
485,28 -> 490,28
461,106 -> 466,106
454,106 -> 459,106
464,109 -> 464,111 -> 463,111 -> 463,118 -> 473,118 -> 473,111 -> 468,111 -> 468,109
478,33 -> 478,37 -> 471,37 -> 471,41 -> 483,41 -> 483,37 -> 482,37 -> 482,33
503,30 -> 508,30
461,87 -> 461,80 -> 461,87 -> 463,87 -> 463,80 -> 463,87 -> 465,87 -> 465,80 -> 465,87 -> 467,87 -> 467,86 -> 467,87
466,90 -> 470,90
460,159 -> 460,154 -> 460,159 -> 462,159 -> 462,151 -> 462,159 -> 464,159 -> 464,152 -> 464,159
466,74 -> 466,71 -> 466,74 -> 468,74 -> 468,64 -> 468,74 -> 470,74 -> 470,65 -> 470,74 -> 472,74 -> 472,68 -> 472,74 -> 474,74 -> 474,68 -> 474,74 -> 476,74 -> 476,71 -> 476,74 -> 478,74 -> 478,65 -> 478,74 -> 480,74 -> 480,73 -> 480,74 -> 482,74 -> 482,64 -> 482,74
460,164 -> 464,164
466,74 -> 466,71 -> 466,74 -> 468,74 -> 468,64 -> 468,74 -> 470,74 -> 470,65 -> 470,74 -> 472,74 -> 472,68 -> 472,74 -> 474,74 -> 474,68 -> 474,74 -> 476,74 -> 476,71 -> 476,74 -> 478,74 -> 478,65 -> 478,74 -> 480,74 -> 480,73 -> 480,74 -> 482,74 -> 482,64 -> 482,74
496,30 -> 501,30
461,87 -> 461,80 -> 461,87 -> 463,87 -> 463,80 -> 463,87 -> 465,87 -> 465,80 -> 465,87 -> 467,87 -> 467,86 -> 467,87
468,137 -> 468,132 -> 468,137 -> 470,137 -> 470,130 -> 470,137 -> 472,137 -> 472,136 -> 472,137 -> 474,137 -> 474,127 -> 474,137 -> 476,137 -> 476,131 -> 476,137 -> 478,137 -> 478,130 -> 478,137 -> 480,137 -> 480,128 -> 480,137 -> 482,137 -> 482,128 -> 482,137 -> 484,137 -> 484,133 -> 484,137
472,96 -> 476,96
510,30 -> 515,30
463,93 -> 467,93
466,74 -> 466,71 -> 466,74 -> 468,74 -> 468,64 -> 468,74 -> 470,74 -> 470,65 -> 470,74 -> 472,74 -> 472,68 -> 472,74 -> 474,74 -> 474,68 -> 474,74 -> 476,74 -> 476,71 -> 476,74 -> 478,74 -> 478,65 -> 478,74 -> 480,74 -> 480,73 -> 480,74 -> 482,74 -> 482,64 -> 482,74
466,74 -> 466,71 -> 466,74 -> 468,74 -> 468,64 -> 468,74 -> 470,74 -> 470,65 -> 470,74 -> 472,74 -> 472,68 -> 472,74 -> 474,74 -> 474,68 -> 474,74 -> 476,74 -> 476,71 -> 476,74 -> 478,74 -> 478,65 -> 478,74 -> 480,74 -> 480,73 -> 480,74 -> 482,74 -> 482,64 -> 482,74
468,137 -> 468,132 -> 468,137 -> 470,137 -> 470,130 -> 470,137 -> 472,137 -> 472,136 -> 472,137 -> 474,137 -> 474,127 -> 474,137 -> 476,137 -> 476,131 -> 476,137 -> 478,137 -> 478,130 -> 478,137 -> 480,137 -> 480,128 -> 480,137 -> 482,137 -> 482,128 -> 482,137 -> 484,137 -> 484,133 -> 484,137
468,137 -> 468,132 -> 468,137 -> 470,137 -> 470,130 -> 470,137 -> 472,137 -> 472,136 -> 472,137 -> 474,137 -> 474,127 -> 474,137 -> 476,137 -> 476,131 -> 476,137 -> 478,137 -> 478,130 -> 478,137 -> 480,137 -> 480,128 -> 480,137 -> 482,137 -> 482,128 -> 482,137 -> 484,137 -> 484,133 -> 484,137
468,137 -> 468,132 -> 468,137 -> 470,137 -> 470,130 -> 470,137 -> 472,137 -> 472,136 -> 472,137 -> 474,137 -> 474,127 -> 474,137 -> 476,137 -> 476,131 -> 476,137 -> 478,137 -> 478,130 -> 478,137 -> 480,137 -> 480,128 -> 480,137 -> 482,137 -> 482,128 -> 482,137 -> 484,137 -> 484,133 -> 484,137
466,74 -> 466,71 -> 466,74 -> 468,74 -> 468,64 -> 468,74 -> 470,74 -> 470,65 -> 470,74 -> 472,74 -> 472,68 -> 472,74 -> 474,74 -> 474,68 -> 474,74 -> 476,74 -> 476,71 -> 476,74 -> 478,74 -> 478,65 -> 478,74 -> 480,74 -> 480,73 -> 480,74 -> 482,74 -> 482,64 -> 482,74
475,99 -> 479,99
466,168 -> 470,168
468,137 -> 468,132 -> 468,137 -> 470,137 -> 470,130 -> 470,137 -> 472,137 -> 472,136 -> 472,137 -> 474,137 -> 474,127 -> 474,137 -> 476,137 -> 476,131 -> 476,137 -> 478,137 -> 478,130 -> 478,137 -> 480,137 -> 480,128 -> 480,137 -> 482,137 -> 482,128 -> 482,137 -> 484,137 -> 484,133 -> 484,137
466,74 -> 466,71 -> 466,74 -> 468,74 -> 468,64 -> 468,74 -> 470,74 -> 470,65 -> 470,74 -> 472,74 -> 472,68 -> 472,74 -> 474,74 -> 474,68 -> 474,74 -> 476,74 -> 476,71 -> 476,74 -> 478,74 -> 478,65 -> 478,74 -> 480,74 -> 480,73 -> 480,74 -> 482,74 -> 482,64 -> 482,74
469,99 -> 473,99
466,74 -> 466,71 -> 466,74 -> 468,74 -> 468,64 -> 468,74 -> 470,74 -> 470,65 -> 470,74 -> 472,74 -> 472,68 -> 472,74 -> 474,74 -> 474,68 -> 474,74 -> 476,74 -> 476,71 -> 476,74 -> 478,74 -> 478,65 -> 478,74 -> 480,74 -> 480,73 -> 480,74 -> 482,74 -> 482,64 -> 482,74
464,109 -> 464,111 -> 463,111 -> 463,118 -> 473,118 -> 473,111 -> 468,111 -> 468,109
499,13 -> 499,16 -> 498,16 -> 498,19 -> 507,19 -> 507,16 -> 501,16 -> 501,13
483,44 -> 483,48 -> 481,48 -> 481,55 -> 489,55 -> 489,48 -> 487,48 -> 487,44
495,26 -> 500,26
460,159 -> 460,154 -> 460,159 -> 462,159 -> 462,151 -> 462,159 -> 464,159 -> 464,152 -> 464,159
468,137 -> 468,132 -> 468,137 -> 470,137 -> 470,130 -> 470,137 -> 472,137 -> 472,136 -> 472,137 -> 474,137 -> 474,127 -> 474,137 -> 476,137 -> 476,131 -> 476,137 -> 478,137 -> 478,130 -> 478,137 -> 480,137 -> 480,128 -> 480,137 -> 482,137 -> 482,128 -> 482,137 -> 484,137 -> 484,133 -> 484,137
466,74 -> 466,71 -> 466,74 -> 468,74 -> 468,64 -> 468,74 -> 470,74 -> 470,65 -> 470,74 -> 472,74 -> 472,68 -> 472,74 -> 474,74 -> 474,68 -> 474,74 -> 476,74 -> 476,71 -> 476,74 -> 478,74 -> 478,65 -> 478,74 -> 480,74 -> 480,73 -> 480,74 -> 482,74 -> 482,64 -> 482,74
499,28 -> 504,28
483,44 -> 483,48 -> 481,48 -> 481,55 -> 489,55 -> 489,48 -> 487,48 -> 487,44
461,87 -> 461,80 -> 461,87 -> 463,87 -> 463,80 -> 463,87 -> 465,87 -> 465,80 -> 465,87 -> 467,87 -> 467,86 -> 467,87
483,44 -> 483,48 -> 481,48 -> 481,55 -> 489,55 -> 489,48 -> 487,48 -> 487,44
461,87 -> 461,80 -> 461,87 -> 463,87 -> 463,80 -> 463,87 -> 465,87 -> 465,80 -> 465,87 -> 467,87 -> 467,86 -> 467,87
469,93 -> 473,93
468,137 -> 468,132 -> 468,137 -> 470,137 -> 470,130 -> 470,137 -> 472,137 -> 472,136 -> 472,137 -> 474,137 -> 474,127 -> 474,137 -> 476,137 -> 476,131 -> 476,137 -> 478,137 -> 478,130 -> 478,137 -> 480,137 -> 480,128 -> 480,137 -> 482,137 -> 482,128 -> 482,137 -> 484,137 -> 484,133 -> 484,137
506,28 -> 511,28
468,137 -> 468,132 -> 468,137 -> 470,137 -> 470,130 -> 470,137 -> 472,137 -> 472,136 -> 472,137 -> 474,137 -> 474,127 -> 474,137 -> 476,137 -> 476,131 -> 476,137 -> 478,137 -> 478,130 -> 478,137 -> 480,137 -> 480,128 -> 480,137 -> 482,137 -> 482,128 -> 482,137 -> 484,137 -> 484,133 -> 484,137
469,166 -> 473,166
460,159 -> 460,154 -> 460,159 -> 462,159 -> 462,151 -> 462,159 -> 464,159 -> 464,152 -> 464,159
466,74 -> 466,71 -> 466,74 -> 468,74 -> 468,64 -> 468,74 -> 470,74 -> 470,65 -> 470,74 -> 472,74 -> 472,68 -> 472,74 -> 474,74 -> 474,68 -> 474,74 -> 476,74 -> 476,71 -> 476,74 -> 478,74 -> 478,65 -> 478,74 -> 480,74 -> 480,73 -> 480,74 -> 482,74 -> 482,64 -> 482,74
468,137 -> 468,132 -> 468,137 -> 470,137 -> 470,130 -> 470,137 -> 472,137 -> 472,136 -> 472,137 -> 474,137 -> 474,127 -> 474,137 -> 476,137 -> 476,131 -> 476,137 -> 478,137 -> 478,130 -> 478,137 -> 480,137 -> 480,128 -> 480,137 -> 482,137 -> 482,128 -> 482,137 -> 484,137 -> 484,133 -> 484,137
466,74 -> 466,71 -> 466,74 -> 468,74 -> 468,64 -> 468,74 -> 470,74 -> 470,65 -> 470,74 -> 472,74 -> 472,68 -> 472,74 -> 474,74 -> 474,68 -> 474,74 -> 476,74 -> 476,71 -> 476,74 -> 478,74 -> 478,65 -> 478,74 -> 480,74 -> 480,73 -> 480,74 -> 482,74 -> 482,64 -> 482,74
461,87 -> 461,80 -> 461,87 -> 463,87 -> 463,80 -> 463,87 -> 465,87 -> 465,80 -> 465,87 -> 467,87 -> 467,86 -> 467,87
`

const sand = '·'

const rock = '█'

type Room struct {
	Grid [][]rune
	MinX int
	MaxX int
	MinY int
	MaxY int
}

func (r *Room) Place(what rune, x, y int) {
	if r.MinX > x || r.MinX == 0 {
		r.MinX = x
	} else if r.MaxX < x {
		r.MaxX = x
	}
	if r.MinY > y || r.MinY == 0 {
		r.MinY = y
	} else if r.MaxY < y {
		r.MaxY = y
	}
	r.Grid[y][x] = what
}

func (r *Room) DrawRocks(sx, sy, ex, ey int) {
	x0, x1 := minMax(sx, ex)
	y0, y1 := minMax(sy, ey)
	// add more rows if needed
	for xr := len(r.Grid); xr <= y1; xr++ {
		r.Grid = append(r.Grid, make([]rune, 1000))
	}
	for y := y0; y <= y1; y++ {
		for x := x0; x <= x1; x++ {
			r.Place(rock, x, y)
		}
	}
}

func (r Room) Print() {
	for y := r.MinY; y <= r.MaxY; y++ {
		for x := r.MinX; x <= r.MaxX; x++ {
			if r.Grid[y][x] == 0 {
				fmt.Print(" ")
			} else {
				fmt.Print(string(r.Grid[y][x]))
			}
		}
		fmt.Printf("   %d", y)
		fmt.Println()
	}
	fmt.Println()
	strlen := len(fmt.Sprintf("%d", r.MaxX))
	for i := 0; i < strlen; i++ {
		for x := r.MinX; x <= r.MaxX; x++ {
			str := fmt.Sprintf("%d", x)
			fmt.Printf("%c", str[i])
		}
		fmt.Println()
	}
}

func main() {
	part1()
	part2()
}

func part1() {
	room := parseInput(input)
	rested := 0
	x, y := 500, 0
	for {
		if y+1 > room.MaxY {
			break // into the abyss!
		}
		// try down
		if room.Grid[y+1][x] == 0 {
			y++
			continue
		}
		// try down-left
		if room.Grid[y+1][x-1] == 0 {
			y++
			x--
			continue
		}
		// try down-right
		if room.Grid[y+1][x+1] == 0 {
			y++
			x++
			continue
		}
		// rest
		room.Grid[y][x] = sand
		room.Place(sand, x, y)
		x, y = 500, 0
		rested++
	}
	fmt.Println("Sand units rested:", rested)
}

func part2() {
	room := parseInput(input)
	room.DrawRocks(0, room.MaxY+2, 999, room.MaxY+2)
	rested := 0
	x, y := 500, 0
	for room.Grid[0][500] == 0 {
		if y+1 > room.MaxY {
			break // into the abyss!
		}
		// try down
		if room.Grid[y+1][x] == 0 {
			y++
			continue
		}
		// try down-left
		if room.Grid[y+1][x-1] == 0 {
			y++
			x--
			continue
		}
		// try down-right
		if room.Grid[y+1][x+1] == 0 {
			y++
			x++
			continue
		}
		// rest
		room.Grid[y][x] = sand
		room.Place(sand, x, y)
		x, y = 500, 0
		rested++
	}
	fmt.Println("Sand units rested:", rested)

}

func parseInput(input string) Room {
	lines := strings.Split(input, "\n")
	rows := make([][]rune, 10)
	for i := range rows {
		rows[i] = make([]rune, 1000)
	}
	room := Room{rows, 0, 0, 0, 0}
	for _, line := range lines {
		coords := strings.Split(line, " -> ")
		for i := 0; i < len(coords)-1; i++ {
			sx, sy := parseCoord(coords[i])
			ex, ey := parseCoord(coords[i+1])
			room.DrawRocks(sx, sy, ex, ey)
		}
	}
	return room
}

func parseCoord(s string) (int, int) {
	coord := strings.Split(s, ",")
	x, err := strconv.Atoi(coord[0])
	if err != nil {
		panic(err)
	}
	y, err := strconv.Atoi(coord[1])
	if err != nil {
		panic(err)
	}
	return x, y
}

func minMax(a, b int) (int, int) {
	if a < b {
		return a, b
	}
	return b, a
}
