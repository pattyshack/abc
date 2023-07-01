; bits 16

;
; Save which drive number did the BIOS boot from.
;
save_boot_drive_id:
  pushf
  pusha

  mov [_boot_drive_id], dl  ; BIOS stores the drive id in dl on startup

  mov si, .msg
  call print_str

  xor dh, dh
  call print_hex16

  mov si, _crlf
  call print_str

  popa
  popf
  ret

.msg:
  db 'Boot drive: ', 0

;
; load sectors [2, 2 + dh) from the boot drive (cylinder 0, head 0) into
; memory at es:bx.  This assumes that save_boot_drive_id is already called.
; (The first sector, i.e., the boot sector, is ignored since it's already
; loaded)
;
load_boot_data:
  pushf
  pusha
  push dx

  ; XXX maybe add retry loop

  ; Reset the boot drive via int=0x13,ah=0x00 to ensure the drive is positioned
  ; correctly.
  ; http://www.ctyme.com/intr/rb-0605.htm
  mov ah, 0x00
  mov dl, [_boot_drive_id]

  int 0x13
  jc .reset_error

  ; Perform read via int=0x13,ah=0x02
  ; http://www.ctyme.com/intr/rb-0607.htm
  mov ah, 0x02
  mov al, dh  ; number of sectors to read
  mov ch, 0  ; cylinder 0
  mov cl, 0x02  ; starting from the 2nd sector
  mov dh, 0x00  ; head 0
  ; dl is already set to [_boot_drive_id]

  int 0x13
  jc .read_error

  pop dx

  mov dl, dh
  xor dh, dh

  ; Extra check to ensure number of sectors read is expected
  cmp al, dl
  jne .read_error

  mov si, ._ok_msg
  call print_str

  call print_hex16

  mov si, _crlf
  call print_str

  popa
  popf
  ret

.reset_error:
  mov si, ._reset_err_msg
  call print_str

  mov dh, ah
  xor dl, dl
  call print_hex16

  mov si, _crlf
  call print_str

  jmp halt

.read_error:
  mov si, ._read_err_msg
  call print_str

  mov dx, ax
  call print_hex16

  mov si, _crlf
  call print_str

  jmp halt

._ok_msg:
  db "Sectors read: ", 0

._reset_err_msg:
  db "Reset drive failed: ", 0

._read_err_msg:
  db "Read failed: ", 0

