# CasaOS - Your Personal Cloud 
<!-- Readme i18n links -->
<!-- > English | [中文](#) | [Français](#) -->

<p align="center">
    <!-- CasaOS Banner -->
    <picture>
        <source media="(prefers-color-scheme: dark)" srcset="https://raw.githubusercontent.com/IceWhaleTech/logo/main/casaos/casaos_banner_dark_night_800x300.png">
        <source media="(prefers-color-scheme: light)" srcset="https://raw.githubusercontent.com/IceWhaleTech/logo/main/casaos/casaos_banner_twilight_blue_800x300.png">
        <img alt="CasaOS" src="https://raw.githubusercontent.com/IceWhaleTech/logo/main/casaos/casaos_banner_twilight_blue_800x300.png">
    </picture>
    <br/>
    <i>Connect with the community, establish autonomy, reduce the cost of SaaS, and MAXIMIZE the potential for a personalized copilot.</i>
    <br/>
    <br/>
    <!-- CasaOS Badges -->
    <a href="https://github.com/IceWhaleTech/CasaOS" target="_blank">
        <img alt="CasaOS Version" src="https://img.shields.io/github/v/release/IceWhaleTech/CasaOS?color=162453&style=flat-square&label=CasaOS" />
    </a>
    <a href="https://github.com/IceWhaleTech/CasaOS/blob/main/LICENSE" target="_blank">
        <img alt="CasaOS License" src="https://img.shields.io/github/license/IceWhaleTech/CasaOS?color=162453&style=flat-square&label=License" />
    </a>
    <a href="https://github.com/IceWhaleTech/CasaOS/pulls" target="_blank">
        <img alt="CasaOS Pull Requests" src="https://img.shields.io/github/issues-pr/IceWhaleTech/CasaOS?color=162453&style=flat-square&label=PRs" />
    </a>
    <a href="https://github.com/IceWhaleTech/CasaOS/issues" target="_blank">
        <img alt="CasaOS Issues" src="https://img.shields.io/github/issues/IceWhaleTech/CasaOS?color=162453&style=flat-square&label=Issues" />
    </a>
    <a href="https://codecov.io/gh/IceWhaleTech/CasaOS" > 
    <img src="https://codecov.io/gh/IceWhaleTech/CasaOS/branch/main/graph/badge.svg?token=l9uMKGlkxM"/> 
    </a>
    <a href="https://github.com/IceWhaleTech/CasaOS/stargazers" target="_blank">
        <img alt="CasaOS Stargazers" src="https://img.shields.io/github/stars/IceWhaleTech/CasaOS?color=162453&style=flat-square&label=Stars" />
    </a>
    <!-- <a href="https://github.com/IceWhaleTech/CasaOS/releases" target="_blank">
    <img alt="CasaOS Downloads" src="https://img.shields.io/github/downloads/IceWhaleTech/CasaOS/total?color=162453&style=flat-square" />
    </a> -->
    <br/>
    <!-- CasaOS Community -->
    <a href="https://discord.gg/knqAbbBbeX" target="_blank">
        <img alt="IceWhale Discord" src="https://img.shields.io/discord/884667213326463016?color=162453&style=flat-square&label=Discord&logo=discord&logoColor=fff" />
    </a>
    <a href="https://github.com/IceWhaleTech/CasaOS/discussions" target="_blank">
        <img alt="CasaOS GitHub Discussions" src="https://img.shields.io/github/discussions/IceWhaleTech/CasaOS?color=162453&style=flat-square&label=Discussions&logo=github" />
    </a>
<!-- ALL-CONTRIBUTORS-BADGE:START - Do not remove or modify this section -->
    <a href="#credits">
        <img alt="All Contributors" src="https://img.shields.io/static/v1?label=All%20Contributors&message=15&color=162453&style=flat-square&logo=Handshake&logoColor=fff" />
    </a>
    <!-- CasaOS YouTube -->
    <a href="https://www.youtube.com/channel/UC2zMrUYT17AJhIl9XWZzT8g" target="_blank">
        <img alt="YouTube Tutorial Views" src="https://img.shields.io/youtube/channel/views/UC2zMrUYT17AJhIl9XWZzT8g?style=flat-square&logo=youtube&logoColor=red&label=YouTube%20Tutorial%20Views" />
    </a>
    <br/>    
    <a href="http://bit.ly/45JQIiL" target="_blank">
        <img alt="twitter ZimaSpace" src="https://img.shields.io/twitter/follow/ZimaSpace?style=flat-square&logo=X&label=Contact%20Us%20%40%20ZimaSpace&labelColor=555&color=555" />
    </a>
    <a href="http://bit.ly/4lgHj7V" target="_blank">
        <img alt="facebook ZimaSpace" src="https://img.shields.io/badge/ZimaSpace-1877F2?style=flat-square&logo=Facebook&logoColor=fff&label=Contact%20Us&labelColor=555&color=162453" />
    </a>
    <br/>
    <!-- CasaOS Links -->
    <a href="https://www.casaos.io" target="_blank">Website</a> |
    <a href="http://demo.casaos.io" target="_blank">Demo</a> |
    <a href="https://github.com/IceWhaleTech/CasaOS" target="_blank">GitHub</a>
    <br/>
    <br/>
    <!-- CasaOS Snapshots -->
    <kbd>
      <picture>
          <source media="(prefers-color-scheme: dark)" srcset="snapshot-dark.jpg">
          <source media="(prefers-color-scheme: light)" srcset="snapshot-light.jpg">
          <img alt="CasaOS Snapshot" src="snapshot-light.jpg">
      </picture>
    </kbd>
</p>

## Migrate from official CasaOS

If CasaOS was installed from IceWhale's official installer, you can migrate to the latest release of this fork in place. Run the command from an SSH session or a directly attached terminal; do not uninstall CasaOS first.

```sh
curl -fsSL https://github.com/alvins82/CasaOS-Install/releases/latest/download/install.sh | sudo bash
```

The installer stops and restarts CasaOS services while upgrading, so back up important data first. After it finishes, verify the installed fork release with:

```sh
cat /var/lib/casaos/fork-release
```

The command should print the fork release tag installed on the system, such as `v0.4.26` for the current release.

Future updates can be installed from the CasaOS dashboard or by running the same command again. Do not use `get.casaos.io/update` after migrating, because it installs IceWhale's upstream component bundle.

## Fork Changelog

- **2026-08-13 — [CasaOS Install v0.4.26](https://github.com/alvins82/CasaOS-Install/releases/tag/v0.4.26):** Published the full fork bundle with [CasaOS v0.4.25](https://github.com/alvins82/CasaOS/releases/tag/v0.4.25), [CasaOS-UI v0.4.28](https://github.com/alvins82/CasaOS-UI/releases/tag/v0.4.28), and [CasaOS-LocalStorage v0.4.22](https://github.com/alvins82/CasaOS-LocalStorage/releases/tag/v0.4.22). The installer verifies the bundled components, compatibility overlay, manifest, and component lock with SHA-256 checksums.
- **2026-08-13 — [CasaOS v0.4.25](https://github.com/alvins82/CasaOS/releases/tag/v0.4.25):** Added authenticated Debian-family system package updates with live progress and completion reconciliation, kept system storage out of merged `/DATA` branches while preserving `/DATA/AppData`, advanced fork update discovery, and documented the fork component releases. PRs: [CasaOS #16](https://github.com/alvins82/CasaOS/pull/16), [CasaOS #17](https://github.com/alvins82/CasaOS/pull/17), [CasaOS #18](https://github.com/alvins82/CasaOS/pull/18), [CasaOS #19](https://github.com/alvins82/CasaOS/pull/19), [CasaOS #20](https://github.com/alvins82/CasaOS/pull/20), [CasaOS-UI #7](https://github.com/alvins82/CasaOS-UI/pull/7), and [CasaOS-LocalStorage #4](https://github.com/alvins82/CasaOS-LocalStorage/pull/4).
- **2026-08-13 — [CasaOS-UI v0.4.28](https://github.com/alvins82/CasaOS-UI/releases/tag/v0.4.28):** Added the persistent Show Search Bar setting, removed sidebar clipping, showed physical disk ownership and accurate filesystem usage, and added the in-dashboard system package update flow with completion reconciliation. PRs: [CasaOS-UI #8](https://github.com/alvins82/CasaOS-UI/pull/8), [CasaOS-UI #9](https://github.com/alvins82/CasaOS-UI/pull/9), [CasaOS-UI #10](https://github.com/alvins82/CasaOS-UI/pull/10).
- **2026-08-13 — [CasaOS-LocalStorage v0.4.22](https://github.com/alvins82/CasaOS-LocalStorage/releases/tag/v0.4.22):** Reported nested filesystem usage accurately and retained physical disk ownership in storage entries. PR: [CasaOS-LocalStorage #5](https://github.com/alvins82/CasaOS-LocalStorage/pull/5).
- **2026-08-12 — [v0.4.25](https://github.com/alvins82/CasaOS-Install/releases/tag/v0.4.25):** Kept system storage out of merged `/DATA` storage while preserving system AppData at `/DATA/AppData`, so installations on flash drives can use external disks for media storage.
- **2026-08-12 — [v0.4.21](https://github.com/alvins82/CasaOS-Install/releases/tag/v0.4.21):** Fixed fresh one-line installs and CasaOS Storage Manager merges on mergerfs 2.40, including persistent multi-disk pools on Ubuntu 26. PRs: [LocalStorage #2](https://github.com/alvins82/CasaOS-LocalStorage/pull/2), [LocalStorage #3](https://github.com/alvins82/CasaOS-LocalStorage/pull/3), [installer #11](https://github.com/alvins82/CasaOS-Install/pull/11).
- **2026-08-12 — [CasaOS-UI v0.4.25](https://github.com/alvins82/CasaOS-UI/releases/tag/v0.4.25):** Added the dedicated merged-storage tab, framed in-page app dialogs like App Store modals, applied the existing CasaOS zoom transition, and retained the close button with responsive full-screen behavior on mobile. PRs: [CasaOS-UI #1](https://github.com/alvins82/CasaOS-UI/pull/1), [CasaOS-UI #3](https://github.com/alvins82/CasaOS-UI/pull/3).
- **2026-08-12 — [v0.4.21](https://github.com/alvins82/CasaOS-Install/releases/tag/v0.4.21):** Added the CasaOS-UI v0.4.21 in-page app launcher so installed apps open inside the dashboard with a close button. Also fixed fresh one-line installs and CasaOS Storage Manager merges on mergerfs 2.40, including persistent multi-disk pools on Ubuntu 26. PRs: [CasaOS-UI #2](https://github.com/alvins82/CasaOS-UI/pull/2), [CasaOS #14](https://github.com/alvins82/CasaOS/pull/14), [CasaOS-Install #12](https://github.com/alvins82/CasaOS-Install/pull/12), [LocalStorage #2](https://github.com/alvins82/CasaOS-LocalStorage/pull/2), [LocalStorage #3](https://github.com/alvins82/CasaOS-LocalStorage/pull/3), [installer #11](https://github.com/alvins82/CasaOS-Install/pull/11).
- **2026-08-12 — [v0.4.20](https://github.com/alvins82/CasaOS-Install/releases/tag/v0.4.20):** Fixed OneDrive identity resolution by falling back to `createdBy.user.displayName` when Microsoft Graph omits `createdBy.user.email`, and returning a clear error when neither field is available. PR: [CasaOS #2530](https://github.com/IceWhaleTech/CasaOS/pull/2530).
- **2026-08-11 — [v0.4.19](https://github.com/alvins82/CasaOS-Install/releases/tag/v0.4.19):** Fixed in-app updates by running the installer outside the CasaOS service process group and preserving upgrade logs. PRs: [CasaOS #9](https://github.com/alvins82/CasaOS/pull/9), [installer #10](https://github.com/alvins82/CasaOS-Install/pull/10).
- **2026-08-11 — [v0.4.18](https://github.com/alvins82/CasaOS-Install/releases/tag/v0.4.18):** Made release naming platform-neutral while retaining automatic distro detection and amd64, arm64, and armv7 packages in one installer. PRs: [CasaOS #8](https://github.com/alvins82/CasaOS/pull/8), [installer #9](https://github.com/alvins82/CasaOS-Install/pull/9).
- **2026-08-11 — [v0.4.17-ubuntu26.3](https://github.com/alvins82/CasaOS-Install/releases/tag/v0.4.17-ubuntu26.3):** Routed dashboard update discovery and installation through CasaOS-Install releases so upstream updates cannot replace fork patches. PRs: [CasaOS #7](https://github.com/alvins82/CasaOS/pull/7), [installer #8](https://github.com/alvins82/CasaOS-Install/pull/8).
- **2026-08-11 — [v0.4.17-ubuntu26.2](https://github.com/alvins82/CasaOS-Install/releases/tag/v0.4.17-ubuntu26.2):** Added Ubuntu 26.04 and Docker 24–29 compatibility, multi-architecture patched components, clean-install and reboot validation, and a single-command installer. PRs: [CasaOS #1](https://github.com/alvins82/CasaOS/pull/1), [AppManagement #1](https://github.com/alvins82/CasaOS-AppManagement/pull/1), [installer #2](https://github.com/alvins82/CasaOS-Install/pull/2).

## Why do you need Personal Cloud?

In 2020, the team noticed three important trends:
- The cost of computing power and storage was decreasing fast.
- A part of cloud computing was moving towards edge computing.
- The issue of consumer data asset ownership and attribution had been ignored.

Based on these trends, the team proposed a thought experiment internally: what if personal clouds were available under $100 in next five years? This personal cloud would provide a low-cost data collaboration solution as a personal data center, storing and managing data for creators and small organizations. A distributed collaborative computing network can be formed by personal servers located around the world. It could also control and connect all smart devices, providing cross-ecosystem local intelligent services.

Furthermore, the personal cloud could combine personal data to train personalized AI assistants. The idea is that this technology would be an effective way to solve the issue of consumer data asset ownership and , as well as provide a more affordable and efficient computing solution for individuals and small organizations.

> If you think what we are doing is valuable. Please **give us a star ⭐** and **fork it 🤞**!

## Features

- Friendly UI designed for home scenarios
  - No code, no forms, intuitive, design for humanity
- Multiple hardware and base system support
  - ZimaBoard, NUC, RPi, old computers, whatever is available.
- Selected apps in the app store, one-click installation
  - Nextcloud, HomeAssistant, AdGuard, Jellyfin, *arr and more!
- Easily install numerous Docker apps
  - Over 100,000 apps from the Docker ecosystem can be easily installed
- Elegant drive and file management
  - What you see is what you get. No technical background required.
- Well-designed system/app widgets
  - What you care about, at a glance. Resource usage, app status, and more!

## Getting Started

CasaOS fully supports ZimaBoard, Intel NUC, and Raspberry Pi. Also, more computers and development boards and fully compatible with Ubuntu, Debian, Raspberry Pi OS, and CentOS with one-liner installation.

### Hardware Compatibility

- amd64 / x86-64
- arm64
- armv7

### System Compatibility

Official Support
- Debian 12 (✅ Tested, Recommended)
- Ubuntu Server 26.04 (✅ Tested by this fork)
- Ubuntu Server 20.04 (✅ Tested)
- Raspberry Pi OS (✅ Tested)

Community Support
- Elementary 6.1 (✅ Tested)
- Armbian 22.04 (✅ Tested)
- Alpine (🚧 Not Fully Tested Yet)
- OpenWrt (🚧 Not Fully Tested Yet)
- ArchLinux (🚧 Not Fully Tested Yet)

### Quick Setup CasaOS

Install this fork with the maintained [CasaOS-Install](https://github.com/alvins82/CasaOS-Install) release so that all component compatibility fixes are included.

Install the latest stable release with one command:

```sh
curl -fsSL https://github.com/alvins82/CasaOS-Install/releases/latest/download/install.sh | sudo bash
```

This runs the same released `install.sh` on every supported operating system and architecture.

Do not use `get.casaos.io` to install this fork; that endpoint installs IceWhaleTech's upstream component bundle.

### Update CasaOS

The CasaOS dashboard checks [CasaOS-Install releases](https://github.com/alvins82/CasaOS-Install/releases) and installs updates from this fork, so upstream updates cannot replace the patched components. Release tags are platform-neutral starting with `v0.4.18`; the installer detects the operating system and architecture at runtime. You can also repeat the installation command above at any time.

Existing `v0.4.17-ubuntu26.2` installations must run the installation command above once to enable the fork-aware updater; future releases can then be applied from the dashboard.

To check the installed fork release from a terminal, run:

```sh
cat /var/lib/casaos/fork-release
```



### Uninstall CasaOS


v0.3.3 or newer

```sh
casaos-uninstall
```

Before v0.3.3

```sh
curl -fsSL https://get.icewhale.io/casaos-uninstall.sh | sudo bash
```

## Community 

The word Casa comes from the Spanish word for "home". Project CasaOS originated as a pre-installed system for the crowdfunded product [ZimaBoard](https://www.zimaboard.com) on Kickstarter.

After looking at many systems and software on the market, the team found no server system designed for home scenarios, sadly true.

So, we set out to build this open-source project to develop CasaOS with our own hands, everyone in the community, and you.

We believe that through community-driven collaborative innovation and open communication with global developers, we can reshape the digital home experience like never before.

**A warm welcome for you to get help or share great ideas in the [Discord](https://discord.gg/knqAbbBbeX)!**

[![Discord Card](https://discordapp.com/api/guilds/884667213326463016/widget.png?style=banner2)](https://discord.gg/knqAbbBbeX)

## Contributing

CasaOS is a community-driven open source project and the people involved are CasaOS users. That means CasaOS will always need contributions from community members just like you!

- See <https://wiki.casaos.io/en/contribute> for ways of contributing to CasaOS
- See <https://wiki.casaos.io/en/contribute/development> if you want to be involved in code contribution specifically


## Credits

Many thanks to everyone who has helped CasaOS so far!

Everyone's contribution is greatly appreciated. ([Emoji Key](https://allcontributors.org/docs/en/emoji-key))

<!-- ALL-CONTRIBUTORS-LIST:START - Do not remove or modify this section -->
<!-- prettier-ignore-start -->
<!-- markdownlint-disable -->
<table>
  <tr>
    <td align="center"><a href="https://github.com/jerrykuku"><img src="https://avatars.githubusercontent.com/u/9485680?v=4?s=100" width="100px;" alt=""/><br /><sub><b>老竭力</b></sub></a><br /><a href="https://github.com/IceWhaleTech/CasaOS/commits?author=jerrykuku" title="Code">💻</a> <a href="https://github.com/IceWhaleTech/CasaOS/commits?author=jerrykuku" title="Documentation">📖</a> <a href="#ideas-jerrykuku" title="Ideas, Planning, & Feedback">🤔</a> <a href="#infra-jerrykuku" title="Infrastructure (Hosting, Build-Tools, etc)">🚇</a> <a href="#maintenance-jerrykuku" title="Maintenance">🚧</a> <a href="#platform-jerrykuku" title="Packaging/porting to new platform">📦</a> <a href="#question-jerrykuku" title="Answering Questions">💬</a> <a href="https://github.com/IceWhaleTech/CasaOS/pulls?q=is%3Apr+reviewed-by%3Ajerrykuku" title="Reviewed Pull Requests">👀</a></td>
    <td align="center"><a href="https://github.com/LinkLeong"><img src="https://avatars.githubusercontent.com/u/13556972?v=4?s=100" width="100px;" alt=""/><br /><sub><b>link</b></sub></a><br /><a href="https://github.com/IceWhaleTech/CasaOS/commits?author=LinkLeong" title="Code">💻</a> <a href="https://github.com/IceWhaleTech/CasaOS/commits?author=LinkLeong" title="Documentation">📖</a> <a href="#ideas-LinkLeong" title="Ideas, Planning, & Feedback">🤔</a> <a href="#infra-LinkLeong" title="Infrastructure (Hosting, Build-Tools, etc)">🚇</a> <a href="#maintenance-LinkLeong" title="Maintenance">🚧</a> <a href="#question-LinkLeong" title="Answering Questions">💬</a> <a href="https://github.com/IceWhaleTech/CasaOS/pulls?q=is%3Apr+reviewed-by%3ALinkLeong" title="Reviewed Pull Requests">👀</a></td>
    <td align="center"><a href="https://github.com/tigerinus"><img src="https://avatars.githubusercontent.com/u/7172560?v=4?s=100" width="100px;" alt=""/><br /><sub><b>太戈</b></sub></a><br /><a href="https://github.com/IceWhaleTech/CasaOS/commits?author=tigerinus" title="Code">💻</a> <a href="https://github.com/IceWhaleTech/CasaOS/commits?author=tigerinus" title="Documentation">📖</a> <a href="#ideas-tigerinus" title="Ideas, Planning, & Feedback">🤔</a> <a href="#infra-tigerinus" title="Infrastructure (Hosting, Build-Tools, etc)">🚇</a> <a href="#maintenance-tigerinus" title="Maintenance">🚧</a> <a href="#mentoring-tigerinus" title="Mentoring">🧑‍🏫</a> <a href="#security-tigerinus" title="Security">🛡️</a> <a href="#question-tigerinus" title="Answering Questions">💬</a> <a href="https://github.com/IceWhaleTech/CasaOS/pulls?q=is%3Apr+reviewed-by%3Atigerinus" title="Reviewed Pull Requests">👀</a></td>
    <td align="center"><a href="https://github.com/Lauren-ED209"><img src="https://avatars.githubusercontent.com/u/8243355?v=4?s=100" width="100px;" alt=""/><br /><sub><b>Lauren</b></sub></a><br /><a href="#ideas-Lauren-ED209" title="Ideas, Planning, & Feedback">🤔</a> <a href="#fundingFinding-Lauren-ED209" title="Funding Finding">🔍</a> <a href="#projectManagement-Lauren-ED209" title="Project Management">📆</a> <a href="#question-Lauren-ED209" title="Answering Questions">💬</a> <a href="https://github.com/IceWhaleTech/CasaOS/commits?author=Lauren-ED209" title="Tests">⚠️</a></td>
    <td align="center"><a href="https://JohnGuan.Cn"><img src="https://avatars.githubusercontent.com/u/3358477?v=4?s=100" width="100px;" alt=""/><br /><sub><b>John Guan</b></sub></a><br /><a href="#blog-JohnGuan" title="Blogposts">📝</a> <a href="#content-JohnGuan" title="Content">🖋</a> <a href="https://github.com/IceWhaleTech/CasaOS/commits?author=JohnGuan" title="Documentation">📖</a> <a href="#ideas-JohnGuan" title="Ideas, Planning, & Feedback">🤔</a> <a href="#eventOrganizing-JohnGuan" title="Event Organizing">📋</a> <a href="#mentoring-JohnGuan" title="Mentoring">🧑‍🏫</a> <a href="#question-JohnGuan" title="Answering Questions">💬</a> <a href="https://github.com/IceWhaleTech/CasaOS/pulls?q=is%3Apr+reviewed-by%3AJohnGuan" title="Reviewed Pull Requests">👀</a></td>
    <td align="center"><a href="https://blog.tippybits.com"><img src="https://avatars.githubusercontent.com/u/17506770?v=4?s=100" width="100px;" alt=""/><br /><sub><b>David Tippett</b></sub></a><br /><a href="https://github.com/IceWhaleTech/CasaOS/commits?author=dtaivpp" title="Documentation">📖</a> <a href="#ideas-dtaivpp" title="Ideas, Planning, & Feedback">🤔</a> <a href="#question-dtaivpp" title="Answering Questions">💬</a></td>
    <td align="center"><a href="https://github.com/zarevskaya"><img src="https://avatars.githubusercontent.com/u/60230221?v=4?s=100" width="100px;" alt=""/><br /><sub><b>Skaya</b></sub></a><br /><a href="#mentoring-zarevskaya" title="Mentoring">🧑‍🏫</a> <a href="#question-zarevskaya" title="Answering Questions">💬</a> <a href="#tutorial-zarevskaya" title="Tutorials">✅</a> <a href="#translation-zarevskaya" title="Translation">🌍</a></td>
  </tr>
  <tr>
    <td align="center"><a href="https://github.com/AuthorShin"><img src="https://avatars.githubusercontent.com/u/4959043?v=4?s=100" width="100px;" alt=""/><br /><sub><b>AuthorShin</b></sub></a><br /><a href="https://github.com/IceWhaleTech/CasaOS/commits?author=AuthorShin" title="Tests">⚠️</a> <a href="https://github.com/IceWhaleTech/CasaOS/issues?q=author%3AAuthorShin" title="Bug reports">🐛</a> <a href="#question-AuthorShin" title="Answering Questions">💬</a> <a href="#ideas-AuthorShin" title="Ideas, Planning, & Feedback">🤔</a></td>
    <td align="center"><a href="https://github.com/baptiste313"><img src="https://avatars.githubusercontent.com/u/93325157?v=4?s=100" width="100px;" alt=""/><br /><sub><b>baptiste313</b></sub></a><br /><a href="#translation-baptiste313" title="Translation">🌍</a></td>
    <td align="center"><a href="https://github.com/DrMxrcy"><img src="https://avatars.githubusercontent.com/u/58747968?v=4?s=100" width="100px;" alt=""/><br /><sub><b>DrMxrcy</b></sub></a><br /><a href="https://github.com/IceWhaleTech/CasaOS/commits?author=DrMxrcy" title="Tests">⚠️</a> <a href="#ideas-DrMxrcy" title="Ideas, Planning, & Feedback">🤔</a> <a href="#question-DrMxrcy" title="Answering Questions">💬</a></td>
    <td align="center"><a href="https://github.com/Joooost"><img src="https://avatars.githubusercontent.com/u/12090673?v=4?s=100" width="100px;" alt=""/><br /><sub><b>Joooost</b></sub></a><br /><a href="#ideas-Joooost" title="Ideas, Planning, & Feedback">🤔</a></td>
    <td align="center"><a href="https://potyarkin.ml"><img src="https://avatars.githubusercontent.com/u/334908?v=4?s=100" width="100px;" alt=""/><br /><sub><b>Vitaly Potyarkin</b></sub></a><br /><a href="#ideas-sio" title="Ideas, Planning, & Feedback">🤔</a></td>
    <td align="center"><a href="https://github.com/bearfrieze"><img src="https://avatars.githubusercontent.com/u/1023813?v=4?s=100" width="100px;" alt=""/><br /><sub><b>Bjørn Friese</b></sub></a><br /><a href="#ideas-bearfrieze" title="Ideas, Planning, & Feedback">🤔</a></td>
    <td align="center"><a href="https://github.com/Protektor-Desura"><img src="https://avatars.githubusercontent.com/u/1195496?v=4?s=100" width="100px;" alt=""/><br /><sub><b>Protektor</b></sub></a><br /><a href="https://github.com/IceWhaleTech/CasaOS/issues?q=author%3AProtektor-Desura" title="Bug reports">🐛</a> <a href="#ideas-Protektor-Desura" title="Ideas, Planning, & Feedback">🤔</a> <a href="#question-Protektor-Desura" title="Answering Questions">💬</a></td>
  </tr>
  <tr>
    <td align="center"><a href="https://github.com/llwaini"><img src="https://avatars.githubusercontent.com/u/59589857?v=4?s=100" width="100px;" alt=""/><br /><sub><b>llwaini</b></sub></a><br /><a href="#projectManagement-llwaini" title="Project Management">📆</a> <a href="https://github.com/IceWhaleTech/CasaOS/commits?author=llwaini" title="Tests">⚠️</a> <a href="#tutorial-llwaini" title="Tutorials">✅</a></td>
    <td align="center"><a href="https://github.com/CorrectRoadH"><img src="https://avatars.githubusercontent.com/u/29306285?v=4?s=100" width="100px;" alt=""/><br /><sub><b>CorrectRoadH</b></sub></a><br /><a href="https://github.com/IceWhaleTech/CasaOS/commits?author=correctroadh" title="Code">💻</a> <a href="https://github.com/IceWhaleTech/CasaOS/commits?author=correctroadh" title="Documentation">📖</a></td>
    <td align="center"><a href="https://github.com/zhanghengxin"><img src="https://avatars.githubusercontent.com/u/24197448?v=4?s=100" width="100px;" alt=""/><br /><sub><b>zhanghengxin</b></sub></a><br /><a href="https://github.com/IceWhaleTech/CasaOS/commits?author=zhanghengxin" title="Code">💻</a> <a href="https://github.com/IceWhaleTech/CasaOS/commits?author=zhanghengxin" title="Documentation">📖</a></td>
  </tr>
</table>

<!-- markdownlint-restore -->
<!-- prettier-ignore-end -->

<!-- ALL-CONTRIBUTORS-LIST:END -->

This project follows the [all-contributors](https://github.com/all-contributors/all-contributors) specification. Contributions of any kind are welcome!

## Changelog

Detailed changes for each release are documented in the [release notes](https://github.com/IceWhaleTech/CasaOS/releases).

---

<p align="center">
    <a href="https://dashboard.trackgit.com/token/l5q8egi92tfhlxd70l2l">
        <img src="https://us-central1-trackgit-analytics.cloudfunctions.net/token/ping/l5q8egi92tfhlxd70l2l" alt="trackgit-views" />
    </a>
</p>
