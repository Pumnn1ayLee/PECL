<p align="center">
  <img src="./assets/pecl-logo.png" width="164" alt="PECL logo" />
</p>

<h1 align="center">PECL</h1>

<p align="center">
  <strong>A clean Windows Minecraft launcher under active refinement</strong>
</p>

<p align="center">
  Versions · Java · Mods · Modpacks · Resource Management
</p>

<p align="center">
  <a href="./README.md">
    <img alt="简体中文" src="https://img.shields.io/static/v1?label=&message=%E7%AE%80%E4%BD%93%E4%B8%AD%E6%96%87&color=0f766e&style=for-the-badge">
  </a>
  <a href="./README.en.md">
    <img alt="English" src="https://img.shields.io/static/v1?label=&message=English&color=2457d6&style=for-the-badge">
  </a>
</p>

<p align="center">
  <a href="https://github.com/Pumnn1ayLee/PECL/releases/latest">
    <img alt="Download Latest" src="https://img.shields.io/static/v1?label=&message=Download%20Latest&color=2457d6&style=for-the-badge">
  </a>
  <a href="https://github.com/Pumnn1ayLee/PECL/releases">
    <img alt="All Releases" src="https://img.shields.io/static/v1?label=&message=All%20Releases&color=6b7280&style=for-the-badge">
  </a>
  <a href="https://github.com/Pumnn1ayLee/PECL/issues">
    <img alt="Issues" src="https://img.shields.io/static/v1?label=&message=Issues&color=f59e0b&style=for-the-badge">
  </a>
</p>

<p align="center">
  <img alt="Latest release" src="https://img.shields.io/github/v/release/Pumnn1ayLee/PECL?display_name=tag&style=flat-square&label=Latest">
  <img alt="Platform Windows" src="https://img.shields.io/badge/Platform-Windows-1b1f2a?style=flat-square">
</p>

<p align="center">
  <img src="./assets/pecl-screenshot-versions.jpg" alt="PECL versions page preview" width="960" />
</p>

---

## Overview

PECL is a Windows Minecraft launcher and resource management hub. The current focus is not just launching the game, but making this core workflow smoother and more reliable:

`Choose an instance -> search resources -> install -> launch the game -> recover from issues`

This public repository is mainly used to:

- publish installers and updater manifests
- present project information and public release notes
- collect public feedback and issue reports

The private development source repository is not mirrored here.

## Core Features

- Minecraft version installation and instance management
- Java detection, managed runtime downloads, auto selection, manual selection, and hide/restore support for external Java
- Forge, Fabric, and OptiFine workflow support
- Browsing, downloading, and managing mods, modpacks, resource packs, shader packs, and data packs
- Resource isolation, task progress sync, installation feedback, and updater delivery flow
- Community browsing and continuously evolving community-facing features

## Current Release And Links

- Latest public release: [v0.3.0](https://github.com/Pumnn1ayLee/PECL/releases/tag/v0.3.0)
- Recommended installer: `PECL_0.3.0_x64-setup.exe`
- Releases page: [GitHub Releases](https://github.com/Pumnn1ayLee/PECL/releases)

## 0.3.0 Highlights

- Fixed the Mod management flicker when switching into the Resources page from other pages
- Added continuous load-more support for modpack, resource pack, and shader pack browsers
- The Mod browser now performs a default search automatically when opened
- Improved Java management by separating PECL-managed and external Java, supporting hide and restore for external Java, and preventing hidden Java from participating in automatic fallback

## Download And Install

- Download the latest Windows installer from the [Releases](https://github.com/Pumnn1ayLee/PECL/releases) page
- The `x64` installer is recommended
- If Windows SmartScreen appears, verify the source before continuing

## Feedback

If you run into issues while using PECL, please open an Issue and include:

- your PECL version
- your Windows version
- the steps you were taking
- screenshots, logs, or error messages when available
