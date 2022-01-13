from gpiozero import LED,Button
from signal import pause
from time import sleep

led = LED(17)
button = Button(6)


#while True:
   # led.on()
   # sleep(1)
  #  led.off()
  #  sleep(1)

button.when_pressed=led.on
button.when_released = led.off
pause()
