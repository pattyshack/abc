; bits 16

;
; Detect memory via int 0x15, ax=0xe820.  The memory map is saved to the
; address stored in [_memory_map_addr], and the length of the map is save
; into _memory_map_count.
;
; For additional detail, see:
;   http://wiki.osdev.org/Detecting_Memory_(x86)
;   http://www.ctyme.com/intr/rb-1741.htm
;   (error status http://www.ctyme.com/intr/rb-1506.htm#Table496)
;

; Constant expected by int 0x15, ax=0xe820 calls.
%define SMAP 0x534d4150

detect_memory:
  pushf
  pusha

  mov si, 0  ; Track number of map entries.
  mov di, [_memory_map_addr]  ; Where to save the first entry of the memory map

  ; 32-bit registers and instructions are usable in real mode.  However, the
  ; memory addressing are different.  So, "mov eax, 123" is safe, but
  ; "mov [eax], 123" is not ...
  mov ebx, 0  ; Ask for the first entry of the map

.read_entry:
  mov eax, 0xe820
  mov ecx, MEMORY_MAP_ENTRY_SIZE  ; Ask for 24 bytes, i.e., a single entry.
  mov edx, SMAP

  ; A map entry is either 20 bytes (in older BIOS) or 24 bytes (4 extra bytes
  ; for extended flags support, starting in ACPI 3.x). In case the interrupt
  ; call returns an older 20 byte entry, we need to pad the last word word
  ; with 1 in order to make it a valid ACPI 3.x entry (in ACPI 3.x, we are
  ; suppose to ignore the entry if bit 0 is set to zero).
  mov [di+20], dword 1

  int 0x15
  jc .int_error

  cmp eax, SMAP
  jne .invalid_entry

  ; Entry must either be 20 bytes ...
  cmp ecx, 20
  je .valid_entry

  ; or 24 bytes
  cmp ecx, MEMORY_MAP_ENTRY_SIZE
  jne .invalid_entry

  ; For 24 byte entries, we need to check the flag's bit 0, and ignore the
  ; entry if it's not set.
  mov ecx, [di+20]
  and ecx, 1
  jz .maybe_read_next_entry  ; reuse di address, and don't inc si

.valid_entry:
  add di, MEMORY_MAP_ENTRY_SIZE  ; next entry's location
  add si, 1  ; increment count

.maybe_read_next_entry:
  cmp ebx, 0
  je .done

  jmp .read_entry

.done:
  cmp si, 0
  je .no_entry

  mov [_memory_map_count], si

  mov si, ._found_msg
  call print_str

  mov dx, [_memory_map_count]
  call print_hex16

  mov si, _crlf
  call print_str

  popa
  popf

  ret

.no_entry:
  mov si, ._no_entry_msg
  call print_str

  jmp halt

.invalid_entry:
  mov si, ._invalid_entry_msg
  call print_str

  ; printing out eax, ebx, ecx is a pain in real mode ...

  jmp halt

.int_error:
  mov si, ._int_error_msg
  call print_str

  mov dh, ah
  mov dl, 0
  call print_hex16

  mov si, _crlf
  call print_str

  jmp halt

._found_msg:
  db 'Number of memory map entries: ', 0

._invalid_entry_msg:
  db 'Failed to detect memory.  Invalid entry found.', 13, 10, 0

._no_entry_msg:
  db 'Failed to detect memory.  No entry found.', 13, 10, 0

._int_error_msg:
  db 'Failed to detect memory.  Interrupt error: ', 0


print_memory_map:
  pushf
  pusha

  mov si, ._msg
  call print_str

  mov di, [_memory_map_addr]
  mov bx, 0

.iter:
  cmp bx, [_memory_map_count]
  jge .done

  mov si, _space
  call print_str
  call print_str

  ; print addr

  mov si, di
  mov cx, 8
  call print_hex_number

  mov si, _space
  call print_str

  ; print len

  mov si, di
  add si, 8
  call print_hex_number


  mov si, _space
  call print_str

  ; print type

  mov si, di
  add si, 16
  mov cx, 4
  call print_hex_number

  mov si, _space
  call print_str

  ; print flag

  mov si, di
  add si, 20
  call print_hex_number

  mov si, _crlf
  call print_str

  add bx, 1
  add di, MEMORY_MAP_ENTRY_SIZE

  jmp .iter

.done:
  popa
  popf
  ret

._msg:
  db 'Memory Map:', 13, 10, 0

._0x:
  db '0x', 0
