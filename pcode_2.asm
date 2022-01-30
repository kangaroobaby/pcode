var a, b, c
push 1
    push 2
    $sum
print "return = %d"
exit 0

    FUNC @sum:
        arg a, b

        push a
        push b
        add
        ret ~
    ENDFUNC
