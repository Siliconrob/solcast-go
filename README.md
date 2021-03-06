[![Build Status](https://travis-ci.org/Siliconrob/solcast-go.svg?branch=master)](https://travis-ci.org/Siliconrob/solcast-go) [![Go Report Card](https://goreportcard.com/badge/github.com/Siliconrob/solcast-go)](https://goreportcard.com/report/github.com/Siliconrob/solcast-go) [![GoDoc](https://godoc.org/github.com/Siliconrob/solcast-go/solcast?status.svg)](https://godoc.org/github.com/Siliconrob/solcast-go/solcast) 

### Live Glitch :sun_with_face:[Demo](https://solcast-demo.glitch.me/):sun_with_face:

# Solcast Go API Client library

Go client library for querying the [Solcast API](https://api.solcast.com.au) 

This module allows a registered users to query the Solcast API [https://api.solcast.com.au](https://api.solcast.com.au).  You will need to register your user account to obtain an API key [https://solcast.com.au/api/register](https://solcast.com.au/api/register/).  Without an API key you will not be able to successfully obtain valid API results.

## Install
```
go get github.com/Siliconrob/solcast-go/solcast
```

## Documentation

Read [GoDoc](https://godoc.org/github.com/Siliconrob/solcast-go/solcast)

### Windows setup Solcast API Key

Setup **System/User** `environment variable`.  Details on advanced editing [StackOverflow superuser walkthrough](https://superuser.com/questions/949560/how-do-i-set-system-environment-variables-in-windows-10)

```
WinKey + R
```

**Copy and Paste** the following text to the **Open:** input text box for the Run Dialog

```
%windir%\System32\rundll32.exe sysdm.cpl,EditEnvironmentVariables
```

That will present this screen

![Run Windows Environment Editor](/imgs/win_launch_environment_editor.png)

Add a `USER` or `SYSTEM` `environment variable` to hold the Solcast API key.  User environment variables will only be available to your particular user, system environment variables are shared for all users on the system

![Add Windows User Solcast API key](/imgs/win_env_user_variable.png)

After you have added the `environment variable` you will see the key listed in the current variables

![Added Windows User Solcast API key](/imgs/win_solcast_variable.png)

**NOTE**: To reference this key you will need to reopen your shell prompt to read these variables again from the system (cmd, command.com, powershell, etc)


### Linux / mac OS

Open a terminal prompt
- mac OS: Spotlight search for `terminal`
- Linux: Open `bash`

```
nano .bash_profile
```

If you do not have `nano` it is a simpler text editor than `vi`.  Use your package manager to download and install or use `vi`.  The preferred package manager for mac OS is [Homebrew](https://brew.sh/) and once installed on your system you can issue similar commands to Linux `apt-get` and `yum` with the `brew` package manager.

Add the **Solcast API Key** to your user profile variables.

![Added mac OS User Solcast API key](/imgs/mac_os_environment_variable.png)

### Latitude and Longitude

- First as stated above you will need an API key to make **valid** API requests to the Solcast system.
- Second for all current library calls you will need a **valid** Lat/Lng coordinate in the [EPSG:4326](http://spatialreference.org/ref/epsg/wgs-84/) format.  If you are familiar with modern web maps you most likely have used the expected format or a decimal point that expresses a position on the Earth.

Clarification as I often forget the coordinate planes of Latitude and Longitude along with bounds.
![Lat/Lng](/imgs/Lat_Long.gif)

[Credits - Learner.org](http://www.learner.org/jnorth/tm/LongitudeIntro.html)

The Solcast API expects **West** for Longitude and **South** for Latitude to be expressed as a negative numbers

Example Locations on the Globe

Name | Latitude | Longitude
--- | --- | ---
Sydney, Australia | -33.865143 | 151.209900
Mumbai, India |‎ 19.228825 | 72.854118
Tokyo, Japan | 35.6895 | 139.69171
Paris, France | 48.864716 | 2.349014
Los Angeles, USA | 34.052235 | -118.243683

### Need help?
* [Documentation](https://solcast.com.au/api/docs/)
* [Forums](https://forums.solcast.com.au)
* [Gitter](https://gitter.im/Solcast/Lobby?utm_source=share-link&utm_medium=link&utm_campaign=share-link)

### How to contribute
 * Fork the repository
 * Add something awesome
 * Create a pull request
 * :sun_with_face: Celebrate :sun_with_face:


### Examples

```go

import (
	solcast "github.com/Siliconrob/solcast-go/solcast"
	datatypes "github.com/Siliconrob/solcast-go/solcast/types"
)

rad_forecasts := solcast.RadiationForecast(datatypes.LatLng{Longitude: -97, Latitude: 32})
```

#### JSON sample results
```
{ forecasts: 
   [ { ghi: 0,
       ghi90: 0,
       ghi10: 0,
       ebh: 0,
       dni: 0,
       dni10: 0,
       dni90: 0,
       dhi: 0,
       air_temp: 10,
       zenith: 103,
       azimuth: -88,
       cloud_opacity: 3,
       period_end: '2017-08-31T19:30:00.0000000Z',
       period: 'PT30M' },
    ...
   ]
}
```  

#### NOTE: 
You can use standard environment variables to hold your API key and not need to pass the optional `apiKey` to the current methods

Environment variable name if you care to work with it directly
```
SOLCAST_API_KEY
```

Accessible through process environment variable directly or with helper API method.


License
-------
License can be found here: [LICENSE](LICENSE)
