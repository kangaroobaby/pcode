print "Hello world"         ; 输出：Hello world

push 1
push 2                      ; 相当于 print("(%d, %d)", 1, 2);
print "(%d, %d)"            ; 输出：(1, 2)

var x
readint "Input: "
pop x                       ; 相当于 x = readint("Input: ");
