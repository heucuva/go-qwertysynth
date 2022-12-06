# go-qwertysynth
Very loose reinterpretation of [spencerharmon/qwertysynth](https://github.com/spencerharmon/qwertysynth) in Go.

Portions (C) [ebitengine.org](https://ebitengine.org/) and [Hajime Hoshi](https://hajimehoshi.com/).

## What is it?

This is a keyboard-operated synthesizer. It's not a very good one, at that.

# What systems does it work on?

* Windows (Windows 2000 or newer)
  * Sound Card
    * WinMM (`WAVE_MAPPER` device)
    * DirectSound (via optional build flag: `directsound`)
    * PulseAudio (via optional build flag: `pulseaudio`) - NOTE: Not recommended except for WSL (Linux) builds!
* Linux
  * Sound Card
    * PulseAudio

## How do I build this thing?

### What you need

For a Windows build, we recommend the following:
* Windows 2000 (or newer) - we used Windows 10 Pro (Windows 10 Version 20H2)
* Visual Studio Code
  * Go extension for VSCode v0.19.0 (or newer) 
  * Go v1.18.1 (or newer)

For a non-Windows (e.g.: Linux) build, we recommend the following:
* Ubuntu 20.04 (or newer) - we used Ubuntu 20.04.1 LTS running in WSL2
* Go v1.18.1 (or newer)
* The following libraries (versions listed are for Ubuntu 20.04):
  * libxcursor-dev (1:1.2.0-2)
  * libxi-dev (2:1.7.10-0ubuntu1)
  * libxinerama-dev (2:1.1.4-2)
  * libxrandr-dev (2:1.5.2-0ubuntu1)
  * libxxf86vm-dev (1:1.1.4-1build1)
  * libglfw3-dev (3.3.2-1)
  * libx11-dev (2:1.6.9-2ubuntu1.2)
  
  On Ubuntu, you can get these libraries with this command:
  ```bash
  sudo apt install libx11-dev libxcursor-dev libxrandr-dev libxinerama-dev libxi-dev libglfw3-dev libxxf86vm-dev -y
  ```
* On WSL2, we recommend using [VcXsrv 1.20.8.1](https://sourceforge.net/projects/vcxsrv/files/vcxsrv/1.20.8.1/).

### How to build (on Windows)

1. First, load the project folder in VSCode.  If this is the first time you've ever opened a Go project, VSCode will splash up a thousand alerts asking to install various things for Go. Allow it to install them before continuing on.
2. Next, open a Terminal for `powershell`.
3. Enter the following commands
   ```powershell
   go mod download
   go build
   ```
   When the command completes, you should now have the go-qwertysynth.exe file - run it with `.\go-qwertysynth.exe play` to start!

### How to build (on Linux)

1. Build the player with the following commands
   ```bash
   go mod download
   go build
   ```
   When the command completes, you should now have the go-qwertysynth file - run it with `./go-qwertysynth play` to start!

NOTE: In order to use PulseAudio, you must have your `PULSE_SERVER` connection string environment variable configured:
* e.g.:
  ```bash
  PULSE_SERVER=tcp:127.0.0.1:4713
  ```
  (*Take note that there are bugs associated with TCP connection strings; see bugs section below*)
  For more information about the `PULSE_SERVER` environment variable, please see the [PulseAudio documentation](https://www.freedesktop.org/wiki/Software/PulseAudio/Documentation/User/ServerStrings/).

## Waveform Configuration

Waveforms are configured using settings strings. An example setting string looks like this:
```text
sine,adsr:1s:80ms:-12.75db:1s,frequency:1.125,amplitude:16
```

The components of the above string are:

- name of the waveform generator to be used (`sine`)
- envelope settings (`adsr`) which are comprised of the following items:
  - attack duration (`1s`)
  - decay duration (`80ms`)
  - sustain level (`-12.75db`)
  - release duration (`1s`)
- additional waveform parameters:
  - the desired default waveform frequency (`frequency`) to be used (`1.125` - in Hertz)
  - the desired default waveform amplitude (`amplitude`) to be used (`16` - as a strict multiplier, but can be specified as dBv attenuation (`db`))

## Supported Waveform Generators and Parameters

Waveforms:

| Identifier | Name |
|----|----|
| `sine`, `sin` | Sine |
| `square`, `sqr` | Square |
| `triangle`, `tri` | Triangle |

Parameters:

| Identifier | Name | Default | Notes |
|----|----|----|----|
| `frequency`<br/>`freq`<br/>`f` | Frequency | C-4 | C-4 = `261.626 Hz`, usually, but depends on configured scale for synth machine used |
| `amplitude`<br/>`amp`<br/>`a` | Amplitude | 1.0 | For amplitude modulator, this is full volume<br/><br/>For frequency modulator, this is 1 full note microtone that the synth machine understands |

## Envelope

Envelopes are based on a simple ADSR mechanism. See [Waveform Configuration](#waveform-configuration) above for more information.

## FM Modulation

When configured (_i.e._: the waveform configuration is not empty or blank), a waveform will be applied to an amplitude modulator waveform as a note microtone (`Microtone`) value.

Microtones are synth machine specific implementations, but are usually defined as rational mini-steps between notes.

## Synth Machines

In order to generate a waveform and sample it for playback, a synth machine must be configured.

Supported synth machines:

| Identifier | Name | Details |
|----|----|----|
| `xm` | FastTracker II | 64 microtones of 1.5625 cents per microtone between notes<br/>12 notes per octave<br/>Central octave is 4<br/>A440 (equal) tuning |
| `it` | Impulse Tracker| 64 microtones of 1.5625 cents per microtone between notes<br/>12 notes per octave<br/>Central octave is 5<br/>A440 (equal) tuning |

Synth machines can also be configured with a tuning system. Available tuning systems include:

| Identifier | Name | Details |
|----|----|----|
| `default` | | Use the default value specified by the machine |
| `equal-A415`<br/>`A415` | A415 (equal) | A-4 is identified to be exactly 415.0 Hz<br/>[Twelve-tone equal temperament](https://en.wikipedia.org/wiki/12_equal_temperament) |
| `equal-A427`<br/>`A427` | A427 (equal) | A-4 is identified to be exactly 427.0 Hz<br/>Twelve-tone equal temperament |
| `equal-A428`<br/>`A428` | A428 (equal) | A-4 is identified to be exactly 428.0 Hz<br/>Twelve-tone equal temperament |
| `equal-A429`<br/>`A429` | A429 (equal) | A-4 is identified to be exactly 429.0 Hz<br/>Twelve-tone equal temperament |
| `equal-A430`<br/>`A430` | A430 (equal) | A-4 is identified to be exactly 430.0 Hz<br/>Twelve-tone equal temperament |
| `equal-A432`<br/>`A432` | A432 (equal) | A-4 is identified to be exactly 432.0 Hz<br/>Twelve-tone equal temperament |
| `equal-A435`<br/>`A435` | A435 (equal) | A-4 is identified to be exactly 435.0 Hz<br/>Twelve-tone equal temperament |
| `equal-A440`<br/>`A440` | A440 (equal) | A-4 is identified to be exactly 440.0 Hz<br/>Twelve-tone equal temperament |
| `equal-A444`<br/>`A444` | A444 (equal) | A-4 is identified to be exactly 444.0 Hz<br/>Twelve-tone equal temperament |
| `equal-A466`<br/>`A466` | A466 (equal) | A-4 is identified to be exactly 466.0 Hz<br/>Twelve-tone equal temperament |
| `equal-scientific`<br/>`scientific` | Scientific (equal) | C-4 is identified to be exactly 256.0 Hz<br/>Twelve-tone equal temperament |
| `equal-53`<br/>`53TET`<br/>`53` | 53-TET (equal) | A-4 is identified to be exactly 440.0 Hz<br/>[53 tone equal temperament](https://en.wikipedia.org/wiki/53_equal_temperament) |
| `just-harmonic`<br/>`harmonic` | Harmonic (just) | C-4 is identified to be 261.626 Hz<br/>Twelve-tone [just intonation](https://en.wikipedia.org/wiki/Just_intonation) |
| `just-pythagorean`<br/>`pythagorean` | Pythagorean (just) | D-4 is identified to be 288.325 Hz<br/>Twelve-tone just intonation |

## Synth Mode

You can access this via the `play` command-line parameter. For more details about settings of this mode, add the `--help` command-line parameter.

Keyboard Legend (for Twelve-tone intonations):

- `Q key`-row starts with C-5 on XM and C-6 on IT
- `A key`-row starts with C-4 on XM and C-5 on IT
- `Z key`-row starts with C-3 on XM and C-4 on IT
  - Note: there are not enough keys present on this row to get a full octave, so key mappings for A# and B are unavailable.
- `Escape` (`Esc`) to quit
- `Page Up` (`Pg Up`) to increase keyboard octave
- `Page Down` (`Pg Dn`) to decrease keyboard octave

Hold note keys to sustain notes; release the note keys to decay them.

Release note keys while holding `Shift` to cut/stop them.

Note: US English keyboard layout works best

## Tracker Mode

Included is a silly little tracker. It has a single portion of the song _The Celebrated Chop Waltz_ by Euphemia Allen (originally published under the pseudonym Arthur de Lulli) attached.

You can access this via the `play tracker` command-line parameters. For more details about settings of this mode, add the `--help` command-line parameter.

## Quirks

| Item | Notes |
|------|-------|
| `windows` `winmm` | Setting the number of channels to more than 2 may cause WinMM and/or qwertysynth to do unusual things. You might be able to get a hardware 4-channel capable card (such as the Aureal Vortex 2 AU8830) to work, but driver inconsistencies and weirdnesses in WinMM will undoubtedly cause needless strife. |
| `pulseaudio` | PulseAudio support is offered through a Pure Go interface originally created by Johann Freymuth, called [jfreymuth/pulse](https://github.com/jfreymuth/pulse). While it seems to work pretty well, it does have some inconsistencies when compared to the FreeDesktop supported C interface. If you see an error about there being a "`missing port in address`" specifically when using a TCP connection string, make sure to append the default port specifier of `:4713` to the end of the `PULSE_SERVER` environment variable. |
| `windows` `directsound` | DirectSound integration is not great code. It works well enough after recent code changes fixing event support, but it's still pretty ugly. |
