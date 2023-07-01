; sleep for ax * 100ms.  ax must be in [1, 600]
busy_sleep:
  pushf
  pusha

  ; do nothing if input is zero
  cmp ax, 0
  je .input_error

  cmp ax, 600
  jg .input_error

  jmp .input_ok

.input_error:
  mov si, ._msg
  call print_str

  mov si, _crlf

  mov dx, ax
  call print_hex16

  mov si, _crlf
  call print_str

  jmp halt

.input_ok:
  ; ax = (ax * 100 + 54) / 55
  mov cx, 100
  mul cx
  add ax, 54  ; round up
  mov cx, 55
  div cx

;
; compute when to stop
;
  mov bx, ax

  ; get system time (# ticks since midnight)
  ; http://www.ctyme.com/intr/rb-2271.htm
  mov ah, 0x00
  int 0x1a

  add bx, dx

  cmp bx, dx
  ja .bx_not_overflow  ; jg is signed, ja is unsigned

  add cx, 1  ; carry one

.bx_not_overflow:

  cmp cx, 0x18
  je .last_hour

  jmp .sleep_until

.last_hour:
  cmp bx, 0x00b0
  jb .sleep_until  ; jle is signed, jbe is unsigned

  ; pass midnight, cx:bx = cx:bx mod 0x1800b0
  mov cx, 0
  sub bx, 0x00b0

;
; Sleep until we have reach the computed stop tick
;

.sleep_until:
  push bx
  push cx

  ; copy sp into si so that we can access stack without popping
  mov si, sp

.sleep_loop:
  mov ah, 0x00
  int 0x1a

  mov bx, [si]  ; stop hour

  ; if stop hour == 24
  cmp bx, 0x18
  je .cmp_midnight  ; handle wrap around

.cmp_hour:
  ; if stop hour < current hour
  cmp bx, cx
  jb .done  ; jl is signed, jb is unsigned

  ; if stop hour > current hour
  cmp bx, cx
  ja .sleep_loop

  mov bx, [si+2]  ; stop minute

  ; if stop minute > current minute
  cmp bx, dx
  ja .sleep_loop

.done:
  pop cx
  pop bx

  popa
  popf
  ret

.cmp_midnight:
  cmp cx, 0
  je .done  ; past midnight

  jmp .cmp_hour

._msg:
  db 'Invalid sleep duration: ', 0

