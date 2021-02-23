# StrongAuras

Provides an overlay for EverQuest .  This app will tail your character's log file and watch it for triggers.  You can configure what indicators that you are interested in. I will try and check in newer versions of `StrongAuras.exe` but that file may be out of date.  If you wish compile the binary by hand for any reason please follow the instructions below.

Example
![Example](img/example.png)

### Building

* Install Go
* Install deps via `go get ./...` in the project folder
* Run the build.bat file

### Configuration

Create a `config.json` file in the same folder as StrongAuras.exe.  It should look something like this:

```json
{
    "eq_folder": "C:\\Users\\Public\\Daybreak Game Company\\Installed Games\\EverQuest\\",
    "Character": "Bossman",
    "Server": "mangler",
    "window_position": {
        "x": 900,
        "y": 200,
        "w": 40,
        "h": 40
    },
    "indicators": [
        {
            "name": "Theif's eyes",
            "default": true,
            "show_when": ["Your battle sense falters"],
            "hide_when": ["You begin to see the weaknesses in your opponents"],
            "duration": 60,
            "x": 0,
            "y": 0,
            "sprite_sheet":"default\\spells01.tga",
            "sprite_box": {
                "x1": 80,
                "y1": 40,
                "x2": 120,
                "y2": 80
            }
        }
    ]
}


```

Settings:

* eq_folder - location on disk where your EverQuest install is located. (note the double slashes)
* Character - Character name that you want to monitor.  This is used to find your log file.
* Server - Server that you play on.  This is also used to find your log file.
* window_position - Positioning of the window if overlay_mode is set to window
* indicators - a list of indicators you wish to display
    * name - the name of the indicator.
    * default - `true` means it will show on load, `false` means it will be hidden on load
    * show_when - list of search text that will toggle the indicator to the visable state. If any log line contains this phrase the indicator will be displayed.
    * hide_when - list of search text that will toggle the indicator to the hidden state. If any log line contains this phrase the indicator will be hidden.
    * duration - Optional.  # Seconds for the indicator to remain hidden. If you would like this indicator to slowly come back you can configure this option of how long you want the indicator hidden for.
    * x - x position (pixels) from the top left corner
    * y - y position (pixels) from the top left corner
    * w - Optional. the width (pixels) if not configured it uses the source image size
    * h - Optional. the height (pixels) if not configured it uses the source image size
    * sprite_sheet - the EQ `.tga` sprite sheet to use for the indicator.
    * sprite_box - Optional. The top left and bottom right corner of the image.  If not provided the whole image is used.
        * For the EQ spell icons you can start from the top left (0,0 offset).  The images are `40`px by `40`px.  That means if you want the 3rd icon in the 2nd row it would be: 
            * "x1": 80,
            * "y1": 40,
            * "x2": 120,
            * "y2": 80