    ![]([34mhttps://cdn.rawgit.com/MichaelMu[0m
    [34mre/git-bug/master/misc/logo/logo-alp[0m
    [34m          ha-flat-bg.svg[0m)

    [32;1m1 git-bug[0;22m
    ────────────────────────────────────

    [![Build Status](https://travis-ci.o
    rg/MichaelMure/git-bug.svg?branch=ma
    ster)](https://travis-ci.org/Michael
    Mure/git-bug)[![Backers on Open Coll
    ective](https://opencollective.com/g
    it-bug/backers/badge.svg)](#backers)
    [![Sponsors on Open Collective](htt
    ps://opencollective.com/git-bug/spon
        sors/badge.svg)](#sponsors)
    [![License: GPL v3](https://img.shie
    lds.io/badge/License-GPLv3+-blue.svg
    )](http://www.gnu.org/licenses/gpl-3
    .0)[![GoDoc](https://godoc.org/githu
    b.com/MichaelMure/git-bug?status.svg
    )](https://godoc.org/github.com/Mich
    aelMure/git-bug)[![Go Report Card](h
    ttps://goreportcard.com/badge/github
    .com/MichaelMure/git-bug)](https://g
    oreportcard.com/report/github.com/Mi
    chaelMure/git-bug)[![Gitter chat](ht
    tps://badges.gitter.im/gitterHQ/gitt
    er.png)](https://gitter.im/the-git-b
                 ug/Lobby)

    [44;3mgit-bug[0;23m is a bug tracker that:
    [32m• [0m[1mfully embed in git[0m: you only need
      your git repository to have a bug
      tracker
    [32m• [0m[1mis distributed[0m: use your normal
      git remote to collaborate, push
      and pull your bugs !
    [32m• [0m[1mworks offline[0m: in a plane or under
      the sea ? keep reading and writing
      bugs
    [32m• [0m[1mprevent vendor locking[0m: your usual
      service is down or went bad ? you
      already have a full backup
    [32m• [0m[1mis fast[0m: listing bugs or opening
      them is a matter of milliseconds
    [32m• [0m[1mdoesn't pollute your project[0m: no
      files are added in your project
    [32m• [0m[1mintegrate with your tooling[0m: use
      the UI you like (CLI, terminal,
      web) or integrate with your
      existing tools through the CLI or
      the GraphQL API
    [32m• [0m[1mbridge with other bug trackers[0m:
      [bridges]([34m#bridges[0m) exist to
      import and soon export to other
      trackers.

    🚧  This is now more than a proof of
    concept, but still not fully stable.
    Expect dragons and unfinished
    business. 🚧

    [32;1m1.1 Install[0;22m

    [31m<details><summary>go[0m
    [31mget</summary>```shell[0m
    [31mgo get -u[0m
    [31mgithub.com/MichaelMure/git-bug[0m
    [31m```
[0m
    [31mIf it's not done already, add golang[0m
    [31mbinary directory in your PATH:
[0m
    [31m```bash[0m
    [31mexport PATH=$PATH:$(go env[0m
    [31mGOROOT)/bin:$(go env GOPATH)/bin[0m
    [31m```</details>[0m[31m<summary>go[0m
    [31mget</summary>[0mgo get[31m<summary>go[0m
    [31mget</summary>[0m```shellgo get -u
    github.com/MichaelMure/git-bug```If
    it's not done already, add golang
    binary directory in your
    PATH:```bashexport PATH=$PATH:$(go
    env GOROOT)/bin:$(go env
    GOPATH)/bin```[31m<details><summary>go[0m
    [31mget</summary>```shell[0m
    [31mgo get -u[0m
    [31mgithub.com/MichaelMure/git-bug[0m
    [31m```
[0m
    [31mIf it's not done already, add golang[0m
    [31mbinary directory in your PATH:
[0m
    [31m```bash[0m
    [31mexport PATH=$PATH:$(go env[0m
    [31mGOROOT)/bin:$(go env GOPATH)/bin[0m
    [31m```</details>[0m

    [31m<details><summary>Pre-compiled[0m
    [31mbinaries</summary>1. Go to the[0m
    [31m[release page](https://github.com/Mi[0m
    [31mchaelMure/git-bug/releases/latest)[0m
    [31mand download the appropriate binary[0m
    [31mfor your system.[0m
    [31m2. Copy the binary anywhere in your[0m
    [31mPATH[0m
    [31m3. Rename the binary to `git-bug`[0m
    [31m(or `git-bug.exe` on windows)
[0m
    [31mThat's all[0m
    [31m!</details>[0m[31m<summary>Pre-compiled[0m
    [31mbinaries</summary>[0mPre-compiled
    binaries[31m<summary>Pre-compiled[0m
    [31mbinaries</summary>[0m1. Go to the
    [release page](https://github.com/Mi
    chaelMure/git-bug/releases/latest)
    and download the appropriate binary
    for your system.2. Copy the binary
    anywhere in your PATH3. Rename the
    binary to `git-bug` (or
    `git-bug.exe` on windows)That's all
    ![31m<details><summary>Pre-compiled[0m
    [31mbinaries</summary>1. Go to the[0m
    [31m[release page](https://github.com/Mi[0m
    [31mchaelMure/git-bug/releases/latest)[0m
    [31mand download the appropriate binary[0m
    [31mfor your system.[0m
    [31m2. Copy the binary anywhere in your[0m
    [31mPATH[0m
    [31m3. Rename the binary to `git-bug`[0m
    [31m(or `git-bug.exe` on windows)
[0m
    [31mThat's all !</details>[0m

    [31m<details><summary>Linux[0m
    [31mpackages</summary>* [Archlinux (AUR)[0m
    [31m](https://aur.archlinux.org/packages[0m
    [31m/?K=git-bug)</details>[0m[31m<summary>Linux[0m
    [31mpackages</summary>[0mLinux
    packages[31m<summary>Linux[0m
    [31mpackages</summary>[0m* [Archlinux (AUR)
    ](https://aur.archlinux.org/packages
    /?K=git-bug)[31m<details><summary>Linux[0m
    [31mpackages</summary>* [Archlinux (AUR)[0m
    [31m](https://aur.archlinux.org/packages[0m
    [31m/?K=git-bug)</details>[0m

    [32;1m1.2 CLI usage[0;22m

    Create a new identity:

    [32;1m┃ [0;22mgit bug user create

    Create a new bug:

    [32;1m┃ [0;22mgit bug add

    Your favorite editor will open to
    write a title and a message.

    You can push your new entry to a
    remote:

    [32;1m┃ [0;22mgit bug push [<remote>]

    And pull for updates:

    [32;1m┃ [0;22mgit bug pull [<remote>]

    List existing bugs:

    [32;1m┃ [0;22mgit bug ls

    Filter and sort bugs using a
    [query]([34mdoc/queries.md[0m):

    [32;1m┃ [0;22mgit bug ls "status:open sort:edit"

    You can now use commands like [44;3mshow[0;23m,
    [44;3mcomment[0;23m, [44;3mopen[0;23m or [44;3mclose[0;23m to display
    and modify bugs. For more details
    about each command, you can run [44;3mgit[0m
    [3;44mbug <command> --help[0;23m or read the
    [command's
    documentation]([34mdoc/md/git-bug.md[0m).

    [32;1m1.3 Interactive terminal UI[0;22m

    An interactive terminal UI is
    available using the command [44;3mgit bug[0m
    [3;44mtermui[0;23m to browse and edit bugs.

    ![Termui recording]([34mtestdata_media/t[0m
    [34mermui_recording.gif[0m)


    [32;1m1.4 Web UI (status: WIP)[0;22m

    You can launch a rich Web UI with
    [44;3mgit bug webui[0;23m.

    ![Web UI screenshot
    1]([34mtestdata_media/webui1.png[0m)
    ![Web UI screenshot
    2]([34mtestdata_media/webui2.png[0m)


    This web UI is entirely packed
    inside the same go binary and serve
    static content through a localhost
    http server.

    The web UI interact with the backend
    through a GraphQL API. The schema is
    available [here]([34mgraphql/[0m).

    [32;1m1.5 Bridges[0;22m

    [92m1.5.1 Importer implementations[0m

    ┌─────────────────┬──────┬─────────┐
    │                 │Github│Launchpad│
    ╞═════════════════╪══════╪═════════╡
    │[1mincremental[0m[31m<br/>[0m(│✔️    │   ❌    │
    │can import more  │      │         │
    │than once)       │      │         │
    ├─────────────────┼──────┼─────────┤
    │[1mwith resume[0m[31m<br/>[0m(│❌    │   ❌    │
    │download only new│      │         │
    │data)            │      │         │
    ├─────────────────┼──────┼─────────┤
    │[1midentities[0m       │✔️    │   ✔️    │
    ├─────────────────┼──────┼─────────┤
    │identities update│❌    │   ❌    │
    ├─────────────────┼──────┼─────────┤
    │[1mbug[0m              │✔️    │   ✔️    │
    ├─────────────────┼──────┼─────────┤
    │comments         │✔️    │   ✔️    │
    ├─────────────────┼──────┼─────────┤
    │comment editions │✔️    │   ❌    │
    ├─────────────────┼──────┼─────────┤
    │labels           │✔️    │   ❌    │
    ├─────────────────┼──────┼─────────┤
    │status           │✔️    │   ❌    │
    ├─────────────────┼──────┼─────────┤
    │title edition    │✔️    │   ❌    │
    ├─────────────────┼──────┼─────────┤
    │[1mautomated test[0m   │❌    │   ❌    │
    │[1msuite[0m            │      │         │
    └─────────────────┴──────┴─────────┘
    [92m1.5.2 Exporter implementations[0m

    Todo !

    [32;1m1.6 Internals[0;22m

    Interested by how it works ? Have a
    look at the [data
    model]([34mdoc/model.md[0m) and the
    [internal
    bird-view]([34mdoc/architecture.md[0m).

    [32;1m1.7 Misc[0;22m

    [32m• [0m[Bash completion]([34m../misc/bash_com[0m
      [34mpletion[0m)
    [32m• [0m[Zsh completion]([34m../misc/zsh_compl[0m
      [34metion[0m)
    [32m• [0m[ManPages]([34mdoc/man[0m)

    [32;1m1.8 Planned features[0;22m

    [32m• [0mmedia embedding
    [32m• [0mexporter to github issue
    [32m• [0mextendable data model to support
      arbitrary bug tracker
    [32m• [0minflatable raptor

    [32;1m1.9 Contribute[0;22m

    PRs accepted. Drop by the [Gitter lo
    bby]([34mhttps://gitter.im/the-git-bug/L[0m
    [34mobby[0m) for a chat or browse the
    issues to see what is worked on or
    discussed.

    Developers unfamiliar with Go may
    try to clone the repository using
    "git clone". Instead, one should
    use:

    [32;1m┃ [0;22mgo get -u
    [32;1m┃ [0;22mgithub.com/MichaelMure/git-bug

    The git repository will then be
    available:

    [32;1m┃ [0;22m[3m[36m# Note that $GOPATH defaults to[0m
    [32;1m┃ [0;22m[3;36m$HOME/go[0m
    [32;1m┃ [0;22m$ [32mcd[0m [34m$GOPATH[0m/src/github.com/Michae
    [32;1m┃ [0;22mlMure/git-bug/

    You can now run [44;3mmake[0;23m to build the
    project, or [44;3mmake install[0;23m to install
    the binary in [44;3m$GOPATH/bin/[0;23m.

    To work on the web UI, have a look
    at [the dedicated
    Readme.]([34mwebui/Readme.md[0m)

    [32;1m1.10 Contributors ❤️ [0;22m

    This project exists thanks to all
    the people who contribute. [31m<a href="[0m
    [31mhttps://github.com/MichaelMure/git-b[0m
    [31mug/graphs/contributors">[0m[31m<img src="ht[0m
    [31mtps://opencollective.com/git-bug/con[0m
    [31mtributors.svg?width=890&button=false[0m
    [31m" />[0m[31m</a>[0m

    [32;1m1.11 Backers[0;22m

    Thank you to all our backers! 🙏
    [[Become a backer]([34mhttps://opencolle[0m
    [34mctive.com/git-bug#backer[0m)]

    [31m<a href="https://opencollective.com/[0m
    [31mgit-bug#backers"[0m
    [31mtarget="_blank">[0m[31m<img src="https://op[0m
    [31mencollective.com/git-bug/tiers/backe[0m
    [31mr.svg?width=890">[0m[31m</a>[0m

    [32;1m1.12 Sponsors[0;22m

    Support this project by becoming a
    sponsor. Your logo will show up here
    with a link to your website.
    [[Become a sponsor]([34mhttps://opencoll[0m
    [34mective.com/git-bug#sponsor[0m)]

    [31m<a href="https://opencollective.com/[0m
    [31mgit-bug/sponsor/0/website"[0m
    [31mtarget="_blank">[0m[31m<img src="https://op[0m
    [31mencollective.com/git-bug/tiers/spons[0m
    [31mor/0/avatar.svg">[0m[31m</a>[0m[31m<a href="https:[0m
    [31m//opencollective.com/git-bug/sponsor[0m
    [31m/1/website" target="_blank">[0m[31m<img src[0m
    [31m="https://opencollective.com/git-bug[0m
    [31m/tiers/sponsor/1/avatar.svg">[0m[31m</a>[0m[31m<a[0m
    [31mhref="https://opencollective.com/git[0m
    [31m-bug/sponsor/2/website"[0m
    [31mtarget="_blank">[0m[31m<img src="https://op[0m
    [31mencollective.com/git-bug/tiers/spons[0m
    [31mor/2/avatar.svg">[0m[31m</a>[0m[31m<a href="https:[0m
    [31m//opencollective.com/git-bug/sponsor[0m
    [31m/3/website" target="_blank">[0m[31m<img src[0m
    [31m="https://opencollective.com/git-bug[0m
    [31m/tiers/sponsor/3/avatar.svg">[0m[31m</a>[0m[31m<a[0m
    [31mhref="https://opencollective.com/git[0m
    [31m-bug/sponsor/4/website"[0m
    [31mtarget="_blank">[0m[31m<img src="https://op[0m
    [31mencollective.com/git-bug/tiers/spons[0m
    [31mor/4/avatar.svg">[0m[31m</a>[0m[31m<a href="https:[0m
    [31m//opencollective.com/git-bug/sponsor[0m
    [31m/5/website" target="_blank">[0m[31m<img src[0m
    [31m="https://opencollective.com/git-bug[0m
    [31m/tiers/sponsor/5/avatar.svg">[0m[31m</a>[0m[31m<a[0m
    [31mhref="https://opencollective.com/git[0m
    [31m-bug/sponsor/6/website"[0m
    [31mtarget="_blank">[0m[31m<img src="https://op[0m
    [31mencollective.com/git-bug/tiers/spons[0m
    [31mor/6/avatar.svg">[0m[31m</a>[0m[31m<a href="https:[0m
    [31m//opencollective.com/git-bug/sponsor[0m
    [31m/7/website" target="_blank">[0m[31m<img src[0m
    [31m="https://opencollective.com/git-bug[0m
    [31m/tiers/sponsor/7/avatar.svg">[0m[31m</a>[0m[31m<a[0m
    [31mhref="https://opencollective.com/git[0m
    [31m-bug/sponsor/8/website"[0m
    [31mtarget="_blank">[0m[31m<img src="https://op[0m
    [31mencollective.com/git-bug/tiers/spons[0m
    [31mor/8/avatar.svg">[0m[31m</a>[0m[31m<a href="https:[0m
    [31m//opencollective.com/git-bug/sponsor[0m
    [31m/9/website" target="_blank">[0m[31m<img src[0m
    [31m="https://opencollective.com/git-bug[0m
    [31m/tiers/sponsor/9/avatar.svg">[0m[31m</a>[0m

    [32;1m1.13 License[0;22m

    Unless otherwise stated, this
    project is released under the
    [GPLv3]([34mLICENSE[0m) or later license ©
    Michael Muré.

    The git-bug logo by [Viktor Teplov](
    [34mhttps://github.com/vandesign[0m) is
    released under the [Creative Commons
    Attribution 4.0 International (CC
    BY 4.0)]([34m../misc/logo/LICENSE[0m)
    license © Viktor Teplov.
