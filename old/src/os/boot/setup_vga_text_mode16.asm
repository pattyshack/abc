; bits 16

; http://www.ctyme.com/intr/rb-0069.htm
; http://wiki.osdev.org/Text_UI
setup_vga_text_mode:
  pushf
  pusha

  mov ah, 0
  mov al, 3  ; mode 3 = 80x25 vga text mode, with video memory at 0xB8000

  int 0x10

  popa
  popf
  ret
