# :zap: Pi Read Meter! :movie_camera:

Raspberry Pi + Camera + pi-read-meter = reading gas meter every month automatically.

Instead of taking the number every month manually on the first day of a month, I made it automatic by one old Raspberry Pi Model B and Raspberry Pi NOIR camera V2.

Perfect use of Go in my opinion! I really enjoy building binary on my Mac, transporting it to Pi and running there. Very convenient!

# How did I do it?

Here I'll go over the steps to set something similar up to yourself. These are notes for future self and for anyone who is interested in it.

**Basic equipment**
* (Gas/water/electricity) meter to read
* Raspberry Pi with memory card and connection to internet
* Raspberry Pi camera or some other camera. It must be possible to capture image using command line command.
* In case of dark room, some low consumption night light or similar might be necessary
* Other computer (can be done on Pi as well) to build Go binaries on (this code right here)

**Setting it up from scratch**
1. Get Raspbian (Lite is fine) running in your Pi: https://www.raspberrypi.org/downloads/raspbian/
2. Boot up the Pi and make sure to have access to the command line.
3. Check your Pi ARM version: `cat /proc/cpuinfo`
4. Clone this repository to a computer with Go installed
5. Change ARM version in `Makefile` if yours is different from the one I used in `build-pi` command.
6. Build the binary for Pi using `make build-pi`
7. When you have the binary, move it to Pi (using `scp` or some other means)
8. Move also the `config.example.json` from this repository to Pi
9. Modify `config.example.json` to have correct information
    1.  `capture_command` for Pi camera is `raspistill`. If you use something else, make sure to change it in the config file.
    2.  `capture_command_args` are passed in to the command. Modify these according to your needs. `%s` in this list of args denotes the file name where the image is saved.
    3.  `file_path` is a full path to the captured image. `%s` here is replaced with date and time.
    4.  `dropbox_token` can be acquired from https://www.dropbox.com/developers/apps/create. Create an app with access to only its directory, this is where the images will be.
10. Run the binary on Pi: `./pi-read-meter-armv6 config.json` (I hope you named your config `config.json`)
11. If all goes well, you should have the image in your Dropbox's app folder (`/Apps/APP_NAME` in Dropbox)
12. To make it run automatically, configure cron using `crontab -e` command. In there,
    1.  Specify cron shell, I used bash like this: `SHELL=/bin/bash`
    2.  Make it run every night: 
        ```
        0 0 * * * /home/pi/pi-read-meter-armv6 config.json >> /home/pi/pi-read-meter.log 2>&1
        ```
        1.  Assuming the binary is at `/home/pi/pi-read-meter-armv6`
        2.  Assuming the config is at `/home/pi/config.json`
        3.  Assuming that logs will be created to `/home/pi/pi-read-meter.log`
        4.  Feel free to change it however you like better
13. Now enjoy the pictures coming to Dropbox!

I plan to add OCR at some point as well, so that Pi would be able to send the data "Pi" itself! :heart_eyes:

# Contributing

Contributions are welcome!

If you are using it locally (in development), then use `config.dev.json` file for specifing the config.
