; bits 16

;
; Initialize code segment register to zero.  Some BIOS load the boot sector
; at 0x07c0:0x0000, while others load the boot sector at 0x0000:0x7c00.
; Since we have specified org 0x7c00 at the beginning of the file, we want
; the latter.  Performing a far jump with absolute address to ensure code
; segement register is set to zero.
;
jmp 0x0000:.init_non_cs_registers

;
; Initialize non-cs segment registers and stack registers to zero.
;
.init_non_cs_registers:
  cli  ; Disable interrupts

  xor ax, ax ; i.e., mov ax, 0

  mov ds, ax  ; data segment
  mov es, ax  ; extra segment
  mov fs, ax  ; no specific use
  mov gs, ax  ; no specific use

  ; Stack registers must be set together with interrupt disable since
  ; interrupts may access the stack while the registers are changing.
  ;
  ; Setup stack before boot sector.
  ; [0x0500, 0x7c00) is free for use. See http://wiki.osdev.org/Memory_Map_(x86)
  mov ss, ax  ; stack segment
  mov bp, boot_sector_entry_point  ; stack base pointer register
  mov sp, bp  ; stack pointer register

  cld  ; clear direction flag.  increment esi / edi instead of decrement

  sti  ; Re-enable interrupts
