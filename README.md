<p align="center">
  <img src="./assets/pecl-logo.png" width="164" alt="PECL logo" />
</p>

<h1 align="center">PECL</h1>

<p align="center">
  <strong>轻量、干净、持续打磨中的 Windows Minecraft 启动器</strong><br />
  <strong>A clean Windows Minecraft launcher under active refinement</strong><br />
  <strong>軽量で洗練された Windows 向け Minecraft ランチャー</strong>
</p>

<p align="center">
  版本 · Java · Mods · 整合包 · 资源管理
</p>

<p align="center">
  <a href="#zh-cn">
    <img alt="简体中文" src="https://img.shields.io/static/v1?label=&message=%E7%AE%80%E4%BD%93%E4%B8%AD%E6%96%87&color=0f766e&style=for-the-badge">
  </a>
  <a href="#en-us">
    <img alt="English" src="https://img.shields.io/static/v1?label=&message=English&color=2457d6&style=for-the-badge">
  </a>
  <a href="#ja-jp">
    <img alt="日本語" src="https://img.shields.io/static/v1?label=&message=%E6%97%A5%E6%9C%AC%E8%AA%9E&color=7c3aed&style=for-the-badge">
  </a>
</p>

<p align="center">
  <a href="https://github.com/Pumnn1ayLee/PECL/releases/latest">
    <img alt="下载最新版 / Download Latest / 最新版をダウンロード" src="https://img.shields.io/static/v1?label=&message=%E4%B8%8B%E8%BD%BD%E6%9C%80%E6%96%B0%E7%89%88%20%2F%20Download%20Latest%20%2F%20%E6%9C%80%E6%96%B0%E7%89%88%E3%82%92%E3%83%80%E3%82%A6%E3%83%B3%E3%83%AD%E3%83%BC%E3%83%89&color=2457d6&style=for-the-badge">
  </a>
  <a href="https://github.com/Pumnn1ayLee/PECL/releases">
    <img alt="全部版本 / All Releases / すべてのリリース" src="https://img.shields.io/static/v1?label=&message=%E5%85%A8%E9%83%A8%E7%89%88%E6%9C%AC%20%2F%20All%20Releases%20%2F%20%E3%81%99%E3%81%B9%E3%81%A6%E3%81%AE%E3%83%AA%E3%83%AA%E3%83%BC%E3%82%B9&color=6b7280&style=for-the-badge">
  </a>
  <a href="https://github.com/Pumnn1ayLee/PECL/issues">
    <img alt="问题反馈 / Issues / フィードバック" src="https://img.shields.io/static/v1?label=&message=%E9%97%AE%E9%A2%98%E5%8F%8D%E9%A6%88%20%2F%20Issues%20%2F%20%E3%83%95%E3%82%A3%E3%83%BC%E3%83%89%E3%83%90%E3%83%83%E3%82%AF&color=f59e0b&style=for-the-badge">
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

<a id="zh-cn"></a>

## 简体中文

### 概览

PECL 是一个面向 Windows 的 Minecraft 启动器与资源管理工作台。当前重点不是单纯“能启动”，而是把下面这条核心链路做得更顺手、更稳定：

`选择实例 -> 搜索资源 -> 安装 -> 启动游戏 -> 出问题可恢复`

这个公开仓库主要用于：

- 发布安装包与更新清单
- 展示项目介绍与公开版本说明
- 收集公开反馈与使用问题

私有开发源码仓库不直接镜像到这里。

### 核心能力

- Minecraft 版本安装与实例管理
- Java 检测、托管下载、自动选择、手动指定与外部 Java 隐藏/恢复
- Forge、Fabric、OptiFine 工作流支持
- Mod、整合包、资源包、光影包、数据包的浏览、下载与管理
- 资源隔离、任务进度同步、安装反馈与更新器发布链路
- 社区浏览与持续扩展中的社区功能入口

### 当前版本与入口

- 最新公开版本：[v0.3.0](https://github.com/Pumnn1ayLee/PECL/releases/tag/v0.3.0)
- 推荐安装包：`PECL_0.3.0_x64-setup.exe`
- 公开版本说明：[docs/releases/v0.3.0.md](./docs/releases/v0.3.0.md)
- 发布页入口：[GitHub Releases](https://github.com/Pumnn1ayLee/PECL/releases)

### 0.3.0 更新重点

- 修复从其他页面切换到资源管理页时 Mod 管理闪烁的问题
- 整合包、资源包、光影包浏览器支持继续下拉加载更多结果
- 打开 Mod 浏览器时会自动执行默认搜索
- 优化 Java 管理：区分 PECL 托管与外部 Java，支持隐藏与恢复外部 Java，隐藏后不再参与自动回退

### 下载与使用

- 从 [Releases](https://github.com/Pumnn1ayLee/PECL/releases) 页面下载最新的 Windows 安装包
- 首选 `x64` 安装程序
- 如果 Windows SmartScreen 弹出提示，请确认来源后再继续

### 反馈

如果你在使用 PECL 时遇到问题，欢迎在这个仓库提交 Issue。建议附上这些信息：

- PECL 版本号
- Windows 版本
- 具体操作步骤
- 截图、日志或报错信息

---

<a id="en-us"></a>

## English

### Overview

PECL is a Windows Minecraft launcher and resource management hub. The current focus is not just launching the game, but making this core workflow smoother and more reliable:

`Choose an instance -> search resources -> install -> launch the game -> recover from issues`

This public repository is mainly used to:

- publish installers and updater manifests
- present project information and public release notes
- collect public feedback and issue reports

The private development source repository is not mirrored here.

### Core Features

- Minecraft version installation and instance management
- Java detection, managed runtime downloads, auto selection, manual selection, and hide/restore support for external Java
- Forge, Fabric, and OptiFine workflow support
- Browsing, downloading, and managing mods, modpacks, resource packs, shader packs, and data packs
- Resource isolation, task progress sync, installation feedback, and updater delivery flow
- Community browsing and continuously evolving community-facing features

### Current Release And Links

- Latest public release: [v0.3.0](https://github.com/Pumnn1ayLee/PECL/releases/tag/v0.3.0)
- Recommended installer: `PECL_0.3.0_x64-setup.exe`
- Public release notes: [docs/releases/v0.3.0.md](./docs/releases/v0.3.0.md)
- Releases page: [GitHub Releases](https://github.com/Pumnn1ayLee/PECL/releases)

### 0.3.0 Highlights

- Fixed the Mod management flicker when switching into the Resources page from other pages
- Added continuous load-more support for modpack, resource pack, and shader pack browsers
- The Mod browser now performs a default search automatically when opened
- Improved Java management by separating PECL-managed and external Java, supporting hide and restore for external Java, and preventing hidden Java from participating in automatic fallback

### Download And Install

- Download the latest Windows installer from the [Releases](https://github.com/Pumnn1ayLee/PECL/releases) page
- The `x64` installer is recommended
- If Windows SmartScreen appears, verify the source before continuing

### Feedback

If you run into issues while using PECL, please open an Issue and include:

- your PECL version
- your Windows version
- the steps you were taking
- screenshots, logs, or error messages when available

---

<a id="ja-jp"></a>

## 日本語

### 概要

PECL は Windows 向けの Minecraft ランチャー兼リソース管理ハブです。単にゲームを起動できるだけでなく、次の主要な流れをより快適で安定したものにすることを重視しています。

`インスタンスを選ぶ -> リソースを検索する -> インストールする -> ゲームを起動する -> 問題発生時に復旧する`

この公開リポジトリの主な用途は次のとおりです。

- インストーラーとアップデーター用マニフェストの配布
- プロジェクト情報と公開リリースノートの案内
- フィードバックと Issue の受付

非公開の開発用ソースリポジトリはここにはミラーされていません。

### 主な機能

- Minecraft バージョンのインストールとインスタンス管理
- Java 検出、PECL 管理ランタイムのダウンロード、自動選択、手動選択、外部 Java の非表示と復元
- Forge、Fabric、OptiFine のワークフロー対応
- Mod、Modpack、リソースパック、シェーダーパック、データパックの閲覧、ダウンロード、管理
- リソース分離、タスク進行状況の同期、インストールフィードバック、アップデーター配信フロー
- コミュニティ閲覧機能と継続的に拡張されるコミュニティ機能

### 現在の公開版とリンク

- 最新の公開版: [v0.3.0](https://github.com/Pumnn1ayLee/PECL/releases/tag/v0.3.0)
- 推奨インストーラー: `PECL_0.3.0_x64-setup.exe`
- 公開リリースノート: [docs/releases/v0.3.0.md](./docs/releases/v0.3.0.md)
- リリース一覧: [GitHub Releases](https://github.com/Pumnn1ayLee/PECL/releases)

### 0.3.0 の主な更新

- 他のページからリソース管理ページへ切り替えた際の Mod 管理画面のちらつきを修正
- Modpack、リソースパック、シェーダーパックのブラウザーで追加読み込みに対応
- Mod ブラウザーを開くと自動でデフォルト検索を実行
- Java 管理を改善し、PECL 管理 Java と外部 Java を区別。外部 Java の非表示と復元に対応し、非表示の Java は自動フォールバックに使われなくなりました

### ダウンロードとインストール

- 最新の Windows インストーラーは [Releases](https://github.com/Pumnn1ayLee/PECL/releases) ページからダウンロードしてください
- `x64` インストーラーの使用を推奨します
- Windows SmartScreen が表示された場合は、配布元を確認してから続行してください

### フィードバック

PECL の利用中に問題が発生した場合は、Issue を作成し、可能であれば次の情報を含めてください。

- PECL のバージョン
- Windows のバージョン
- 再現手順
- スクリーンショット、ログ、エラーメッセージ
