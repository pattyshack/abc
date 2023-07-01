; bits 16

;
; Check to ensure we have enough low/conventional memory to load everything.
; Here we assume we're working on a modern system with plenty (>> 1MB) of ram.
;
real_mode_memory_check:
  pushf
  pusha

  ; http://wiki.osdev.org/Detecting_Memory_(x86)#Detecting_Low_Memory
  ; suggest checking the carry flag, but http://www.ctyme.com/intr/rb-0598.htm
  ; indicates int 0x12 does not return any value besides ax.
  int 0x12

  ; 639k of base memory (i.e., maximum conventional memory size)
  ; https://en.wikipedia.org/wiki/Conventional_memory
  cmp ax, 639
  jl .error

  mov si, ._ok_msg
  call print_str

  mov dx, ax
  call print_hex16

  mov si, _crlf
  call print_str

  popa
  popf
  ret

.error:
  mov si, ._err_msg
  call print_str

  mov dx, ax
  call print_hex16

  mov si, _crlf
  call print_str

  jmp halt

._ok_msg:
  db "Base mem: ", 0

._err_msg:
  db "Failed mem check: ", 0
