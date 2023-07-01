; bits 16

; print null-terminated string stored in si
print_str:
  pushf
  pusha

  cld  ; make sure we're incrementing on lodsb

  mov bx, 0

.iter:
  lodsb
  or al, al  ; this sets zf to 1 if the result is zero
  jz .done ; if zf is set

  mov ah, 0x0e
  int 0x10  ; http://www.ctyme.com/intr/rb-0106.htm

  jmp .iter

.done:
  popa
  popf
  ret

; print dl's value via BIOS (without 0x prefix)
print_hex_char:
  pushf
  pusha

  mov ah, 0x0e
  mov bx, 0

  mov al, dl
  shr al, 4
  call .format_hex
  int 0x10

  mov al, dl
  and al, 0x0f
  call .format_hex
  int 0x10

  popa
  popf
  ret

.format_hex:
  cmp al, 10
  jge .else

  add al, '0'
  ret
.else:
  add al, ('A'-10)
  ret

; print dl's value via BIOS (with 0x prefix)
print_hex8:
  pushf
  pusha

  mov si, _0x
  call print_str

  call print_hex_char

  popa
  popf
  ret

; print dx's value via BIOS
print_hex16:
  pushf
  pusha
  push dx

  mov ah, 0x0e
  mov bx, 0

  mov al, '0'
  int 0x10

  mov al, 'x'
  int 0x10

  mov dl, dh
  call print_hex_char

  pop dx
  call print_hex_char

  popa
  popf
  ret

; print edx's value via BIOS
print_hex32:
  pushf
  pusha

  push dx
  shr edx, 16
  push dx

  mov ah, 0x0e
  mov bx, 0

  mov al, '0'
  int 0x10

  mov al, 'x'
  int 0x10

  shr dx, 8
  call print_hex_char

  pop dx
  call print_hex_char

  mov si, sp
  mov dx, [si]
  shr dx, 8
  call print_hex_char

  pop dx
  call print_hex_char

  popa
  popf
  ret

