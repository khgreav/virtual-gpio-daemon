#!/usr/bin/env python3
from pynput import keyboard
import gpiod
import sys

CHIP = "/dev/gpiochip0"
OUTPUT_LINE = 0

chip = gpiod.Chip(CHIP)
out = chip.get_line(OUTPUT_LINE)
out.request(consumer="out", type=gpiod.LINE_REQ_DIR_OUT, default_vals=[0])

def on_press(key):
    if key == keyboard.Key.space:
        out.set_value(1)

def on_release(key):
    if key == keyboard.Key.space:
        out.set_value(0)
    if key == keyboard.Key.esc:
        sys.exit(0)

with keyboard.Listener(on_press=on_press, on_release=on_release) as listener:
    listener.join()
out.release()
