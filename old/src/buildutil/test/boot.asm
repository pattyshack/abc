; Basic boot sector
[org 0x7c00]

mov bx, msg
call print_string

sleep:
  hlt
  jmp sleep

%include "src/buildutil/test/print_string.asm"

msg:
  db "OK",0

times 510-($-$$) db 0
dw 0xaa55
