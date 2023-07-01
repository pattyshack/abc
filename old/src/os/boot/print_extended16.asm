print_reg:
  ; pushfd/pushad instead pushf/pusha since we want the full 32 bit registers
  pushfd
  pushad

  push esi
  push edx

  mov edx, eax
  call print_hex32
  mov si, _space
  call print_str

  mov edx, ebx
  call print_hex32
  mov si, _space
  call print_str

  mov edx, ecx
  call print_hex32
  mov si, _space
  call print_str

  pop edx  ; print edx
  call print_hex32
  mov si, _space
  call print_str

  pop edx  ; print si
  call print_hex32
  mov si, _space
  call print_str

  mov edx, edi
  call print_hex32
  mov si, _crlf
  call print_str

  popad
  popfd
  ret


; print si's content via BIOS (assuming little endian), cx specifies number of
; bytes to print
print_hex_number:
  pushf
  pusha

  push si

  mov si, _0x
  call print_str

  pop si

  mov ax, 0
  add si, cx
  sub si, 1

.iter:
  cmp ax, cx
  jge .done

  mov dl, [si]
  call print_hex_char

  sub si, 1
  add ax, 1
  jmp .iter

.done:
  popa
  popf
  ret


; print si's content via BIOS, cx specifies number of bytes to print
print_hex_bytes:
  pushf
  pusha

  mov ax, 0

.iter:
  cmp ax, cx
  jge .done

  mov dl, [si]
  call print_hex_char

  add si, 1
  add ax, 1
  jmp .iter

.done:
  popa
  popf
  ret

