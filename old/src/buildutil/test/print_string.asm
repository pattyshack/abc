; print string via BIOS
; input: bx points to null-terminated string
print_string:
  push ax
  push bx

  mov ah, 0x0e ; BIOS scrolling teletype func

.iter:
  mov al, [bx]

  cmp al, 0
  je .done

  int 0x10

  add bx, 1
  jmp .iter

.done:
  pop bx
  pop ax
  ret
