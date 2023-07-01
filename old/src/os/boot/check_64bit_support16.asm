; bits 16

; Info on CPUID:
;   http://www.sandpile.org/x86/cpuid.htm
;   http://wiki.osdev.org/CPUID
;   https://en.wikipedia.org/wiki/FLAGS_register
;   https://en.wikipedia.org/wiki/CPUID
;
; Unlike int, cpuid can be called in protected/long mode. We'll do the
; minimum here and only check for 64 bit mode.

; bit 21 (0x200000) of EFLAGS indicates if we can use the cpuid instruction.
; This bit is modifiable iff cpuid instruction is supported.
check_cpuid_available:
  pushfd
  pusha

  ; read EFLAGS' value into eax, ebx
  pushfd
  pop eax
  mov ebx, eax

  ; flip eax's cpuid flag bit and load it into EFLAGS
  xor eax, 1 << 21
  push eax
  popfd

  ; load the EFLAG into eax again to see if the cpuid flag bit changed
  pushfd
  pop eax

  ; if cpuid is not supported, the bit in eax will be revert back to the
  ; original value, i.e., ebx
  cmp eax, ebx
  je .cupid_not_supported

  mov si, ._ok_msg
  call print_str

  popa
  popfd
  ret

.cupid_not_supported:
  mov si, ._failed_msg
  call print_str

  jmp halt

._ok_msg:
  db 'CPUID supported', 13, 10, 0

._failed_msg:
  db 'CPUID not supported', 13, 10, 0

; Detect if the cpu can support x64_64
;
; http://wiki.osdev.org/Setting_Up_Long_Mode
; http://www.sandpile.org/x86/cpuid.htm#level_8000_0001h
check_64bit_support:
  pushf
  pusha

  ; Get maximum supported extended level
  mov eax, 0x80000000
  cpuid

  cmp eax, 0x80000001
  jb .not_supported  ; 0x80000001 is not available

  mov eax, 0x80000001
  cpuid

  ; bit 29 of edx tells us if long mode is available
  test edx, 1 << 29
  jz .not_supported

  mov si, ._supported_msg
  call print_str

  popa
  popf
  ret

.not_supported:
  mov si, ._unsupported_msg
  call print_str

  jmp halt

._supported_msg:
  db 'CPU supports long mode', 13, 10, 0

._unsupported_msg:
  db 'CPU does not support long mode', 13, 10, 0
