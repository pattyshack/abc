; BIOS loads the boot sector at 0x7c00.
; See http://wiki.osdev.org/Memory_Map_(x86) for basic layout
org 0x7c00
bits 16  ; x86 always boots in real mode

;
; Memory map:
;
; [0x0000, 0x0400)
;     real mode interrupt vector table (unusable)
; [0x0400, 0x0500)
;     BIOS data area (unusable)
; [0x0500, _gbt_addr)
;     memory map (via detect_memory); should use boot sector's global variables
;     _memory_map_addr and _memory_map_count to access the table.
; [_gdt_addr, 0x????)
;     gdt
; [_page_table, 0x????)  - _page_table should be 4k aligned
; [0x????, 0x7c00)
;     stack (grows down)
; [0x7c00, 0x7e00)
;     boot sector
; [0x7e00, 0x7e00 + 512 * SECOND_STAGE_SECTORS)
;     second-stage boot sectors
; [0x????, 0x9fc00)
;     free for use
;
; For memory beyond 0x9fc00, use memory map table (via detect memory).
;


%define SECOND_STAGE_SECTORS 4

; Boot loader's entry point is just the first byte of the sector.
boot_sector_entry_point:

%include "src/os/boot/init_registers16.asm"

;
; Now that the stack have been setup, we can make calls
;

; Virtualbox defaults to this already, but setting it never hurts
call setup_vga_text_mode

; This should be early in the boot process in case dl gets clobbered.
call save_boot_drive_id

; Make sure we have space to load additional boot data.
call real_mode_memory_check

; Read SECOND_STAGE_SECTORS from boot disk, and load that into 0x7e00.
; [0x7e00, 0x9fc00) is free for use. See http://wiki.osdev.org/Memory_Map_(x86)
mov dh, SECOND_STAGE_SECTORS
mov bx, 0x7e00
call load_boot_data

;
; Time to move off the boot sector.
;
; Since the boot sector is in [0x7c00, 0x7e00), and the newly loaded sectors
; are in [0x7e00, ...), the two memory ranges are adjacent.  Hence, the label
; address offset are identical, and the boot sector can referr to the second
; sector via label (and vice versa).
;
jmp second_stage_entry_point  ; can also use "0x0000:0x7e00"

;
; Code
;

%include "src/os/boot/boot_disk16.asm"
%include "src/os/boot/halt16.asm"
%include "src/os/boot/print_basic16.asm"
%include "src/os/boot/real_mode_memory_check16.asm"
%include "src/os/boot/setup_vga_text_mode16.asm"

;
; Shared global variables
;

; Which drive did the BIOS boot from.  Initialized by save_boot_drive_id call.
_boot_drive_id:
  db 0

%define MEMORY_MAP_ENTRY_SIZE 24

_space:
  db ' ', 0

_crlf:
  db 13, 10, 0  ; BIOS printing requires explicit \r

_0x:
  db '0x', 0

;
; End of boot sector.  The last two bytes on the boot sector must end with
; 0xaa55.
;

times 510-($-$$) db 0
dw 0xaa55

; -----------------------------------------------------------------------------

; This sector is loaded into 0x7e00.
second_sector:

;
; Code
;

%include "src/os/boot/a20_16.asm"
%include "src/os/boot/check_64bit_support16.asm"
%include "src/os/boot/detect_memory16.asm"
%include "src/os/boot/print_extended16.asm"
%include "src/os/boot/sleep16.asm"

;
; Shared global variables
;

_second_stage_msg:
  db "Entered second stage booting", 13, 10, 0

align 4

; Num entries (not byte count) in the memory map.  Initialized by detect_memory
; call.
_memory_map_count:
  dw 0  ; initialized by detect_memory call

; Where to save the memory map.  The map itself is initialized by call to
; detect_memory.  Each entry is a packed struct
;
; struct MemoryMapEntry {
;   uint64 base_addr;
;   uint64 len;  // in bytes
;   uint32 type;
;   unint32 flags;
; };
;
; See http://wiki.osdev.org/Detecting_Memory_(x86) and
; http://www.ctyme.com/intr/rb-1741.htm
_memory_map_addr:
  dd 0x0500  ; 8 byte pointer
  dd 0

; gdt description structure.  See
; http://wiki.osdev.org/Global_Descriptor_Table
_gdt_size:
  dw 0  ; in bytes.  Initialized by setup_gdt

_gdt_addr:
  dd 0  ; 8 byte pointer.  Initialized by setup_gdt call
  dd 0

; Unlike the boot sector entry point, the second stage entry point can be
; anywhere.
second_stage_entry_point:
  mov si, _second_stage_msg
  call print_str

; We want to use more than 1MB of memory ...
call enable_a20

; Find out how much memory do we have.
call detect_memory
call print_memory_map

; Assert 64 bit mode is available.  The kernel will be in 64 bit.
call check_cpuid_available
call check_64bit_support

;call setup_gdt

;; jumping into long mode follows the same procedure as jumping into protected
;; mode.
;; http://wiki.osdev.org/Journey_To_The_Protected_Land
;; http://wiki.osdev.org/Setting_Up_Long_Mode
;switch_to_long_mode:
;  ; Disable paging.  We don't really need to do this since we never setup
;  ; paging for protected mode, but it doesn't hurt to run a few instructions.
;  ; See https://en.wikipedia.org/wiki/Control_register
;  mov eax, cr0
;  and eax 0x7ffffff  ; set bit 31 to zero
;  mov cr0, eax
;
;  cli  ; disable interrupt
;  ; load gdt
;  ; set cr0
;  ; jump to new long mode code
;


;
; To be continue ...
;

; TODO: switch to graphics mode.  For now, stick to vga text mode at 0xb8000.
; Info:
;   http://wiki.osdev.org/Getting_VBE_Mode_Info
;   http://wiki.osdev.org/User:Omarrx024/VESA_Tutorial
;   http://wiki.osdev.org/Drawing_In_Protected_Mode

jmp halt

; VBoxManage requires the generate bin file to be multiples of 512 bytes ...
times (512*(SECOND_STAGE_SECTORS+1))-($-$$) db 0

