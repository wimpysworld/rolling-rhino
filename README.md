# Rolling Rhino

Convert Ubuntu into a *"rolling release"* that tracks the `devel` series.

<h1 align="center">
  <img src=".github/logo.png" alt="Rolling Rhino" />
  <br />
  Rolling Rhino
</h1>

<p align="center"><b>Simple shell script to make Ubuntu track the `devel` series.</b></p>
<!-- <div align="center"><img src=".github/screenshot.png" alt="Rolling Rhino Screenshot" /></div> -->
<p align="center">Made with 💝 for <img src="https://assets.ubuntu.com/v1/cb22ba5d-favicon-16x16.png" align="top" width="24" /></p>

## Introduction

Rolling Rhino is a simple tool to convert Ubuntu Desktop, and the official
desktop flavours, that has been installed from a daily image into a
*"rolling release"* by opting into and tracking the `devel` series. 

Rolling Rhino is intended for Ubuntu developers and experienced Ubuntu users
who want to install Ubuntu once and the track all development updates with
automatic tracking of subsequent series.

## Caveats

If you use Rolling Rhino to opt-in to `devel` series you're assuming support
of your system, including taking care of PPA migrations, cleaning
obsolete/orphaned packages and **actively participating in any issue resolution
for problems you may encounter** via [Launchpad](https://launchpad.net) using
tools such are `apport` and `ubuntu-bug`.

### Origins of Rolling Rhino

[Ubuntu Podcast](https://ubuntupodcast.org) had feedback about making Ubuntu
a rolling release, something we discussed during the main segment of
[S13E12 - Red Sky in the Morning](https://ubuntupodcast.org/2020/06/11/s13e12-red-sky-in-the-morning/)
and then covered again based on listener feedback during
[S13E14 - Ace of Spades](https://ubuntupodcast.org/2020/06/25/s13e14-ace-of-spades/).
During episode S13E14 guest presenter [Stuart Langridge](https://twitter.com/sil)
proposed *"Ubuntu Rolling Rhino"* as the name for a rolling Ubuntu release
along with some ideas as to how it could be implemented. [Sergio Schvezov](https://twitter.com/sergiusens) then
[followed up via Twitter reminding us that the `devel` series exists in Ubuntu](https://twitter.com/sergiusens/status/1276479711372292096).
This inspired me to create this `rolling-rhino` tool to somewhat implement
Stuart's idea by taking advantage of the `devel` series.

## Installation

Clone this repository:

```
git clone https://github.com/wimpysworld/rolling-rhino.git
```

## Usage

  * Install Ubuntu Desktop, or one of the desktop flavours, from a daily image.
    * [Ubuntu Desktop Daily Build](http://cdimage.ubuntu.com/daily-live/current/)
    * [Kubuntu Daily Build](http://cdimage.ubuntu.com/kubuntu/daily-live/current/)
    * [Lubuntu Daily Build](http://cdimage.ubuntu.com/lubuntu/daily-live/current/)
    * [Ubuntu Budgie Daily Build](http://cdimage.ubuntu.com/ubuntu-budgie/daily-live/current/)
    * [Ubuntu Kylin Daily Build](http://cdimage.ubuntu.com/ubuntukylin/daily-live/current/)
    * [Ubuntu MATE Daily Build](http://cdimage.ubuntu.com/ubuntu-mate/daily-live/current/)
    * [Ubuntu Studio Daily Build](http://cdimage.ubuntu.com/ubuntustudio/dvd/current/)
    * [Xubuntu Daily Build](http://cdimage.ubuntu.com/xubuntu/daily-live/current/)
  * Use `rolling-rhino` to convert it to a rolling release.

```
./rolling-rhino
```

Which will output something like this:

```
tbc
```

### All the options

Here are the full usage instructions:

```
tbc
```

## Credits

  * Thanks to [Stuart Langridge](https://twitter.com/sil) for [naming the project and proposing the idea]().
  * Thanks to [Sergio Schvezov](https://twitter.com/sergiusens) for [reminding me the `devel` series exists in Ubuntu](https://twitter.com/sergiusens/status/1276479711372292096).
  * Thanks to [RickAndTired](https://twitter.com/RickAndTired) for [answering the call for help](https://twitter.com/RickAndTired/status/1276729643068911618) and [making the Rolling Rhino logo](https://github.com/RickAndTired/Artwork).

## TODO

  - [ ] Detect system is running an Ubuntu Development Branch.
  - [ ] Detect desktop meta packages.
  - [ ] Detect PPAs.
  - [ ] Detect `sources.list` is not already tracking `devel`.
  - [ ] Create clean `sources.list` that tracks `devel`.
  - [ ] Use `yad` to [create a UI](https://sanana.kiev.ua/index.php/yad)