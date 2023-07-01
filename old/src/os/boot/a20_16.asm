; bits 16

;
; Due to historical reasons, we need to ensure A20 line is enabled in order
; to use unlock upper memory ...
;   http://wiki.osdev.org/A20_Line
;   http://www.win.tue.nl/~aeb/linux/kbd/A20.html
;

enable_a20:
  pushf
  push ax
  push bx
  push cx
  push dx
  push si

  ; skip enabling a20 if it's already enabled
  call is_a20_enabled
  test ax, 1
  jnz .done

  ; try to enable a20 via bios first
  call enable_a20_bios

  ; 100ms
  mov ax, 1
  call busy_sleep

  call is_a20_enabled
  test ax, 1
  jnz .done

  ; bios failed, try to enable a20 via keyboard controller
  call enable_a20_keyboard

  ; 200ms
  mov ax, 2
  call busy_sleep

  call is_a20_enabled
  test ax, 1
  jnz .done

  ; keyboard controller also failed, finally try to enable a20 via fast gate
  call enable_a20_fast

  ; 200ms
  mov ax, 2
  call busy_sleep

  call is_a20_enabled
  test ax, 1
  jz .error

.done:
  call print_a20_status

  pop si
  pop dx
  pop cx
  pop bx
  pop ax
  popf

  ret

.error:
  mov si, ._err_msg
  call print_str

  jmp halt

._err_msg:
  db 'Failed to enable A20', 13, 10, 0


; [0x7dfe] = 0xaa55 (i.e., the magic number at the end of the boot sector)
%define A20_TEST_ADDR 0x7dfe
%define WRAP_OFFSET 0x0010

;
; When A20 is disabled, memory address wraps around:
;     0x0000:X and 0xffff:(X+0x0010) refers to the same location
;
; input: none
; return: ax = 1 if a20 is enable else 0
;
is_a20_enabled:
  pushf
  push fs
  push gs
  push bx
  push cx

  cli  ; disable interrupt

  ; set fs to 0x0000 and gs to 0xffff
  xor ax, ax  ; ax = 0
  mov fs, ax

  not ax  ; ax = 0xffff
  mov gs, ax

  ; save values referenced by test addrs.
  mov cl, [fs:A20_TEST_ADDR]
  mov bl, [gs:(A20_TEST_ADDR+WRAP_OFFSET)]

  mov byte [fs:A20_TEST_ADDR], 0xab
  mov byte [gs:(A20_TEST_ADDR+WRAP_OFFSET)], 0xba

  cmp byte [fs:A20_TEST_ADDR], 0xab
  jne .disabled

  ; enabled
  mov ax, 1

.done:
  ; restore value to test addrs
  mov [fs:A20_TEST_ADDR], cl
  mov [gs:(A20_TEST_ADDR+WRAP_OFFSET)], bl

  sti  ; re-enable interrupt

  pop cx
  pop bx
  pop gs
  pop fs
  popf

  ret

.disabled:
  mov ax, 0

  jmp .done

print_a20_status:
  pushf
  push ax
  push dx
  push si

  call is_a20_enabled

  mov si, ._msg
  call print_str

  mov dx, ax
  call print_hex16

  mov si, _crlf
  call print_str

  pop si
  pop dx
  pop ax
  popf

  ret

._msg:
  db 'A20 Enabled: ', 0

;
; A20 via BIOS ---------------------------------------------------------------
;

; Enable A20 via bios.  input: none / return: none
; http://www.ctyme.com/intr/rb-1336.htm
enable_a20_bios:
  pushf
  push ax
  push dx
  push si

  mov ax, 0x2401
  int 0x15
  jc .error

  mov si, ._ok_msg
  call print_str

.done:
  mov dh, ah
  mov dl, 0
  call print_hex16

  mov si, _crlf
  call print_str

  pop si
  pop dx
  pop ax
  popf

  ret

.error:
  mov si, ._err_msg
  call print_str

  jmp .done

._ok_msg:
  db 'A20 enabled via BIOS: ', 0

._err_msg:
  db 'Failed to Enable A20 via BIOS: ', 0

; Disable A20 via bios.  input: none / return: none
; http://www.ctyme.com/intr/rb-1335.htm
disable_a20_bios:
  pushf
  push ax
  push dx
  push si

  mov ax, 0x2400
  int 0x15
  jc .error

  mov si, ._ok_msg
  call print_str

.done:
  mov dh, ah
  mov dl, 0
  call print_hex16

  mov si, _crlf
  call print_str

  pop si
  pop dx
  pop ax
  popf

  ret

.error:
  mov si, ._err_msg
  call print_str

  jmp .done

._ok_msg:
  db 'Disabled A20 via BIOS: ', 0

._err_msg:
  db 'Failed to Disable A20 via BIOS: ', 0

; Print A20 bios status.  input: none / return: none
; http://www.ctyme.com/intr/rb-1337.htm
print_a20_bios_status:
  pushf
  push ax
  push dx
  push si

  mov ax, 0x2402
  int 0x15

  jc .error

  mov si, ._ok_msg
  call print_str

.done:
  mov dx, ax
  call print_hex16

  mov si, _crlf
  call print_str

  pop si
  pop dx
  pop ax
  popf

  ret

.error:
  mov si, ._err_msg
  call print_str

  jmp .done

._ok_msg:
  db 'A20 BIOS status: ', 0

._err_msg:
  db 'Failed to get A20 BIOS status: ', 0

;
; A20 via fast gate -------------------------------------------------------
; For additional info, see http://www.win.tue.nl/~aeb/linux/kbd/A20.html,
; section "A20 control via System Control Port A"
;

enable_a20_fast:
  pushf
  push ax
  push si

  ; port 0x92 = System Control Port A
  ; http://wiki.osdev.org/I/O_Ports
  in al, 0x92

  ; The second bit is the A20 bit
  test al, 2
  jnz .done

  ; We want the first bit to be zero since setting it to one will cause a fast
  ; reset.
  or al, 2  ; set A20 bit
  and al, 0xfe

  out 0x92, al

.done:
  mov si, ._msg
  call print_str

  pop si
  pop ax
  popf

  ret

._msg:
  db 'A20 enabled via fast gate', 13, 10, 0

disable_a20_fast:
  pushf
  push ax
  push si

  ; port 0x92 = System Control Port A
  ; http://wiki.osdev.org/I/O_Ports
  in al, 0x92

  ; The second bit is the A20 bit
  test al, 2
  jz .done

  ; We want the first bit to be zero since setting it to one will cause a fast
  ; reset.  Setting the second bit to zero disables A20.
  and al, 0xfc

  out 0x92, al

.done:
  mov si, ._msg
  call print_str

  pop si
  pop ax
  popf

  ret

._msg:
  db 'A20 fast gate disabled', 13, 10, 0

; Print A20 fast gate status.  input: none / return: none
print_a20_fast_status:
  pushf
  push ax
  push dx
  push si

  ; port 0x92 = System Control Port A
  ; http://wiki.osdev.org/I/O_Ports
  in al, 0x92

  and al, 0x02 ; 2nd bit is the A20 status

  mov si, ._msg
  call print_str

  mov dl, al
  mov dh, 0
  call print_hex16

  mov si, _crlf
  call print_str

  pop si
  pop dx
  pop ax
  popf
  ret

._msg:
  db 'A20 Fast status: ', 0

;
; A20 via keyboard controller -------------------------------------------------
; For additional info, see
;   http://wiki.osdev.org/I/O_Ports
;   http://wiki.osdev.org/%228042%22_PS/2_Controller
;
; 0x64 is the keyboard status/command port, while 0x60 is the data port
;

enable_a20_keyboard:
  pushf
  push ax
  push si

  cli  ; disable interrupts

  ; disable keyboard
  call _keyboard_wait_until_input_processed
  mov al, 0xad
  out 0x64, al

  ; read controller output port
  call _keyboard_wait_until_input_processed
  mov al, 0xd0
  out 0x64, al

  ; write byte from keyboard to al
  call _keyboard_wait_until_output_ready
  in al, 0x60
  push ax  ; save for use later


  ; write next byte to keyboard controller output port
  call _keyboard_wait_until_input_processed
  mov al, 0xd1
  out 0x64, al

  call _keyboard_wait_until_input_processed
  pop ax
  or al, 2  ; 2nd bit controls A20
  out 0x60, al  ; port 0x60 is the data port

  ; re-enable keyboard
  call _keyboard_wait_until_input_processed
  mov al, 0xae
  out 0x64, al

  sti  ; re-enable interrupts

  mov si, ._msg
  call print_str

  pop si
  pop ax
  popf
  ret

._msg:
  db 'A20 enabled via keyboard controller', 13, 10, 0

disable_a20_keyboard:
  pushf
  push ax
  push si

  cli  ; disable interrupts

  ; disable keyboard
  call _keyboard_wait_until_input_processed
  mov al, 0xad
  out 0x64, al

  ; read controller output port
  call _keyboard_wait_until_input_processed
  mov al, 0xd0
  out 0x64, al

  ; write byte from keyboard to al
  call _keyboard_wait_until_output_ready
  in al, 0x60
  push ax  ; save for use later


  ; write next byte to keyboard controller output port
  call _keyboard_wait_until_input_processed
  mov al, 0xd1
  out 0x64, al

  call _keyboard_wait_until_input_processed
  pop ax
  and al, 0xfd  ; 2nd bit controls A20
  out 0x60, al  ; port 0x60 is the data port

  ; re-enable keyboard
  call _keyboard_wait_until_input_processed
  mov al, 0xae
  out 0x64, al

  sti  ; re-enable interrupts

  mov si, ._msg
  call print_str

  pop si
  pop ax
  popf
  ret

._msg:
  db 'A20 keyboard disabled', 13, 10, 0

print_a20_keyboard_status:
  pushf
  push ax
  push dx
  push si

  cli  ; disable interrupts

  ; disable keyboard
  call _keyboard_wait_until_input_processed
  mov al, 0xad
  out 0x64, al

  ; read controller output port
  call _keyboard_wait_until_input_processed
  mov al, 0xd0
  out 0x64, al

  ; write byte from keyboard to al
  call _keyboard_wait_until_output_ready
  in al, 0x60
  push ax  ; save for use later

  ; re-enable keyboard
  call _keyboard_wait_until_input_processed
  mov al, 0xae
  out 0x64, al

  sti  ; re-enable interrupts

  pop ax

  and al, 0x02 ; 2nd bit is the A20 status

  mov si, ._msg
  call print_str

  mov dl, al
  mov dh, 0
  call print_hex16

  mov si, _crlf
  call print_str

  pop si
  pop dx
  pop ax
  popf
  ret

._msg:
  db 'A20 keyboard status: ', 0

; Wait until the keyboard is done processing the previous input
; (helper function, assumes interrupts are disabled.)
_keyboard_wait_until_input_processed:
  ; http://wiki.osdev.org/I/O_Ports
  in al, 0x64

  ; status zero at bit 1 means output buffer is empty
  ; http://wiki.osdev.org/%228042%22_PS/2_Controller
  test al, 2
  jnz _keyboard_wait_until_input_processed

  ret

; Wait until there is something to read from keyboard's output
; (helper function, assumes interrupts are disabled.)
_keyboard_wait_until_output_ready:
  ; http://wiki.osdev.org/I/O_Ports
  in al, 0x64

  ; status zero at bit 0 means output buffer is empty
  ; http://wiki.osdev.org/%228042%22_PS/2_Controller
  test al, 1
  jz _keyboard_wait_until_output_ready

  ret

