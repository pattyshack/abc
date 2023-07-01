; bits 16

halt:
  mov si, ._halt_msg
  call print_str

.halt_loop:
  hlt  ; halt
  jmp .halt_loop

._halt_msg:
  db "HALT", 0
